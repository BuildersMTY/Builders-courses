package main

import (
	"bufio"
	"fmt"
)

// Request represents a parsed HTTP/1.1 request read from a client socket.
type Request struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    []byte
}

// parseRequestLine reads the first line of the HTTP request from the reader
// (e.g. "GET /path HTTP/1.1\r\n"), trims the trailing \r\n, splits by space
// into exactly 3 parts, and returns (method, path, version).
// Return an error if the line cannot be read or doesn't have exactly 3 parts.
func parseRequestLine(reader *bufio.Reader) (method, path, version string, err error) {
	// TODO: implement
	return "", "", "", nil
}

// parseHeaders reads header lines from the reader until it encounters an
// empty line (the CRLF that separates headers from body). Each header has
// the format "Key: Value\r\n". Use strings.SplitN(line, ":", 2) and
// strings.TrimSpace on both parts. Malformed headers (no ":") should be
// skipped without error — just continue to the next line.
func parseHeaders(reader *bufio.Reader) (map[string]string, error) {
	// TODO: implement
	return map[string]string{}, nil
}

// parseBody checks if headers contains "Content-Length". If present, parse
// the value as int with strconv.Atoi, allocate a byte slice of that size,
// and read exactly that many bytes with io.ReadFull(reader, body).
// If no Content-Length or value is 0, return nil without error.
func parseBody(reader *bufio.Reader, headers map[string]string) ([]byte, error) {
	// TODO: implement
	return nil, nil
}

// ParseRequest orchestrates the three parsing phases.
// This function is pre-implemented — implement the three helpers above.
func ParseRequest(reader *bufio.Reader) (*Request, error) {
	method, path, version, err := parseRequestLine(reader)
	if err != nil {
		return nil, fmt.Errorf("request line: %w", err)
	}
	headers, err := parseHeaders(reader)
	if err != nil {
		return nil, fmt.Errorf("headers: %w", err)
	}
	body, err := parseBody(reader, headers)
	if err != nil {
		return nil, fmt.Errorf("body: %w", err)
	}
	return &Request{
		Method:  method,
		Path:    path,
		Version: version,
		Headers: headers,
		Body:    body,
	}, nil
}
