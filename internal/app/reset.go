package app

import "os"

// Reset clears all screenshots and resets the state
func (a *App) Reset() error {
	for _, screenshot := range a.screenshots {
		if err := os.Remove(screenshot); err != nil {
			return err
		}
	}
	a.screenshots = make([]string, 0)
	return nil
}
