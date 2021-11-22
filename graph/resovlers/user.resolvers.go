package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graph_jwt/graph/model"
)

var Users []*model.User

func (r *mutationResolver) CreateUser(ctx context.Context, name string, email string) (*model.User, error) {
	User := &model.User{
		ID:    fmt.Sprint(len(Users) + 1),
		Name:  name,
		Email: email,
	}
	Users = append(Users, User)
	return User, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, name string, email string) (*model.User, error) {
	var User *model.User
	for _, user := range Users {
		if user.ID == id {
			User = user
			break
		}
	}
	if User == nil {
		return nil, fmt.Errorf("User not found")
	}
	User.Name = name
	User.Email = email
	return User, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.User, error) {
	var User *model.User
	for i, user := range Users {
		if user.ID == id {
			User = user
			Users = append(Users[:i], Users[i+1:]...)
			break
		}
	}
	if User == nil {
		return nil, fmt.Errorf("User not found")
	}
	return User, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	var User *model.User
	for _, user := range Users {
		if user.ID == id {
			User = user
			break
		}
	}
	if User == nil {
		return nil, fmt.Errorf("User not found")
	}
	return User, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return Users, nil
}
