package c2

import (
	"io"
	"os"
)

var Fs FileSystem = osFS{}

type FileSystem interface {
	Open(name string) (File, error)
	Stat(name string) (os.FileInfo, error)
	Create(name string) (File, error)
}

type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
	Stat() (os.FileInfo, error)
}

type AltFile struct {
	Delegate interface{}
}

func (self AltFile) Read(b []byte) (int, error) {
	if r, ok := self.Delegate.(io.Reader); ok {
		return r.Read(b)
	}
	panic("not a reader")
}

func (self AltFile) Write(b []byte) (int, error) {
	if r, ok := self.Delegate.(io.Writer); ok {
		return r.Write(b)
	}
	panic("not a writer")
}

// osFS implements fileSystem using the local disk.
type osFS struct{}

func (osFS) Open(name string) (File, error)        { return os.Open(name) }
func (osFS) Stat(name string) (os.FileInfo, error) { return os.Stat(name) }
func (osFS) Create(name string) (File, error)      { return os.Create(name) }

var AltFs = OnFileSystem{
	OnOpen:   Fs.Open,
	OnStat:   Fs.Stat,
	OnCreate: Fs.Create,
}

type OnFileSystem struct {
	OnOpen   func(name string) (File, error)
	OnStat   func(name string) (os.FileInfo, error)
	OnCreate func(name string) (File, error)
}

func (self OnFileSystem) Open(name string) (File, error) {
	return self.OnOpen(name)
}

func (self OnFileSystem) Stat(name string) (os.FileInfo, error) {
	return self.OnStat(name)
}

func (self OnFileSystem) Create(name string) (File, error) {
	return self.OnCreate(name)
}
