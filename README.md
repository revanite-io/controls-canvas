# Controls Canvas

A terminal-based UI tool for interactively selecting and managing cloud security controls.

## Features

- Interactive TUI for browsing and selecting cloud security controls
- Support for Common Cloud Controls (CCC) framework
- Real-time YAML preview of selected controls
- Dynamic window resizing with responsive layout
- Efficient filtering and navigation

## Installation

### macOS with Homebrew

```bash
brew tap revanite-io/tap
brew install controls-canvas
```

### Using Go

```bash
go install github.com/revanite-io/controls-canvas@latest
```

### Direct Download

1. Visit the [releases page](https://github.com/revanite-io/controls-canvas/releases)
2. Download the archive for your platform:
   - macOS ARM64: `controls-canvas_Darwin_arm64.tar.gz`
   - macOS Intel: `controls-canvas_Darwin_x86_64.tar.gz`
   - Linux: `controls-canvas_Linux_x86_64.tar.gz`
   - Windows: `controls-canvas_Windows_x86_64.zip`
3. Extract the archive
4. Move the `controls-canvas` binary to a directory in your PATH (optional)

## Usage

```bash
controls-canvas
```