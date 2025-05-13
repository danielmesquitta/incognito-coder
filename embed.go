package root

import "embed"

//go:embed all:frontend/dist
var Assets embed.FS

//go:embed .env*
var Env embed.FS
