package internal

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"regexp"
)

// Compiles the template located at path. Once the template has been created,
// a corresponding file in the output folder will be created and written.
func (bh *BlogHead) compile(p string) ([]byte, error) {
	// Get dependencies for the template and save to the BlogHead
	templates, err := bh.gatherTemplates(p)
	if err != nil {
		return nil, err
	}

	bh.saveDependencies(p, templates...)

	templates = append(templates, p)

	// Create a new named template from the html file
	t, err := template.New("html").ParseFiles(templates...)
	if err != nil {
		return nil, err
	}

	data, err := getTemplateData(p)
	if err != nil {
		return nil, err
	}

	if data != nil {
		// Set the data file as a dependency of the current page
		bh.saveDependencies(p, p[:len(p)-5]+"_meta.json")
	}

	var b []byte
	buf := bytes.NewBuffer(b)
	if err := t.Execute(buf, data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Takes a text file as input and parses the text to determine what
// templates are used in the file. Returns a string slice containing
// the file path of each template
func (bh *BlogHead) gatherTemplates(p string) ([]string, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	re := regexp.MustCompile("{{\\s*template\\s*\"([-_./\\w ]+)\"\\s*([.$\\w]+)?\\s*}}")

	text, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	matches := re.FindAllStringSubmatch(string(text), -1)

	filenames := []string{}
	for _, match := range matches {
		if len(match) > 1 {
			templateFile := path.Join(bh.tmplDir, match[1])
			filenames = append(filenames, templateFile)
		}
	}

	return filenames, nil
}

// Creates a file and any directories on the path that don't currently exist.
// Returns the newly opened file. The caller must close the file
func createFile(p string) (*os.File, error) {
	dir := path.Dir(p)
	if err := os.MkdirAll(dir, 0744); err != nil {
		return nil, err
	}

	f, err := os.Create(p)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// Look for a file with the name <filename>_meta.json
// This will contain data to be used in the template, if any
func getTemplateData(p string) (map[string]interface{}, error) {
	// Trim .html extension from name
	if f, err := os.Open(p[:len(p)-5] + "_meta.json"); err == nil {
		defer f.Close()

		d, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}
		data := make(map[string]interface{})
		if err := json.Unmarshal(d, &data); err != nil {
			return nil, err
		}

		return data, nil
	} else if !os.IsNotExist(err) {
		return nil, err
	}

	// Returns nil by default if no data file exists
	return nil, nil
}

// Trims the base path from the path p.
// If p does not start with base, then p is returned
func trimPath(base string, p string) string {
	if len(base) >= len(p) {
		return p
	}

	if p[:len(base)] == base {
		return p[len(base):]
	}

	return p
}
