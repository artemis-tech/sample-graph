// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.
package graph

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Resolver struct{}

// Users stands for the in-memory list of users
var Users []*User

func (r *mutationResolver) Signup(ctx context.Context, input NewUser) (string, error) {
	user := &User{
		ID:        uuid.New().String(),
		Name:      input.Name,
		Surename:  input.Surename,
		CreatedAt: time.Now().String(),
	}

	// removes oldest user if users length is over 30
	if len(Users) > 30 {
		Users = Users[1:29]
	}

	// appends new user to global (in-memory) users list
	Users = append(Users, user)

	return user.ID, nil
}

func (r *queryResolver) User(ctx context.Context, ID string) (*User, error) {
	var user *User
	for _, current := range Users {
		if current.ID == ID {
			user = current
		}
	}

	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
	return Users, nil
}

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
