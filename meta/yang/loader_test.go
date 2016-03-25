package yang

import (
	"log"
	"meta"
	"testing"
)

func LoadSampleModule(t *testing.T) *meta.Module {
	m, err := LoadModuleCustomImport(TestDataRomancingTheStone, nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	return m
}

func TestLoadLoader(t *testing.T) {
	//yyDebug = 4
	LoadSampleModule(t)
}

func TestFindByPath(t *testing.T) {
	m := LoadSampleModule(t)
	groupings := m.GetGroupings().(*meta.MetaContainer)
	found := meta.FindByPath(groupings, "team")
	log.Println("found", found)
	AssertIterator(t, m.DataDefs())
	if teams := meta.FindByPathWithoutResolvingProxies(m, "game/teams"); teams != nil {
		if def := meta.FindByPathWithoutResolvingProxies(teams.(meta.MetaList), "team"); def != nil {
			if team, isContainer := def.(*meta.Container); isContainer {
				AssertFindGrouping(t, team)
			} else {
				t.Error("Found team but it's not a container type")
			}
		} else {
			t.Error("FindByPathWithoutResolvingProxies Could not find ../teams/team")
		}
		AssertProxies(t, teams.(meta.MetaList))
	} else {
		t.Error("Could not find game/teams")
	}
}

func AssertIterator(t *testing.T, defs meta.MetaList) {
	i := meta.NewMetaListIterator(defs, false)
	game := i.NextMeta()
	if game == nil {
		t.Error("first and only child:game not found in module defs")
	} else if game.GetIdent() != "game" {
		t.Error("expected 'game' child but got ", game.GetIdent())
	} else {
		t.Log("Iterator passed")
	}
}

func AssertFindGrouping(t *testing.T, team *meta.Container) {
	if uses := team.GetFirstMeta(); uses != nil {
		if grouping := uses.(*meta.Uses).FindGrouping("team"); grouping != nil {
			t.Log("Found team grouping")
		} else {
			t.Error("Could not find 'team' grouping in 'team' container")
		}
	} else {
		t.Error("Could not find uses child in team")
	}
}

func AssertProxies(t *testing.T, teams meta.MetaList) {
	if def := meta.FindByPath(teams, "team"); def != nil {
		i := meta.NewMetaListIterator(def.(meta.MetaList), true)
		t.Log("first team child", i.NextMeta().GetIdent())
		i = meta.NewMetaListIterator(def.(meta.MetaList), true)
		c := meta.FindByIdent(i, "color")
		t.Log("color", c)

		if members := meta.FindByPath(def.(meta.MetaList), "members"); members != nil {
			t.Log("Found members from grouping")
		} else {
			t.Error("team grouping didn't resolve")
		}
	} else {
		t.Error("FindByPath could not find ../teams/team")
	}
}
