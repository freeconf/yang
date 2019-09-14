package parser

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/freeconf/yang/c2"

	"github.com/freeconf/yang/meta"
)

func LoadModuleCustomImport(data string, submoduleLoader meta.Loader) (*meta.Module, error) {
	m, err := parseModule(data, nil, meta.AllFeaturesOn(), submoduleLoader)
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

func LoadModuleWithFeatures(source meta.StreamSource, yangfile string, rev string, features meta.FeatureSet) (*meta.Module, error) {
	m, err := loadModule(source, nil, yangfile, rev, features)
	if err != nil {
		return nil, err
	}
	return m, meta.Validate(m)
}

func LoadModule(source meta.StreamSource, yangfile string) (*meta.Module, error) {
	return LoadModuleWithFeatures(source, yangfile, "", meta.AllFeaturesOn())
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

func parseModule(data string, parent *meta.Module, features meta.FeatureSet, submoduleLoader meta.Loader) (*meta.Module, error) {
	l := lex(string(data), submoduleLoader)
	l.parent = parent
	l.featureSet = features
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

func loadModule(source meta.StreamSource, parent *meta.Module, yangfile string, rev string, features meta.FeatureSet) (*meta.Module, error) {
	// TODO: Use rev
	res, err := source.OpenStream(yangfile, ".yang")
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, c2.NewErrC(yangfile+" resource not found", 404)
	}
	if closer, ok := res.(io.Closer); ok {
		defer closer.Close()
	}
	data, err := ioutil.ReadAll(res)
	if err != nil {
		return nil, err
	}
	return parseModule(string(data), parent, features, submoduleLoader(source))
}

func submoduleLoader(source meta.StreamSource) meta.Loader {
	return func(parent *meta.Module, submodName string, rev string, features meta.FeatureSet) (*meta.Module, error) {
		return loadModule(source, parent, submodName, rev, features)
	}
}
