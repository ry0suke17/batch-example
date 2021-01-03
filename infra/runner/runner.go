package runner

import "context"

// Runner represent interface for runner.
type Runner interface {
	Run(ctx context.Context)
}
