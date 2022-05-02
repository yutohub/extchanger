package main

import (
	"extchanger/converter"
	"flag"
	"fmt"
	"os"
)

var (
	from string
	to   string
	dir  string
)

func init() {
	flag.StringVar(&from, "from", "jpg", "From which extension.")
	flag.StringVar(&to, "to", "png", "To which extension.")
	flag.StringVar(&dir, "dir", ".", "Directory with images.")
}

func main() {
	flag.Parse()
	// Get conv structure
	conv, err := converter.NewConv(from, to, dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Do conversion
	if err := conv.Do(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
