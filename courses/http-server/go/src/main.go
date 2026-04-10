package main

import "log"

// main wires the server with routes, static fallback and middlewares.
// You don't need to modify this file — just implement the stubs in the other
// .go files and run `go run .`
func main() {
	port := ":8080"

	router := NewRouter()

	router.Handle("GET", "/health", func(w ResponseWriter, r *Request) {
		w.SetHeader("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	router.Handle("POST", "/echo", func(w ResponseWriter, r *Request) {
		w.SetHeader("Content-Type", "text/plain")
		w.Write(r.Body)
	})

	router.Fallback = NewStaticHandler("./static/")

	handlerChain := ApplyMiddlewares(router, LoggerMiddleware, CORSMiddleware)

	server := &Server{
		Addr:    port,
		Handler: handlerChain,
	}

	log.Printf("HTTP server listening on %s", port)
	if err := server.Start(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
