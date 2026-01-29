# Predator TUI

A terminal UI for managing power profiles on Acer Predator laptops.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white)

## Features

- Switch between power profiles (quiet, balanced, performance)
- Visual indicator for the currently active profile
- Keyboard navigation

## Installation

```bash
go install github.com/higorprado/predator-tui@latest
```

Or build from source:

```bash
git clone https://github.com/higorprado/predator-tui.git
cd predator-tui
go build -o predator-tui
```

## Usage

```bash
./predator-tui
```

### Controls

| Key | Action |
|-----|--------|
| `↑`/`↓` | Navigate profiles |
| `Enter` | Apply selected profile |
| `q`/`Esc` | Quit |

## Requirements

- Acer Predator laptop with platform profile support
- Linux with `/sys/firmware/acpi/platform_profile` interface

## License

MIT
