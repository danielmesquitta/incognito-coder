package app

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"time"

	"github.com/kbinani/screenshot"
)

const tmpDir = "tmp"

// CaptureScreenshot captures the screen and saves it
func (a *App) CaptureScreenshot() error {
	if len(a.screenshots) >= 5 {
		return fmt.Errorf("maximum number of screenshots (5) reached")
	}

	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return err
	}

	// Create screenshots directory if it doesn't exist
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		return err
	}

	// Generate unique filename
	filename := filepath.Join(
		tmpDir,
		fmt.Sprintf("screenshot_%d.png", time.Now().Unix()),
	)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		return err
	}

	a.screenshots = append(a.screenshots, filename)
	return nil
}
