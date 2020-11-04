package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage:", filepath.Base(os.Args[0]), "in_file out_file")
		os.Exit(1)
	}
	in := os.Args[1]
	err := render(in)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func render(in string) error {
	f, err := template.ParseFiles(in, "templates/port-component.svg", "templates/port-file.svg")
	if err != nil {
		return fmt.Errorf("reading templates: %w", err)
	}
	return f.Execute(os.Stdout, "")
}
