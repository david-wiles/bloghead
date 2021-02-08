package internal

import (
	"github.com/fsnotify/fsnotify"
	"testing"
)

func TestBlogHead_Add(t *testing.T) {
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
		typ  string
		p    string
		name string
	}
	tests := []struct {
		name    string
		fields  fields
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
			if err := bh.Add(tt.args.typ, tt.args.p, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBlogHead_Create(t *testing.T) {
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
		typ  string
		name string
	}
	tests := []struct {
		name    string
		fields  fields
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
			if err := bh.Create(tt.args.typ, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBlogHead_Save(t *testing.T) {
	type fields struct {
		Root       string
		Output     string
		tmplDir    string
		configFile string
		config     *BlogConfig
		templates  map[string][]string
		watcher    *fsnotify.Watcher
	}
	tests := []struct {
		name    string
		fields  fields
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
			if err := bh.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBlogHead_addDefaultMeta(t *testing.T) {
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
		page string
	}
	tests := []struct {
		name    string
		fields  fields
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
			if err := bh.addDefaultMeta(tt.args.page); (err != nil) != tt.wantErr {
				t.Errorf("addDefaultMeta() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBlogHead_addNewArticle(t *testing.T) {
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
		name string
		p    string
	}
	tests := []struct {
		name    string
		fields  fields
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
			if err := bh.addNewArticle(tt.args.name, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("addNewArticle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBlogHead_addNewPage(t *testing.T) {
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
		name string
		p    string
	}
	tests := []struct {
		name    string
		fields  fields
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
			if err := bh.addNewPage(tt.args.name, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("addNewPage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBlogHead_createBlueprint(t *testing.T) {
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
		name string
	}
	tests := []struct {
		name    string
		fields  fields
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
			if err := bh.createBlueprint(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("createBlueprint() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBlogHead_createContentFile(t *testing.T) {
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
			if err := bh.createContentFile(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("createContentFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
