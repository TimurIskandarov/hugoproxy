package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"

	"test/binary"
	"test/counter"
	"test/geo"
	"test/graph"

	"github.com/go-chi/chi"
	"github.com/mohae/deepcopy"
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
	// r.Post("/api/address/geocode", geoService.GeocodeHandler)

	// go WorkerTest()

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

func WorkerTest() {
	t := time.NewTicker(5 * time.Second)
	var b byte = 0

	var totalNodes int = 4
	avl := binary.GenerateTree(totalNodes)
	bin := deepcopy.Copy(avl).(*binary.AVLTree)

	for {
		select {
		case tick := <-t.C:
			contentCounter := counter.GetCounterPage(tick, b)
			err := os.WriteFile("/app/static/tasks/_index.md", []byte(contentCounter), 0644)
			if err != nil {
				log.Println(err)
			}
			b++

			binary.SetRandomNode(avl, bin)
			contentAVL := binary.GetAVLPage(avl, bin)
			err = os.WriteFile("/app/static/tasks/binary.md", []byte(contentAVL), 0644)
			if err != nil {
				log.Println(err)
			}

			totalNodes++
			if totalNodes == 100 {
				totalNodes = 4
			}

			contentGraph := graph.GetGraphPage()
			err = os.WriteFile("/app/static/tasks/graph.md", []byte(contentGraph), 0644)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
