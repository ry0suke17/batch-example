package chainmiddleware

import (
	"context"
	"batch-example/infra/runner"
)

// Chain summarizes multiple interceptors. execute left-to-right order.
func Chain(middles ...runner.Middleware) runner.Middleware {
	n := len(middles)

	if n > 1 {
		lastI := n - 1
		return func(ctx context.Context, runnable runner.Runnable, info *runner.Info) error {
			var (
				chainRunnable runner.Runnable
				curI          int
			)

			chainRunnable = func(ctx2 context.Context) error {
				if curI == lastI {
					return runnable(ctx2)
				}
				curI++
				return middles[curI](ctx2, chainRunnable, info)
			}

			return middles[0](ctx, chainRunnable, info)
		}
	}

	if n == 1 {
		return middles[0]
	}

	// n == 0; nil を返さないための Dummy interceptor。
	return func(ctx context.Context, runnable runner.Runnable, info *runner.Info) error {
		return runnable(ctx)
	}
}
