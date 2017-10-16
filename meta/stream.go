package meta

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	if err != nil || str == "" {
		return nil, err
	}
	return strings.NewReader(str), nil
}

type Stream interface {
	StreamSink
	StreamSource
}

type StreamSink interface {
	WriteStream(resourceId string, ext string, copy DataStream) error
}

type CacheSource struct {
	Local    Stream
	Upstream StreamSource
}

func (self CacheSource) OpenStream(resourceId string, ext string) (DataStream, error) {
	s, err := self.Local.OpenStream(resourceId, ext)
	if s == nil {
		u, err := self.Upstream.OpenStream(resourceId, ext)
		if err != nil || u == nil {
			return nil, err
		}
		if err := self.Local.WriteStream(resourceId, ext, u); err != nil {
			return nil, err
		}
		s, err = self.Local.OpenStream(resourceId, ext)
	}
	return s, err
}

func (src *FileStreamSource) OpenStream(resourceId string, ext string) (DataStream, error) {
	path := fmt.Sprint(src.Root, "/", resourceId, ext)
	stream, err := os.Open(path)
	if os.IsNotExist(err) {
		return nil, nil
	}
	return stream, err
}

func (src *FileStreamSource) WriteStream(resourceId string, ext string, copy DataStream) error {
	path := fmt.Sprint(src.Root, "/", resourceId, ext)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	all, err := ioutil.ReadAll(copy)
	if err != nil {
		return err
	}
	data := all
	for {
		n, err := f.Write(data)
		if err != nil {
			return err
		}
		if n == len(data) {
			break
		}
		data = data[n:]
	}
	return nil
}
