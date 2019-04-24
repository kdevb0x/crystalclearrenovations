// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package templates

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"

	"github.com/flosch/pongo2"
)

// Project represents model for completed projects in the database.
type PastProjects struct {
	WorkType string   `db:"type_of_work"` // house, bussiness, flooring ect.
	Comments []string `db:"comments"`
	Images   []string `db:"images"` // filenames of images
}

func LoadTemplateSet(filename string) *pongo2.TemplateSet {
	var t = &TemplateLoader{Context: make(pongo2.Context)}
	return pongo2.NewSet("index", t)
}

type TemplateLoader struct {
	Context pongo2.Context
}

func (l *templateLoader) Abs(base, name string) string {
	return filepath.Join(base, name)
}

func (l *templateLoader) Get(path string) (io.Reader, error) {
	return getTemplate(path)
}

func getTemplate(filepath string) (io.Reader, error) {
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(f)
	return r, nil
}

func (l *templateLoader) AddTag(name string, action interface{}) error {
	if _, ok := l.context[name]; !ok {
		l.context[name] = action
		return nil
	}
	return fmt.Errorf("error: cant add %s, tag already exists by that name", name)
}
