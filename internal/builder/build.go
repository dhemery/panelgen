package builder

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"dhemery.com/panelgen/internal/panel"
	"dhemery.com/panelgen/internal/svg"
)

func Generate() error {
	contents := map[string][]byte{}

	for moduleSlug, p := range buildPanels() {
		imagePath := filepath.Join(imageDir, moduleSlug+".svg")
		imageBytes, err := encode(p.ImageSvg())
		if err != nil {
			return err
		}
		contents[imagePath] = imageBytes
		for frameSlug, frameSvg := range p.FrameSvgs() {
			framePath := filepath.Join(frameDir, moduleSlug, frameSlug+".svg")
			frameBytes, err := encode(frameSvg)
			if err != nil {
				return err
			}
			contents[framePath] = frameBytes
		}
	}

	for path, content := range contents {
		if !outdated(path, content) {
			fmt.Println("up to date:", path)
			continue
		}
		fmt.Println("outdated :", path)
		if err := write(path, content); err != nil {
			return err
		}
	}
	return nil
}

func List() {
	paths := []string{}
	for moduleSlug, modulePanel := range buildPanels() {
		imagePath := filepath.Join(imageDir, moduleSlug+".svg")
		paths = append(paths, imagePath)
		for _, c := range modulePanel.Controls {
			for frameSlug := range c.Frames {
				framePath := filepath.Join(frameDir, moduleSlug, frameSlug+".svg")
				paths = append(paths, framePath)
			}
		}
	}
	fmt.Println(strings.Join(paths, " "))
}

var builtPanels map[string]*panel.Panel

func buildPanels() map[string]*panel.Panel {
	if builtPanels == nil {
		builtPanels = map[string]*panel.Panel{}
		for slug, buildPanel := range panel.Builders() {
			builtPanels[slug] = buildPanel()
		}
	}
	return builtPanels
}

const (
	buildDir = "_build"
)

var (
	frameDir = filepath.Join(buildDir, "frames")
	imageDir = filepath.Join(buildDir, "images")
)

func encode(s svg.Svg) ([]byte, error) {
	b := &bytes.Buffer{}
	b.Write([]byte(xml.Header))
	e := xml.NewEncoder(b)
	e.Indent("", "    ")
	err := e.Encode(s)
	return b.Bytes(), err

}

func write(path string, content []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	w, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = w.Write(content)
	return err
}

func outdated(path string, content []byte) bool {
	b, err := os.ReadFile(path)
	if err != nil {
		return true
	}
	return bytes.Compare(b, content) != 0
}
