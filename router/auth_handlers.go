package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"api.proddx.com/storage"
	"api.proddx.com/tokens"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func login(storage storage.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(loginRequest)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			fmt.Println("Error:", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if &req.Email == nil || &req.Password == nil {
			fmt.Println("Error:", "email and password are required")
			http.Error(w, "email and password are required", http.StatusBadRequest)
			return
		}
		record, err := storage.Find(req.Email)
		if err != nil {
			fmt.Println("Error:", err.Error())
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if !checkPasswordHash(req.Password, record.UserPassword) {
			fmt.Println("Error:", "Incorrect password")
			http.Error(w, "Incorrect password", http.StatusUnauthorized)
			return
		}
		token, err := tokens.New(record.ID.String())
		if err != nil {
			fmt.Println("Error:", err.Error())
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		res := map[string]string{"token": token}
		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

func register(userStorage storage.User, companyStorage storage.Company) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req registrationRequest
		var err error
		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			fmt.Println("Error", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if &req.Email == nil || &req.Name == nil || &req.Password == nil {
			errString := "name, email and password are required"
			fmt.Println("Error:", errString)
			http.Error(w, errString, http.StatusBadRequest)
			return
		}

		u := user{
			ID:        uuid.NewV4().String(),
			Email:     req.Email,
			CreatedAt: time.Now(),
		}
		if u.Password, err = hashPassword(req.Password); err != nil {
			fmt.Println("Error:", err.Error())
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		userModel := userToStorage(&u)
		if err = userStorage.Save(userModel); err != nil {
			fmt.Println("Error:", err.Error())
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		comp := company{
			ID:        uuid.NewV4().String(),
			UserID:    u.ID,
			Name:      req.Name,
			Email:     req.Email,
			Logo:      "",
			CreatedAt: time.Now(),
		}
		compModel := companyToStorage(&comp)
		if err = companyStorage.Save(compModel); err != nil {
			fmt.Println("Error:", err.Error())
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(comp)
	}
}
