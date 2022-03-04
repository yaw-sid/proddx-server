package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"api.proddx.com/storage"
	uuid "github.com/satori/go.uuid"
)

func TestInsertCompany(t *testing.T) {
	compReq := companyRequest{
		UserID: uuid.NewV4().String(),
		Name:   "Company One",
		Email:  "company@domain.com",
		Logo:   "https://proddx.com/company-one/logo.png",
	}
	compReqJSON, err := json.Marshal(compReq)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	userStore := new(storage.UserMemoryStore)
	companyStore := new(storage.CompanyMemoryStore)
	productStore := new(storage.ProductMemoryStore)
	reviewStore := new(storage.ReviewMemoryStore)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodPost, "/companies", bytes.NewBuffer(compReqJSON))
	router := New(userStore, companyStore, productStore, reviewStore)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusCreated {
		t.Fatal("Expected route POST /companies to be valid")
	}
	var compRes company
	if err := json.Unmarshal(w.Body.Bytes(), &compRes); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	record, err := companyStore.Find(compRes.ID)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if record.ID.String() != compRes.ID {
		t.Errorf("Record ID inconsistency: %s -%s", record.ID.String(), compRes.ID)
	}
}

func TestListCompanies(t *testing.T) {
	userStore := new(storage.UserMemoryStore)
	companyStore := new(storage.CompanyMemoryStore)
	productStore := new(storage.ProductMemoryStore)
	reviewStore := new(storage.ReviewMemoryStore)

	id := uuid.NewV4().String()
	cm := &storage.CompanyModel{
		ID:            uuid.FromStringOrNil(id),
		CompanyUserID: uuid.NewV4().String(),
		CompanyName:   "Company One",
		Email:         "company@domain.com",
		Logo:          "https://proddx.com/company-one/logo.png",
		CreatedAt:     time.Now(),
	}
	if err := companyStore.Save(cm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/companies", nil)
	router := New(userStore, companyStore, productStore, reviewStore)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatal("Expected route GET /companies to be valid")
	}
	var compRes []company
	if err := json.Unmarshal(w.Body.Bytes(), &compRes); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	records, err := companyStore.List()
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if len(compRes) != len(records) {
		t.Errorf("Error: %s: %d -%d", "Number of records inconsistency", len(compRes), len(records))
	}
}

func TestFindCompany(t *testing.T) {
	userStore := new(storage.UserMemoryStore)
	companyStore := new(storage.CompanyMemoryStore)
	productStore := new(storage.ProductMemoryStore)
	reviewStore := new(storage.ReviewMemoryStore)

	id := uuid.NewV4().String()
	cm := &storage.CompanyModel{
		ID:            uuid.FromStringOrNil(id),
		CompanyUserID: uuid.NewV4().String(),
		CompanyName:   "Company One",
		Email:         "company@domain.com",
		Logo:          "https://proddx.com/company-one/logo.png",
		CreatedAt:     time.Now(),
	}
	if err := companyStore.Save(cm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	route := fmt.Sprintf("/companies/%s", id)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, route, nil)
	router := New(userStore, companyStore, productStore, reviewStore)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatal(fmt.Sprintf("Expected route GET %s to be valid", route))
	}
	var compRes company
	if err := json.Unmarshal(w.Body.Bytes(), &compRes); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	record, err := companyStore.Find(id)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if compRes.ID != record.ID.String() {
		t.Errorf("Error: %s: %s - %s", "Record ID inconsistency", compRes.ID, record.ID.String())
	}
}

func TestUpdateCompany(t *testing.T) {
	userStore := new(storage.UserMemoryStore)
	companyStore := new(storage.CompanyMemoryStore)
	productStore := new(storage.ProductMemoryStore)
	reviewStore := new(storage.ReviewMemoryStore)

	id := uuid.NewV4().String()
	cm := &storage.CompanyModel{
		ID:            uuid.FromStringOrNil(id),
		CompanyUserID: uuid.NewV4().String(),
		CompanyName:   "Company One",
		Email:         "company@domain.com",
		Logo:          "https://proddx.com/company-one/logo.png",
		CreatedAt:     time.Now(),
	}
	if err := companyStore.Save(cm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	compReq := companyRequest{
		Logo: "https://proddx.com/company-one/logo-2.png",
	}
	compReqJSON, err := json.Marshal(compReq)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	route := fmt.Sprintf("/companies/%s", id)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodPut, route, bytes.NewBuffer(compReqJSON))
	router := New(userStore, companyStore, productStore, reviewStore)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatal(fmt.Sprintf("Expected route PUT %s to be valid", route))
	}
	var compRes company
	if err := json.Unmarshal(w.Body.Bytes(), &compRes); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	record, err := companyStore.Find(id)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if compRes.Logo != record.Logo {
		t.Errorf("Error: %s: %s - %s", "Record Logo inconsistency", compRes.Logo, record.Logo)
	}
}

func TestDeleteCompany(t *testing.T) {
	userStore := new(storage.UserMemoryStore)
	companyStore := new(storage.CompanyMemoryStore)
	productStore := new(storage.ProductMemoryStore)
	reviewStore := new(storage.ReviewMemoryStore)

	id := uuid.NewV4().String()
	cm := &storage.CompanyModel{
		ID:            uuid.FromStringOrNil(id),
		CompanyUserID: uuid.NewV4().String(),
		CompanyName:   "Company One",
		Email:         "company@domain.com",
		Logo:          "https://proddx.com/company-one/logo.png",
		CreatedAt:     time.Now(),
	}
	if err := companyStore.Save(cm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	route := fmt.Sprintf("/companies/%s", id)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodDelete, route, nil)
	router := New(userStore, companyStore, productStore, reviewStore)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusNoContent {
		t.Fatal(fmt.Sprintf("Expected route DELETE %s to be valid", route))
	}
	if _, err := companyStore.Find(id); err == nil {
		t.Errorf("Error: %s", "Record failed to delete")
	}
}
