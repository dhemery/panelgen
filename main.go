package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"

	"dhemery.com/panelgen/module"
)

func main() {
	for _, m := range module.All {
		fmt.Println("Rendering", m.Slug())
		fpPath := filepath.Join("out", m.Slug(), m.Slug()+".svg")
		if err := write(fpPath, m.Faceplate()); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		imagePath := filepath.Join("out", "image", m.Slug()+".svg")
		if err := write(imagePath, m.Image()); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		for path, svg := range m.Controls() {
			cPath := filepath.Join("out", m.Slug(), path+".svg")
			if err := write(cPath, svg); err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}

func write(path string, data interface{}) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	w, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() { _ = w.Close }()
	e := xml.NewEncoder(w)
	e.Indent("", "   ")
	return e.Encode(data)
}
