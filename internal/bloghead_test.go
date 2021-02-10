package internal

import (
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	"os"
	"path"
	"sort"
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

func unwrap(val interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	} else {
		return val
	}
}

// Tests whether two maps contain the same key->value pairs
func mapsAreEqual(m1, m2 map[string][]string) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v1 := range m1 {
		// Loop over each value in v and compare
		v2, ok := m2[k]
		if !ok {
			return false
		}

		sort.Strings(v1)
		sort.Strings(v2)

		for i := range v2 {
			if v1[i] != v2[i] {
				return false
			}
		}
	}

	return true
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
				want := `<!DOCTYPE html>
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
	cwd := unwrap(os.Getwd()).(string)

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
		{
			name:   "plain HTML file outside template directory",
			fields: makeTestBH("basic"),
			args: args{
				p:    path.Join(cwd, "../testdata/basic/index.html"),
				info: unwrap(os.Stat(path.Join(cwd, "../testdata/basic/index.html"))).(os.FileInfo),
			},
			want: true,
		},
		{
			name:   "HTML file within template directory",
			fields: makeTestBH("basic"),
			args: args{
				p:    path.Join(cwd, "../testdata/basic/.templates/head.html"),
				info: unwrap(os.Stat(path.Join(cwd, "../testdata/basic/.templates/head.html"))).(os.FileInfo),
			},
			want: false,
		},
		{
			name:   "non-HTML file outside template directory",
			fields: makeTestBH("basic"),
			args: args{
				p:    path.Join(cwd, "../testdata/basic/index_meta.json"),
				info: unwrap(os.Stat(path.Join(cwd, "../testdata/basic/index_meta.json"))).(os.FileInfo),
			},
			want: false,
		},
		{
			name:   "non-HTML file within template directory",
			fields: makeTestBH("basic"),
			args: args{
				p:    path.Join(cwd, "../testdata/basic/.templates/README"),
				info: unwrap(os.Stat(path.Join(cwd, "../testdata/basic/.templates/README"))).(os.FileInfo),
			},
			want: false,
		},
		{
			name:   "Directory in root folder",
			fields: makeTestBH("basic"),
			args: args{
				p:    path.Join(cwd, "../testdata/basic/.templates/some-folder"),
				info: unwrap(os.Stat(path.Join(cwd, "../testdata/basic/some-folder"))).(os.FileInfo),
			},
			want: false,
		},
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
	existingDeps := map[string][]string{
		"test1": []string{"file1", "file2"},
		"test2": []string{"file2"},
		"test3": []string{"file1", "file2", "file3"},
	}

	type args struct {
		p         string
		templates []string
	}
	tests := []struct {
		name   string
		fields BHFields
		args   args
		want   map[string][]string
	}{
		{
			name:   "An empty list will add no dependencies",
			fields: makeTestBH("basic"),
			args: args{
				p:         "test-path",
				templates: []string{},
			},
			want: map[string][]string{},
		},
		{
			name:   "A list of a single dependent will add a single dependency",
			fields: makeTestBH("basic"),
			args: args{
				p:         "test-path",
				templates: []string{"test-dep"},
			},
			want: map[string][]string{
				"test-dep": []string{"test-path"},
			},
		},
		{
			name:   "Multiple dependencies will add multiple entries",
			fields: makeTestBH("basic"),
			args: args{
				p:         "test-path",
				templates: []string{"test1", "test2", "test3"},
			},
			want: map[string][]string{
				"test1": []string{"test-path"},
				"test2": []string{"test-path"},
				"test3": []string{"test-path"},
			},
		},
		{
			name: "Pre-existing dependencies will not be duplicated",
			fields: BHFields{
				templates: existingDeps,
			},
			args: args{
				p:         "file1",
				templates: []string{"test1", "test3"},
			},
			want: existingDeps,
		},
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
			bh.saveDependencies(tt.args.p, tt.args.templates...)
			if !mapsAreEqual(bh.templates, tt.want) {
				t.Errorf("Resulting dependencies are not equal. want = %+v, got = %+v", tt.want, bh.templates)
			}
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
