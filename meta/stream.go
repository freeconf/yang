package meta

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/c2stack/c2g/c2"
)

type DataStream interface {
	Read(p []byte) (n int, err error)
}

type StreamSource interface {
	OpenStream(streamId string, ext string) (DataStream, error)
}

func PathStreamSource(path string) StreamSource {
	dirs := strings.Split(path, ":")
	sources := make([]StreamSource, len(dirs))
	for i, dir := range dirs {
		sources[i] = &FileStreamSource{Root: dir}
		i++
	}
	return &MulticastStreamSource{sources}
}

func MultipleSources(s ...StreamSource) *MulticastStreamSource {
	return &MulticastStreamSource{Sources: s}
}

type MulticastStreamSource struct {
	Sources []StreamSource
}

func (s *MulticastStreamSource) OpenStream(resourceId string, ext string) (DataStream, error) {
	for _, source := range s.Sources {
		found, err := source.OpenStream(resourceId, ext)
		if found != nil {
			return found, err
		}
	}
	return nil, os.ErrNotExist
}

type FileStreamSource struct {
	Root string
}

func NewCwdSource() StreamSource {
	cwd, _ := os.Getwd()
	return &FileStreamSource{Root: cwd}
}

type StringSource struct {
	Streamer StringStreamer
}

type StringStreamer func(resource string) (string, error)

type stringStream strings.Reader

func (s *StringSource) OpenStream(resourceId string, ext string) (DataStream, error) {
	str, err := s.Streamer(resourceId)
	if err != nil {
		return nil, err
	}
	return strings.NewReader(str), nil
}

func (src *FileStreamSource) OpenStream(resourceId string, ext string) (DataStream, error) {
	path := fmt.Sprint(src.Root, "/", resourceId, ext)
	stream, err := os.Open(path)
	if os.IsNotExist(err) {
		return nil, err
	}
	return stream, err
}

type FsError struct {
	Msg string
}

func (e *FsError) Error() string {
	return e.Msg
}

type CachingStreamSource struct {
	Stream StreamSource
	Dir    string
	fs     c2.FileSystem
}

func (self CachingStreamSource) OpenStream(resource string, ext string) (DataStream, error) {
	fname := self.Dir + "/" + resource + ext
	if _, err := self.fs.Stat(fname); err == nil {
		return self.fs.Open(fname)
	}
	ds, err := self.Stream.OpenStream(resource, ext)
	if err != nil {
		return nil, err
	}
	if ds != nil {
		file, ferr := self.fs.Create(fname)
		if ferr != nil {
			return nil, ferr
		}
		ds = &cachingStream{
			delegate: ds,
			file:     file,
		}
	}
	return ds, nil
}

type cachingStream struct {
	delegate DataStream
	file     io.WriteCloser
}

func (self *cachingStream) Close() error {
	if c, ok := self.delegate.(io.Closer); ok {
		c.Close()
	}
	return self.file.Close()
}

func (self *cachingStream) Read(b []byte) (int, error) {
	n, err := self.delegate.Read(b)
	if err != nil && n > 0 {
		if _, werr := self.file.Write(b[:n]); werr != nil {
			return 0, werr
		}
	}
	return n, err
}
