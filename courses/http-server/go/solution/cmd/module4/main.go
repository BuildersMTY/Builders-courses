package main

import (
	"fmt"
	"log"
)

// Para compilar este demo, se tienen que incluir los otros archivos del paquete main:
// go run cmd/module4/main.go router.go response.go request.go server.go

// mockResponseWriter permite probar nuestro router sin levantar un socket red.
type mockResponseWriter struct {
	body   []byte
	status int
}

func (m *mockResponseWriter) Write(b []byte) (int, error) {
	m.body = append(m.body, b...)
	return len(b), nil
}
func (m *mockResponseWriter) WriteHeader(status int) {
	m.status = status
}
func (m *mockResponseWriter) SetHeader(k, v string) {}

func main() {
	log.Println("Inciando Test Aislado - Módulo 4: Router")

	router := NewRouter()

	// Simulamos registro de ruta
	router.Handle("GET", "/ping", func(w ResponseWriter, r *Request) {
		w.WriteHeader(200)
		w.SetHeader("Content-Type", "text/plain")
		w.Write([]byte("pong!"))
	})

	log.Println("=> Enviando llamada artificial: GET /ping")
	r1 := &Request{Method: "GET", Path: "/ping"}
	w1 := &mockResponseWriter{}
	router.ServeHTTP(w1, r1)
	fmt.Printf("   Resultado -> Status: %d, Body: %s\n\n", w1.status, string(w1.body))

	log.Println("=> Enviando llamada artificial: POST /ping (Debe retornar 405)")
	r2 := &Request{Method: "POST", Path: "/ping"}
	w2 := &mockResponseWriter{}
	router.ServeHTTP(w2, r2)
	fmt.Printf("   Resultado -> Status: %d, Body: %s\n\n", w2.status, string(w2.body))

	log.Println("=> Enviando llamada artificial: GET /desconocido (Debe retornar 404)")
	r3 := &Request{Method: "GET", Path: "/desconocido"}
	w3 := &mockResponseWriter{}
	router.ServeHTTP(w3, r3)
	fmt.Printf("   Resultado -> Status: %d, Body: %s\n", w3.status, string(w3.body))
}
