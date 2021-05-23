package main

import (
	"encoding/xml"
	"os"
	"path/filepath"

	"dhemery.com/panelgen/module"
)

func main() {
	for _, m := range module.Modules {
		_ = write(m.Faceplate())
		_ = write(m.Image())
		for _, c := range m.Controls() {
			_ = write(c)
		}
	}
}

func write(s module.Slugger) error {
	path := filepath.Join("out", s.Slug()+".svg")
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	w, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() { _ = w.Close }()
	e := xml.NewEncoder(w)
	return e.Encode(s)
}
