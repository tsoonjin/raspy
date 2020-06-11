package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"github.com/tsoonjin/raspy/internal/gql"
	"github.com/tsoonjin/raspy/internal/gql/models"
)

type Resolver struct{}

func (r *mutationResolver) AddPage(ctx context.Context, src string) (*models.Page, error) {
	panic("not implemented")
}

func (r *mutationResolver) AddPages(ctx context.Context, src []string) ([]*models.Page, error) {
	panic("not implemented")
}

func (r *queryResolver) Page(ctx context.Context, url string) (*models.Page, error) {
	panic("not implemented")
}

// Mutation returns gql.MutationResolver implementation.
func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
