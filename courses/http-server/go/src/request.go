package main

import "bufio"

// Request represents a parsed HTTP/1.1 request read from a client socket.
type Request struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    []byte
}

// ParseRequest reads an HTTP/1.1 request from the buffered reader and
// returns a *Request. It must parse:
//   - The request line (e.g. "GET / HTTP/1.1")
//   - All headers until an empty line (CRLF CRLF)
//   - The body, only if Content-Length is set and > 0, reading exactly that
//     many bytes with io.ReadFull.
//
// Malformed headers should be ignored (skip the line, don't crash).
// Return an error only on unrecoverable I/O or malformed request line.
func ParseRequest(reader *bufio.Reader) (*Request, error) {
	// TODO: implement
	return &Request{Headers: map[string]string{}}, nil
}
