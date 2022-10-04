package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/stneto1/gqlgen-example/pkg/graph/generated"
	"github.com/stneto1/gqlgen-example/pkg/graph/model"
	"github.com/stneto1/gqlgen-example/pkg/middlewares"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	cookie, err := gc.Cookie("cookie")
	if err != nil {
		return nil, fmt.Errorf("Unauthorized")
	}

	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("%d", len(r.TodoList)+1),
		User: &model.User{
			ID:   "1",
			Name: cookie,
		},
	}

	r.TodoList = append(r.TodoList, todo)

	return todo, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (bool, error) {
	gc, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return false, err
	}

	if input.Email == "mail@mail.com" && input.Password == "123" {
		gc.SetCookie("cookie", "stneto1", 3600, "/", "localhost", false, true)
		return true, nil
	}

	return false, fmt.Errorf("invalid credentials")
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.TodoList, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
