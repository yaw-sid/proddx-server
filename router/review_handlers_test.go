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

func TestInsertReview(t *testing.T) {
	req := reviewRequest{
		CompanyID: uuid.NewV4().String(),
		ProductID: uuid.NewV4().String(),
		Comment:   "Lorem ipsum dolor sit amet",
		Rating:    3,
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
	r, _ := http.NewRequest(http.MethodPost, "/reviews", bytes.NewBuffer(reqJSON))
	router := New(userStore, companyStore, productStore, reviewStore)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusCreated {
		t.Fatal("Expected route POST /reviews to be valid")
	}
	var res review
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	record, err := reviewStore.Find(res.ID)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if record.ID.String() != res.ID {
		t.Errorf("Record ID inconsistency: %s -%s", record.ID.String(), res.ID)
	}
}

func TestListReviews(t *testing.T) {
	userStore := new(storage.UserMemoryStore)
	companyStore := new(storage.CompanyMemoryStore)
	productStore := new(storage.ProductMemoryStore)
	reviewStore := new(storage.ReviewMemoryStore)

	id := uuid.NewV4().String()
	rm := &storage.ReviewModel{
		ID:        uuid.FromStringOrNil(id),
		CompanyID: uuid.NewV4(),
		ProductID: uuid.NewV4(),
		Comment:   "Lorem ipsum dolor sit amet",
		Rating:    3,
		CreatedAt: time.Now(),
	}
	if err := reviewStore.Save(rm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/reviews", nil)
	router := New(userStore, companyStore, productStore, reviewStore)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatal("Expected route GET /reviews to be valid")
	}
	var res []review
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	records, err := reviewStore.List("", "")
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if len(res) != len(records) {
		t.Errorf("Error: %s: %d -%d", "Number of records inconsistency", len(res), len(records))
	}
}

func TestFindReview(t *testing.T) {
	userStore := new(storage.UserMemoryStore)
	companyStore := new(storage.CompanyMemoryStore)
	productStore := new(storage.ProductMemoryStore)
	reviewStore := new(storage.ReviewMemoryStore)

	id := uuid.NewV4().String()
	rm := &storage.ReviewModel{
		ID:        uuid.FromStringOrNil(id),
		CompanyID: uuid.NewV4(),
		ProductID: uuid.NewV4(),
		Comment:   "Lorem ipsum dolor sit amet",
		Rating:    3,
		CreatedAt: time.Now(),
	}
	if err := reviewStore.Save(rm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	route := fmt.Sprintf("/reviews/%s", id)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, route, nil)
	router := New(userStore, companyStore, productStore, reviewStore)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatal(fmt.Sprintf("Expected route GET %s to be valid", route))
	}
	var res review
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	record, err := reviewStore.Find(id)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if res.ID != record.ID.String() {
		t.Errorf("Error: %s: %s - %s", "Record ID inconsistency", res.ID, record.ID.String())
	}
}

func TestUpdateReview(t *testing.T) {
	userStore := new(storage.UserMemoryStore)
	companyStore := new(storage.CompanyMemoryStore)
	productStore := new(storage.ProductMemoryStore)
	reviewStore := new(storage.ReviewMemoryStore)

	id := uuid.NewV4().String()
	rm := &storage.ReviewModel{
		ID:        uuid.FromStringOrNil(id),
		CompanyID: uuid.NewV4(),
		ProductID: uuid.NewV4(),
		Comment:   "Lorem ipsum dolor sit amet",
		Rating:    3,
		CreatedAt: time.Now(),
	}
	if err := reviewStore.Save(rm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	req := reviewRequest{
		Comment: "Lorem ipsum dolor sit amet consectetur",
	}
	reqJSON, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	route := fmt.Sprintf("/reviews/%s", id)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodPut, route, bytes.NewBuffer(reqJSON))
	router := New(userStore, companyStore, productStore, reviewStore)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatal(fmt.Sprintf("Expected route PUT %s to be valid", route))
	}
	var res review
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	record, err := reviewStore.Find(id)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if res.Comment != record.Comment {
		t.Errorf("Error: %s: %s - %s", "Record Comment inconsistency", res.Comment, record.Comment)
	}
}

func TestDeleteReview(t *testing.T) {
	userStore := new(storage.UserMemoryStore)
	companyStore := new(storage.CompanyMemoryStore)
	productStore := new(storage.ProductMemoryStore)
	reviewStore := new(storage.ReviewMemoryStore)

	id := uuid.NewV4().String()
	rm := &storage.ReviewModel{
		ID:        uuid.FromStringOrNil(id),
		CompanyID: uuid.NewV4(),
		ProductID: uuid.NewV4(),
		Comment:   "Lorem ipsum dolor sit amet",
		Rating:    3,
		CreatedAt: time.Now(),
	}
	if err := reviewStore.Save(rm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	route := fmt.Sprintf("/reviews/%s", id)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodDelete, route, nil)
	router := New(userStore, companyStore, productStore, reviewStore)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusNoContent {
		t.Fatal(fmt.Sprintf("Expected route DELETE %s to be valid", route))
	}
	if _, err := reviewStore.Find(id); err == nil {
		t.Errorf("Error: %s", "Record failed to delete")
	}
}
