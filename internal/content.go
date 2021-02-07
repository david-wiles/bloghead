package internal

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// Saves the current state of the site
func (bh *BlogHead) Save() error {
	return SaveConfig(bh.config, bh.configFile)
}

// Create a new template at the specified location
// This will create a new 'blueprint' file which will be copied
// when a new page is added using the 'add' command.
func (bh *BlogHead) Create(typ, name string) (err error) {
	switch typ {
	case "blueprint":
		err = bh.createBlueprint(name)
	default:
		errorStr := `Unknown type %v. Valid types are:

blueprint - creates a new named blueprint to initialize pages with
`
		err = errors.New(fmt.Sprintf(errorStr, typ))
	}
	if err := bh.Save(); err != nil {
		return err
	}
	return err
}

func (bh *BlogHead) createBlueprint(name string) error {
	// Check if the blueprint already exists
	if _, ok := bh.config.Blueprints[name]; ok {
		return nil
	}

	bp := path.Join(bh.tmplDir, "blueprints", name+".html")

	// Ensure blueprints directory exists
	if err := os.MkdirAll(path.Join(bh.tmplDir, "blueprints"), 0744); err != nil {
		return err
	}

	f, err := os.Create(bp)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write the required name of the template to the blueprint file
	if _, err := f.WriteString("{{ define \"html\" }}\n{{ end }}"); err != nil {
		return err
	}

	bh.config.Blueprints[name] = bp

	return nil
}

// Adds a new page of 'typ' at 'path'
func (bh *BlogHead) Add(typ, name, p string) (err error) {
	switch typ {
	case "page":
		err = bh.addNewPage(name, path.Join(bh.Root, p+".html"))
	default:
		errorStr := `Unknown type %v. Valid types are:

page - creates a new page at the specified path using the specified blueprint`
		err = errors.New(fmt.Sprintf(errorStr, typ))
	}
	return err
}

func (bh *BlogHead) addNewPage(name, p string) error {
	// Check that a blueprint with the name exists
	if _, ok := bh.config.Blueprints[name]; !ok {
		return errors.New("Could not find a blueprint named " + name + ". Did you remember to create it first?\n")
	}

	// Ensure that the directory exists
	if err := os.MkdirAll(path.Dir(p), 0744); err != nil {
		return err
	}

	// Check that a page doesn't already exist at the path
	_, err := os.Stat(p)
	if err == nil {
		return errors.New("Cannot create a page at " + p + ": already exists.")
	} else if !os.IsNotExist(err) {
		return err
	}

	// Copy the blueprint page to the new path
	html, err := ioutil.ReadFile(bh.config.Blueprints[name])
	if err != nil {
		return err
	}

	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(html); err != nil {
		return err
	}
	return nil
}
