package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

// Para ejecutar de forma aislada este archivo usando las funciones definidas main:
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

	// ParseRequest está definido en request.go
	req, err := ParseRequest(reader)
	if err != nil {
		log.Fatalf("Error parseando request: %v", err)
	}

	fmt.Println("--- RESULTADO DEL PARSEO ---")
	fmt.Printf("Método:  %s\n", req.Method)
	fmt.Printf("Ruta:    %s\n", req.Path)
	fmt.Printf("Versión: %s\n", req.Version)
	fmt.Println("Headers:")
	for k, v := range req.Headers {
		fmt.Printf("  %s: %s\n", k, v)
	}
	fmt.Printf("Body:    %s\n", string(req.Body))
}
