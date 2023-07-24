package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)

// Gzip handler for to be used as middleware
type GzipHandler struct {
}

// a wrapper struct that wraps http.ResponseWrite and gzip.Writer
// and combines their core method
type WrapperdResponseWriter struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

func (g *GzipHandler) GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			// create an instance of WrappedResponseWriter
			rw.Header().Set("Content-Encoding", "gzip")
			wrw := NewWrappedResponseWriter(rw)
			next.ServeHTTP(wrw, r)
			defer wrw.Flush()
			return
		}
		next.ServeHTTP(rw, r)
	})
}


func NewWrappedResponseWriter(rw http.ResponseWriter) *WrapperdResponseWriter {
	gw := gzip.NewWriter(rw)
	return &WrapperdResponseWriter{rw: rw, gw: gw}
}


func (wr *WrapperdResponseWriter) Header() http.Header {
	return wr.rw.Header()
}



func (wr *WrapperdResponseWriter) Write(d []byte) (int, error) {
	return wr.gw.Write(d)
}


func (wr *WrapperdResponseWriter) WriteHeader(statusCode int) {
	wr.rw.WriteHeader(statusCode)
}


func (wr *WrapperdResponseWriter) Flush() {
	wr.gw.Flush()
	wr.gw.Close()
}