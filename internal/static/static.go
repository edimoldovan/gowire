package static

import (
	"compress/gzip"
	"embed"
	"fmt"
	"gowire/config"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/andybalholm/brotli"
)

//go:embed assets
var embeddeAssets embed.FS

// this stuff below is needed for gzip compression
type gzipResponseWriter struct {
	http.ResponseWriter
	*gzip.Writer
}

type brotliResponseWriter struct {
	http.ResponseWriter
	*brotli.Writer
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w gzipResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w gzipResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w brotliResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func supportsEncoding(r *http.Request, encoding string) bool {
	return strings.Contains(r.Header.Get("Accept-Encoding"), encoding)
}

func customFileServer(fs http.FileSystem) http.Handler {
	fileServer := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if supportsEncoding(r, "br") && strings.Contains(r.Header.Get("Accept-Encoding"), "br") {
			w.Header().Set("Content-Encoding", "br")
			br := brotli.NewWriter(w)
			defer br.Close()
			w = brotliResponseWriter{Writer: br, ResponseWriter: w}
		} else if filepath.Ext(r.URL.Path) == ".css" && strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			w.Header().Set("Cache-Control", "public, max-age=86400")
			w.Header().Set("Content-Encoding", "gzip")
			gz := gzip.NewWriter(w)
			defer gz.Close()
			w = gzipResponseWriter{Writer: gz, ResponseWriter: w}
		}
		fileServer.ServeHTTP(w, r)
	})
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	file := fmt.Sprintf("internal/server/static%s", r.URL.Path)
	http.ServeFile(w, r, file)
}

func SetupStaticServer(mux *http.ServeMux) http.Handler {
	var stripPrefix string
	var staticServer http.Handler

	if config.IsDevelopment() {
		log.Println("Serving static files from disk")
		stripPrefix = "/assets/"
		mux.HandleFunc("GET /assets/{path...}", serveStatic)
	} else {
		staticServer = customFileServer(http.FS(embeddeAssets))
		stripPrefix = "/"
		mux.HandleFunc("GET /assets/{path...}", http.StripPrefix(stripPrefix, staticServer).ServeHTTP)
	}

	return customFileServer(http.FS(embeddeAssets))
}
