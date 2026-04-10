package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// HandlerFunc es la firma que deben implementar los middlewares y controladores.
// Nota: en este módulo 1 usamos interfaces vacías o directamente un dummy,
// pero lo prepararemos para no tener que reescribir la estructura más adelante.
// Para el Módulo 1 simplemente leemos bytes crudos.

// Server representa nuestro servidor HTTP/TCP.
type Server struct {
	Addr    string
	Handler Handler // Generalizamos para usar un Router u otro componente HTTP
}

// Start arranca el listener TCP y el loop de aceptación de conexiones.
func (s *Server) Start() error {
	// 1. Abrir socket TCP
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return fmt.Errorf("error al abrir el puerto %s: %w", s.Addr, err)
	}
	defer listener.Close()

	log.Printf("Módulo 1 - Servidor escuchando en %s", s.Addr)

	// 2. Accept loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error aceptando conexión: %v", err)
			continue
		}

		// 3. Manejar conexión de forma concurrente
		go s.handleConnection(conn)
	}
}

// handleConnection lee requests de la conexión tcp.
func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	// Usamos bufio.Reader para parsear el streaming TCP cómodamente
	reader := bufio.NewReader(conn)
	req, err := ParseRequest(reader)
	// En caso de que haya error de parseo, creamos al ResponseWriter temprano por robustez (aunque aquí abortemos)
	if err != nil {
		log.Printf("Error parseando request: %v", err)
		w := NewResponseWriter(conn)
		WriteError(w, 400, "Bad Request")
		return
	}

	w := NewResponseWriter(conn)
	// log.Printf("Nuevo request (Módulo 4 en adelante es manejado en middlewares): %s %s", req.Method, req.Path)

	// Delegar la petición al router / handler enlazado
	if s.Handler != nil {
		s.Handler.ServeHTTP(w, req)
	} else {
		// Log solo por debug en caso de que el sistema esté mal configurado.
		log.Println("WARNING: No handler configurado en Server")
		WriteError(w, 404, "Not Found")
	}
}
