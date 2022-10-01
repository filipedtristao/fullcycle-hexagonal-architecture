package handler

import (
	"github.com/filipedtristao/hexagonal-architecture/application"
	"github.com/filipedtristao/hexagonal-architecture/adapters/dto"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
	"encoding/json"
)

func MakeProductHandler(r *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	r.Handle(
		"/products/{id}", 
		n.With(negroni.Wrap(getProduct(service))),
	).Methods("GET", "OPTIONS")

	r.Handle(
		"/products", 
		n.With(negroni.Wrap(createProduct(service))),
	).Methods("POST", "OPTIONS")

	r.Handle(
		"/products/{id}/enable", 
		n.With(negroni.Wrap(enableProduct(service))),
	).Methods("POST", "OPTIONS")

	r.Handle(
		"/products/{id}/disable", 
		n.With(negroni.Wrap(disableProduct(service))),
	).Methods("POST", "OPTIONS")
}

func getProduct(service application.ProductServiceInterface) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		
		err = json.NewEncoder(w).Encode(product)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	})
}

func createProduct(service application.ProductServiceInterface) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var productDto dto.Product
		err := json.NewDecoder(r.Body).Decode(&productDto)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(JsonError(err.Error()))

			return
		}

		product, err := service.Create(productDto.Name, productDto.Price)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(JsonError(err.Error()))

			return
		}

		err = json.NewEncoder(w).Encode(product)
		
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(JsonError(err.Error()))

			return
		}
	})
}

func enableProduct(service application.ProductServiceInterface) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		result, err := service.Enable(product)
		
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(JsonError(err.Error()))

			return
		}
		
		err = json.NewEncoder(w).Encode(result)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(JsonError(err.Error()))

			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

func disableProduct(service application.ProductServiceInterface) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		result, err := service.Disable(product)
		
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(JsonError(err.Error()))

			return
		}
		
		err = json.NewEncoder(w).Encode(result)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(JsonError(err.Error()))

			return
		}

		w.WriteHeader(http.StatusOK)
	})
}