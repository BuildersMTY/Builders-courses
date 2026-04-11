package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Request representa una petición HTTP leída desde el cliente.
type Request struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    []byte
}

// parseRequestLine lee la primera línea del request HTTP (ej: "GET / HTTP/1.1\r\n"),
// la recorta y divide en 3 partes: method, path, version.
func parseRequestLine(reader *bufio.Reader) (method, path, version string, err error) {
	reqLine, err := reader.ReadString('\n')
	if err != nil {
		return "", "", "", fmt.Errorf("error leyendo request line: %w", err)
	}
	reqLine = strings.TrimRight(reqLine, "\r\n")

	parts := strings.Split(reqLine, " ")
	if len(parts) != 3 {
		return "", "", "", fmt.Errorf("formato inválido de request line: %q", reqLine)
	}
	return parts[0], parts[1], parts[2], nil
}

// parseHeaders lee las líneas de headers hasta encontrar una línea vacía (CRLF).
// Cada header tiene formato "Key: Value". Headers malformados se ignoran.
func parseHeaders(reader *bufio.Reader) (map[string]string, error) {
	headers := make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("error leyendo headers: %w", err)
		}
		line = strings.TrimRight(line, "\r\n")

		// Una línea vacía indica el final de los headers (\r\n\r\n)
		if line == "" {
			break
		}

		// Los headers tienen el formato "Key: Value"
		headerParts := strings.SplitN(line, ":", 2)
		if len(headerParts) != 2 {
			continue // Header mal formado, se ignora por robustez
		}

		key := strings.TrimSpace(headerParts[0])
		val := strings.TrimSpace(headerParts[1])
		headers[key] = val
	}
	return headers, nil
}

// parseBody lee el body del request si existe Content-Length en los headers.
// Usa io.ReadFull para leer exactamente esos bytes.
func parseBody(reader *bufio.Reader, headers map[string]string) ([]byte, error) {
	clStr, ok := headers["Content-Length"]
	if !ok {
		return nil, nil
	}
	cl, err := strconv.Atoi(clStr)
	if err != nil || cl <= 0 {
		return nil, nil
	}
	body := make([]byte, cl)
	_, err = io.ReadFull(reader, body)
	if err != nil {
		return nil, fmt.Errorf("error leyendo el body: %w", err)
	}
	return body, nil
}

// ParseRequest orquesta las tres fases de parseo.
// Esta función ya está implementada — implementa los tres helpers de arriba.
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
