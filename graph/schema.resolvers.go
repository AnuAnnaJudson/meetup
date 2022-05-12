package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/AnuAnnaJudson/meetup/graph/generated"
	"github.com/AnuAnnaJudson/meetup/graph/model"
	"github.com/AnuAnnaJudson/meetup/newmodels"
)

func (r *meetupResolver) User(ctx context.Context, obj *newmodels.Meetup) (*newmodels.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*newmodels.Meetup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*newmodels.Meetup, error) {
	panic(fmt.Errorf("not implemented"))
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type meetupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
