package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
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
	case "template":
		err = bh.createTemplate(name)
	case "blueprint":
		err = bh.createBlueprint(name)
	default:
		errorStr := `Unknown type %v. Valid types are:

blueprint - creates a new named blueprint to initialize pages with
template  - create a blank template in the templates directory 
`
		err = errors.New(fmt.Sprintf(errorStr, typ))
	}
	if err := bh.Save(); err != nil {
		return err
	}
	return err
}

// Create a template with the specified name
// Simply generates the required markup so the compiler will be
// able to correctly define the templates, the user must create the
// HTML manually
func (bh *BlogHead) createTemplate(name string) error {
	tmplName := name + ".html"
	f, err := createFile(path.Join(bh.tmplDir, tmplName))
	if err != nil {
		return err
	}
	_ = f.Close()
	return nil
}

func (bh *BlogHead) createBlueprint(name string) error {
	// Check if the blueprint already exists
	if _, ok := bh.config.Blueprints[name]; ok {
		return nil
	}

	bp := path.Join(bh.tmplDir, "blueprints", name+".html")

	f, err := createFile(bp)
	if err != nil {
		return err
	}
	_ = f.Close()

	bh.config.Blueprints[name] = bp

	return nil
}

// Adds a new page of 'typ' at 'path'
func (bh *BlogHead) Add(typ, name, bp string) (err error) {
	switch typ {
	case "page":
		err = bh.addNewPage(bp, path.Join(bh.Root, name+".html"))
	case "article":
		err = bh.addNewArticle(bp, path.Join(bh.Root, name+".html"))
	default:
		errorStr := `Unknown type %v. Valid types are:

page - creates a new page at the specified path using the specified blueprint
article - a blog post used in a sequence of posts
`
		err = errors.New(fmt.Sprintf(errorStr, typ))
	}
	if err := bh.Save(); err != nil {
		return err
	}
	return err
}

// Creates a new generic web page based on the named template
func (bh *BlogHead) addNewPage(bp, name string) error {
	// Check that a blueprint with the bp exists
	// If bp is an empty string, we should skip this and initialize an empty page
	if _, ok := bh.config.Blueprints[bp]; bp != "" && !ok {
		return errors.New("Could not find a blueprint named " + bp + ". Did you remember to create it first?\n")
	}

	// Ensure that the directory exists
	if err := os.MkdirAll(path.Dir(name), 0744); err != nil {
		return err
	}

	// Check that a page doesn't already exist at the path
	_, err := os.Stat(name)
	if err == nil {
		return errors.New("Cannot create a page at " + name + ": already exists.")
	} else if !os.IsNotExist(err) {
		return err
	}

	html := []byte("")
	if bp != "" {
		// Copy the blueprint page to the new path
		html, err = ioutil.ReadFile(bh.config.Blueprints[bp])
		if err != nil {
			return err
		}
	}

	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(html); err != nil {
		return err
	}
	return nil
}

// Creates a new article based on the named template and adds the
// page to a list of articles. When the site is published, this article's
// content is added to a feed.xml file entry
func (bh *BlogHead) addNewArticle(bp, name string) error {
	// Copy the html page from the template
	if err := bh.addNewPage(bp, name); err != nil {
		return err
	}

	// Create a new meta.json for the page with date entries
	if err := bh.addDefaultMeta(name); err != nil {
		return err
	}

	if err := bh.createContentFile(name); err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// Add the page to articles list
	bh.config.Articles = appendUnique(bh.config.Articles, "."+trimPath(cwd, name))
	return nil
}

func (bh *BlogHead) createContentFile(name string) error {
	// Create empty file for content in .data
	f, err := createFile(path.Join(bh.tmplDir, ".data", path.Base(name), "content.html"))
	if err != nil {
		return err
	}
	_ = f.Close()
	return nil
}

func (bh *BlogHead) addDefaultMeta(page string) error {
	type defaultMeta struct {
		Title   string `json:"title"`
		Updated string `json:"updated"`
		Link    string `json:"link"`
	}

	title := path.Base(page)
	meta := &defaultMeta{
		title[:len(title)-5],
		time.Now().Format(time.RFC3339),
		path.Join(bh.config.Domain, trimPath(bh.Root, page)),
	}

	b, err := json.Marshal(meta)
	if err != nil {
		return err
	}

	f, err := os.Create(page[:len(page)-5] + "_meta.json")
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(b); err != nil {
		return err
	}

	return nil
}
