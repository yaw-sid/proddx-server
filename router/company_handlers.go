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

func insertCompany(storage storage.Company) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody := new(companyRequest)
		if err := json.NewDecoder(r.Body).Decode(reqBody); err != nil {
			fmt.Println("Marshalling error:", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !(reqBody.UserID != "" && reqBody.Name != "" && reqBody.Email != "") {
			fmt.Println("Error: user_id, name and email are required")
			http.Error(w, "user_id, name and email are required", http.StatusBadRequest)
			return
		}

		comp := companyFromTransport(reqBody)
		comp.ID = uuid.NewV4().String()
		comp.CreatedAt = time.Now()
		model := companyToStorage(comp)
		if err := storage.Save(model); err != nil {
			fmt.Println("Storage error:", err.Error())
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(*comp)
	}
}

func listCompanies(storage storage.Company) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		records, err := storage.List()
		if err != nil {
			fmt.Println("Storage error:", err.Error())
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if len(records) == 0 {
			fmt.Println("Error:", "No companies found")
			http.Error(w, "No companies found", http.StatusNotFound)
			return
		}
		var resp []company
		for _, record := range records {
			resp = append(resp, *companyFromStorage(&record))
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func findCompany(storage storage.Company) http.HandlerFunc {
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
		resp := companyFromStorage(record)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(*resp)
	}
}

func updateCompany(storage storage.Company) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(companyRequest)
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

		comp := companyFromTransport(req)
		comp.ID = id
		model := companyToStorage(comp)
		if err := storage.Save(model); err != nil {
			fmt.Println("Storage error:", err.Error())
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		comp = companyFromStorage(model)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(*comp)
	}
}

func deleteCompany(storage storage.Company) http.HandlerFunc {
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
