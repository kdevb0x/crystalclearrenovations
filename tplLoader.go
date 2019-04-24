// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package main

import (
	"io"
	"io/ioutil"
	"path/filepath"
)

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
	return getTemplate(path)
}

func (t *TemplateHandler) FetchTemplate(filename string) ([]byte, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return f, nil

}
