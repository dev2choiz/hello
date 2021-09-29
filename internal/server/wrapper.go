package server

import (
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"net/http"
	"strings"
)

// GetWrappedServer a grpc server to serve http and grpc
func GetWrappedServer(server *grpc.Server, mux *http.ServeMux, conf *Config) *http.Server {
	wrappedServer := grpcweb.WrapServer(
		server,
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true
		}),
	)

	var handler http.Handler
	handler = &CHandler{
		config: conf,
		wrappedGrpcServer: wrappedServer,
		grpcServer: server,
		mux: mux,
	}

	if !conf.WithTLS {
		handler = h2c.NewHandler(handler, &http2.Server{})
	}
	return &http.Server{Handler: handler}
}

type CHandler struct {
	config            *Config
	wrappedGrpcServer *grpcweb.WrappedGrpcServer
	grpcServer        *grpc.Server
	mux               *http.ServeMux
}

// ServeHTTP handle grpc, web-grpc ant http requests
func (h CHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ct := req.Header.Get("Content-Type")
	if req.ProtoMajor == 2 && strings.Contains(ct, "application/grpc") {
		h.grpcServer.ServeHTTP(resp, req)
	} else if h.wrappedGrpcServer.IsAcceptableGrpcCorsRequest(req) || h.wrappedGrpcServer.IsGrpcWebRequest(req) {
		h.wrappedGrpcServer.ServeHTTP(resp, req)
	} else {
		h.mux.ServeHTTP(resp, req)
	}
}
