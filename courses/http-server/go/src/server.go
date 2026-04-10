package main

import (
	"bufio"
	"net"
)

// Server represents your HTTP/TCP server.
// Addr is the listen address (e.g. ":8080").
// Handler is the component that will handle each parsed request.
type Server struct {
	Addr    string
	Handler Handler
}

// Start opens a TCP listener on Server.Addr and runs the accept loop.
// For each accepted connection, spawn a goroutine that calls
// s.handleConnection(conn). If Accept returns an error for a single
// connection, log it and continue — do NOT kill the whole server.
//
// Return a non-nil error only if the listener itself cannot be opened.
func (s *Server) Start() error {
	// TODO: implement
	return nil
}

// handleConnection reads a request from the connection, parses it with
// ParseRequest, wraps the connection in a ResponseWriter, and delegates to
// s.Handler.ServeHTTP. On parse error, respond with 400 Bad Request.
// Always close the connection when done.
func (s *Server) handleConnection(conn net.Conn) {
	// TODO: implement
	_ = bufio.NewReader(conn)
}
