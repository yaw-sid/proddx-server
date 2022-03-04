package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"api.proddx.com/storage"
	uuid "github.com/satori/go.uuid"
)

func TestLogin(t *testing.T) {
	email := "user@example.com"
	password := "password"
	hash, err := hashPassword(password)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	model := storage.UserModel{
		ID:           uuid.NewV4(),
		Email:        email,
		UserPassword: hash,
		CreatedAt:    time.Now(),
	}
	userStore := new(storage.UserMemoryStore)
	companyStore := new(storage.CompanyMemoryStore)
	productStore := new(storage.ProductMemoryStore)
	reviewStore := new(storage.ReviewMemoryStore)
	if err = userStore.Save(&model); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	reqBody := loginRequest{
		Email:    email,
		Password: password,
	}
	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(reqJSON))
	router := New(userStore, companyStore, productStore, reviewStore)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected route POST /login to be valid: %d - %s", w.Code, string(w.Body.Bytes()))
	}
	var res map[string]string
	if err = json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if res["token"] == "" {
		t.Fatal("Error: No token returned")
	}
}

func TestRegister(t *testing.T) {
	req := registrationRequest{
		Name:     "Company One",
		Email:    "company@domain.com",
		Password: "password",
	}
	reqJSON, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	userStore := new(storage.UserMemoryStore)
	companyStore := new(storage.CompanyMemoryStore)
	productStore := new(storage.ProductMemoryStore)
	reviewStore := new(storage.ReviewMemoryStore)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewReader(reqJSON))
	router := New(userStore, companyStore, productStore, reviewStore)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected route POST /register to be valid: %d - %s", w.Code, string(w.Body.Bytes()))
	}
	var comp company
	if err = json.Unmarshal(w.Body.Bytes(), &comp); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if comp.Name != req.Name || comp.Email != req.Email {
		t.Error("Error:", "Record inconsistency")
	}
}
