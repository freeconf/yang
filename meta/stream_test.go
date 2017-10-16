package meta_test

import "testing"
import "os"
import "github.com/c2stack/c2g/meta"
import "github.com/c2stack/c2g/c2"

func TestFileMetaStream(t *testing.T) {
	m := &meta.FileStreamSource{Root: "."}
	s, err := m.OpenStream("bogus", ".txt")
	if s != nil {
		t.Error("expected no stream")
	}
	if err != nil {
		t.Error("expected no err")
	}
}

func TestCachingStream(t *testing.T) {
	tmpDir := "./.var/TestCachingStream"
	os.RemoveAll(tmpDir)
	os.MkdirAll(tmpDir, 0777)
	c := meta.CacheSource{
		Local:    &meta.FileStreamSource{Root: tmpDir},
		Upstream: &meta.FileStreamSource{Root: "."},
	}
	noExist, err := c.OpenStream("bogus", ".txt")
	if noExist != nil {
		t.Fatal("found resource that shouldn't exist")
	}
	if err != nil {
		t.Fatal("expected no err")
	}
	exists, err := c.OpenStream("stream_test", ".go")
	if exists == nil {
		t.Fatal("found not resource")
	}
	if err != nil {
		t.Fatal("expected no err")
	}
	c2.DiffFiles(t, "stream_test.go", tmpDir+"/stream_test.go")
}
