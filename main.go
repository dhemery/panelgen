package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"

	"dhemery.com/panelgen/panel"
	boosterstage "dhemery.com/panelgen/panel/booster-stage"
	"dhemery.com/panelgen/panel/cubic"
)

func main() {
	type buildFunc func() *panel.Panel

	builders := map[string]buildFunc{}

	builders["booster-stage"] = boosterstage.Panel
	builders["cubic"] = cubic.Panel

	for slug, buildFn := range builders {
		export(slug, buildFn())
	}
}

const (
	buildDir = "_build"
)

var (
	frameDir     = filepath.Join(buildDir, "controls")
	faceplateDir = filepath.Join(buildDir, "faceplates")
	imageDir     = filepath.Join(buildDir, "images")
)

func export(moduleSlug string, p *panel.Panel) error {
	faceplatePath := filepath.Join(faceplateDir, moduleSlug+".svg")
	if err := write(faceplatePath, p.FaceplateSvg()); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return err
	}

	imagePath := filepath.Join(imageDir, "image", moduleSlug+".svg")
	if err := write(imagePath, p.ImageSvg()); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return err
	}

	for frameSlug, frameSvg := range p.FrameSvgs() {
		framePath := filepath.Join(frameDir, moduleSlug, frameSlug+".svg")
		if err := write(framePath, frameSvg); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			return err
		}
	}
	return nil
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
	w.Write([]byte(xml.Header))
	e := xml.NewEncoder(w)
	e.Indent("", "    ")
	return e.Encode(data)
}
