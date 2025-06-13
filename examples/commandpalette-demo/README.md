# Command Palette Demo

This demo application showcases the functionality of the command palette component from the Bobatea library.

## Features

- 🎯 Interactive command palette with fuzzy search
- ⌨️  Keyboard navigation (arrow keys, enter, escape)
- 🔍 Real-time command filtering
- 📝 Support for commands with and without parameters
- 🎨 Beautiful styling with Lipgloss
- 📱 Responsive layout

## Running the Demo

### Prerequisites

- Go 1.21 or later
- Optional: [VHS](https://github.com/charmbracelet/vhs) for generating demo recordings

### Quick Start

```bash
# Run the demo
make run

# Or directly with go
go run .
```

### Usage

- Press `Ctrl+K` to open the command palette
- Type to filter commands by name or description
- Use arrow keys to navigate
- Press `Enter` to execute a command
- Press `Escape` to close the palette
- Press `Ctrl+C` or `q` to quit

## Available Commands

The demo includes these sample commands:

- `/hello` - Say hello
- `/echo <message>` - Echo a message (needs parameters)
- `/time` - Show current time
- `/clear` - Clear the screen
- `/theme <light|dark>` - Change theme (needs parameters)
- `/help` - Show help
- `/quit` - Exit the application
- `/random` - Generate random number
- `/status` - Show app status
- `/version` - Show version info

## Demo Recordings

Generate animated GIF demonstrations of the command palette:

```bash
# Install VHS (if not already installed)
make install-vhs

# Generate all demo GIFs
make demos

# Generate specific demos
make demo-basic        # Basic usage
make demo-filtering    # Command filtering
make demo-navigation   # Keyboard navigation
```

The generated GIFs will be saved in the `demos/` directory:

- `basic-usage.gif` - Shows opening palette, filtering, and executing commands
- `command-filtering.gif` - Demonstrates the search/filter functionality
- `navigation.gif` - Shows keyboard navigation and selection

## Implementation Details

The demo application demonstrates:

1. **Command Palette Integration**: How to embed the command palette in a Bubble Tea application
2. **Event Handling**: Processing command selection events
3. **State Management**: Managing application state with command execution
4. **UI Layout**: Creating an overlay interface with the palette
5. **Responsive Design**: Adapting to different terminal sizes

## Code Structure

```
.
├── main.go              # Main application
├── demos/              # VHS tape files and generated GIFs
│   ├── basic-usage.tape
│   ├── command-filtering.tape
│   └── navigation.tape
├── Makefile           # Build and demo scripts
└── README.md          # This file
```

## Building

```bash
# Build executable
make build

# Run tests
make test

# Clean generated files
make clean
```

## Integration Example

Here's how to integrate the command palette into your own Bubble Tea application:

```go
import "github.com/bobatea/pkg/commandpalette"

// Create configuration
config := commandpalette.CommandPaletteConfig{
    Commands: []commandpalette.Command{
        {Name: "/save", Description: "Save file", Usage: "/save", NeedsParameters: false},
        // ... more commands
    },
    Width:       60,
    Height:      15,
    Title:       "Commands",
    Placeholder: "Search...",
}

// Create palette
palette := commandpalette.New(config)

// In your Update method
if palette.IsVisible() {
    palette, cmd = palette.Update(msg)
    return model, cmd
}

// Handle command events
case commandpalette.CommandSelectedMsg:
    // Execute command without parameters
case commandpalette.CommandWithParamsMsg:
    // Execute command with parameters

// In your View method
if palette.IsVisible() {
    return palette.View() // Show as overlay
}
```

## Contributing

Feel free to enhance the demo with additional features or create more sophisticated examples!
