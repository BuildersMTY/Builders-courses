package main

import (
	"log"
	"time"
)

// Middleware es una función que envuelve un Handler para interceptar o modificar requests/responses.
type Middleware func(Handler) Handler

// loggingResponseWriter intercepta ResponseWriter para capturar el HTTP status code (para el Logger).
type loggingResponseWriter struct {
	ResponseWriter
	statusCode int
}

func (w *loggingResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *loggingResponseWriter) Write(b []byte) (int, error) {
	// Si Write se llama sin haber llamado WriteHeader, asumimos HTTP 200 OK internamente
	if w.statusCode == 0 {
		w.statusCode = 200
	}
	return w.ResponseWriter.Write(b)
}

// LoggerMiddleware imprime en consola los detalles de la petición: METHOD PATH STATUS_CODE DURATION
func LoggerMiddleware(next Handler) Handler {
	return HandlerFunc(func(w ResponseWriter, r *Request) {
		start := time.Now()

		// Envolvemos el response writer original para interceptar el status code resultante
		lw := &loggingResponseWriter{
			ResponseWriter: w,
			statusCode:     0,
		}

		// Delegar al siguiente handler
		next.ServeHTTP(lw, r)

		// Una vez completado, registramos el log
		duration := time.Since(start)
		log.Printf("%s %s %d %v", r.Method, r.Path, lw.statusCode, duration)
	})
}

// CORSMiddleware añade cabeceras para permitir peticiones cross-origin (CORS).
func CORSMiddleware(next Handler) Handler {
	return HandlerFunc(func(w ResponseWriter, r *Request) {
		w.SetHeader("Access-Control-Allow-Origin", "*")
		w.SetHeader("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.SetHeader("Access-Control-Allow-Headers", "Content-Type")

		// Manejo rápido para preflight request. No procesamos hacia el controller si es OPTIONS.
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// ApplyMiddlewares encadena una serie de middlewares al Handler base en el orden adecuado.
func ApplyMiddlewares(base Handler, middlewares ...Middleware) Handler {
	// Se aplican en orden inverso para que el primero de la lista envuelva todo y se ejecute primero.
	for i := len(middlewares) - 1; i >= 0; i-- {
		base = middlewares[i](base)
	}
	return base
}
