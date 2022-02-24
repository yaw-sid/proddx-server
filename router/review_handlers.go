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

func insertReview(storage storage.Review) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(reviewRequest)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			fmt.Println("Error:", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if req.CompanyID == "" || req.ProductID == "" || req.Comment == "" || req.Rating == 0 {
			fmt.Println("Error: company_id, product_id, comment and rating are required")
			http.Error(w, "company_id, product_id, comment and rating are required", http.StatusBadRequest)
			return
		}

		rev := reviewFromTransport(req)
		rev.ID = uuid.NewV4().String()
		rev.CreatedAt = time.Now()
		model := reviewToStorage(rev)
		if err := storage.Save(model); err != nil {
			fmt.Println("Error:", err.Error())
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(*rev)
	}
}

func listReviews(storage storage.Review) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		records, err := storage.List()
		if err != nil {
			fmt.Println("Error:", err.Error())
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if len(records) == 0 {
			fmt.Println("Error:", "No reviews found")
			http.Error(w, "No reviews found", http.StatusNotFound)
			return
		}
		var resp []review
		for _, record := range records {
			resp = append(resp, *reviewFromStorage(&record))
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func findReview(storage storage.Review) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := httprouter.ParamsFromContext(r.Context())
		id := params.ByName("id")
		if _, err := uuid.FromString(id); err != nil {
			fmt.Println("Error:", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		record, err := storage.Find(id)
		if err != nil {
			fmt.Println("Error:", err.Error())
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		resp := reviewFromStorage(record)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(*resp)
	}
}

func updateReview(storage storage.Review) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(reviewRequest)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			fmt.Println("Error:", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		params := httprouter.ParamsFromContext(r.Context())
		id := params.ByName("id")
		if _, err := uuid.FromString(id); err != nil {
			fmt.Println("Error:", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		rev := reviewFromTransport(req)
		rev.ID = id
		model := reviewToStorage(rev)
		if err := storage.Save(model); err != nil {
			fmt.Println("Error:", err.Error())
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		rev = reviewFromStorage(model)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(*rev)
	}
}

func deleteReview(storage storage.Review) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := httprouter.ParamsFromContext(r.Context())
		id := params.ByName("id")
		if _, err := uuid.FromString(id); err != nil {
			fmt.Println("Error:", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := storage.Delete(id); err != nil {
			fmt.Println("Error:", err.Error())
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusNoContent)
	}
}
