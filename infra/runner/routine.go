package runner

import (
	"context"
	"time"
)

// Routine represent runner that run it periodically.
type Routine struct {
	runnable     Runnable
	interval     time.Duration
	errorHandler func(error)
}

// Option represent option for Routine.
type Option func(*Routine)

// WithInterval specify interval for run.
func WithInterval(interval time.Duration) Option {
	return func(r *Routine) {
		r.interval = interval
	}
}

// WithErrorHandler specify error handler.
func WithErrorHandler(handler func(error)) Option {
	return func(r *Routine) {
		r.errorHandler = handler
	}
}

// NewRoutine create Routine.
func NewRoutine(
	runnable Runnable,
	opts ...Option,
) (*Routine, error) {
	r := &Routine{
		interval:     10 * time.Second,
		runnable:     runnable,
		errorHandler: func(error) { /* noop */ },
	}
	for _, opt := range opts {
		opt(r)
	}
	return r, nil
}

// Run execute the process periodically.
func (r *Routine) Run(ctx context.Context) {
	t := time.NewTicker(r.interval)
	defer t.Stop()

	for {
		err := r.runnable(ctx)
		if err != nil {
			r.errorHandler(err)
		}

		select {
		case <-ctx.Done():
			return
		case <-t.C:
		}
	}
}
