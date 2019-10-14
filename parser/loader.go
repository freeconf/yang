package parser

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/source"

	"github.com/freeconf/yang/meta"
)

// LoadModule loads YANG file with all features on and standard submodule loader
func LoadModule(source source.Opener, yangfile string) (*meta.Module, error) {
	return LoadModuleWithOptions(source, yangfile, Options{})
}

// RequireModule loads YANG file with all features on and standard submodule loader and
// panic if YANG file is not available.
func RequireModule(source source.Opener, yangfile string) *meta.Module {
	m, err := LoadModule(source, yangfile)
	if err != nil {
		panic(fmt.Sprintf("Could not load module %s : %s", yangfile, err))
	}
	return m
}

// Options is for non-standard options when loading YANG files
type Options struct {

	// Features when you know want to control what features are on of off
	Features meta.FeatureSet

	// Revision to lock into a specific revision
	Revision string
}

// LoadModuleFromString parses YANG from a string, not a file.
func LoadModuleFromString(source source.Opener, yang string) (*meta.Module, error) {
	return LoadModuleFromStringWithOptions(source, yang, Options{})
}

// LoadModuleFromStringWithOptions parses YANG from a string, not a file but allows custom options
func LoadModuleFromStringWithOptions(source source.Opener, yang string, options Options) (*meta.Module, error) {
	options = setDefaults(options)
	m, err := parseModule(yang, nil, options.Features, submoduleLoader(source))
	if err != nil {
		return nil, err
	}
	return m, meta.Validate(m)
}

// LoadModuleWithOptions is like LoadModule but with more control on process
func LoadModuleWithOptions(source source.Opener, yangfile string, options Options) (*meta.Module, error) {
	options = setDefaults(options)
	m, err := loadModule(source, nil, yangfile, options.Revision, options.Features, submoduleLoader(source))
	if err != nil {
		return nil, err
	}
	return m, meta.Validate(m)
}

func parseModule(data string, parent *meta.Module, featureSet meta.FeatureSet, loader meta.Loader) (*meta.Module, error) {
	l := lex(string(data), loader)
	l.parent = parent
	l.featureSet = featureSet
	errcode := yyParse(l)
	if errcode != 0 || l.lastError != nil {
		if l.lastError == nil {
			l.lastError = fmt.Errorf("Error parsing, code %d", errcode)
		}
		return nil, l.lastError
	}

	m := l.stack.Peek().(*meta.Module)
	return m, nil
}

func loadModule(source source.Opener, parent *meta.Module, yangfile string, rev string, featureSet meta.FeatureSet, loader meta.Loader) (*meta.Module, error) {
	// TODO: Use rev
	res, err := source(yangfile, ".yang")
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fc.NotFoundError(yangfile + " resource not found")
	}
	if closer, ok := res.(io.Closer); ok {
		defer closer.Close()
	}
	data, err := ioutil.ReadAll(res)
	if err != nil {
		return nil, err
	}
	return parseModule(string(data), parent, featureSet, loader)
}

func submoduleLoader(source source.Opener) meta.Loader {
	return func(parent *meta.Module, submodName string, rev string, featureSet meta.FeatureSet, submoduleLoader meta.Loader) (*meta.Module, error) {
		return loadModule(source, parent, submodName, rev, featureSet, submoduleLoader)
	}
}

func setDefaults(o Options) Options {
	if o.Features == nil {
		o.Features = meta.AllFeaturesOn()
	}
	return o
}
