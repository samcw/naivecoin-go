package http

import "net/http"

type HttpServer struct {
}

func handleBlocks(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

}

func (hs *HttpServer) initHttpServer(port uint16) {
	http.HandleFunc("/blocks", handleBlocks)
}
