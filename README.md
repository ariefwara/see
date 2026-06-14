# see — Minimal Image Previewer

A lightweight desktop image viewer written in Go using the [Fyne](https://fyne.io/) GUI toolkit. Opens an image file in a scrollable window with a single command.

## Usage

```bash
see <image-file>
```

### Examples

```bash
see photo.png
see ~/Pictures/screenshot.jpg
see /absolute/path/to/image.gif
```

## Supported Formats

| Format | Extension |
|--------|-----------|
| PNG    | `.png`    |
| JPEG   | `.jpg`, `.jpeg` |
| GIF    | `.gif`    |

## Installation

### Quick install (macOS / Linux)

```bash
bash -c "$(curl -fsSL https://raw.githubusercontent.com/ariefwara/see/main/install.sh)"
```

### Quick install (Windows PowerShell)

```powershell
powershell -ExecutionPolicy Bypass -c "irm https://raw.githubusercontent.com/ariefwara/see/main/install.ps1 | iex"
```

The install script will:
1. Detect your OS and architecture
2. Download a pre-built binary from the latest release (if available)
3. Install it to `/usr/local/bin/see` (macOS/Linux) or `$USERPROFILE\go\bin\see.exe` (Windows)
4. Fall back to building from source if no release exists

> Set `SEE_BIN` environment variable to customize the install path.

### From source

```bash
git clone https://github.com/ariefwara/see.git && cd see
go build -o see .
```

Or install directly to `$GOPATH/bin`:

```bash
go install .
```

## Controls

| Key | Action |
|-----|--------|
| `←` / `→` | Navigate to previous / next image in the same directory |
| `Esc` | Close window |

Images in the directory are sorted alphabetically.

## Features

- **Keyboard navigation** — browse through images without leaving the app
- **Drag-to-scroll** — pan around large images that exceed the window size
- **Input validation** — checks file existence, extension, and image integrity before opening
- **Zero web dependencies** — no HTML, CSS, or JavaScript required
- **Small binary** — standalone, no external runtime needed

## Error Handling

| Scenario | Output |
|----------|--------|
| No argument | `Usage: see <image-file>` |
| File not found | `Error: file 'xxx' not found` |
| Unsupported format | `Error: unsupported format '...'` |
| Corrupt / invalid image | `Error: 'xxx' is not a valid image file` |

## Tech Stack

- **Go** — standard library + `image`, `image/png`, `image/jpeg`, `image/gif`
- **Fyne v2** — cross-platform GUI toolkit
- **License**: MIT
