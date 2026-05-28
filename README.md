# Ponyfetch

A lightweight, cross-platform CLI system information utility written in Go. Like `fastfetch` or `neofetch`, but it features cute pony ASCII arts instead of boring OS logos.

## Features
* **Cross-platform support:** Native system statistics parsing for both **macOS** and **Linux (Arch, Debian, etc.)** without any external dependencies.
* **Customizable CLI:** Choose your favorite pony and size using built-in flags.

## Installation

### From Source
Make sure you have Go installed, then clone and build the project:
```bash
git clone https://github.com/Krev3tka/Ponyfetch.git
cd Ponyfetch
go build -o ponyfetch cmd/main.go
./ponyfetch
```

## Usage
```Bash
ponyfetch [options]
```

Available Flags:
| Short | Long | Description | Default Value |
| :--- | :--- | :--- | :--- |
| `-p` | `--pony` | Pony name for ASCII-art | `twilight` |
| `-s` | `--size` | Art size (`little` or `normal`) | `normal` |
| `-h` | `--help` | Show manual for utility | |
| `-c` | `--color` | Art color (all ANSI colors, like `red`, `blue`, `green`, etc.)| reset |
| `-f` | `--filename` | File name or path to custom ASCII art, requires .txt format | |

#### Short flags are not working now


### Usage examples

```Bash
# Run with default Twilight art
ponyfetch

# Display Fluttershy in a smaller format
ponyfetch -p fluttershy -s little

# Trigger the smart fallback (displays Twilight Sparkle)
ponyfetch --pony=unknown_pony_name
```

## Roadmap
- [x] Default and custom ASCII arts
- [x] Dynamic padding support for art and info printing
- [ ] Full system and hardware info parsing
- [ ] Configuration with .toml files

License
This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.
