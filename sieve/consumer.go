package sieve

import (
	"context"
	"log"
	"runtime"
	"sync"
	"time"
)

type consumer struct {
	producer Producer

	mu       sync.Mutex
	progress Progress

	done chan struct{}
	err  error
}

func New(producer Producer) Interface {
	if producer == nil {
		log.Fatal("nil producer")
	}
	return &consumer{
		producer: producer,
		done:     make(chan struct{}),
	}
}

func (c *consumer) Start(ctx context.Context) {
	go c.start(ctx)
}

func (c *consumer) start(ctx context.Context) {
	var (
		workers = make([]*worker, runtime.NumCPU())
		errs    []<-chan error
	)

	// start producer
	tasks := c.producer.Produce(ctx)

	// start workers
	for i := 0; i < len(workers); i++ {
		workers[i] = &worker{
			id: i,
		}
		err := workers[i].start(tasks)
		errs = append(errs, err)
	}
	done := merge(errs...)

	defer func() {
		// cleanup producer
		if err := c.producer.Close(); err != nil && c.err == nil {
			c.err = err
		}

		// wait/progress signal
		close(c.done)
	}()

	// main loop is propagating tasks
	// to workers
	for {
		select {
		// cancel by user or timeout
		case <-ctx.Done():
			for _, w := range workers {
				w.stop()
			}
		// continue even if task fails but log taks err
		case err, ok := <-done:
			if !ok {
				return
			}
			if err != nil {
				// TODO: maybe log to file
				log.Println("task err: ", err.Error())
			}
			// set progress
			p := c.getProgress()
			c.setProgress(Progress{actual: p.actual + 1, total: c.producer.TasksCount()})
		}
	}
}

func (c *consumer) Wait(report func(Progress)) {
	if report == nil {
		report = func(Progress) {}
	}

	// start with 50 milisecond after first progress report
	// continue with 500 milisends
	duration := time.Duration(50 * time.Millisecond)
	for {
		select {
		case _, ok := <-c.done:
			if !ok {
				report(c.getProgress())
			}
			return
		case <-time.After(duration):
			report(c.getProgress())
		}
		duration = 1000 * time.Millisecond
	}
}

func (c *consumer) Error() error {
	return c.err
}

func (c *consumer) setProgress(p Progress) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.progress = p
}

func (m *consumer) getProgress() Progress {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.progress
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
