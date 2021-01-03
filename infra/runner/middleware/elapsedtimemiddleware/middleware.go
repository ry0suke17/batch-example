package elapsedtimemiddleware

import (
	"batch-example/infra/runner"
	exptime "batch-example/infra/time"
	"context"
	"log"
	"time"
)

// Clock represents interface for clock.
type Clock interface {
	// Now return current time.
	Now() time.Time
}

// Arguments represents args for middleware.
type Arguments struct {
	clock Clock
}

// Option represents opt for middleware.
type Option func(*Arguments)

// WithClock specify Clock.
func WithClock(clock Clock) Option {
	return func(a *Arguments) {
		a.clock = clock
	}
}

// New generate middleware that outputs elapsed time
func New(
	opts ...Option,
) runner.Middleware {
	args := &Arguments{
		clock: exptime.NewClock(),
	}
	for _, opt := range opts {
		opt(args)
	}

	return func(ctx context.Context, runnable runner.Runnable, info *runner.Info) (err error) {
		startTime := args.clock.Now()
		defer func() {
			log.Printf("elpased time: %s", time.Since(startTime)*time.Microsecond)
		}()
		return runnable(ctx)
	}
}
