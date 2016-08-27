package sieve_test

import (
	"context"
	"errors"
	"testing"
	"time"

	. "github.com/melias122/engine/sieve"
)

type task struct {
}

func (t *task) Run() error {
	time.Sleep(50 * time.Millisecond)
	return nil
}

func (t *task) Cancel() {}

type errtask struct {
	task
}

func (t *errtask) Run() error {
	time.Sleep(50 * time.Millisecond)
	return errors.New("error")
}

type producer struct {
	tasks []Task
}

func NewProducer(t Task, n int) *producer {
	tasks := make([]Task, n)
	for i := range tasks {
		tasks[i] = t
	}
	return &producer{tasks}
}

func (p *producer) Start(ctx context.Context) <-chan Task {
	tasks := make(chan Task)
	go func() {
		defer close(tasks)
		for _, t := range p.tasks {
			select {
			case <-ctx.Done():
				return
			case tasks <- t:
			}
		}
	}()
	return tasks
}

func (p *producer) Count() int   { return len(p.tasks) }
func (p *producer) Close() error { return nil }

func run(t *testing.T, ctx context.Context, p Producer) error {
	progress := make(chan Progress)
	s, err := New(p, progress)
	if err != nil {
		t.Fatal(err)
	}
	s.Start(ctx)
	for {
		select {
		case <-s.Done():
			return s.Error()
		case <-progress:
		}
	}
}

func TestNormal(t *testing.T) {
	err := run(t, context.Background(), NewProducer(&task{}, 100))
	if err != nil {
		t.Fatal(err)
	}
}

func TestCanceled(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	err := run(t, ctx, NewProducer(&task{}, 100))
	if err != nil {
		t.Fatal(err)
	}
}

func TestError(t *testing.T) {
	err := run(t, context.Background(), NewProducer(&errtask{}, 100))
	if err == nil {
		t.Fatal("nil err")
	}
}

func BenchmarkSieve(b *testing.B) {
	p := NewProducer(&errtask{}, 100)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s, err := New(p, nil)
		if err != nil {
			b.Fatal(err)
		}
		s.Start(context.Background())
		<-s.Done()
	}
}
