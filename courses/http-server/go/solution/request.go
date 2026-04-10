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

// ParseRequest lee de un bufio.Reader y construye un objeto Request.
// Parsea la Request Line, los Headers y el Body (si existe Content-Length).
func ParseRequest(reader *bufio.Reader) (*Request, error) {
	req := &Request{
		Headers: make(map[string]string),
	}

	// 1. Request Line (ej: GET / HTTP/1.1)
	reqLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("error leyendo request line: %w", err)
	}
	reqLine = strings.TrimRight(reqLine, "\r\n")

	parts := strings.Split(reqLine, " ")
	if len(parts) != 3 {
		return nil, fmt.Errorf("formato inválido de request line: %q", reqLine)
	}
	req.Method = parts[0]
	req.Path = parts[1]
	req.Version = parts[2]

	// 2. Headers
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
		req.Headers[key] = val
	}

	// 3. Body
	// Revisamos si existe el header Content-Length para saber cuántos bytes leer
	if clStr, ok := req.Headers["Content-Length"]; ok {
		cl, err := strconv.Atoi(clStr)
		if err == nil && cl > 0 {
			req.Body = make([]byte, cl)
			// io.ReadFull asegura que leemos exactamente 'cl' bytes
			_, err = io.ReadFull(reader, req.Body)
			if err != nil {
				return nil, fmt.Errorf("error leyendo el body: %w", err)
			}
		}
	}

	return req, nil
}
