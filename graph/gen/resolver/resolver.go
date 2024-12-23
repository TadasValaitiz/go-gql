package resolver

import (
	db "go-gql-demo/graph/static/db"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// After implementation is done, this file will not be modified
type Resolver struct {
	db.Database
}
