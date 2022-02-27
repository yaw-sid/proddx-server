package storage

import (
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
)

func TestCompanyMemorySave(t *testing.T) {
	id := uuid.NewV4().String()
	cm := &CompanyModel{
		ID:            uuid.FromStringOrNil(id),
		CompanyUserID: uuid.NewV4().String(),
		CompanyName:   "Company One",
		Email:         "company@domain.com",
		Logo:          "https://proddx.com/company-one/logo.png",
		CreatedAt:     time.Now(),
	}
	storage := new(CompanyMemoryStore)
	if err := storage.Save(cm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if _, err := storage.Find(id); err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestCompanyMemoryList(t *testing.T) {
	id := uuid.NewV4().String()
	cm := &CompanyModel{
		ID:            uuid.FromStringOrNil(id),
		CompanyUserID: uuid.NewV4().String(),
		CompanyName:   "Company One",
		Email:         "company@domain.com",
		Logo:          "https://proddx.com/company-one/logo.png",
		CreatedAt:     time.Now(),
	}
	storage := new(CompanyMemoryStore)
	if err := storage.Save(cm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	cms, err := storage.List()
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	if len(cms) != 1 {
		t.Errorf("Error: %s - %d", "Wrong number of records", len(cms))
	}
}

func TestCompanyMemoryFind(t *testing.T) {
	id := uuid.NewV4().String()
	cm := &CompanyModel{
		ID:            uuid.FromStringOrNil(id),
		CompanyUserID: uuid.NewV4().String(),
		CompanyName:   "Company One",
		Email:         "company@domain.com",
		Logo:          "https://proddx.com/company-one/logo.png",
		CreatedAt:     time.Now(),
	}
	storage := new(CompanyMemoryStore)
	if err := storage.Save(cm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	record, err := storage.Find(id)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if record.ID != cm.ID {
		t.Errorf("Error: %s: %s - %s", "Record ID inconsistency", record.ID.String(), cm.ID.String())
	}
}

func TestCompanyMemoryDelete(t *testing.T) {
	id := uuid.NewV4().String()
	cm := &CompanyModel{
		ID:            uuid.FromStringOrNil(id),
		CompanyUserID: uuid.NewV4().String(),
		CompanyName:   "Company One",
		Email:         "company@domain.com",
		Logo:          "https://proddx.com/company-one/logo.png",
		CreatedAt:     time.Now(),
	}
	storage := new(CompanyMemoryStore)
	if err := storage.Save(cm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if err := storage.Delete(id); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if _, err := storage.Find(id); err == nil {
		t.Errorf("Error: %s", "Record was not deleted")
	}
}

func TestProductMemorySave(t *testing.T) {
	id := uuid.NewV4().String()
	pm := &ProductModel{
		ID:          uuid.FromStringOrNil(id),
		CompanyID:   uuid.NewV4(),
		ProductName: "Product One",
		FeedbackURL: "https://proddx.com/products/111/reviews",
		Rating:      4,
		CreatedAt:   time.Now(),
	}
	storage := new(ProductMemoryStore)
	if err := storage.Save(pm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if _, err := storage.Find(id); err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestProductMemoryList(t *testing.T) {
	id := uuid.NewV4().String()
	pm := &ProductModel{
		ID:          uuid.FromStringOrNil(id),
		CompanyID:   uuid.NewV4(),
		ProductName: "Product One",
		FeedbackURL: "https://proddx.com/products/111/reviews",
		Rating:      4,
		CreatedAt:   time.Now(),
	}
	storage := new(ProductMemoryStore)
	if err := storage.Save(pm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	records, err := storage.List()
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if len(records) != 1 {
		t.Errorf("Error: %s - %d", "Wrong number of records", len(records))
	}
}

func TestProductMemoryFind(t *testing.T) {
	id := uuid.NewV4().String()
	pm := &ProductModel{
		ID:          uuid.FromStringOrNil(id),
		CompanyID:   uuid.NewV4(),
		ProductName: "Product One",
		FeedbackURL: "https://proddx.com/products/111/reviews",
		Rating:      4,
		CreatedAt:   time.Now(),
	}
	storage := new(ProductMemoryStore)
	if err := storage.Save(pm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	record, err := storage.Find(id)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	if record.ID != pm.ID {
		t.Errorf("Error: %s: %s - %s", "Record ID inconsistency", record.ID.String(), pm.ID.String())
	}
}

func TestProductMemoryDelete(t *testing.T) {
	id := uuid.NewV4().String()
	pm := &ProductModel{
		ID:          uuid.FromStringOrNil(id),
		CompanyID:   uuid.NewV4(),
		ProductName: "Product One",
		FeedbackURL: "https://proddx.com/products/111/reviews",
		Rating:      4,
		CreatedAt:   time.Now(),
	}
	storage := new(ProductMemoryStore)
	if err := storage.Save(pm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if err := storage.Delete(id); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if _, err := storage.Find(id); err == nil {
		t.Errorf("Error: %s", "Record not deleted")
	}
}

func TestReviewMemorySave(t *testing.T) {
	id := uuid.NewV4().String()
	rm := &ReviewModel{
		ID:        uuid.FromStringOrNil(id),
		CompanyID: uuid.NewV4(),
		ProductID: uuid.NewV4(),
		Comment:   "Lorem ipsum dolor sit amet",
		Rating:    4,
		CreatedAt: time.Now(),
	}
	storage := new(ReviewMemoryStore)
	if err := storage.Save(rm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if _, err := storage.Find(id); err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestReviewMemoryList(t *testing.T) {
	id := uuid.NewV4().String()
	rm := &ReviewModel{
		ID:        uuid.FromStringOrNil(id),
		CompanyID: uuid.NewV4(),
		ProductID: uuid.NewV4(),
		Comment:   "Lorem ipsum dolor sit amet",
		Rating:    4,
		CreatedAt: time.Now(),
	}
	storage := new(ReviewMemoryStore)
	if err := storage.Save(rm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	records, err := storage.List()
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if len(records) != 1 {
		t.Errorf("Error: %s - %d", "Wrong number of records", len(records))
	}
}

func TestReviewMemoryFind(t *testing.T) {
	id := uuid.NewV4().String()
	rm := &ReviewModel{
		ID:        uuid.FromStringOrNil(id),
		CompanyID: uuid.NewV4(),
		ProductID: uuid.NewV4(),
		Comment:   "Lorem ipsum dolor sit amet",
		Rating:    4,
		CreatedAt: time.Now(),
	}
	storage := new(ReviewMemoryStore)
	if err := storage.Save(rm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	record, err := storage.Find(id)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if record.ID != rm.ID {
		t.Errorf("Error: %s: %s - %s", "Record ID inconsistency", record.ID.String(), rm.ID.String())
	}
}

func TestReviewMemoryDelete(t *testing.T) {
	id := uuid.NewV4().String()
	rm := &ReviewModel{
		ID:        uuid.FromStringOrNil(id),
		CompanyID: uuid.NewV4(),
		ProductID: uuid.NewV4(),
		Comment:   "Lorem ipsum dolor sit amet",
		Rating:    4,
		CreatedAt: time.Now(),
	}
	storage := new(ReviewMemoryStore)
	if err := storage.Save(rm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if err := storage.Delete(id); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if _, err := storage.Find(id); err == nil {
		t.Errorf("Error: %s", "Record not deleted")
	}
}
