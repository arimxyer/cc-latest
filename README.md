# aic

AI Coding Agent Changelog Viewer - fetch the latest changelog entries for popular AI coding assistants.

## Supported Tools

| Source | Command | Tool |
|--------|---------|------|
| `claude` | `aic claude` | Claude Code (Anthropic) |
| `codex` | `aic codex` | Codex CLI (OpenAI) |
| `opencode` | `aic opencode` | OpenCode (SST) |
| `gemini` | `aic gemini` | Gemini CLI (Google) |
| `copilot` | `aic copilot` | Copilot CLI (GitHub) |

## Installation

### Homebrew (macOS/Linux)

```bash
brew install arimxyer/tap/aic
```

### Scoop (Windows)

```bash
scoop bucket add arimxyer https://github.com/arimxyer/scoop-bucket
scoop install aic
```

### Go

```bash
go install github.com/arimxyer/aic@latest
```

### From source

```bash
git clone https://github.com/arimxyer/aic
cd aic
go build -o aic
```

## Usage

```bash
aic <source> [flags]
```

### Examples

```bash
aic claude                    # Latest Claude Code changelog
aic codex -json               # Latest Codex changelog as JSON
aic opencode -list            # List all OpenCode versions
aic gemini -version 0.1.0     # Specific Gemini CLI version
aic copilot -md               # Latest Copilot changelog as markdown
```

## Flags

| Flag | Description |
|------|-------------|
| `-json` | Output as JSON |
| `-md` | Output as markdown |
| `-list` | List all available versions |
| `-version <ver>` | Fetch specific version |
| `-v` | Show aic version |
| `-h` | Show help |

## Output Examples

### Plain text (default)

```
$ aic claude
Claude Code 2.0.73
----------------------------------------
  * Added clickable `[Image #N]` links that open attached images
  * Fixed slow input history cycling
  * Improved theme picker UI
```

### JSON output

```
$ aic codex -json
{
  "version": "0.0.1",
  "changes": [
    "Initial release",
    "Added support for..."
  ]
}
```

### List versions

```
$ aic opencode -list
0.2.0
0.1.9
0.1.8
...
```

## License

MIT
