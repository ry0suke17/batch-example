package runner

import "context"

// Worker represent interface for worker
type Worker interface {
	Do(context.Context) error
}
