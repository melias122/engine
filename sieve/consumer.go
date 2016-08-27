package sieve

import (
	"context"
	"errors"
	"runtime"
	"sync"
)

type Consumer interface {
	Start(context.Context)
	Done() <-chan struct{}
	Error() error
}

type consumer struct {
	once sync.Once

	producer Producer
	workers  []*worker

	progress Progress
	report   chan<- Progress

	done chan struct{}
	err  error
}

func New(producer Producer, progress chan<- Progress) (Consumer, error) {
	return newConsumer(producer, progress)
}

func newConsumer(producer Producer, report chan<- Progress) (Consumer, error) {
	if producer == nil {
		return nil, errors.New("producer can't be nil")
	}
	return &consumer{
		producer: producer,
		workers:  make([]*worker, runtime.NumCPU()),
		progress: Progress{total: producer.Count()},
		report:   report,
		done:     make(chan struct{}),
	}, nil
}

func (c *consumer) Start(ctx context.Context) {
	c.once.Do(func() {
		go c.start(ctx)
	})
}

func (c *consumer) start(ctx context.Context) {

	// start producer
	tasks := c.producer.Start(ctx)

	// start workers
	var errs []<-chan error
	for i := 0; i < len(c.workers); i++ {
		c.workers[i] = &worker{id: i}
		errs = append(errs, c.workers[i].start(tasks))
	}
	workers := merge(errs...)

	defer func() {
		// stop workers
		for _, w := range c.workers {
			w.stop()
		}
		// drain everything that left
		for range workers {
		}

		// cleanup producer
		err := c.producer.Close()
		c.setError(err)

		// report done
		close(c.done)
	}()

	// main loop is propagating tasks
	// to workers
	for {
		select {
		// cancel by user or timeout
		case <-ctx.Done():
			return

		// continue even if task fails but log taks err
		case err, ok := <-workers:

			switch {
			// cancel on err
			case err != nil:
				c.setError(err)
				return
			// cancel on chan close
			case !ok:
				return
			}

			// increase progress
			c.progress.actual++

			// report actual progress
			// if receiver is too slow
			// drop report and continue
			select {
			case c.report <- c.progress:
			default:
			}
		}
	}
}

func (c *consumer) Done() <-chan struct{} {
	return c.done
}

func (c *consumer) Error() error {
	return c.err
}

func (c *consumer) setError(err error) {
	if c.err == nil && err != nil {
		c.err = err
	}
}

type Progress struct {
	actual int
	total  int
}

func (p *Progress) Relative() float64 {
	total := p.total
	if total == 0 {
		total = 1
	}
	return float64(p.actual) / float64(total)
}

func (p *Progress) Absolute() (actual, total int) {
	return p.actual, p.total
}
