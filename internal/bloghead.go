package internal

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
)

// BlogHead keeps track of everything going on during the site's build process.
type BlogHead struct {
	Root    string
	Output  string
	tmplDir string

	// Configuration file for the site created by this bloghead
	// This file also stores the state of the site
	configFile string
	config     *BlogConfig

	// Templates is a map of each template and the templates is is used in.
	// When running in watch mode, this is used to determine which files to watch
	templates map[string][]string

	// The filesystem watcher used when running with the watch option
	// Does not have a value unless the watch option is set
	watcher *fsnotify.Watcher
}

func FromEnv() *BlogHead {
	rootStr := viper.GetString("root")
	outStr := viper.GetString("output")

	rootPath, err := filepath.Abs(rootStr)
	if err != nil {
		panic(err)
	}

	outPath, err := filepath.Abs(outStr)
	if err != nil {
		panic(err)
	}

	config, err := ReadConfig(viper.ConfigFileUsed())
	if err != nil {
		panic(err)
	}

	return &BlogHead{
		Root:       rootPath,
		Output:     outPath,
		tmplDir:    path.Join(rootPath, ".templates"),
		configFile: viper.ConfigFileUsed(),
		config:     config,
		templates:  make(map[string][]string),
	}
}

func Init(root, output, filename string) error {
	// Write a blank config file
	if err := SaveConfig(&BlogConfig{
		Root:       root,
		Output:     output,
		Blueprints: make(map[string]string),
	}, filename); err != nil {
		return err
	}
	return nil
}

// Start compiling pages found in the root directory
// Ignores the directory named '.templates'
func (bh *BlogHead) Start() error {
	return filepath.Walk(bh.Root, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		absPath, err := filepath.Abs(p)
		if err != nil {
			return err
		}

		if bh.isHTMLPage(absPath, info) {
			if err := bh.compile(absPath); err != nil {
				return err
			}
		}

		return nil
	})
}

// Watch initializes the filesystem watcher for all files found
// in the root directory, including the '.templates' directory.
// On a file change, the file is rebuilt along with all files which
// use the changed template. The site is created before the watcher
// is initialized
func (bh *BlogHead) Watch() error {
	// Build all files
	if err := bh.Start(); err != nil {
		return err
	}

	// Watch files for changes
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	bh.watcher = watcher

	// Register listeners on each file
	if err := filepath.Walk(bh.Root, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if err := bh.watcher.Add(p); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	// When a filesystem change is detected, handle the event
	done := make(chan bool)
	go bh.watchFiles()
	<-done

	return nil
}

func (bh *BlogHead) watchFiles() {
	for {
		select {
		case event, ok := <-bh.watcher.Events:
			if !ok {
				// If the watcher was closed, return
				return
			} else {
				if event.Op&fsnotify.Write == fsnotify.Write {
					p, err := filepath.Abs(event.Name)
					if err != nil {
						println(err.Error())
					}

					// Rewrite all pages dependent on the modified file
					if err := bh.walkDependencies(p, func(p string) error {
						// If the trimmed path is equal to the original path, then the
						// page is not in the template directory
						if trimPath(bh.tmplDir, p) == p {
							if err := bh.compile(p); err != nil {
								return err
							}
						}
						return nil
					}); err != nil {
						println(err.Error())
					}
				}
			}
		case err, ok := <-bh.watcher.Errors:
			if !ok {
				// If the watcher was closed, return
				return
			} else {
				println(err.Error())
			}
		}
	}
}

// Determine if the file at the path p should be processed as a page
// The conditions are:
//   1: has the .html file extension
//   2: is not a directory
//   3: is not in the templates directory or one of its subdirectories
func (bh *BlogHead) isHTMLPage(p string, info os.FileInfo) bool {
	return len(info.Name()) > 5 &&
		p[len(p)-5:] == ".html" &&
		!info.IsDir() &&
		// If the trimmed path is equal to the original path,
		// then the template directory is not a parent directory of the file
		trimPath(bh.tmplDir, p) == p
}

func (bh *BlogHead) saveDependencies(p string, templates ...string) {
	// Add entries for each dependency in the templates map
	for _, tmpl := range templates {
		if list, ok := bh.templates[tmpl]; ok {
			bh.templates[tmpl] = appendUnique(list, p)
		} else {
			bh.templates[tmpl] = []string{p}
		}
	}
}

// Walk through each page dependent page on p and call the walkFn
// for each page, including p
func (bh *BlogHead) walkDependencies(p string, walkFn func(p string) error) error {
	// TODO detect circular dependencies
	if deps, ok := bh.templates[p]; ok {
		for _, dep := range deps {
			if err := walkFn(dep); err != nil {
				return err
			}
			// Walk through all dependencies of dep
			if err := bh.walkDependencies(dep, walkFn); err != nil {
				return err
			}
		}
	}

	if err := walkFn(p); err != nil {
		return err
	}

	return nil
}

// Adds the value to the list if it doesn't already exist
func appendUnique(list []string, val string) []string {
	for _, el := range list {
		if el == val {
			return list
		}
	}
	return append(list, val)
}
