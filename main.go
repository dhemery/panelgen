package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	inputDir    = "input"
	outputDir   = "output"
	templateDir = "templates"
)

func main() {
	err := filepath.Walk(filepath.Join(inputDir), renderSvgFiles)
	if err != nil {
		_ = fmt.Errorf("%w", err)
	}
}

func renderSvgFiles(path string, info os.FileInfo, err error) error {
	rel, err := filepath.Rel(inputDir, path)
	outPath := filepath.Join(outputDir, rel)
	fmt.Println("Will render", path, "to", outPath)
	return nil
}
