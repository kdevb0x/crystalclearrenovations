// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package main

import (
	"bufio"
	"errors"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const ErrMalformedOrNilRequest = localErr("error: malformed or nil http request")

// var ErrMalformedOrNilRequest = errors.New("error: malformed or nil http request")
type localErr string

func (l localErr) Error() string {
	return string(l)
}

type SessionServer struct {
	Server      *http.Server
	HostAddr    string
	Router      *mux.Router
	ReqHandler  RequestHandler
	RespWriter  http.ResponseWriter
	lastRequest *struct {
		timeRecvd int64 // time.Time.UnixNano()
		body      []byte
	}
}

type RequestHandler interface {
	HandleRequest(r *http.Request) error
}

type RequestReader interface {
	ReadRequest(from *net.Conn) (*bufio.Reader, error)
}

func NewSessionServer(hostAddr string, altmux *mux.Router) *SessionServer {
	s := &SessionServer{
		Server: &http.Server{},
		// HostAddr: hostAddr,

		// ReqHandler and RespWriter are nil at this point as they
		// change on per-connection basis (we have no connection yet).
	}
	s.Server.Addr = hostAddr
	if altmux != nil {
		s.Router = altmux
		return s
	}
	s.Router = mux.NewRouter()
	return s

}

func (s *SessionServer) AddRoute(url string, handler http.Handler) error {
	if handler != nil {
		url = strings.ToLower(url)
		if strings.Contains(strings.ToLower(url), "://") {
			url = strings.TrimLeft(url, "://")
		}

		s.Router.Handle(url, handler)
		return nil
	}
	return errors.New("ERROR: cannot add route REASON: nil handler")
}

func (s *SessionServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ListenAndServe(s.Server.Addr, s.Router)
}

func (s *SessionServer) HandleRequest(r *http.Request) error {
	return nil
}

/*
// Original Version

func (s *SessionServer) HandleRequest(r *http.Request) error {
	if r == nil {
		return ErrMalformedOrNilRequest
	}
	body, err := r.GetBody()
	if err != nil {
		return err
	}
	defer body.Close()
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	type lastrequest struct {
		timeRecvd int64
		body      []byte
	}

	last := &lastrequest{
		timeRecvd: time.Now().UnixNano(),
		body:      bodyBytes,
	}
	s.lastRequest = (*struct {
		timeRecvd int64
		body      []byte
	})(last)
	return nil
}
*/
