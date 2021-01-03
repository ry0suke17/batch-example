package runner

import (
	"context"
	"fmt"
	"sync"
)

// Group represents group of runners.
type Group struct {
	elms []Runner
}

// NewGroup generates Group.
func NewGroup(
	elms ...Runner,
) *Group {
	return &Group{
		elms: elms,
	}
}

// Add add runner to group.
func (g *Group) Add(
	elm Runner,
) {
	g.elms = append(g.elms, elm)
}

// Run executes proccess.
func (g *Group) Run(
	ctx context.Context,
) error {
	if len(g.elms) == 0 {
		return fmt.Errorf("len(g.elms)=0")
	}
	var wg sync.WaitGroup
	for _, elm := range g.elms {
		elm := elm

		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			elm.Run(ctx)
		}()
	}

	wg.Wait()
	return nil
}
