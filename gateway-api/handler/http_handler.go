package handler

import (
	"encoding/json"
	"net/http"
)

type GatewayApiHandler struct {
}

func NewHandler() *GatewayApiHandler {
	return &GatewayApiHandler{}
}

func (h *GatewayApiHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/hello", h.sayHello)
	mux.HandleFunc("GET /api/v1/hello/nomames", h.noMames)

}

func (h *GatewayApiHandler) sayHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("nice")

}

func (h *GatewayApiHandler) noMames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("wooooow!")

}
