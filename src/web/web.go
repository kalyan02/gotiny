package web

import (
	"fmt"
	"net/http"
)

type Connection struct {
	Response http.ResponseWriter
	Request  *http.Request
}

func (connection *Connection) Write(content string) {
	fmt.Fprintf(connection.Response, content)
}

type WebServer struct {
	server *http.Server
	mux    *http.ServeMux
}

func (webserver *WebServer) Route(path string, handler func(connection *Connection)) {
	webserver.mux.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		handler(&Connection{Request: request, Response: writer})
	})
}

func (webserver *WebServer) Start() {
	webserver.server.ListenAndServe()
}

func NewWebServer() *WebServer {
	s := &WebServer{}
	s.server = &http.Server{
		Addr: ":8080",
	}
	s.mux = http.DefaultServeMux
	handler := new(http.Handler)
	*handler = s.mux
	s.server.Handler = *handler
	return s
}