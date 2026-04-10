package main

// Middleware wraps a Handler to intercept or modify requests/responses.
type Middleware func(Handler) Handler

// loggingResponseWriter decorates a ResponseWriter to capture the final
// status code so LoggerMiddleware can print it after the chain completes.
type loggingResponseWriter struct {
	ResponseWriter
	statusCode int
}

// WriteHeader records the status code and delegates to the wrapped writer.
func (w *loggingResponseWriter) WriteHeader(code int) {
	// TODO: implement
	w.ResponseWriter.WriteHeader(code)
}

// LoggerMiddleware logs METHOD PATH STATUS DURATION after the downstream
// handler returns. Use a loggingResponseWriter to observe the status code.
func LoggerMiddleware(next Handler) Handler {
	return HandlerFunc(func(w ResponseWriter, r *Request) {
		// TODO: implement
		next.ServeHTTP(w, r)
	})
}

// CORSMiddleware sets Access-Control-* headers for every response and
// short-circuits OPTIONS preflights with an immediate 200 (no downstream
// call).
func CORSMiddleware(next Handler) Handler {
	return HandlerFunc(func(w ResponseWriter, r *Request) {
		// TODO: implement
		next.ServeHTTP(w, r)
	})
}

// ApplyMiddlewares chains middlewares around base. The first middleware
// in the list must be the outermost one (executes first on the way in,
// last on the way out). Tip: iterate the slice backwards.
func ApplyMiddlewares(base Handler, middlewares ...Middleware) Handler {
	// TODO: implement
	return base
}
