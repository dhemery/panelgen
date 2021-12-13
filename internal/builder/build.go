package builder

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"dhemery.com/panelgen/internal/panel"
)

func List() {
	slugs := strings.Join(allSlugs(), " ")
	fmt.Println(slugs)
}

func BuildAll() error {
	return Build(allSlugs())
}

func Build(slugs []string) error {
	builders := panel.Builders()
	for _, slug := range slugs {
		build, ok := builders[slug]
		if !ok {
			return errors.New("No such module: " + slug)
		}
		fmt.Println("Building", slug)
		if err := export(slug, build()); err != nil {
			return err
		}
	}
	return nil
}

func allSlugs() []string {
	slugs := []string{}
	for slug, _ := range panel.Builders() {
		slugs = append(slugs, slug)
	}
	return slugs
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
		return err
	}

	imagePath := filepath.Join(imageDir, "image", moduleSlug+".svg")
	if err := write(imagePath, p.ImageSvg()); err != nil {
		return err
	}

	for frameSlug, frameSvg := range p.FrameSvgs() {
		framePath := filepath.Join(frameDir, moduleSlug, frameSlug+".svg")
		if err := write(framePath, frameSvg); err != nil {
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
