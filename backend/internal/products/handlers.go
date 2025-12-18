package products

import (
	"log"
	"net/http"

	"github.com/danielsonng/ecomgo/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 1. Call service
	products := []string{"Hello", "Wood"}
	// 2. Return JSON in a HTTP Res

	json.Write(w, http.StatusOK, products)
}
