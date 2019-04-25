// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type ErrMalformedOrNilRequest struct {
	s string
}

func (e *ErrMalformedOrNilRequest) Error() string {
	e.s = string("error: malformed or nil http request")
	return e.s
}

type SessionServer struct {
	HostAddr   string
	Router     *mux.Router
	ReqHandler RequestHandler
	RespWriter http.ResponseWriter
}

type RequestHandler interface {
	HandleRequest(r *http.Request) error
}

func NewSessionServer(hostAddr string) *SessionServer {
	return &SessionServer{
		HostAddr: hostAddr,
		Router:   mux.NewRouter(),
		// ReqHandler and RespWriter are nil at this point as they
		// change on per-connection basis (we have no connection yet).
	}
}

func (s *SessionServer) HandleRequest(r *http.Request) error {
	if r == nil {
		return ErrMalformedOrNilRequest{}
	}
	return nil
}
