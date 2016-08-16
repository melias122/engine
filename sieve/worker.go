package sieve

import "sync"

type worker struct {
	id int
	mu sync.Mutex
	t  Task
}

func (w *worker) stop() {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.t != nil {
		w.t.Cancel()
	}
}

func (w *worker) set(t Task) Task {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.t = t
	return w.t
}

func (w *worker) start(tasks <-chan Task) <-chan error {
	ch := make(chan error)
	go func() {
		defer close(ch)
		for t := range tasks {
			ch <- w.set(t).Run()
		}
	}()
	return ch
}

func merge(chans ...<-chan error) <-chan error {
	var (
		new = make(chan error)
		wg  = &sync.WaitGroup{}
	)
	wg.Add(len(chans))
	go func() {
		defer close(new)
		wg.Wait()
	}()
	for _, ch := range chans {
		go func(ch <-chan error) {
			defer wg.Done()
			for err := range ch {
				new <- err
			}
		}(ch)
	}
	return new
}
