package graph

import "github.com/msartini/graphql-go/graph/model"

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	videos []*model.Video
}
