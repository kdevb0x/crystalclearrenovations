// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package main

import (
	"net/http"
	"os"
	"testing"

	client "github.com/gorilla/http"
)

func TestServerInit(t *testing.T) {
	s := InitServer(":8000")
	defer func() {
		s.KillChan <- struct{}{}
	}()

	go func() {
		for err := range s.Start() {
			if err != nil {
				if err != http.ErrServerClosed {
					t.Logf("error: %s", err)
					t.Fail()
				}
			}
		}
	}()

	status, err := client.Get(os.Stdout, "localhost:8000")
	if err != nil {
		t.Logf("status: %d, error: %s", status, err)
		t.Fail()
	}

}
