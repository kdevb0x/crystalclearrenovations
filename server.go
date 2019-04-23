// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

// +build GODEBUG=tls13=1

package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type session struct {
	request        *http.Request // current request
	responseWriter http.ResponseWriter
}

type Server struct {
	http.Server
	ConnListener http.ResponseWriter
	KillChan     chan struct{} // tells server to drop connections and shutdown
}

func InitServer(addr string) *Server {
	s := &Server{KillChan: make(chan struct{}, 1)}
	r := buildRoutes()
	s.Handler = r
	s.Addr = addr
	s.WriteTimeout = 15 * time.Second
	s.ReadTimeout = 15 * time.Second
	return s
}

func buildRoutes() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	http.Handle("/", r)
	return r

}

func (s *Server) Start(rwTimeouts ...time.Duration) <-chan error {
	var errChan = make(chan error, 1)

	// haltfunc listens for the signal to shutdown server
	haltfunc := func() {
		for {
			select {
			case <-s.KillChan:
				err := s.Shutdown(nil)
				if err != nil {
					errChan <- err
				}
			default:
				continue
			}
		}
	}

	// serveFunc enters the mainloop
	var serveFunc = func() {
		errChan <- s.ListenAndServe()
	}
	if rwTimeouts != nil {
		s.ReadTimeout = rwTimeouts[0]
		s.WriteTimeout = rwTimeouts[1]
	}

	go haltfunc()
	go serveFunc()

	return errChan
}
