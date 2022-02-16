package mocks

import (
	"github.com/marklude/go-sovtech/graph"
	"github.com/marklude/go-sovtech/graph/generated"
)

func NewResolver() generated.Config {
	r := graph.Resolver{}

	return generated.Config{
		Resolvers: &r,
	}
}
