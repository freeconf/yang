package meta

import "testing"
import "os"
import "github.com/c2stack/c2g/c2"
import "bytes"

func TestCachingStreamSource(t *testing.T) {
	shouldNotCache := c2.AltFs{
		OnCreate: func(name string) (File, error) {
			t.Error("Should not create file " + name)
			return nil, nil
		},
	}
	shouldCache := c2.AltFs{
		OnStat: os.Stat,
		OnCreate: func(name string) (File, error) {
			return c2.AltFile{bytes.NewBuffer()}, nil
		},
	}
	upstream := FileStreamSource{Root: "."}
	c := CachingStreamSource{
		Dir:    ".",
		Stream: upstream,
		fs:     shouldNotCache,
	}
	c.OpenStream("stream_test", ".go")
	c.fs = shouldCache
	c.OpenStream("stream_test", ".gone")
}
