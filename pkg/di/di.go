package di

import "github.com/sarulabs/di/v2"

func Build(defs ...di.Def) (di.Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	if err := builder.Add(defs...); err != nil {
		panic(err)
	}

	return builder.Build(), nil
}
