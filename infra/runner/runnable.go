package runner

import "context"

// Runnable represent function for Worker.
type Runnable func(
	context.Context,
) error
