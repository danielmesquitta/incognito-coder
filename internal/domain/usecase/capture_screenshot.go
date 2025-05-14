package usecase

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"time"

	"github.com/danielmesquitta/incognito-coder/internal/config/env"
	"github.com/kbinani/screenshot"
)

type CaptureScreenshot struct {
	e *env.Env
}

func NewCaptureScreenshot(e *env.Env) *CaptureScreenshot {
	return &CaptureScreenshot{
		e: e,
	}
}

// CaptureScreenshot captures the screen and saves it
func (c *CaptureScreenshot) Execute(screenshots []string) ([]string, error) {
	if len(screenshots) >= 5 {
		return nil, fmt.Errorf("maximum number of screenshots (5) reached")
	}

	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return nil, err
	}

	// Create screenshots directory if it doesn't exist
	if err := os.MkdirAll(c.e.TmpDir, 0755); err != nil {
		return nil, err
	}

	// Generate unique filename
	filename := filepath.Join(
		c.e.TmpDir,
		fmt.Sprintf("screenshot_%d.png", time.Now().Unix()),
	)
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		return nil, err
	}

	screenshots = append(screenshots, filename)
	return screenshots, nil
}
