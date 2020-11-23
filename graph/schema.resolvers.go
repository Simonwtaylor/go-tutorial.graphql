package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"sync"

	"github.com/Simonwtaylor/go-tutorial.graphql/graph/generated"
	"github.com/Simonwtaylor/go-tutorial.graphql/graph/model"
)

var mutex = &sync.Mutex{}

func (r *mutationResolver) PostMessage(ctx context.Context, user string, text string) (*model.Message, error) {
	mutex.Lock()
	for _, ch := range r.userChannels {
		ch <- user
	}
	mutex.Unlock()

	m := model.Message{
		ID:   "1",
		Text: text,
		User: user,
	}

	mutex.Lock()
	for _, ch := range r.messageChannels {
		ch <- &m
	}

	return &m, nil
}

func (r *queryResolver) Messages(ctx context.Context) ([]*model.Message, error) {
	return []*model.Message{
		{
			ID:   "1",
			Text: "aidsjaoisd",
			User: "asuhdaushd",
		},
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MessagePost(ctx context.Context, user string) (<-chan *model.Message, error) {
	// Create new channel for request
	messages := make(chan *model.Message, 1)
	mutex.Lock()
	r.messageChannels[user] = messages
	mutex.Unlock()

	// Delete channel when done
	go func() {
		<-ctx.Done()
		mutex.Lock()
		delete(r.messageChannels, user)
		mutex.Unlock()
	}()

	return messages, nil
}

func (r *subscriptionResolver) UserJoined(ctx context.Context, user string) (<-chan string, error) {
	users := make(chan string, 1)
	mutex.Lock()
	r.userChannels[user] = users
	mutex.Unlock()

	go func() {
		<-ctx.Done()
		mutex.Lock()
		delete(r.userChannels, user)
		mutex.Unlock()
	}()

	// mutex.Lock()
	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
