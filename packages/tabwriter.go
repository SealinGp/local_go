package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintf(w, "a\tb\tc\n")
	fmt.Fprintf(w, "aa\tbb\tcc\n")
	fmt.Fprintf(w, "aaa\t\n")
	fmt.Fprintf(w, "aaaa\tdddd\teeee\n")
	w.Flush()
}
