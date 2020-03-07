// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.
package graph

import (
	"context"
	"fmt"
)

type Resolver struct{}

func (r *mutationResolver) Signup(ctx context.Context, input NewUser) (string, error) {
	return "", fmt.Errorf("not implemented")
}

func (r *queryResolver) User(ctx context.Context, id string) (*User, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
	var results []*User
	return results, fmt.Errorf("not implemented")
}

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
