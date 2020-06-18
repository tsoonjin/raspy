package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	dbm "github.com/tsoonjin/raspy/internal/orm/models"
    tf "github.com/tsoonjin/raspy/internal/gql/resolvers/transformations"
	"github.com/tsoonjin/raspy/internal/gql/models"
	"github.com/tsoonjin/raspy/internal/gql"
	"github.com/tsoonjin/raspy/internal/orm"
)

type Resolver struct{
    ORM *orm.ORM
}

func (r *mutationResolver) AddPage(ctx context.Context, src string) (*models.Page, error) {
	panic("not implemented")
}

func (r *mutationResolver) AddPages(ctx context.Context, src []string) ([]*models.Page, error) {
	panic("not implemented")
}

func (r *queryResolver) Page(ctx context.Context, url string) (*models.Page, error) {
    return getPage(r, url)
}

func (r *queryResolver) Pages(ctx context.Context) ([]*models.Page, error) {
	panic("not implemented")
}

// Mutation returns gql.MutationResolver implementation.
func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func getPage(r *queryResolver, url string) (*models.Page, error) {
    db := r.ORM.DB.New()
    dbRecord := &dbm.Page{}
    db = db.Where("src=?", url).First(&dbRecord)
    if rec, err := tf.DBPageToGQLPage(dbRecord); err != nil {
        return rec, err
    } else {
        return rec, db.Error
    }
}
