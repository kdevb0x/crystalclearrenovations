// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package main

import (
	"os"
	"testing"
)

var hostAddr = os.Args[len(os.Args)-1]

func TestHandleRequest(t *testing.T) {
	s := NewSessionServer(hostAddr)
	if s.HostAddr != hostAddr {
		t.Fail()
	}
}
