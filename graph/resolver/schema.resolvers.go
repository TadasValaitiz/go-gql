package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.61

import (
	"context"
	"fmt"
	"go-gql-demo/graph/generated"
	"go-gql-demo/graph/model"
)

// UpdateOrganization is the resolver for the updateOrganization field.
func (r *mutationResolver) UpdateOrganization(ctx context.Context, id string, value model.OrganizationUpdate) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: UpdateOrganization - updateOrganization"))
}

// MyProfile is the resolver for the myProfile field.
func (r *queryResolver) MyProfile(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented: MyProfile - myProfile"))
}

// Organizations is the resolver for the organizations field.
func (r *queryResolver) Organizations(ctx context.Context) ([]*model.Organization, error) {
	panic(fmt.Errorf("not implemented: Organizations - organizations"))
}

// Organization is the resolver for the organization field.
func (r *queryResolver) Organization(ctx context.Context, id string) (*model.Organization, error) {
	// Create a string pointer for the alias
	alias := "test-alias"
	// Return a dummy organization for testing
	return &model.Organization{
		Name:  "Test Organization",
		Alias: &alias,
	}, nil
}

// Studies is the resolver for the studies field.
func (r *queryResolver) Studies(ctx context.Context) ([]*model.Study, error) {
	panic(fmt.Errorf("not implemented: Studies - studies"))
}

// Study is the resolver for the study field.
func (r *queryResolver) Study(ctx context.Context, id string) (*model.Study, error) {
	panic(fmt.Errorf("not implemented: Study - study"))
}

// Sites is the resolver for the sites field.
func (r *queryResolver) Sites(ctx context.Context) ([]*model.Site, error) {
	panic(fmt.Errorf("not implemented: Sites - sites"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }