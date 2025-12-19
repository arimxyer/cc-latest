# cc-latest

Fetch the latest Claude Code changelog entry from the command line.

## Installation

### Homebrew (macOS/Linux)

```bash
brew install arimxyer/tap/cc-latest
```

### Scoop (Windows)

```bash
scoop bucket add arimxyer https://github.com/arimxyer/scoop-bucket
scoop install cc-latest
```

### Go

```bash
go install github.com/arimxyer/cc-latest@latest
```

### From source

```bash
git clone https://github.com/arimxyer/cc-latest
cd cc-latest
go build -o cc-latest
```

## Usage

```bash
cc-latest              # Latest entry as plain text
cc-latest -json        # Latest entry as JSON
cc-latest -md          # Latest entry as raw markdown
cc-latest -version 2.0.70  # Specific version
cc-latest -list        # List all versions
```

## Flags

| Flag | Description |
|------|-------------|
| `-json` | Output as JSON |
| `-md` | Output raw markdown |
| `-version <X.X.X>` | Fetch specific version |
| `-list` | List all available versions |
| `-v` | Show cc-latest version |
| `-h` | Show help |

## Examples

### Plain text (default)

```
$ cc-latest
Claude Code 2.0.73
----------------------------------------
  * Added clickable `[Image #N]` links that open attached images
  * Fixed slow input history cycling
  * Improved theme picker UI
```

### JSON output

```
$ cc-latest -json
{
  "version": "2.0.73",
  "changes": [
    "Added clickable `[Image #N]` links...",
    "Fixed slow input history cycling..."
  ]
}
```

### Specific version

```
$ cc-latest -version 2.0.70
Claude Code 2.0.70
----------------------------------------
  * ...
```

## License

MIT
