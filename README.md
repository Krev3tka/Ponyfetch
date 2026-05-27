# Ponyfetch

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
