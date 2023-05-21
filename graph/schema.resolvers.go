package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"

	"github.com/ZihxS/go-graphql/graph/model"
)

// User is the resolver for the user field.
func (r *meetupResolver) User(ctx context.Context, obj *model.Meetup) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Meetups is the resolver for the meetups field.
func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetup, error) {
	panic(fmt.Errorf("not implemented: Meetups - meetups"))
}

// Meetup returns MeetupResolver implementation.
func (r *Resolver) Meetup() MeetupResolver { return &meetupResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type meetupResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
