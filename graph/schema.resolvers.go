package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"

	"bsm.com/gqlgengofirestore/graph/model"
	"bsm.com/gqlgengofirestore/src/service"
)

// CreateCustomer is the resolver for the createCustomer field.
func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.CustomerReq) (*model.Customer, error) {
	return service.CreateCustomer(input), nil
}

// Customers is the resolver for the customers field.
func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	return service.FechCustomers(), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
