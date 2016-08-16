package sieve

import (
	"context"
	"testing"
	"time"
)

type task struct {
}

func (t *task) Run() error {
	// time.Sleep(10 * time.Millisecond)
	return nil
}

func (t *task) Cancel() {}

type producer struct{}

func (p *producer) Produce(ctx context.Context) <-chan Task {
	tasks := make(chan Task)
	go func() {
		defer close(tasks)
		for i := 0; i < 100; i++ {
			select {
			case <-ctx.Done():
				return
			case tasks <- &task{}:
			}
		}
	}()
	return tasks
}

func (p *producer) TasksCount() int { return 100 }
func (p *producer) Close() error    { return nil }

func TestSieve(t *testing.T) {
	ctx := context.Background()
	s := New(&producer{})
	s.Start(ctx)
	s.Wait(nil)
	if s.Error() != nil {
		t.Fail()
	}
}

func TestCancelation(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Millisecond)
	s := New(&producer{})
	s.Start(ctx)
	s.Wait(nil)
	if s.Error() != nil {
		t.Fail()
	}
}

func BenchmarkSieve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := New(&producer{})
		s.Start(context.Background())
		s.Wait(nil)
	}
}
