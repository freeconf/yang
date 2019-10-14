package source

import (
	"os"
	"testing"

	"github.com/freeconf/yang/fc"
)

func TestCache(t *testing.T) {
	tmpDir := "./.var/TestCachingStream"

	os.RemoveAll(tmpDir)
	os.MkdirAll(tmpDir, 0777)
	opener := Cached(Dir("."), DirCache(tmpDir))
	noExist, err := opener("bogus", ".txt")
	if noExist != nil {
		t.Fatal("found resource that shouldn't exist")
	}
	if err != nil {
		t.Fatal("expected no err")
	}
	exists, err := opener("cache_test", ".go")
	if exists == nil {
		t.Fatal("found not resource")
	}
	if err != nil {
		t.Fatal("expected no err")
	}
	fc.DiffFiles(t, "cache_test.go", tmpDir+"/cache_test.go")
}
