package meta

type Loader func(parent *Module, name string, rev string, features FeatureSet, loader Loader) (*Module, error)
