package yang

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/c2stack/c2g/meta"
)

type ModuleLoader func(name string) (*meta.Module, error)

func LoadModuleCustomImport(data string, loader ModuleLoader) (*meta.Module, error) {
	l := lex(string(data), loader)
	err_code := yyParse(l)
	if err_code != 0 || l.lastError != nil {
		if l.lastError == nil {
			// Developer - Find out why there's no error
			msg := fmt.Sprint("Error parsing, code ", string(err_code))
			l.lastError = &yangError{msg}

		}
		return nil, l.lastError
	}

	d := l.stack.Peek()
	return d.(*meta.Module), nil
}

var gYangPath meta.StreamSource

func YangPath() meta.StreamSource {
	if gYangPath == nil {
		path := os.Getenv("YANGPATH")
		if len(path) == 0 {
			panic("Environment variable YANGPATH not set")
		}
		gYangPath = meta.PathStreamSource(path)
	}
	return gYangPath
}

func RequireModule(source meta.StreamSource, yangfile string) *meta.Module {
	m, err := LoadModule(source, yangfile)
	if err != nil {
		panic(fmt.Sprintf("Could not load module %s : %s", yangfile, err))
	}
	return m
}

func LoadModule(source meta.StreamSource, yangfile string) (*meta.Module, error) {
	if res, err := source.OpenStream(yangfile, ".yang"); err != nil {
		return nil, err
	} else {
		if closer, ok := res.(io.Closer); ok {
			defer closer.Close()
		}
		if data, err := ioutil.ReadAll(res); err != nil {
			return nil, err
		} else {
			return LoadModuleCustomImport(string(data), moduleLoader(source))
		}
	}
}

func moduleLoader(source meta.StreamSource) ModuleLoader {
	return func(submodName string) (*meta.Module, error) {
		return LoadModule(source, submodName)
	}
}

func RequireModuleFromString(source meta.StreamSource, yangStr string) *meta.Module {
	m, err := LoadModuleFromString(source, yangStr)
	if err != nil {
		panic(err)
	}
	return m
}

func LoadModuleFromString(source meta.StreamSource, yangStr string) (*meta.Module, error) {
	return LoadModuleCustomImport(yangStr, moduleLoader(source))
}
