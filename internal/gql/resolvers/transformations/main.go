package transformations

import (
  "strconv"

  gql "github.com/tsoonjin/raspy/internal/gql/models"
  dbm "github.com/tsoonjin/raspy/internal/orm/models"
)

// DBUserToGQLUser transforms [user] db input to gql type
func DBPageToGQLPage(i *dbm.Page) (o *gql.Page, err error) {
  o = &gql.Page{
    ID:        strconv.FormatInt(int64(i.ID),   10),
    Src:       i.Src,
  }
  return o, err
}
