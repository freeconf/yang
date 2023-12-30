package meta

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/freeconf/yang/fc"
)

// responsiblities:
//
// 1.) resolve expands all "uses" with selected "grouping" which includes
// refinements and augments.  Once this is complete the grouping, augments and
// refinement statements are no longer useful and can be removed from the schema
// tree.
//
// 2.) process imports which triggers whole new, recursive chain of processing. This
// is a form of resolving because imports are really just a way of grouping groupings into
// separate files
//
// The order and recursive calls in this module have been through extensive analysis to
// fix bugs and match expectations in RFC. For a diagram to understand the flow of
// this module and look for possible flaws, see
//
//	https://docs.google.com/drawings/d/14eozearBNP_vReNOMDGZmXnnC12RI-YFhWknRH0FSDI
func resolve(m *Module) error {
	r := &resolver{
		builder:        &Builder{},
		inProgressUses: make(map[*Grouping]*usesResolved),
		loadedModules:  make(map[string]*Module),
	}
	if err := r.module(m); err != nil {
		return err
	}
	if err := r.fillInRecursiveDefs(m); err != nil {
		return err
	}
	return nil
}

type usesUnresolved struct {
	parent   HasDataDefinitions
	resolved *usesResolved
	uses     *Uses
}

type usesResolved struct {
	grouping *Grouping
	defs     []Definition
}

type resolver struct {
	builder        *Builder
	inProgressUses map[*Grouping]*usesResolved
	unresolvedUses []*usesUnresolved
	loadedModules  map[string]*Module
	trace          bool
}

func (r *resolver) module(y *Module) error {
	r.loadedModules[y.ident] = y
	if y.featureSet != nil {
		if err := y.featureSet.Initialize(y); err != nil {
			return err
		}
	}

	// exand all includes
	if err := r.copyOverIncludes(y, y.includes); err != nil {
		return err
	}

	// expand all imports first because local uses may reference groupings in other files.
	if len(y.imports) > 0 {
		// imports were indexed by module name, but now that we know the
		// prefix, we need to reindex them
		byName := y.imports
		y.imports = make(map[string]*Import, len(byName))
		for _, i := range byName {
			if i.loader == nil {
				return fmt.Errorf("%s - no module loader defined", i.moduleName)
			}
			if i.prefix == "" {
				return fmt.Errorf("%s - prefix required on import", i.moduleName)
			}
			var err error
			var rev string
			if i.rev != nil {
				rev = i.rev.Ident()
			}
			if loaded, found := r.loadedModules[i.moduleName]; found {
				i.module = loaded
			} else {
				i.module, err = i.loader(nil, i.moduleName, rev, i.parent.featureSet, i.loader)
				if err != nil {
					return fmt.Errorf("%s - %s", i.moduleName, err)
				}
				// recurse
				if err = r.module(i.module); err != nil {
					return err
				}
			}

			// imports were originally added by module name, but now that we know the
			// prefix, we need to re-add them with proper key: prefix
			y.imports[i.Prefix()] = i
		}
	}

	//
	// now we can go into definitions and resolve "uses"
	//

	if _, err := r.enter(y); err != nil {
		return err
	}

	// expand and deviate AFTER uses because the targets we want to change
	// might not exist until after uses are expanded
	for _, a := range y.Augments() {

		// augments might have uses too
		if _, err := r.addDefinitions(a, a.popDataDefinitions()); err != nil {
			return err
		}

		if err := r.expandAugment(a, y); err != nil {
			return err
		}
	}

	for _, d := range y.Deviations() {
		if err := r.applyDeviation(y, d); err != nil {
			return err
		}
	}

	return nil
}

func (r *resolver) enter(d Definition) ([]Definition, error) {
	var err error

	if a, valid := d.(*Rpc); valid {
		if i := a.Input(); i != nil {
			if _, err := r.enter(i); err != nil {
				return nil, err
			}
		}

		if o := a.Output(); o != nil {
			if _, err := r.enter(o); err != nil {
				return nil, err
			}
		}
		return nil, nil
	}

	if hasCases, valid := d.(*Choice); valid {
		for _, cident := range hasCases.CaseIdents() {
			c := hasCases.cases[cident]
			if _, err := r.addDefinitions(c, c.popDataDefinitions()); err != nil {
				return nil, err
			}
		}
		return nil, nil
	}

	// we process actions and notification BEFORE containers and lists because
	// actions and notifications are be added as part of resolving uses in
	// datadefs lists and they are resolved then.
	if hasActions, valid := d.(HasActions); valid {
		for _, a := range hasActions.Actions() {
			if _, err := r.enter(a); err != nil {
				return nil, err
			}
		}
	}

	if hasNotification, valid := d.(HasNotifications); valid {
		for _, n := range hasNotification.Notifications() {
			if _, err := r.enter(n); err != nil {
				return nil, err
			}
		}
	}

	var added []Definition
	if hasDefs, valid := d.(HasDataDefinitions); valid {
		if added, err = r.addDefinitions(hasDefs, hasDefs.popDataDefinitions()); err != nil {
			return nil, err
		}
	}

	return added, nil
}

func (r *resolver) copyOverIncludes(main *Module, includes []*Include) error {
	for _, i := range includes {
		if i.loader == nil {
			return errors.New("no module loader defined")
		}
		var err error
		var rev string
		if i.rev != nil {
			rev = i.rev.Ident()
		}
		sub, err := i.loader(i.parent, i.subName, rev, i.parent.featureSet, i.loader)
		if err != nil {
			return errors.New(i.subName + " - " + err.Error())
		}
		if err := r.copyOverSubmoduleData(main, sub); err != nil {
			return err
		}
	}
	return nil
}

func (r *resolver) copyOverSubmoduleData(main *Module, sub *Module) error {
	// issue #50 - we need to copy over items instead of inserting directly
	// because not everything from submodule should come over
	for _, def := range sub.dataDefs {
		if err := main.addDataDefinition(def); err != nil {
			return err
		}
	}
	for _, n := range sub.notifications {
		if err := main.addNotification(n); err != nil {
			return err
		}
	}
	for _, a := range sub.actions {
		if err := main.addAction(a); err != nil {
			return err
		}
	}
	for _, i := range sub.identities {
		main.identities[i.ident] = i
	}
	for _, f := range sub.features {
		main.features[f.ident] = f
	}
	for _, e := range sub.extensionDefs {
		main.extensionDefs[e.ident] = e
	}
	for _, t := range sub.typedefs {
		main.typedefs[t.ident] = t
	}
	for _, g := range sub.groupings {
		main.groupings[g.ident] = g
	}
	for _, i := range sub.imports {
		main.imports[i.moduleName] = i
	}
	main.extensions = append(main.extensions, sub.extensions...)
	main.augments = append(main.augments, sub.augments...)
	main.deviations = append(main.deviations, sub.deviations...)
	return r.copyOverIncludes(main, sub.includes)
}

func (r *resolver) applyDeviation(y *Module, d *Deviation) error {
	target := Find(y, d.Ident())
	if target == nil {
		return fmt.Errorf("could not find target for deviation %s", d.Ident())
	}
	if d.NotSupported {
		switch target.(type) {
		case *Rpc:
			actions := target.Parent().(HasActions).Actions()
			delete(actions, target.Ident())
		case *Notification:
			notifs := target.Parent().(HasNotifications).Notifications()
			delete(notifs, target.Ident())
		default:
			hasDDefs := target.Parent().(HasDataDefinitions)
			existing := hasDDefs.popDataDefinitions()
			for _, candidate := range existing {
				if candidate != target {
					if err := hasDDefs.addDataDefinition(candidate); err != nil {
						return err
					}
				}
			}
		}
		return nil
	}
	// 7.20.3.2. in RFC details restrictions on deviations. I assume
	// violations are errors, not silent ignores.
	hasDets, _ := target.(HasDetails)
	hasType, _ := target.(Leafable)
	hasListDets, _ := target.(HasListDetails)
	if d.Add != nil {
		if d.Add.configPtr != nil {
			if hasDets.IsConfigSet() {
				return fmt.Errorf("config already set on %s", d.Ident())
			}
			hasDets.setConfig(*(d.Add).configPtr)
		}
		if d.Add.mandatoryPtr != nil {
			if hasDets.IsMandatorySet() {
				return fmt.Errorf("mandatory already set on %s", d.Ident())
			}
			hasDets.setMandatory(*(d.Add).mandatoryPtr)
		}
		if d.Add.maxElementsPtr != nil {
			if hasListDets.IsMaxElementsSet() {
				return fmt.Errorf("max-elements already set on %s", d.Ident())
			}
			hasListDets.setMaxElements(*(d.Add).maxElementsPtr)
		}
		if d.Add.minElementsPtr != nil {
			if hasListDets.IsMinElementsSet() {
				return fmt.Errorf("min-elements already set on %s", d.Ident())
			}
			hasListDets.setMinElements(*(d.Add).minElementsPtr)
		}
		for _, must := range d.Add.musts {
			target.(HasMusts).addMust(must)
		}
		if d.Add.units != "" {
			if hasType.Units() != "" {
				return fmt.Errorf("units already set on %s", d.Ident())
			}
			hasType.setUnits(d.Add.units)
		}
		if d.Add.HasDefault() {
			if hasType.HasDefault() {
				return fmt.Errorf("default already set on %s", d.Ident())
			}
			for _, deflt := range d.Add.Default() {
				hasType.addDefault(deflt)
			}
		}
		for _, unique := range d.Add.unique {
			target.(*List).unique = append(target.(*List).unique, unique)
		}
		for _, must := range d.Add.musts {
			target.(HasMusts).addMust(must)
		}
	}
	if d.Replace != nil {
		if d.Replace.configPtr != nil {
			if !hasDets.IsConfigSet() {
				return fmt.Errorf("config not set on %s", d.Ident())
			}
			hasDets.setConfig(*(d.Replace).configPtr)
		}
		if d.Replace.mandatoryPtr != nil {
			if !hasDets.IsMandatorySet() {
				return fmt.Errorf("mandatory not set on %s", d.Ident())
			}
			hasDets.setMandatory(*(d.Replace).mandatoryPtr)
		}
		if d.Replace.maxElementsPtr != nil {
			if !hasListDets.IsMaxElementsSet() {
				return fmt.Errorf("max-elements not set on %s", d.Ident())
			}
			hasListDets.setMaxElements(*(d.Replace).maxElementsPtr)
		}
		if d.Replace.minElementsPtr != nil {
			if !hasListDets.IsMinElementsSet() {
				return fmt.Errorf("min-elements not set on %s", d.Ident())
			}
			hasListDets.setMinElements(*(d.Replace).minElementsPtr)
		}
		if d.Replace.units != "" {
			if hasType.Units() == "" {
				return fmt.Errorf("units not set on %s", d.Ident())
			}
			hasType.setUnits(d.Replace.units)
		}
		if d.Replace.HasDefault() {
			if !hasType.HasDefault() {
				return fmt.Errorf("default not set on %s", d.Ident())
			}
			defaults := d.Replace.Default()
			if v, valid := hasType.(HasDefaultValues); valid {
				v.setDefault(defaults)
			} else if len(defaults) > 1 {
				return fmt.Errorf("only supports single default %s", d.Ident())
			} else {
				hasType.(HasDefaultValue).setDefault(defaults[0])
			}
		}
	}
	if d.Delete != nil {
		if d.Delete.units != "" {
			if hasType.Units() == d.Delete.units {
				return fmt.Errorf("cannot delete units '%s' != '%s' on %s",
					d.Delete.units, hasType.Units(), d.Ident())
			}
			hasType.setUnits("")
		}
		if d.Delete.HasDefault() {
			if hasType.DefaultValue() == d.Delete.DefaultValue() {
				return fmt.Errorf("cannot delete units '%s' != '%s' on %s",
					d.Delete.Default(), hasType.DefaultValue(),
					d.Ident())
			}
			hasType.clearDefault()
		}
		for _, unique := range d.Delete.unique {
			found := false
			var uniques [][]string
			for _, candidate := range target.(*List).unique {
				if isArrayStringEqual(unique, candidate) {
					found = true
				} else {
					uniques = append(uniques, candidate)
				}
			}
			if !found {
				return fmt.Errorf("unique entry %s not found on %s",
					strings.Join(unique, " "), d.Ident())
			}
			target.(*List).unique = uniques
		}
		for _, must := range d.Delete.musts {
			found := false
			var musts []*Must
			for _, candidate := range target.(HasMusts).Musts() {
				if candidate.Expression() == must.Expression() {
					found = true
				} else {
					musts = append(musts, candidate)
				}
			}
			if !found {
				return fmt.Errorf("must entry %s not found on %s",
					must.Expression(), d.Ident())
			}
			target.(HasMusts).setMusts(musts)
		}

	}
	return nil
}

func isArrayStringEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Strings(a)
	sort.Strings(b)
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// this methods adds data definitions (container, list, leaf...) back into
// their parent but resolves all "uses" while doing so. This operation is
// recursive so the module is the last container to have all it's children
// complete.
//
//	           M
//	         a1  b1
//	       c2      d3
//	order would be:
//	   Enter M, Enter A1, Enter c2, Leave c2, Leave a2, Enter b1,
//	   Enter d3, Leave d3, Leave b1, Leave M
func (r *resolver) addDefinitions(x HasDataDefinitions, defs []Definition) ([]Definition, error) {
	var added []Definition
	for _, def := range defs {
		more, err := r.addDataDefinition(x, def)
		if err != nil {
			return nil, err
		}
		if len(more) > 0 {
			added = append(added, more...)
		}
	}

	return added, nil
}

func (r *resolver) addDataDefinition(parent HasDataDefinitions, child Definition) ([]Definition, error) {
	if hasIf, valid := child.(HasIfFeatures); valid {
		if on, err := checkFeature(hasIf); err != nil || !on {
			return nil, err
		}
	}

	if u, isUses := child.(*Uses); isUses {
		return r.expandUses(parent, u)
	}
	if r.trace {
		fc.Debug.Printf("ADD %s/%s", parent.Ident(), child.Ident())
	}

	if err := parent.addDataDefinition(child); err != nil {
		return nil, err
	}

	if _, err := r.enter(child); err != nil {
		return nil, err
	}

	return []Definition{child}, nil
}

func (r *resolver) expandUses(parent HasDataDefinitions, u *Uses) ([]Definition, error) {
	g, err := r.findGrouping(u)
	if err != nil {
		return nil, err
	}

	var added []Definition
	resolved, recursive := r.inProgressUses[g]
	if recursive {
		return r.delayRecursiveUses(parent, u, resolved)
	}

	if r.trace {
		fc.Debug.Printf("USE %s:%s", parent.Ident(), u.Ident())
	}

	resolved = &usesResolved{grouping: g}
	r.inProgressUses[g] = resolved

	// resolve all children
	groupDefs := r.cloneDefs(parent, g.DataDefinitions(), u.when)
	more, err := r.addDefinitions(parent, groupDefs) // recurse
	if err != nil {
		return nil, err
	}
	if len(more) > 0 {
		added = append(added, more...)
	}

	// copy in any actions or notifications unresolved, they will be resolved
	// in caller loop
	for _, a := range g.Actions() {
		hasActions, validActions := parent.(HasActions)
		if !validActions {
			return nil, fmt.Errorf("cannot add %s. %s does not allow actions", u.ident, SchemaPath(u))
		}
		rpc := a.clone(parent).(*Rpc)
		if err = hasActions.addAction(rpc); err != nil {
			return nil, err
		}
		if _, err := r.enter(rpc); err != nil {
			return nil, err
		}
	}
	for _, a := range g.Notifications() {
		hasNotifs, validNotifs := parent.(HasNotifications)
		if !validNotifs {
			return nil, fmt.Errorf("cannot add %s. %s does not allow notifications", u.ident, SchemaPath(u))
		}
		notify := a.clone(parent).(*Notification)
		if err = hasNotifs.addNotification(notify); err != nil {
			return nil, err
		}
		if _, err := r.enter(notify); err != nil {
			return nil, err
		}
	}

	if err := r.applyRefinements(u, parent); err != nil {
		return nil, err
	}

	for _, a := range u.augments {
		err = r.expandAugment(a, parent)
		if err != nil {
			return nil, err
		}
		if len(more) > 0 {
			added = append(added, more...)
		}
	}

	if r.trace {
		fc.Debug.Printf("!USE %s:%s", parent.Ident(), u.Ident())
	}
	delete(r.inProgressUses, g)
	resolved.defs = added

	return added, nil
}

func (r *resolver) delayRecursiveUses(parent HasDataDefinitions, u *Uses, resolved *usesResolved) ([]Definition, error) {
	if r.trace {
		fc.Debug.Printf("QUE %s:%s", parent.Ident(), u.Ident())
	}

	// detected a recursive uses so resolve this uses later once the root
	// uses has bee completely resolved.
	r.unresolvedUses = append(r.unresolvedUses, &usesUnresolved{parent, resolved, u})
	// fmt.Printf("%s : %s <= %s \n", u.ident, SchemaPath(master), SchemaPath(parent))

	// used as a placeholder to be replaced at end when all unresolved uses are resolved
	if err := parent.addDataDefinition(u); err != nil {
		return nil, err
	}
	// this uses has the potential of being in list of resolved defs.  if so, then it would
	// trigger another pass when resolving at end.
	return []Definition{u}, nil
}

// replace all the *Uses that were detected as recursive and left as placeholder with the
// now resolved list of definitions.  There's a chance the resolved list might also have
// placeholders so loop until all placeholders are replaced.
func (r *resolver) fillInRecursiveDefs(root *Module) error {
	for len(r.unresolvedUses) > 0 {
		if r.trace {
			fc.Debug.Printf("DEQUE %d items", len(r.unresolvedUses))
		}
		unresolved := r.unresolvedUses
		r.unresolvedUses = make([]*usesUnresolved, 0)
		for _, entry := range unresolved {
			if r.trace {
				fc.Debug.Printf("USE %s:%s", entry.parent.Ident(), entry.uses.Ident())
			}
			existing := entry.parent.popDataDefinitions()
			replaced := false
			for _, def := range existing {
				if def == entry.uses {
					if r.trace {
						fc.Debug.Printf("RPL %s.%s", entry.parent.Ident(), def.Ident())
					}
					replaced = true
					for _, subdef := range entry.resolved.defs {
						// watch for unrelated, unresolved uses in list of "resolved" defs that
						// will required another pass
						if u, isUses := subdef.(*Uses); isUses {
							if r.trace {
								fc.Debug.Printf("delayed: resubmitting %s.%s", entry.parent.Ident(), subdef.Ident())
							}
							subr := findResolved(unresolved, u)
							r.unresolvedUses = append(r.unresolvedUses, &usesUnresolved{entry.parent, subr, u})
						}
						if err := entry.parent.addDataDefinitionWithoutOwning(subdef); err != nil {
							return err
						}
					}
				} else {
					if err := entry.parent.addDataDefinitionWithoutOwning(def); err != nil {
						return err
					}
				}
			}
			if !replaced {
				return fmt.Errorf("did not resolve %s", SchemaPathNoModule(entry.uses))
			}
		}
	}
	return nil
}

func findResolved(unresolved []*usesUnresolved, target *Uses) *usesResolved {
	for _, e := range unresolved {
		if e.uses == target {
			return e.resolved
		}
	}
	// bug in freeconf somewhere, cannot be from bad yang file, so panic here
	panic(fmt.Sprintf("could not find resolved list for %s", SchemaPath(target)))
}

func (r *resolver) cloneDefs(parent HasDataDefinitions, defs []Definition, when *When) []Definition {
	copy := make([]Definition, len(defs))
	for i, d := range defs {
		copy[i] = d.(cloneable).clone(parent).(Definition)
		if when != nil {
			copy[i].(HasWhen).setWhen(when)
		}
	}
	return copy
}

func (r *resolver) findGrouping(y *Uses) (*Grouping, error) {
	prefix, ident := splitIdent(y.Ident())

	// From RFC
	//   A reference to an unprefixed type or grouping, or one that uses the
	//   prefix of the current module, is resolved by locating the matching
	//   "typedef" or "grouping" statement among the immediate substatements
	//   of each ancestor statement.
	// this means if prefix is local module, then ignore it and follow chain
	searchHeirarcy := (prefix == "")
	var module *Module
	if !searchHeirarcy {
		m, isExternal, err := findModuleAndIsExternal(y, prefix)
		if err != nil {
			return nil, err
		}
		if !isExternal {
			searchHeirarcy = true
		} else {
			module = m
		}
	}

	var found *Grouping
	if searchHeirarcy {
		var p Definition
		p = y
		for p != nil {
			if ptd, ok := p.(HasGroupings); ok {
				if found = ptd.Groupings()[ident]; found != nil {
					return found, nil
				}
			}
			p = p.getOriginalParent()
			if p != nil {
				// issue #50 - submodules can reference types in parent and in any
				// other submodule w/o prefix
				if m, isModule := p.(*Module); isModule && m.belongsTo != nil {
					p = m.Parent().(Definition)
				}
			}
		}
	} else {
		found = module.Groupings()[ident]
	}

	if found == nil {
		return nil, fmt.Errorf("%s - %s group not found", SchemaPath(y), y.Ident())
	}
	return found, nil
}

func (r *resolver) applyRefinements(u *Uses, parent Definition) error {
	for _, refine := range u.refines {
		if on, err := checkFeature(refine); !on || err != nil {
			return err
		}
		target := Find(parent.(HasDataDefinitions), refine.Ident())
		if target == nil {
			return fmt.Errorf("%s:could not find target for refine %s", SchemaPath(u), refine.Ident())
		}
		if err := r.refine(target, refine); err != nil {
			return err
		}
	}
	return nil
}

func (r *resolver) refine(target Definition, y *Refine) error {
	if y.desc != "" {
		r.builder.Description(target, y.desc)
	}
	if y.ref != "" {
		r.builder.Reference(target, y.ref)
	}
	if y.HasDefault() {
		r.builder.Defaults(target, y.Default())
	}
	if y.configPtr != nil {
		r.builder.Config(target, *y.configPtr)
	}
	if y.mandatoryPtr != nil {
		r.builder.Mandatory(target, *y.mandatoryPtr)
	}
	if y.maxElementsPtr != nil {
		r.builder.MaxElements(target, *y.maxElementsPtr)
	}
	if y.minElementsPtr != nil {
		r.builder.MinElements(target, *y.minElementsPtr)
	}
	if y.unboundedPtr != nil {
		r.builder.UnBounded(target, *y.unboundedPtr)
	}
	for _, m := range y.Musts() {
		h, valid := target.(HasMusts)
		if !valid {
			r.builder.setErr(fmt.Errorf("%T does not support must", target))
		} else {
			h.addMust(m.clone(target).(*Must))
		}
	}
	return r.builder.LastErr
}

func (r *resolver) expandAugment(y *Augment, parent Meta) error {
	if on, err := checkFeature(y); !on || err != nil {
		return err
	}

	// RFC7950 Sec 7.17
	// "The target node MUST be either a container, list, choice, case, input,
	//   output, or notification node."
	target := Find(parent.(HasDataDefinitions), y.ident)
	if target == nil {
		return fmt.Errorf("%s - augment target is not found %s", SchemaPath(y), y.ident)
	}

	targetChoice, targetIsChoice := target.(*Choice)
	for _, orig := range y.DataDefinitions() {
		var err error
		d := orig.(cloneable).clone(target).(Definition)
		if targetIsChoice {
			if cs, isCase := d.(*ChoiceCase); isCase {
				if err = targetChoice.addCase(cs); err != nil {
					return err
				}
				_, err = r.enter(cs)
			} else {
				// add implied case
				cs := r.builder.Case(target, d.Ident())
				_, err = r.addDataDefinition(cs, d)
			}
		} else if parentDef, hasDefs := target.(HasDataDefinitions); hasDefs {
			_, err = r.addDataDefinition(parentDef, d)
		} else {
			// TODO: Support RCP Input and Output, choice cases as targets
			return fmt.Errorf("%T not a recognizable parent for ", parent)
		}
		if err != nil {
			return err
		}
	}

	for _, orig := range y.Actions() {
		d := orig.clone(target).(Definition)
		if err := target.(HasActions).addAction(d.(*Rpc)); err != nil {
			return err
		}
		if _, err := r.enter(d); err != nil {
			return err
		}
	}

	for _, orig := range y.Notifications() {
		d := orig.clone(target).(Definition)
		if err := target.(HasNotifications).addNotification(d.(*Notification)); err != nil {
			return err
		}
		if _, err := r.enter(d); err != nil {
			return err
		}
	}

	return nil
}
