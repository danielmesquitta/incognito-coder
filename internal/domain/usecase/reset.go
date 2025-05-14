package usecase

import (
	"context"
	"os"
)

type Reset struct{}

func NewReset() *Reset {
	return &Reset{}
}

// Reset clears all screenshots and resets the state
func (r *Reset) Execute(ctx context.Context, screenshots []string) error {
	for _, screenshot := range screenshots {
		if err := os.Remove(screenshot); err != nil {
			return err
		}
	}
	return nil
}
