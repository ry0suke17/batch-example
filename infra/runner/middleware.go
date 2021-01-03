package runner

import "context"

// Middleware represent function for Runnable middleware.
type Middleware func(
	context.Context,
	Runnable,
	*Info,
) error

// NewRunnable create Runnable.
func (m Middleware) NewRunnable(
	w Worker,
	info *Info,
) Runnable {
	return func(ctx context.Context) error {
		return m(ctx, w.Do, info)
	}
}
