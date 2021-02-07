package internal

import (
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"testing"
)

type BHFields struct {
	Root       string
	Output     string
	tmplDir    string
	configFile string
	config     *BlogConfig
	templates  map[string][]string
	watcher    *fsnotify.Watcher
}

func makeTestBH(testName string) BHFields {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return BHFields{
		Root:       path.Join(cwd, "../testdata", testName),
		Output:     path.Join(cwd, "../testdata/.output", testName),
		tmplDir:    path.Join(cwd, "../testdata", testName, ".templates"),
		configFile: path.Join(cwd, ".bloghead_test"),
		config: &BlogConfig{
			Root:       "../testdata/" + testName,
			Output:     "../testdata/.output/" + testName,
			Blueprints: make(map[string]string),
		},
		templates: make(map[string][]string),
		watcher:   nil,
	}
}

func TestBlogHead_Start(t *testing.T) {
	tests := []struct {
		name    string
		fields  BHFields
		wantErr bool
		want    func() bool
	}{
		{
			name:    "Should not do anything if root directory is empty",
			fields:  makeTestBH("empty"),
			wantErr: false,
			want: func() bool {
				// Expect the output directory to be empty
				files, _ := ioutil.ReadDir("../testdata/.output/empty")
				return len(files) == 0
			},
		},
		{
			name:    "Creates a basic html page from template and data",
			fields:  makeTestBH("basic"),
			wantErr: false,
			want: func() bool {
				// Expect the output directory to have a single file matching
				// the test index.html string
				want := `
<!DOCTYPE html>
<html lang="en">

    <head>
        <meta charset="UTF-8">
        <title>Hello, World!</title>
    </head>

<body>
<h1>
    Hello, World!
</h1>
</body>
</html>
`
				files, _ := ioutil.ReadDir("../testdata/.output/basic")
				if len(files) != 1 {
					return false
				}
				if got, err := ioutil.ReadFile("../testdata/.output/basic/index.html"); err == nil {
					return string(got) == want
				}
				return false
			},
		},
	}
	for _, tt := range tests {
		// Create output directory
		if err := os.MkdirAll(tt.fields.Output, 0744); err != nil {
			panic(err)
		}

		t.Run(tt.name, func(t *testing.T) {
			bh := &BlogHead{
				Root:       tt.fields.Root,
				Output:     tt.fields.Output,
				tmplDir:    tt.fields.tmplDir,
				configFile: tt.fields.configFile,
				config:     tt.fields.config,
				templates:  tt.fields.templates,
				watcher:    tt.fields.watcher,
			}
			if err := bh.Start(); (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := tt.want(); !got {
				t.Errorf("Want condition not satisfied")
			}
		})
	}
	// Remove all generated files
	t.Cleanup(func() {
		for _, tt := range tests {
			if err := os.RemoveAll(tt.fields.Output); err != nil {
				println(err.Error())
			}
		}
	})
}

func TestBlogHead_isHTMLPage(t *testing.T) {
	type args struct {
		p    string
		info os.FileInfo
	}
	tests := []struct {
		name   string
		fields BHFields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bh := &BlogHead{
				Root:       tt.fields.Root,
				Output:     tt.fields.Output,
				tmplDir:    tt.fields.tmplDir,
				configFile: tt.fields.configFile,
				config:     tt.fields.config,
				templates:  tt.fields.templates,
				watcher:    tt.fields.watcher,
			}
			if got := bh.isHTMLPage(tt.args.p, tt.args.info); got != tt.want {
				t.Errorf("isHTMLPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlogHead_saveDependencies(t *testing.T) {
	type args struct {
		p         string
		templates []string
	}
	tests := []struct {
		name   string
		fields BHFields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//bh := &BlogHead{
			//	Root:       tt.fields.Root,
			//	Output:     tt.fields.Output,
			//	tmplDir:    tt.fields.tmplDir,
			//	configFile: tt.fields.configFile,
			//	config:     tt.fields.config,
			//	templates:  tt.fields.templates,
			//	watcher:    tt.fields.watcher,
			//}
		})
	}
}

func TestBlogHead_walkDependencies(t *testing.T) {
	type args struct {
		p      string
		walkFn func(p string) error
	}
	tests := []struct {
		name    string
		fields  BHFields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bh := &BlogHead{
				Root:       tt.fields.Root,
				Output:     tt.fields.Output,
				tmplDir:    tt.fields.tmplDir,
				configFile: tt.fields.configFile,
				config:     tt.fields.config,
				templates:  tt.fields.templates,
				watcher:    tt.fields.watcher,
			}
			if err := bh.walkDependencies(tt.args.p, tt.args.walkFn); (err != nil) != tt.wantErr {
				t.Errorf("walkDependencies() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFromEnv(t *testing.T) {
	tests := []struct {
		name string
		want *BlogHead
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromEnv(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInit(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Init(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_appendUnique(t *testing.T) {
	type args struct {
		list []string
		val  string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appendUnique(tt.args.list, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
