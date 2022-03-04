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

func TestInsertProduct(t *testing.T) {
	req := productRequest{
		CompanyID:   uuid.NewV4().String(),
		Name:        "Product One",
		FeedbackURL: "https://proddx.com/product-one/reviews",
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
	r, _ := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(reqJSON))
	router := New(userStore, companyStore, productStore, reviewStore)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusCreated {
		t.Fatal("Expected route POST /products to be valid")
	}
	var res product
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	record, err := productStore.Find(res.ID)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if record.ID.String() != res.ID {
		t.Errorf("Record ID inconsistency: %s -%s", record.ID.String(), res.ID)
	}
}

func TestListProducts(t *testing.T) {
	userStore := new(storage.UserMemoryStore)
	companyStore := new(storage.CompanyMemoryStore)
	productStore := new(storage.ProductMemoryStore)
	reviewStore := new(storage.ReviewMemoryStore)

	id := uuid.NewV4().String()
	pm := &storage.ProductModel{
		ID:          uuid.FromStringOrNil(id),
		CompanyID:   uuid.NewV4(),
		ProductName: "Product One",
		FeedbackURL: "https://proddx.com/product-one/reviews",
		Rating:      4,
		CreatedAt:   time.Now(),
	}
	if err := productStore.Save(pm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/products", nil)
	router := New(userStore, companyStore, productStore, reviewStore)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatal("Expected route GET /products to be valid")
	}
	var res []product
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	records, err := productStore.List()
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if len(res) != len(records) {
		t.Errorf("Error: %s: %d -%d", "Number of records inconsistency", len(res), len(records))
	}
}

func TestFindProduct(t *testing.T) {
	userStore := new(storage.UserMemoryStore)
	companyStore := new(storage.CompanyMemoryStore)
	productStore := new(storage.ProductMemoryStore)
	reviewStore := new(storage.ReviewMemoryStore)

	id := uuid.NewV4().String()
	pm := &storage.ProductModel{
		ID:          uuid.FromStringOrNil(id),
		CompanyID:   uuid.NewV4(),
		ProductName: "Product One",
		FeedbackURL: "https://proddx.com/product-one/reviews",
		Rating:      4,
		CreatedAt:   time.Now(),
	}
	if err := productStore.Save(pm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	route := fmt.Sprintf("/products/%s", id)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, route, nil)
	router := New(userStore, companyStore, productStore, reviewStore)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected route GET %s to be valid", route)
	}
	var res product
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	record, err := productStore.Find(id)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if res.ID != record.ID.String() {
		t.Errorf("Error: %s: %s - %s", "Record ID inconsistency", res.ID, record.ID.String())
	}
}

func TestUpdateProduct(t *testing.T) {
	userStore := new(storage.UserMemoryStore)
	companyStore := new(storage.CompanyMemoryStore)
	productStore := new(storage.ProductMemoryStore)
	reviewStore := new(storage.ReviewMemoryStore)

	id := uuid.NewV4().String()
	pm := &storage.ProductModel{
		ID:          uuid.FromStringOrNil(id),
		CompanyID:   uuid.NewV4(),
		ProductName: "Product One",
		FeedbackURL: "https://proddx.com/product-one/reviews",
		Rating:      4,
		CreatedAt:   time.Now(),
	}
	if err := productStore.Save(pm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	req := productRequest{
		Name: "Product Two",
	}
	reqJSON, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	route := fmt.Sprintf("/products/%s", id)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodPut, route, bytes.NewBuffer(reqJSON))
	router := New(userStore, companyStore, productStore, reviewStore)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected route PUT %s to be valid", route)
	}
	var res product
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	record, err := productStore.Find(id)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if res.Name != record.ProductName {
		t.Errorf("Error: %s: %s - %s", "Record Name inconsistency", res.Name, record.ProductName)
	}
}

func TestDeleteProduct(t *testing.T) {
	userStore := new(storage.UserMemoryStore)
	companyStore := new(storage.CompanyMemoryStore)
	productStore := new(storage.ProductMemoryStore)
	reviewStore := new(storage.ReviewMemoryStore)

	id := uuid.NewV4().String()
	pm := &storage.ProductModel{
		ID:          uuid.FromStringOrNil(id),
		CompanyID:   uuid.NewV4(),
		ProductName: "Product One",
		FeedbackURL: "https://proddx.com/product-one/reviews",
		Rating:      4,
		CreatedAt:   time.Now(),
	}
	if err := productStore.Save(pm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	route := fmt.Sprintf("/products/%s", id)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodDelete, route, nil)
	router := New(userStore, companyStore, productStore, reviewStore)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusNoContent {
		t.Fatal(fmt.Sprintf("Expected route DELETE %s to be valid", route))
	}
	if _, err := productStore.Find(id); err == nil {
		t.Errorf("Error: %s", "Record failed to delete")
	}
}
