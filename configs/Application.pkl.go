// Code generated from Pkl module `Application`. DO NOT EDIT.
package configs

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Application struct {
	// The URL to listen on
	Host string `pkl:"host"`

	// The port to listen on
	Port uint16 `pkl:"port"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Application
func LoadFromPath(ctx context.Context, path string) (ret *Application, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Application
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Application, error) {
	var ret Application
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
