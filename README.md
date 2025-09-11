# Controls Canvas

A terminal-based UI tool for interactively selecting and managing import references for a Gemara Layer 2 catalog.

## Features

- Interactive TUI for browsing a Gemara-compatible controls catalog to create import references
- Support for Common Cloud Controls (CCC) framework
- Real-time YAML preview of selected controls
- Dynamic window resizing with responsive layout

## Usage

Select the catalog you'd like to import control references from (Currently the only supported catalog to import from is FINOS Common Cloud Controls).

<img width="1920" height="225" alt="Screenshot 2025-09-11 at 04 48 56" src="https://github.com/user-attachments/assets/3c5ed7bb-2368-4f3e-819c-b60ef2122231" />

You will be prompted to name your catalog, then a menu of capabilities will be displayed. Press `enter` to add any capabilities that are applicable to the service you are creating a catalog for. You will see the relevant identifiers added to your catalog on the right side of the screen. Press `delete` to remove a capability and the corresponding threats and controls. Threats and controls that are referenced by multiple capabilities will be managed appropriately to avoid duplication or errant removal.

<img width="1916" height="1076" alt="Screenshot 2025-09-11 at 04 43 40" src="https://github.com/user-attachments/assets/6c68f4e3-75db-439b-80b0-7b11311dc35d" />

Press `spacebar` when you are finished. The catalog will be prominently displayed for you to review. If more changes are needed, press `n` to return to the canvas. If you are finished, press `y` to write the catalog to a file in your current local directory: `./output.yaml`

<img width="1916" height="877" alt="Screenshot 2025-09-11 at 04 47 24" src="https://github.com/user-attachments/assets/86ddc40d-fd36-44a5-ace8-5b3dcf74c941" />

You may press `q` on any screen to exit the program.

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
