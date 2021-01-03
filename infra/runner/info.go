package runner

import "time"

// Info represent information for Runner.
type Info struct {
	WorkerName string
	Timeout    time.Duration
}
