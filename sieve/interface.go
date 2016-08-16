package sieve

import "context"

type Interface interface {
	Start(context.Context)
	Wait(func(Progress))
	Error() error
}

type Producer interface {
	// Produce must return Task interface or nil if does not have more tasks
	Produce(context.Context) <-chan Task
	// TasksCount return how many tasks will producer make
	TasksCount() int
	// Close serves as cleanup function a returns error if it has any
	Close() error
}

type Task interface {
	Run() error
	Cancel()
}
