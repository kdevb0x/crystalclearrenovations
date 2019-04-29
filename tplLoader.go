// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

// path to render templates for server
const renderPath = "./assets/static/"

type TemplateLoader interface {
	Abs(base, name string) string
	Get(path string) (io.Reader, error)
}

type TemplateHandler struct {
	Dir string
}

func (t *TemplateHandler) Abs(base, name string) string {
	return filepath.Join(base, name)
}

func (t *TemplateHandler) Get(path string) (io.Reader, error) {
	b, err := t.FetchTemplateBytes(nil, path)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(b)
	return r, nil
}

// TODO: Figure out a better interplay between this and TemplateLoader interface.
func (t *TemplateHandler) FetchTemplateBytes(loader TemplateLoader, filename string) ([]byte, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return f, nil

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// TODO: Finish this; Load and parse templates, and render to w.
}
