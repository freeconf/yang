package source

import (
	"embed"
	_ "embed"
	"testing"
)

func TestDirNotFound(t *testing.T) {
	m := Dir(".")
	s, err := m("bogus", ".txt")
	if s != nil {
		t.Error("expected no stream")
	}
	if err != nil {
		t.Error("expected no err")
	}
}

//go:embed testdata/*.yang
var nothing embed.FS

func TestEmbedDir_Nothing(t *testing.T) {
	s := EmbedDir(nothing, "testdata")
	y, err := s("nothing", ".yang")
	if err != nil {
		t.Fatalf("expected to find 'nothing.yang', got: %v", err)
	}
	if y == nil {
		t.Fatalf("expected reader")
	}
}

func TestEmbedDir_NotFound(t *testing.T) {
	s := EmbedDir(nothing, "testdata")
	y, err := s("notfound", ".yang")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if y != nil {
		t.Fatalf("expected no reader, got: %v", y)
	}
}

func TestEmbedDir_SpecialNothing(t *testing.T) {
	s := EmbedDir(nothing, "testdata")
	y, err := s("specialnothing", ".yang")
	if err != nil {
		t.Fatalf("expected to find 'nothing.yang', got: %v", err)
	}
	if y == nil {
		t.Fatalf("expected reader")
	}
}
