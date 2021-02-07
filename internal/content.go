package internal

import (
	"os"
	"path"
)

// Create a new template at the specified location
// This will create a new 'blueprint' file which will be copied
// when a new page is added using the 'add' command.
func (bh *BlogHead) Create(typ, name string) error {
	return bh.createBlueprint(name)
}

func (bh *BlogHead) createBlueprint(name string) error {
	// Check if the blueprint already exists
	if _, ok := bh.config.Blueprints[name]; ok {
		return nil
	}

	bp := path.Join(bh.tmplDir, "blueprints", name+".html")

	f, err := os.Create(bp)
	if err != nil {
		return err
	}

	// Write the required name of the template to the blueprint file
	if _, err := f.WriteString("{{ define \"html\" }}\n{{ end }}"); err != nil {
		return err
	}

	bh.config.Blueprints[name] = bp

	return nil
}
