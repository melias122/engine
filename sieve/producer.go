package sieve

import "context"

type Producer interface {
	// Produce must return Task interface or nil if does not have more tasks
	Start(context.Context) <-chan Task
	// Count return how many tasks will producer make
	Count() int
	// Close serves as cleanup function a returns error if it has any
	Close() error
}

type Task interface {
	Run() error
	Cancel()
}
