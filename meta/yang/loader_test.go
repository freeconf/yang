package yang

import (
	"log"
	"testing"

	"github.com/c2stack/c2g/meta"
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
	found, err := meta.FindByPath(groupings, "team")
	if err != nil {
		t.Error(err)
	}
	log.Println("found", found)
	AssertIterator(t, m.DataDefs())
	if teams, err := meta.FindByPathWithoutResolvingProxies(m, "game/teams"); err != nil {
		t.Error(err)
	} else if teams != nil {
		if def, err := meta.FindByPathWithoutResolvingProxies(teams.(meta.MetaList), "team"); err != nil {
			t.Error(err)
		} else if def != nil {
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
	i := meta.Children(defs, false)
	game, err := i.Next()
	if err != nil {
		t.Error(err)
	} else if game == nil {
		t.Error("first and only child:game not found in module defs")
	} else if game.GetIdent() != "game" {
		t.Error("expected 'game' child but got ", game.GetIdent())
	} else {
		t.Log("Iterator passed")
	}
}

func AssertFindGrouping(t *testing.T, team *meta.Container) {
	if uses := team.GetFirstMeta(); uses != nil {
		if grouping, err := uses.(*meta.Uses).FindGrouping("team"); err != nil {
			t.Error(err)
		} else if grouping != nil {
			t.Log("Found team grouping")
		} else {
			t.Error("Could not find 'team' grouping in 'team' container")
		}
	} else {
		t.Error("Could not find uses child in team")
	}
}

func AssertProxies(t *testing.T, teams meta.MetaList) {
	if def, err := meta.FindByPath(teams, "team"); err != nil {
		t.Error(err)
	} else if def != nil {
		i := meta.Children(def.(meta.MetaList), true)
		if m, err := i.Next(); err != nil {
			t.Error(err)
		} else {
			t.Log("first team child", m.GetIdent())
		}
		i = meta.Children(def.(meta.MetaList), true)
		if c, err := meta.FindByIdent(i, "color"); err != nil {
			t.Error(err)
		} else {
			t.Log("color", c)

		}

		if members, err := meta.FindByPath(def.(meta.MetaList), "members"); err != nil {
			t.Error(err)
		} else if members != nil {
			t.Log("Found members from grouping")
		} else {
			t.Error("team grouping didn't resolve")
		}
	} else {
		t.Error("FindByPath could not find ../teams/team")
	}
}
