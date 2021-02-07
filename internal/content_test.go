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
		name string
		path string
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
			if err := bh.Add(tt.args.typ, tt.args.name, tt.args.path); (err != nil) != tt.wantErr {
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
