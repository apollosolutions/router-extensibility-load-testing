package graph

import (
	"context"
	"subgraph-a/graph/model"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	return &model.User{
		ID:   "1",
		Name: "Tom",
	}, nil
}

// Query returns QueryResolver implementation.
// func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
