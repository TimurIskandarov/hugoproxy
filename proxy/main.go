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

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	rp := NewReverseProxy("hugo", "1313")

	r.Use(rp.ReverseProxy)

	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from API"))
	})

	go WorkerTest()

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

const content = `
---
title: Test
---

# Добро пожаловать

Данный сайт создан на основе go-hugo

Здесь нужно будет выполнить несколько задач

tick: %v

meow-meow: %v

`

func WorkerTest() {
	t := time.NewTicker(15 * time.Second)
	var b byte = 0
	for {
		select {
		case tick := <-t.C:
			err := os.WriteFile("/app/static/_index.md", []byte(fmt.Sprintf(content, tick.Format("15:04:05"), b)), 0644)
			if err != nil {
				log.Println(err)
			}
			b++
		}
	}
}
