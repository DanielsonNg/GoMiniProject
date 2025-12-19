package products

import (
	"log"
	"net/http"
	"strconv"

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
	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, products)
}

func (h *handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		// Handle the error (e.g., return a 400 Bad Request)
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	product, err := h.service.GetProductById(r.Context(), idInt)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, product)
}
