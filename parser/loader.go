package parser

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/freeconf/yang/c2"
	"github.com/freeconf/yang/source"

	"github.com/freeconf/yang/meta"
)

func LoadModuleCustomImport(data string, submoduleLoader meta.Loader) (*meta.Module, error) {
	m, err := parseModule(data, nil, meta.AllFeaturesOn(), submoduleLoader)
	if err != nil {
		return nil, err
	}
	return m, meta.Validate(m)
}

func YangPath() source.Opener {
	path := os.Getenv("YANGPATH")
	if len(path) == 0 {
		panic("Environment variable YANGPATH not set")
	}
	return source.Path(path)
}

func RequireModule(source source.Opener, yangfile string) *meta.Module {
	m, err := LoadModule(source, yangfile)
	if err != nil {
		panic(fmt.Sprintf("Could not load module %s : %s", yangfile, err))
	}
	return m
}

func LoadModuleWithFeatures(source source.Opener, yangfile string, rev string, features meta.FeatureSet) (*meta.Module, error) {
	m, err := loadModule(source, nil, yangfile, rev, features)
	if err != nil {
		return nil, err
	}
	return m, meta.Validate(m)
}

func LoadModule(source source.Opener, yangfile string) (*meta.Module, error) {
	return LoadModuleWithFeatures(source, yangfile, "", meta.AllFeaturesOn())
}

func RequireModuleFromString(source source.Opener, yangStr string) *meta.Module {
	m, err := LoadModuleFromString(source, yangStr)
	if err != nil {
		panic(err)
	}
	return m
}

func LoadModuleFromString(source source.Opener, yangStr string) (*meta.Module, error) {
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

func loadModule(source source.Opener, parent *meta.Module, yangfile string, rev string, features meta.FeatureSet) (*meta.Module, error) {
	// TODO: Use rev
	res, err := source(yangfile, ".yang")
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

func submoduleLoader(source source.Opener) meta.Loader {
	return func(parent *meta.Module, submodName string, rev string, features meta.FeatureSet) (*meta.Module, error) {
		return loadModule(source, parent, submodName, rev, features)
	}
}
