package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"api.proddx.com/storage"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

func insertProduct(storage storage.Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(productRequest)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			fmt.Println("Marshalling error:", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if req.CompanyID == "" || req.Name == "" {
			fmt.Println("Error: company_id and name are required")
			http.Error(w, "company_id, name and feedback_url are required", http.StatusBadRequest)
			return
		}

		prod := productFromTransport(req)
		prod.ID = uuid.NewV4().String()
		prod.FeedbackURL = fmt.Sprintf("https://review.proddx.com/%s", prod.ID)
		prod.Rating = 0
		prod.CreatedAt = time.Now()
		model := productToStorage(prod)
		if err := storage.Save(model); err != nil {
			fmt.Println("Storage error:", err.Error())
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(*prod)
	}
}

func listProducts(storage storage.Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		records, err := storage.List()
		if err != nil {
			fmt.Println("Storage error:", err.Error())
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if len(records) == 0 {
			fmt.Println("Error:", "No products found")
			http.Error(w, "No products found", http.StatusNotFound)
			return
		}
		var resp []product
		for _, record := range records {
			resp = append(resp, *productFromStorage(&record))
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func findProduct(storage storage.Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := httprouter.ParamsFromContext(r.Context())
		id := params.ByName("id")
		if _, err := uuid.FromString(id); err != nil {
			fmt.Println("ID Error:", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		record, err := storage.Find(id)
		if err != nil {
			fmt.Println("Storage error:", err.Error())
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		resp := productFromStorage(record)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(*resp)
	}
}

func updateProduct(storage storage.Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(productRequest)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			fmt.Println("Marshalling error:", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		params := httprouter.ParamsFromContext(r.Context())
		id := params.ByName("id")
		if _, err := uuid.FromString(id); err != nil {
			fmt.Println("ID Error:", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		prod := productFromTransport(req)
		prod.ID = id
		model := productToStorage(prod)
		if err := storage.Save(model); err != nil {
			fmt.Println("Storage error:", err.Error())
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		prod = productFromStorage(model)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(*prod)
	}
}

func deleteProduct(storage storage.Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := httprouter.ParamsFromContext(r.Context())
		id := params.ByName("id")
		if _, err := uuid.FromString(id); err != nil {
			fmt.Println("ID Error:", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := storage.Delete(id); err != nil {
			fmt.Println("Storage error:", err.Error())
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusNoContent)
	}
}
