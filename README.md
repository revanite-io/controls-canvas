# Controls Canvas

A terminal-based UI tool for interactively selecting and managing cloud security controls.

## Features

- Interactive TUI for browsing and selecting cloud security controls
- Support for Common Cloud Controls (CCC) framework
- Real-time YAML preview of selected controls
- Dynamic window resizing with responsive layout
- Efficient filtering and navigation

## Installation

### Using Go

```bash
go install github.com/revanite-io/controls-canvas@latest
```

### From Releases

Download the latest binary for your platform from the [releases page](https://github.com/revanite-io/controls-canvas/releases).

## Usage

```bash
controls-canvas
```

### Navigation

- `↑/↓`: Navigate through items
- `Enter`: Select a capability
- `Backspace/x`: Deselect a capability
- `Space`: Generate output
- `/`: Filter items
- `Esc`: Clear filter/Exit
- `Ctrl+C`: Quit

## Building from Source

```bash
git clone https://github.com/revanite-io/controls-canvas.git
cd controls-canvas
go build
```

## Requirements

- Go 1.23 or later
- Terminal with support for ANSI escape sequences
- Minimum terminal width of 80 characters recommended

## Output

The tool generates an `output.yaml` file containing:
- Selected capabilities
- Associated threats
- Mapped controls

## Contributing

Contributions are welcome! Please read our [contributing guidelines](CONTRIBUTING.md) before submitting pull requests.

## License

[Apache License 2.0](LICENSE) 