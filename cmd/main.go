package main

import (
	"flag"
	"ponyfetch/internal/fetch"
)

func main() {
	var ponyName string
	var ponySize string

	flag.StringVar(&ponyName, "p", "twilight", "Name of the pony for ASCII art")
	flag.StringVar(&ponyName, "pony", "twilight", "Name of the pony for ASCII art")

	flag.StringVar(&ponySize, "s", "normal", "Size of the pony ASCII art (normal, little)")
	flag.StringVar(&ponySize, "size", "normal", "Size of the pony ASCII art (normal, little)")

	flag.Parse()

	fetch.PrintFetch(ponyName, ponySize)
}
