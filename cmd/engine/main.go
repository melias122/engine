package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/melias122/engine/csv"
)

var (
	input = flag.String("input", "", "input file to process")
	n     = flag.Int("n", 5, "n dimension")
	m     = flag.Int("m", 35, "m dimension")
)

func main() {
	flag.Parse()

	file, err := os.Open(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "engine: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	parser := csv.NewParser(file, *n, *m)
	k, err := parser.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "engine: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(k)

	flag.Usage()
}
