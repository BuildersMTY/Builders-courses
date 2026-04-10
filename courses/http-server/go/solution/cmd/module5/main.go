package main

import (
	"log"
	"net"
	"time"
)

// Para compilar este demo correr comandos en el root module:
// go run cmd/module5/main.go static.go response.go server.go request.go router.go

func main() {
	port := ":8080"

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer listener.Close()

	// Configuramos el server y router que redirigirá las peticiones
	server := &Server{Addr: port}
	router := NewRouter()

	// Configuramos que lo que no tenga ruta especifica, lo sirva nuestro StaticHandler
	router.Fallback = NewStaticHandler("./static")
	server.Handler = router

	log.Printf("Módulo 5 - Sirviendo archivos estáticos en %s", port)
	log.Println("Prueba con:")
	log.Println("  curl -v http://localhost:8080/")
	log.Println("  curl -v http://localhost:8080/style.css")
	log.Println("  curl -v --path-as-is http://localhost:8080/../../etc/passwd  (Debe dar 403/Forbidden)")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go func(c net.Conn) {
			c.SetDeadline(time.Now().Add(5 * time.Second))
			// El handleConnection inicial del Server llamará al router como delegador
			server.handleConnection(c)
		}(conn)
	}
}
