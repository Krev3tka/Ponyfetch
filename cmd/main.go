package main

import (
	"flag"
	"ponyfetch/internal/fetch"
)

func main() {
	var ponyName string
	var ponySize string
	var color string
	var customArtFromFile string

	flag.StringVar(&ponyName, "pony", "twilight", "Name of the pony for ASCII art")

	flag.StringVar(&ponySize, "size", "normal", "Size of the pony ASCII art (normal, little)")

	flag.StringVar(&color, "color", "reset", "Color of ASCII art (red, green, yellow, blue, purple, cyan, black and reset for default)")

	flag.StringVar(&customArtFromFile, "filename", "", "Path or name of file, whish contains custom ASCII art. Requires \".txt\" format")

	flag.Parse()

	fetch.PrintFetch(ponyName, ponySize, color, customArtFromFile)
}
