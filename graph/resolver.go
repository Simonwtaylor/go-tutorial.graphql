package graph

import "github.com/Simonwtaylor/go-tutorial.graphql/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver -
type Resolver struct {
	messageChannels map[string]chan *model.Message
	userChannels    map[string]chan string
}

// NewResolver -
func NewResolver() *Resolver {
	return &Resolver{
		messageChannels: map[string]chan *model.Message{},
		userChannels:    map[string]chan string{},
	}
}
