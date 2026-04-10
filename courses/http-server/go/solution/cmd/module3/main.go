package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// Este main aislado demuestra el manejo de ResponseWriter.
// Ejecutar: go run cmd/module3/main.go response.go

func main() {
	port := ":8080"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer listener.Close()

	log.Printf("Módulo 3 - Respondiendo dinámicamente en %s", port)
	log.Println("Prueba con:")
	log.Println("  curl -v http://localhost:8080/")
	log.Println("  curl -v http://localhost:8080/404")
	log.Println("  curl -v -X POST http://localhost:8080/")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go func(c net.Conn) {
			defer c.Close()
			c.SetDeadline(time.Now().Add(5 * time.Second))

			// Instanciamos el escritor
			w := NewResponseWriter(c)

			buf := make([]byte, 1024)
			n, err := c.Read(buf)
			if err != nil || n == 0 {
				WriteError(w, 400, "Bad Request")
				return
			}
			req := string(buf[:n])

			// Demostración condicional según la ruta cruda
			if strings.HasPrefix(req, "GET /404 ") {
				WriteError(w, 404, "Página no encontrada")
				return
			}
			if strings.HasPrefix(req, "GET /500 ") {
				WriteError(w, 500, "Error interno")
				return
			}
			if strings.HasPrefix(req, "POST ") {
				WriteError(w, 405, "Método no permitido (sólo GET)")
				return
			}

			// 200 OK dinámico automático
			w.SetHeader("Content-Type", "text/html")
			html := "<h1>¡Respuesta generada dinámicamente (Módulo 3)!</h1>"
			w.Write([]byte(html))

			fmt.Println("Respuesta enviada con éxito.")
		}(conn)
	}
}
