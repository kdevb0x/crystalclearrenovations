// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
)

var ErrNotFound = errors.New("error: key not found")

// Project represents model for completed projects in the database.
type PastProjects struct {
	WorkType string   `db:"type_of_work"` // house, bussiness, flooring ect.
	Comments []string `db:"comments"`
	Images   []string `db:"images"` // filenames of images
}

func LoadTemplateSet(filename string) {
}

type TemplateSet struct {
	Loader TemplateLoader
}

// Context is a map holding usable tagsets.
type context map[string]interface{}

type Value interface{}

type TemplateLoader interface {
	Abs(base, name string) string
	Get(path string) (io.Reader, error)
}

type Context interface {
	Value(key string) interface{}
	Keys() []string
	DumpAll() []struct {
		key   string
		value Value
	}
	AddTag(key string, value interface{})
}

func (c *context) Value(key string) Value {
	if val, ok := *c[key]; ok {
		return val
	}
	return ErrNotFound
}

func (c *context) Keys() []string {
	var keys []string
	for k, _ := range *c {
		keys = append(keys, k)
	}
	return keys
}

func (c *context) DumpAll() []struct {
	key   string
	value Value
} {
	type tag struct {
		key   string
		value Value
	}
	var tags = make([]tag, len(*c))
	for k, v := range *c {
		tags = append(tags, tag{k, v})
	}
	return tags
}
func (l *TemplateSet) Abs(base, name string) string {
	return filepath.Join(base, name)
}

func (l *TemplateSet) Get(path string) (io.Reader, error) {
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

func (l *TemplateSet) AddTag(name string, action interface{}) error {
	if _, ok := l.context[name]; !ok {
		l.context[name] = action
		return nil
	}
	return fmt.Errorf("error: cant add %s, tag already exists by that name", name)
}
