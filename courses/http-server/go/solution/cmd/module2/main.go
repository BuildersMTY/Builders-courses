package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

// Para ejecutar de forma aislada este archivo usando las funciones definidas:
// go run cmd/module2/main.go request.go

func main() {
	// Simulamos la data que llegaría por el socket TCP en una petición POST real
	rawRequest := "POST /echo HTTP/1.1\r\n" +
		"Host: localhost\r\n" +
		"User-Agent: test-agent\r\n" +
		"Content-Length: 11\r\n" +
		"\r\n" +
		"hello world"

	// Usamos un bufio.Reader para simular la lectura de red
	reader := bufio.NewReader(strings.NewReader(rawRequest))

	log.Printf("Iniciando parseo de la trama...\n")

	// Paso 1: Parsear la request line
	method, path, version, err := parseRequestLine(reader)
	if err != nil {
		log.Fatalf("Error parseando request line: %v", err)
	}
	fmt.Printf("Request Line: %s %s %s\n", method, path, version)

	// Paso 2: Parsear los headers
	headers, err := parseHeaders(reader)
	if err != nil {
		log.Fatalf("Error parseando headers: %v", err)
	}
	fmt.Println("Headers:")
	for k, v := range headers {
		fmt.Printf("  %s: %s\n", k, v)
	}

	// Paso 3: Parsear el body
	body, err := parseBody(reader, headers)
	if err != nil {
		log.Fatalf("Error parseando body: %v", err)
	}
	fmt.Printf("Body: %s\n", string(body))

	// También se puede usar ParseRequest (orquestador) directamente:
	fmt.Println("\n--- ParseRequest completo ---")
	reader2 := bufio.NewReader(strings.NewReader(rawRequest))
	req, err := ParseRequest(reader2)
	if err != nil {
		log.Fatalf("Error con ParseRequest: %v", err)
	}
	fmt.Printf("Método: %s, Ruta: %s, Body: %s\n", req.Method, req.Path, string(req.Body))
}
