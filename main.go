package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"

	"dhemery.com/panelgen/module/cubic"
	"dhemery.com/panelgen/panel"
)

func main() {
	panels := []*panel.Panel{
		cubic.Panel(),
	}

	for _, p := range panels {
		path := filepath.Join("out", p.Slug, p.Slug+".svg")
		if err := write(path, p.Faceplate()); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		path = filepath.Join("out", "image", p.Slug+".svg")
		if err := write(path, p.Image()); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		for slug, svg := range p.Frames() {
			path = filepath.Join("out", p.Slug, slug+".svg")
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
