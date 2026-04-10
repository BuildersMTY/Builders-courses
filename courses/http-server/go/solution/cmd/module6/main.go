package main

import (
	"log"
	"net"
	"time"
)

// main corre el demo aislado de Middlewares (Módulo 6).
// Para compilar y correr (desde root del modulo):
// go run cmd/module6/main.go handler.go response.go server.go request.go router.go static.go

func main() {
	port := ":8080"

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer listener.Close()

	// 1. Configuramos el Componente Principal (ej. Router)
	router := NewRouter()
	router.Handle("GET", "/middleware-test", func(w ResponseWriter, r *Request) {
		time.Sleep(100 * time.Millisecond) // Simula procesamiento
		w.SetHeader("Content-Type", "text/plain")
		w.Write([]byte("Middleware y Router Funcionando Juntos!"))
	})

	// 2. Encadenamos middlewares
	// La petición pasará primero al Logger -> CORS -> Router
	handler := ApplyMiddlewares(router, LoggerMiddleware, CORSMiddleware)

	// 3. Montar el server con la cadena de handler resultante
	server := &Server{
		Addr:    port,
		Handler: handler,
	}

	log.Printf("Módulo 6 - Testeando Middleware Chain en %s", port)
	log.Println("Prueba a disparar request para ver logs de métricas y CORS Headers:")
	log.Println("  curl -v http://localhost:8080/middleware-test")
	log.Println("  curl -v -X OPTIONS http://localhost:8080/middleware-test")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go func(c net.Conn) {
			c.SetDeadline(time.Now().Add(5 * time.Second))
			server.handleConnection(c)
		}(conn)
	}
}
