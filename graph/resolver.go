//go:generate go run github.com/99designs/gqlgen generate
package graph

import (
	"github.com/aberyotaro/gql_sample_api/graph/model"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB    *gorm.DB
	todos []*model.Todo
}
