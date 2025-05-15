# Incognito Coder

## Overview

Incognito Coder offers a range of features designed to assist users during coding interviews. A key characteristic of the tool is its stealth; it remains undetectable during screen sharing sessions on platforms like Zoom and is reportedly not flagged by browser-based coding assessment tools such as HackerRank and CodeSignal.

## Technology:

1. Wails v2
2. Golang v1.24
3. React v18

## Structure

### Backend

```
.
├── bin/
│   └── install.sh        # Install Go tools
├── cmd/
│   └── app/
│       └── main.go       # Main entry point for the app
├── embed.go              # Embedding files for the app
├── go.mod
├── go.sum
├── internal/
│   ├── app/              # Wails app
│   ├── config/           # General configuration (env loading)
│   ├── domain/
│   │   ├── entity/       # Entities
│   │   └── usecase/      # Use-cases
│   ├── pkg/              # Utilities
│   └── provider/         # External integrations (OpenAI, etc.)
└── tmp/
    └── .gitkeep          # Placeholder for empty directory
```

## Product requirements

### Shortcuts

1. Print screen: Ctrl + Alt + P
2. Toggle window visibility: Ctrl + Alt + V
3. Generate solution (Send the screenshots to OpenAI): Ctrl + Alt + Enter
4. Reset (delete screenshots): Ctrl + Alt + R
5. Move to the left: Ctrl + Alt + Arrow Left
6. Move to the right: Ctrl + Alt + Arrow Right
7. Move to the top: Ctrl + Alt + Arrow Up
8. Move to the bottom: Ctrl + Alt + Arrow Down

### Window

1. Window must be translucent
2. Size: 1024x768
3. The window must not be visible by screen recordings by using OS-specific flags to disable screen capture (e.g., `SetWindowSubclass` on Windows, `NSWindowCollectionBehaviorCanJoinAllSpaces` on macOS).

### Solution

1. Use OpenAI API o4-mini to generate the solution
2. The solution must have a "My thoughts", "Code Solution" and "Complexity" section (providing the time and space complexity of the solution and explaining why)

### Overall

1. This must work on Windows, Linux and MacOS
