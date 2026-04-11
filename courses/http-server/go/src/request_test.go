package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestParseRequestLine(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		method  string
		path    string
		version string
		wantErr bool
	}{
		{
			name:    "simple GET",
			input:   "GET /hello HTTP/1.1\r\n",
			method:  "GET",
			path:    "/hello",
			version: "HTTP/1.1",
		},
		{
			name:    "POST with path",
			input:   "POST /echo HTTP/1.1\r\n",
			method:  "POST",
			path:    "/echo",
			version: "HTTP/1.1",
		},
		{
			name:    "root path",
			input:   "GET / HTTP/1.1\r\n",
			method:  "GET",
			path:    "/",
			version: "HTTP/1.1",
		},
		{
			name:    "malformed line",
			input:   "INVALID\r\n",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.input))
			method, path, version, err := parseRequestLine(reader)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if method != tt.method {
				t.Errorf("method = %q, want %q", method, tt.method)
			}
			if path != tt.path {
				t.Errorf("path = %q, want %q", path, tt.path)
			}
			if version != tt.version {
				t.Errorf("version = %q, want %q", version, tt.version)
			}
		})
	}
}

func TestParseHeaders(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    map[string]string
		wantErr bool
	}{
		{
			name:  "two headers",
			input: "Host: localhost\r\nContent-Type: text/plain\r\n\r\n",
			want:  map[string]string{"Host": "localhost", "Content-Type": "text/plain"},
		},
		{
			name:  "no headers",
			input: "\r\n",
			want:  map[string]string{},
		},
		{
			name:  "malformed header skipped",
			input: "Host: localhost\r\nBADHEADER\r\nAccept: */*\r\n\r\n",
			want:  map[string]string{"Host": "localhost", "Accept": "*/*"},
		},
		{
			name:  "header with colon in value",
			input: "Location: http://example.com:8080\r\n\r\n",
			want:  map[string]string{"Location": "http://example.com:8080"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.input))
			headers, err := parseHeaders(reader)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			for k, v := range tt.want {
				if headers[k] != v {
					t.Errorf("headers[%q] = %q, want %q", k, headers[k], v)
				}
			}
			if len(headers) != len(tt.want) {
				t.Errorf("got %d headers, want %d", len(headers), len(tt.want))
			}
		})
	}
}

func TestParseBody(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		headers map[string]string
		want    string
		wantNil bool
	}{
		{
			name:    "with content-length",
			input:   "hello",
			headers: map[string]string{"Content-Length": "5"},
			want:    "hello",
		},
		{
			name:    "no content-length",
			input:   "",
			headers: map[string]string{},
			wantNil: true,
		},
		{
			name:    "content-length zero",
			input:   "",
			headers: map[string]string{"Content-Length": "0"},
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.input))
			body, err := parseBody(reader, tt.headers)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tt.wantNil {
				if body != nil {
					t.Errorf("body = %q, want nil", body)
				}
				return
			}
			if string(body) != tt.want {
				t.Errorf("body = %q, want %q", string(body), tt.want)
			}
		})
	}
}

func TestParseRequest(t *testing.T) {
	input := "POST /echo HTTP/1.1\r\nHost: localhost\r\nContent-Length: 5\r\n\r\nhello"
	reader := bufio.NewReader(strings.NewReader(input))
	req, err := ParseRequest(reader)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if req.Method != "POST" {
		t.Errorf("method = %q, want POST", req.Method)
	}
	if req.Path != "/echo" {
		t.Errorf("path = %q, want /echo", req.Path)
	}
	if req.Version != "HTTP/1.1" {
		t.Errorf("version = %q, want HTTP/1.1", req.Version)
	}
	if req.Headers["Host"] != "localhost" {
		t.Errorf("Host header = %q, want localhost", req.Headers["Host"])
	}
	if string(req.Body) != "hello" {
		t.Errorf("body = %q, want hello", string(req.Body))
	}
}
