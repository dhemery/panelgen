package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage:", filepath.Base(os.Args[0]), "in_file out_file")
		os.Exit(1)
	}
	fmt.Println(os.Args[1:])
}
