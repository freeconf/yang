package get

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestSync(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "sync-test")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		os.RemoveAll(tempDir)
	}()
	err = syncDir("../../../yang", tempDir)
	if err != nil {
		t.Fatal(err)
	}
	files, _ := ioutil.ReadDir(tempDir)
	if len(files) < 1 {
		t.Fatal("expected yang files to be copied")
	}
}
