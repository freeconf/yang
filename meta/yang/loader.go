package yang

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/c2stack/c2g/c2"

	"github.com/c2stack/c2g/meta"
)

func LoadModuleCustomImport(data string, submoduleLoader meta.Loader) (*meta.Module, error) {
	m, err := parseModule(data, nil, submoduleLoader)
	if err != nil {
		return nil, err
	}
	return m, meta.Validate(m)
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
	m, err := loadModule(source, nil, yangfile, "")
	if err != nil {
		return nil, err
	}
	return m, meta.Validate(m)
}

func RequireModuleFromString(source meta.StreamSource, yangStr string) *meta.Module {
	m, err := LoadModuleFromString(source, yangStr)
	if err != nil {
		panic(err)
	}
	return m
}

func LoadModuleFromString(source meta.StreamSource, yangStr string) (*meta.Module, error) {
	return LoadModuleCustomImport(yangStr, submoduleLoader(source))
}

func parseModule(data string, parent *meta.Module, submoduleLoader meta.Loader) (*meta.Module, error) {
	l := lex(string(data), submoduleLoader)
	l.parent = parent
	err_code := yyParse(l)
	if err_code != 0 || l.lastError != nil {
		if l.lastError == nil {
			l.lastError = c2.NewErr(fmt.Sprint("Error parsing, code ", string(err_code)))
		}
		return nil, l.lastError
	}

	m := l.stack.Peek().(*meta.Module)
	return m, nil
}

func loadModule(source meta.StreamSource, parent *meta.Module, yangfile string, rev string) (*meta.Module, error) {
	res, err := source.OpenStream(yangfile, ".yang")
	if err != nil {
		return nil, err
	}
	if closer, ok := res.(io.Closer); ok {
		defer closer.Close()
	}
	data, err := ioutil.ReadAll(res)
	if err != nil {
		return nil, err
	}
	return parseModule(string(data), parent, submoduleLoader(source))
}

func submoduleLoader(source meta.StreamSource) meta.Loader {
	return func(parent *meta.Module, submodName string, rev string) (*meta.Module, error) {
		return loadModule(source, parent, submodName, rev)
	}
}
