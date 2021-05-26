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
		path := filepath.Join("out", m.Slug, m.Slug+".svg")
		if err := write(path, m.Faceplate); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		path = filepath.Join("out", "image", m.Slug+".svg")
		if err := write(path, m.Overlay); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		for slug, svg := range m.Frames {
			path = filepath.Join("out", m.Slug, slug+".svg")
			if err := write(path, svg); err != nil {
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
	e.Indent("", "    ")
	return e.Encode(data)
}
