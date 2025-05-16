## 1. Overview

- **Project:** Incognito Coder
- **Purpose:** Assist users during coding interviews by capturing screenshots, sending them to OpenAI, and displaying stealthy hints and solutions without being detected in screen sharing or browser-based coding assessments.
- **Supported Platforms:** Windows, Linux, macOS

---

## 2. Technology Stack

- **Backend:** Go (Golang v1.24.1) with Wails v2
- **Frontend:** React v18.2.0

When generating code:

- Always target Go 1.24 syntax and module conventions.
- Use Wails v2 APIs for backend–frontend bindings.
- Write React components in TypeScript (per existing codebase) targeting React v18.

---

## 3. Project Structure

```text
.
├── bin/
│   └── install.sh        # Install Go tools
├── cmd/
│   └── app/
│       └── main.go       # Main entry point for the app
├── embed.go              # Embedding static assets
├── go.mod
├── go.sum
├── internal/
│   ├── app/              # Wails app initialization and bindings
│   ├── config/           # Environment and config loading
│   ├── domain/
│   │   ├── entity/       # Core entities and data models
│   │   └── usecase/      # Business logic and use-cases
│   ├── pkg/              # Reusable utilities and helpers
│   └── provider/         # External integrations (OpenAI, file I/O, screenshot capture)
└── tmp/
    └── .gitkeep          # Placeholder for screenshots directory
```

When Copilot generates new files or functions:

- Place CLI and execution code under `cmd/`.
- Implement business logic in `internal/domain/usecase/` and models in `internal/domain/entity/`.
- Add external API clients under `internal/provider/`.
- Keep utilities in `internal/pkg/`.

---
