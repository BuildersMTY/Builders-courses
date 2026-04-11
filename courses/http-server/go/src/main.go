package main

import (
	"log"
)

func main() {
	port := ":8080"

	// 1. Creamos el enrutador principal
	router := NewRouter()

	// 2. Ruta GET /health -> {"status":"ok"}
	router.Handle("GET", "/health", func(w ResponseWriter, r *Request) {
		w.SetHeader("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	// 3. Ruta POST /echo -> retorne el body que recibió
	router.Handle("POST", "/echo", func(w ResponseWriter, r *Request) {
		w.SetHeader("Content-Type", "text/plain")
		w.SetHeader("X-Echo-Server", "BuildersMTY")
		w.Write(r.Body)
	})

	// 4. Configurar el Fallback para servir archivos estáticos desde ./static/
	router.Fallback = NewStaticHandler("./static/")

	// 5. Aplicar la Cadena de Middlewares (Logger + CORS) al Router
	handlerChain := ApplyMiddlewares(router, LoggerMiddleware, CORSMiddleware)

	// 6. Iniciar el Servidor TCP con el Handler configurado
	server := &Server{
		Addr:    port,
		Handler: handlerChain,
	}

	log.Printf("Iniciando HTTP/1.1 Reference Server en %s...", port)
	log.Println("- Directorio Estático Fallback: ./static/")
	log.Println("- Endpoints: GET /health, POST /echo")

	if err := server.Start(); err != nil {
		log.Fatalf("Error crítico del servidor: %v", err)
	}
}
