package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"test/geo"
	// "test/worker"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	rp := NewReverseProxy("hugo", "1313")

	r.Use(rp.ReverseProxy)

	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from API"))
	})

	geoService := geo.New()
	r.Post("/api/address/search", geoService.SearchHandler)
	r.Post("/api/address/geocode", geoService.GeocodeHandler)

	// go worker.Tasks()

	http.ListenAndServe(":8080", r)
}

type ReverseProxy struct {
	host string
	port string
}

func NewReverseProxy(host, port string) *ReverseProxy {
	return &ReverseProxy{
		host: host,
		port: port,
	}
}

func (rp *ReverseProxy) ReverseProxy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			next.ServeHTTP(w, r)
			return
		}
		link := fmt.Sprintf("http://%s:%s", rp.host, rp.port)
		uri, _ := url.Parse(link)

		if uri.Host == r.Host {
			next.ServeHTTP(w, r)
			return
		}
		r.Header.Set("Reverse-Proxy", "true")

		proxy := &httputil.ReverseProxy{
			Rewrite: func(req *httputil.ProxyRequest) {
				req.SetURL(uri)
				req.Out.URL.Scheme = uri.Scheme // http
				req.Out.URL.Host = uri.Host     // hugo:1313
				req.Out.URL.Path = r.URL.Path   // /{endpoint}
				req.Out.Host = uri.Host         // hugo:1313

				// редирект https://stackoverflow.com/questions/45869688/redirects-return-http-multiple-response-writeheader-calls
				// rh := http.RedirectHandler(link + r.URL.Path, http.StatusMovedPermanently)
				// rh.ServeHTTP(w, r)
			},
		}
		proxy.ServeHTTP(w, r)
	})
}
