package main

import (
	"fmt"
	"log"
	"net"
)

// Este main aislado demuestra los conceptos del Módulo 1: TCP Foundation.
// Se puede correr usando `go run main.go` dentro de este directorio.
func main() {
	port := ":8080"

	// Abrir socket
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error abriendo socket TCP en %s: %v", port, err)
	}
	defer listener.Close()

	log.Printf("Servidor TCP básico escuchando en %s", port)

	for {
		// Aceptar la conexión bloqueando la goroutine actual hasta que un cliente se conecte
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error en Accept(): %v", err)
			continue
		}

		// Crear una goroutine para concurrencia
		go func(c net.Conn) {
			defer c.Close()
			log.Printf("Nueva conexión desde: %s", c.RemoteAddr().String())

			// Leer bytes crudos
			buf := make([]byte, 2048)
			n, err := c.Read(buf)
			if err != nil {
				log.Printf("Error leyendo: %v", err)
				return
			}

			fmt.Printf("DATA RECIBIDA:\n%s\n", string(buf[:n]))

			// Escribir de vuelta
			c.Write([]byte("HTTP/1.1 200 OK\r\n\r\n¡Hola desde Módulo 1!"))
		}(conn)
	}
}
