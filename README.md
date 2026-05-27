# Ponyfetch

goida goida zov svo svo zzzzzzz zoooov
A lightweight, cross-platform CLI system information utility written in Go. Like `fastfetch` or `neofetch`, but it features cute pony ASCII arts instead of boring OS logos.

## Features
* **Cross-platform support:** Native system statistics parsing for both **macOS** and **Linux (Arch, Debian, etc.)** without any external dependencies.
* **Customizable CLI:** Choose your favorite pony and size using built-in flags.

## Installation

### From Source
Make sure you have Go installed, then clone and build the project:
```bash
git clone [https://github.com/Krev3tka/Ponyfetch.git](https://github.com/Krev3tka/Ponyfetch.git)
cd Ponyfetch
go build -o ponyfetch cmd/main.go
mv ponyfetch /usr/local/bin/ # Optional: move to PATH
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

```Bash
# Run with default Twilight art
ponyfetch

# Display Fluttershy in a smaller format
ponyfetch -p fluttershy -s little

# Trigger the smart fallback (displays Twilight Sparkle)
ponyfetch --pony=unknown_pony_name
```

📜 LicenseThis project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.
