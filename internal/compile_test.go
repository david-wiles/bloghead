package internal

import (
	"github.com/fsnotify/fsnotify"
	"reflect"
	"testing"
)

func TestBlogHead_compile(t *testing.T) {
	type fields struct {
		Root       string
		Output     string
		tmplDir    string
		configFile string
		config     *BlogConfig
		templates  map[string][]string
		watcher    *fsnotify.Watcher
	}
	type args struct {
		p string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
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
			got, err := bh.compile(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("compile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlogHead_gatherTemplates(t *testing.T) {
	type fields struct {
		Root       string
		Output     string
		tmplDir    string
		configFile string
		config     *BlogConfig
		templates  map[string][]string
		watcher    *fsnotify.Watcher
	}
	type args struct {
		p string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
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
			got, err := bh.gatherTemplates(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("gatherTemplates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gatherTemplates() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTemplateData(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getTemplateData(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTemplateData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTemplateData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trimPath(t *testing.T) {
	type args struct {
		base string
		p    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimPath(tt.args.base, tt.args.p); got != tt.want {
				t.Errorf("trimPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
