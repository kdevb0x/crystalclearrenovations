// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package main

import (
	"net/http"
)

type session struct {
	request        *http.Request // current request
	responseWriter http.ResponseWriter
}

type Server struct {
	Addr         string
	ConnListener http.ResponseWriter
}

func InitServer(addr string) *Server {
	l := NewConnListe
}
