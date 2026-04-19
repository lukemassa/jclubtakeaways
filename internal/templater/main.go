package templater

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Templater struct {
	srcDir string
}

func New(srcDir string) Templater {
	return Templater{
		srcDir: srcDir,
	}
}

func (t Templater) getAllTemplates() ([]*template.Template, error) {
	ret := make([]*template.Template, 0)
	entries, err := os.ReadDir(t.srcDir)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if entry.Name() == "base.html" {
			continue
		}

		t, err := template.ParseFiles(filepath.Join(t.srcDir, "base.html"), filepath.Join(t.srcDir, entry.Name()))
		if err != nil {
			return nil, err
		}
		for _, tmplName := range t.Templates() {
			if tmplName.Name() != entry.Name() {
				continue
			}
			// Print the name of the template
			if !strings.HasSuffix(tmplName.Name(), ".html") {
				continue
			}
			ret = append(ret, tmplName)
		}
	}
	return ret, nil
}

func (t Templater) Write(dir string) error {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	allTemplates, err := t.getAllTemplates()
	if err != nil {
		return err
	}
	for _, tmpl := range allTemplates {
		outputFile, err := os.Create(filepath.Join(dir, tmpl.Name()))

		if err != nil {
			return fmt.Errorf("error creating output file: %v", err)
		}
		defer outputFile.Close()

		// Render the template (without any data, or you can pass data as needed)
		err = tmpl.Execute(outputFile, nil)
		if err != nil {
			return fmt.Errorf("error rendering template: %v", err)
		}
	}

	return nil
}

func (t Templater) getTemplateFromURL(url string) (*template.Template, error) {

	if !strings.HasSuffix(url, ".html") {
		return nil, errors.New("must end with .html")
	}
	if strings.Count(url, "/") != 1 {
		return nil, errors.New("does not look like path")
	}
	fullPath := filepath.Join(t.srcDir, url)
	_, err := os.Stat(fullPath)
	if err != nil {
		return nil, err
	}
	tmpl, err := template.ParseFiles(filepath.Join(t.srcDir, "base.html"), fullPath)
	if err != nil {
		return nil, err
	}
	for _, tmplName := range tmpl.Templates() {
		if tmplName.Name() != filepath.Base(fullPath) {
			continue
		}
		return tmplName, nil
	}
	return nil, errors.New("problem getting template")
}

func (t Templater) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl, err := t.getTemplateFromURL(r.URL.Path)
	if err != nil {
		log.Printf("404: %s: %v", r.URL.Path, err)
		http.NotFound(w, r)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
