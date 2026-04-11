package main

import "testing"

func TestApplyMiddlewares(t *testing.T) {
	var order []string

	mw1 := func(next Handler) Handler {
		return HandlerFunc(func(w ResponseWriter, r *Request) {
			order = append(order, "mw1-before")
			next.ServeHTTP(w, r)
			order = append(order, "mw1-after")
		})
	}
	mw2 := func(next Handler) Handler {
		return HandlerFunc(func(w ResponseWriter, r *Request) {
			order = append(order, "mw2-before")
			next.ServeHTTP(w, r)
			order = append(order, "mw2-after")
		})
	}
	base := HandlerFunc(func(w ResponseWriter, r *Request) {
		order = append(order, "base")
	})

	chain := ApplyMiddlewares(base, mw1, mw2)

	// Use a mock writer that satisfies ResponseWriter
	mock := &mockTestWriter{}
	chain.ServeHTTP(mock, &Request{})

	expected := []string{"mw1-before", "mw2-before", "base", "mw2-after", "mw1-after"}
	if len(order) != len(expected) {
		t.Fatalf("got %d calls %v, want %d %v", len(order), order, len(expected), expected)
	}
	for i, v := range expected {
		if order[i] != v {
			t.Errorf("order[%d] = %q, want %q", i, order[i], v)
		}
	}
}

// mockTestWriter satisfies the ResponseWriter interface for unit testing.
type mockTestWriter struct{}

func (m *mockTestWriter) Write(b []byte) (int, error)    { return len(b), nil }
func (m *mockTestWriter) WriteHeader(statusCode int)      {}
func (m *mockTestWriter) SetHeader(key, value string)     {}
