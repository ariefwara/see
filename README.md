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

### Prerequisites

- Go 1.24+
- [Fyne system dependencies](https://docs.fyne.io/started/) (Xcode / Command Line Tools on macOS, gcc / MinGW on Windows, GTK development packages on Linux)

### From source

```bash
git clone <this-repo> && cd see
go build -o see .
```

Or install directly to `$GOPATH/bin`:

```bash
go install .
```

> On macOS, if `go install` puts the binary in `$HOME/go/bin`, make sure that directory is in your `PATH`.

## Features

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
