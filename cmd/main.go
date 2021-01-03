package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
	"batch-example/infra/graceful"
	"batch-example/infra/runner"
	"batch-example/infra/runner/middleware/chainmiddleware"
	"batch-example/infra/runner/middleware/elapsedtimemiddleware"
)

var (
	interval = flag.Duration("interval", 5*time.Second, "interval represents execution interval.")
	timeout  = flag.Duration("timeout", 10*time.Second, "timeout represents timeout of the process.")
)

type worker struct{}

func (w worker) Do(context context.Context) error {
	log.Printf("Do!")
	return nil
}

func do() error {
	flag.Parse()

	group := runner.NewGroup()

	middleware := chainmiddleware.Chain(
		elapsedtimemiddleware.New(),
	)
	runnable := middleware.NewRunnable(
		worker{},
		&runner.Info{
			WorkerName: "example",
			Timeout:    *timeout,
		},
	)
	routine, err := runner.NewRoutine(
		runnable,
		runner.WithInterval(*interval),
	)
	if err != nil {
		return fmt.Errorf("failed start: %v", err)
	}

	group.Add(routine)
	ctx := graceful.NewContext()
	return group.Run(ctx)
}

func main() {
	if err := do(); err != nil {
		log.Printf("err=%v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
