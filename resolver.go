package main

import (
	"context"

	"Ptncafe.Golang.GraphQLTest/generated"
	"Ptncafe.Golang.GraphQLTest/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) GetComment(ctx context.Context, input models.InputGetComment) (*models.CommentDto, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) QueryComment(ctx context.Context) ([]*models.CommentDto, error) {
	panic("not implemented")
}
