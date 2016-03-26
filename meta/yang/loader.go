package yang

import (
	"fmt"
	"io/ioutil"
	"os"
	"github.com/blitter/meta"
)

type ImportModule func(into *meta.Module, name string) (e error)

func LoadModuleCustomImport(data string, importer ImportModule) (*meta.Module, error) {
	l := lex(string(data), importer)
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

func moduleCopy(dest *meta.Module, src *meta.Module) {
	iters := []meta.MetaIterator{
		meta.NewMetaListIterator(src.GetGroupings(), false),
		meta.NewMetaListIterator(src.GetTypedefs(), false),
		meta.NewMetaListIterator(src.DataDefs(), false),
		meta.NewMetaListIterator(src.GetNotifications(), false),
	}
	for _, iter := range iters {
		for iter.HasNextMeta() {
			dest.AddMeta(iter.NextMeta())
		}
	}
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

func LoadModule(source meta.StreamSource, yangfile string) (*meta.Module, error) {
	if res, err := source.OpenStream(yangfile); err != nil {
		return nil, err
	} else {
		defer meta.CloseResource(res)
		if data, err := ioutil.ReadAll(res); err != nil {
			return nil, err
		} else {
			return LoadModuleCustomImport(string(data), ModuleImporter(source))
		}
	}
}

func ModuleImporter(source meta.StreamSource) ImportModule {
	return func(main *meta.Module, submodName string) (suberr error) {
		var sub *meta.Module
		// TODO: Performance - cache modules
		//subFname := fmt.Sprint(submodName, ".yang")
		if sub, suberr = LoadModule(source, submodName); suberr != nil {
			return suberr
		}
		moduleCopy(main, sub)
		return nil
	}
}

func LoadModuleFromString(source meta.StreamSource, yangStr string) (*meta.Module, error) {
	return LoadModuleCustomImport(yangStr, ModuleImporter(source))
}
