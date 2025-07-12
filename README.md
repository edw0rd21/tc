# tc - Terminal Clipboard Manager

A lightweight CLI tool for macOS that monitors your clipboard history and lets you easily access, copy, and manage past clipboard content from the terminal.

## Features

- 🕐 **Background monitoring** - Continuously watches your clipboard in the background
- 📋 **History management** - Stores up to 100 clipboard items with timestamps
- 🔍 **Easy browsing** - List recent items with formatted output
- 📤 **Quick copying** - Copy any historical item back to your clipboard
- 🧹 **History clearing** - Clean up stored history when needed
- ⚡ **Fast & lightweight** - Minimal resource usage with efficient storage

## Installation

### Prerequisites
- macOS (required for clipboard integration)
- Go 1.24.5 or later

### Build from Source
```bash
git clone https://github.com/edw0rd21/tc.git
cd tc
go build -o tc .
```

### Install to PATH
```bash
# Move to a directory in your PATH
sudo mv tc /usr/local/bin/
```

## Quick Start

1. **Start the background daemon** (monitors clipboard changes):
   ```bash
   tc daemon
   ```

2. **View recent clipboard items**:
   ```bash
   tc list
   ```

3. **Copy a specific item back to clipboard**:
   ```bash
   tc copy 2
   ```

## Commands

### `tc daemon`
Start the background clipboard watcher. This should be running to capture clipboard changes.

```bash
tc daemon
```

**Tip**: Add this to your shell profile (`.zshrc`, `.bashrc`) to start automatically:
```bash
# Start tc daemon in background if not already running
if ! pgrep -f "tc daemon" > /dev/null; then
    tc daemon &
fi
```

### `tc list`
Display clipboard history items.

```bash
# Show last 10 items (default)
tc list

# Show last 5 items
tc list -n 5

# Show specific item by index
tc list 3

# Show full content without truncation
tc list --full

# Show raw content only (no formatting)
tc list --raw
```

**Options:**
- `-n, --count <number>` - Number of items to show (default: 10)
- `-f, --full` - Show full content without truncation
- `-r, --raw` - Show only raw clipboard content

### `tc copy`
Copy a specific item from history back to your clipboard.

```bash
# Copy item #2 to clipboard
tc copy 2

# Copy with preview of what was copied
tc copy 2 --preview

# Copy silently (no output)
tc copy 2 --quiet
```

**Options:**
- `-l, --limit <number>` - Maximum number of items to search through (default: 100)
- `-p, --preview` - Preview the content being copied
- `-q, --quiet` - Suppress output messages

### `tc clear`
Clear all stored clipboard history.

```bash
# Clear history with confirmation
tc clear

# Clear history silently
tc clear --quiet
```

**Options:**
- `-q, --quiet` - Suppress output messages

## Examples

### Basic Workflow
```bash
# Start daemon (run once)
tc daemon &

# Copy some text, then view history
tc list
# Output:
# 1➤ [14:32:15] Hello world
# 2➤ [14:31:45] https://github.com/edw0rd21/tc
# 3➤ [14:30:22] git commit -m "Add new feature"

# Copy item #2 back to clipboard
tc copy 2
# Output: Copied item 2 to clipboard

# View specific item in full
tc list 2 --full
# Output: 2➤ [14:31:45] https://github.com/edw0rd21/tc
```

### Advanced Usage
```bash
# Show only raw content (useful for piping)
tc list 1 --raw | pbcopy

# Preview what you're copying
tc copy 3 --preview

# Search through more items
tc copy 5 --limit 200
```

## Configuration

### Storage Location
Clipboard history is stored in `~/.tc/history.json`

### Limits
- **Maximum items**: 100 (configurable in code)
- **Polling interval**: 500ms
- **Display truncation**: 80 characters (use `--full` to see complete content)

## Architecture

```
tc/
├── cmd/                    # CLI commands
│   ├── root.go            # Main command setup
│   ├── daemon.go          # Background watcher
│   ├── list.go            # List command
│   ├── copy.go            # Copy command
│   └── clear.go           # Clear command
├── internal/
│   ├── clipboard/         # Business logic
│   │   └── manager.go     # Clipboard management
│   ├── daemon/            # Background processes
│   │   └── watcher.go     # Clipboard monitoring
│   └── storage/           # Data persistence
│       └── storage.go     # JSON file storage
└── main.go                # Entry point
```

## How It Works

1. **Daemon Process**: Continuously monitors clipboard changes every 500ms
2. **Storage**: Saves clipboard history as JSON in `~/.tc/history.json`
3. **Deduplication**: Avoids storing duplicate consecutive items
4. **LIFO Ordering**: Newest items appear first (index 1)
5. **Automatic Cleanup**: Maintains only the most recent 100 items

## Troubleshooting

### Daemon Not Capturing Changes
```bash
# Check if daemon is running
pgrep -f "tc daemon"

# Restart daemon
pkill -f "tc daemon"
tc daemon &
```

### Permission Issues
```bash
# Check .tc directory permissions
ls -la ~/.tc/

# Fix permissions if needed
chmod 755 ~/.tc/
chmod 644 ~/.tc/history.json
```

### Empty History
- Ensure the daemon is running: `tc daemon &`
- Copy something to test: `echo "test" | pbcopy`
- Check history: `tc list`

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Make your changes
4. Add tests if applicable
5. Commit your changes: `git commit -m "Add feature"`
6. Push to the branch: `git push origin feature-name`
7. Submit a pull request

## License

MIT License - see LICENSE file for details.

## Related Projects

- [pbcopy/pbpaste](https://ss64.com/osx/pbcopy.html) - macOS built-in clipboard utilities
- [copyq](https://github.com/hluk/CopyQ) - Cross-platform clipboard manager with GUI
- [clipy](https://github.com/Clipy/Clipy) - macOS clipboard extension app

---

**Made with ❤️ for terminal productivity**