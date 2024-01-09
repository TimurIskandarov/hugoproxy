package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewReverseProxy(t *testing.T) {
	type args struct {
		host string
		port string
	}
	tests := []struct {
		name string
		args args
		want *ReverseProxy
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				host: "hugo",
				port: "1313",
			},
			want: &ReverseProxy{"hugo", "1313"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReverseProxy(tt.args.host, tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReverseProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseProxy_ReverseProxy(t *testing.T) {
	type fields struct {
		host string
		port string
	}
	type args struct {
		next http.Handler
		w    *httptest.ResponseRecorder
		r    *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{
			name: "NotFound",
			fields: fields{
				host: "localhost",
				port: "1313",
			},
			args: args{
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("web page http://localhost:1313/test non-existent"))
				}),
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/test", nil),
			},
			want: http.StatusNotFound,
		},
		{
			name: "Redirect",
			fields: fields{
				host: "localhost",
				port: "1313",
			},
			args: args{
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("redirect to web page http://localhost:1313/tasks"))
				}),
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/tasks", nil),
			},
			want: http.StatusMovedPermanently,
		},
		{
			name: "StatusOK",
			fields: fields{
				host: "localhost",
				port: "1313",
			},
			args: args{
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("web page http://localhost:8080/api existent"))
				}),
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/api", nil),
			},
			want: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rp := &ReverseProxy{
				host: tt.fields.host,
				port: tt.fields.port,
			}
			handler := rp.ReverseProxy(tt.args.next)
			handler.ServeHTTP(tt.args.w, tt.args.r)
			if tt.args.w.Code != tt.want {
				t.Errorf("ReverseProxy.ReverseProxy() = %v, want %v", tt.args.w.Code, tt.want)
			}
		})
	}
}
