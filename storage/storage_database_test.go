package storage

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
	uuid "github.com/satori/go.uuid"
)

func TestCompanyDatabaseSave(t *testing.T) {
	id := uuid.NewV4().String()
	cm := &CompanyModel{
		ID:            uuid.FromStringOrNil(id),
		CompanyUserID: uuid.NewV4().String(),
		CompanyName:   "Company One",
		Email:         "company@domain.com",
		Logo:          "https://proddx.com/company-one/logo.png",
		CreatedAt:     time.Now(),
	}
	conn, err := pgx.Connect(context.Background(), "postgres://novi:novi@localhost:5432/proddx_db?sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to database: %s", err.Error())
	}
	storage := &CompanyDatabase{Conn: conn}
	if err = storage.Save(cm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if _, err := storage.Find(id); err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	if err = storage.Delete(id); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
}

func TestCompanyDatabaseList(t *testing.T) {
	id := uuid.NewV4().String()
	cm := &CompanyModel{
		ID:            uuid.FromStringOrNil(id),
		CompanyUserID: uuid.NewV4().String(),
		CompanyName:   "Company One",
		Email:         "company@domain.com",
		Logo:          "https://proddx.com/company-one/logo.png",
		CreatedAt:     time.Now(),
	}
	conn, err := pgx.Connect(context.Background(), "postgres://novi:novi@localhost:5432/proddx_db?sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to database: %s", err.Error())
	}
	storage := &CompanyDatabase{Conn: conn}
	if err = storage.Save(cm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	cms, err := storage.List()
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	if len(cms) != 1 {
		t.Errorf("Error: %s - %d", "Wrong number of records", len(cms))
	}
	if err = storage.Delete(id); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
}

func TestCompanyDatabaseFind(t *testing.T) {
	id := uuid.NewV4().String()
	cm := &CompanyModel{
		ID:            uuid.FromStringOrNil(id),
		CompanyUserID: uuid.NewV4().String(),
		CompanyName:   "Company One",
		Email:         "company@domain.com",
		Logo:          "https://proddx.com/company-one/logo.png",
		CreatedAt:     time.Now(),
	}
	conn, err := pgx.Connect(context.Background(), "postgres://novi:novi@localhost:5432/proddx_db?sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to database: %s", err.Error())
	}
	storage := &CompanyDatabase{Conn: conn}
	if err = storage.Save(cm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	record, err := storage.Find(id)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if record.ID != cm.ID {
		t.Errorf("Error: %s: %s - %s", "Record ID inconsistency", record.ID.String(), cm.ID.String())
	}
	if err = storage.Delete(id); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
}

func TestCompanyDatabaseDelete(t *testing.T) {
	id := uuid.NewV4().String()
	cm := &CompanyModel{
		ID:            uuid.FromStringOrNil(id),
		CompanyUserID: uuid.NewV4().String(),
		CompanyName:   "Company One",
		Email:         "company@domain.com",
		Logo:          "https://proddx.com/company-one/logo.png",
		CreatedAt:     time.Now(),
	}
	conn, err := pgx.Connect(context.Background(), "postgres://novi:novi@localhost:5432/proddx_db?sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to database: %s", err.Error())
	}
	storage := &CompanyDatabase{Conn: conn}
	if err = storage.Save(cm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if err = storage.Delete(id); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if _, err := storage.Find(id); err == nil {
		t.Errorf("Error: %s", "Record was not deleted")
	}
}

func TestProductDatabaseSave(t *testing.T) {
	companyID := uuid.NewV4().String()
	cm := &CompanyModel{
		ID:            uuid.FromStringOrNil(companyID),
		CompanyUserID: uuid.NewV4().String(),
		CompanyName:   "Company One",
		Email:         "company@domain.com",
		Logo:          "https://proddx.com/company-one/logo.png",
		CreatedAt:     time.Now(),
	}
	productID := uuid.NewV4().String()
	pm := &ProductModel{
		ID:          uuid.FromStringOrNil(productID),
		CompanyID:   uuid.FromStringOrNil(companyID),
		ProductName: "Product One",
		FeedbackURL: "https://proddx.com/product-one/reviews",
		Rating:      0,
		CreatedAt:   time.Now(),
	}
	conn, err := pgx.Connect(context.Background(), "postgres://novi:novi@localhost:5432/proddx_db?sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to database: %s", err.Error())
	}
	companyStorage := &CompanyDatabase{Conn: conn}
	if err = companyStorage.Save(cm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	productStorage := &ProductDatabase{Conn: conn}
	if err = productStorage.Save(pm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if _, err = productStorage.Find(productID); err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	if err = productStorage.Delete(productID); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if err = companyStorage.Delete(companyID); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
}

func TestProductDatabaseList(t *testing.T) {
	companyID := uuid.NewV4().String()
	cm := &CompanyModel{
		ID:            uuid.FromStringOrNil(companyID),
		CompanyUserID: uuid.NewV4().String(),
		CompanyName:   "Company One",
		Email:         "company@domain.com",
		Logo:          "https://proddx.com/company-one/logo.png",
		CreatedAt:     time.Now(),
	}
	productID := uuid.NewV4().String()
	pm := &ProductModel{
		ID:          uuid.FromStringOrNil(productID),
		CompanyID:   uuid.FromStringOrNil(companyID),
		ProductName: "Product One",
		FeedbackURL: "https://proddx.com/product-one/reviews",
		Rating:      0,
		CreatedAt:   time.Now(),
	}
	conn, err := pgx.Connect(context.Background(), "postgres://novi:novi@localhost:5432/proddx_db?sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to database: %s", err.Error())
	}
	companyStorage := &CompanyDatabase{Conn: conn}
	if err = companyStorage.Save(cm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	productStorage := &ProductDatabase{Conn: conn}
	if err = productStorage.Save(pm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	records, err := productStorage.List()
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if len(records) != 1 {
		t.Errorf("Error: Wrong number of returned records - %d", len(records))
	}
	if err = productStorage.Delete(productID); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if err = companyStorage.Delete(companyID); err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestProductDatabaseFind(t *testing.T) {
	companyID := uuid.NewV4().String()
	cm := &CompanyModel{
		ID:            uuid.FromStringOrNil(companyID),
		CompanyUserID: uuid.NewV4().String(),
		CompanyName:   "Company One",
		Email:         "company@domain.com",
		Logo:          "https://proddx.com/company-one/logo.png",
		CreatedAt:     time.Now(),
	}
	productID := uuid.NewV4().String()
	pm := &ProductModel{
		ID:          uuid.FromStringOrNil(productID),
		CompanyID:   uuid.FromStringOrNil(companyID),
		ProductName: "Product One",
		FeedbackURL: "https://proddx.com/product-one/reviews",
		Rating:      0,
		CreatedAt:   time.Now(),
	}
	conn, err := pgx.Connect(context.Background(), "postgres://novi:novi@localhost:5432/proddx_db?sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to database: %s", err.Error())
	}
	companyStorage := &CompanyDatabase{Conn: conn}
	if err = companyStorage.Save(cm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	productStorage := &ProductDatabase{Conn: conn}
	if err = productStorage.Save(pm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	record, err := productStorage.Find(productID)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if record.ID != pm.ID {
		t.Errorf("Error: Record ID inconsistency: %s - %s", record.ID.String(), pm.ID.String())
	}
	if err = productStorage.Delete(productID); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if err = companyStorage.Delete(companyID); err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestProductDatabaseDelete(t *testing.T) {
	companyID := uuid.NewV4().String()
	cm := &CompanyModel{
		ID:            uuid.FromStringOrNil(companyID),
		CompanyUserID: uuid.NewV4().String(),
		CompanyName:   "Company One",
		Email:         "company@domain.com",
		Logo:          "https://proddx.com/company-one/logo.png",
		CreatedAt:     time.Now(),
	}
	productID := uuid.NewV4().String()
	pm := &ProductModel{
		ID:          uuid.FromStringOrNil(productID),
		CompanyID:   uuid.FromStringOrNil(companyID),
		ProductName: "Product One",
		FeedbackURL: "https://proddx.com/product-one/reviews",
		Rating:      0,
		CreatedAt:   time.Now(),
	}
	conn, err := pgx.Connect(context.Background(), "postgres://novi:novi@localhost:5432/proddx_db?sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to database: %s", err.Error())
	}
	companyStorage := &CompanyDatabase{Conn: conn}
	if err = companyStorage.Save(cm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	productStorage := &ProductDatabase{Conn: conn}
	if err = productStorage.Save(pm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if err = productStorage.Delete(productID); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if _, err = productStorage.Find(productID); err == nil {
		t.Errorf("Error: %s", "Record failed to delete")
	}
	if err = companyStorage.Delete(companyID); err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestReviewDatabaseSave(t *testing.T) {
	companyID := uuid.NewV4().String()
	cm := &CompanyModel{
		ID:            uuid.FromStringOrNil(companyID),
		CompanyUserID: uuid.NewV4().String(),
		CompanyName:   "Company One",
		Email:         "company@domain.com",
		Logo:          "https://proddx.com/company-one/logo.png",
		CreatedAt:     time.Now(),
	}
	productID := uuid.NewV4().String()
	pm := &ProductModel{
		ID:          uuid.FromStringOrNil(productID),
		CompanyID:   uuid.FromStringOrNil(companyID),
		ProductName: "Product One",
		FeedbackURL: "https://proddx.com/product-one/reviews",
		Rating:      0,
		CreatedAt:   time.Now(),
	}
	reviewID := uuid.NewV4().String()
	rm := &ReviewModel{
		ID:        uuid.FromStringOrNil(reviewID),
		CompanyID: uuid.FromStringOrNil(companyID),
		ProductID: uuid.FromStringOrNil(productID),
		Comment:   "Lorem ipsum dolor sit amet",
		Rating:    3,
		CreatedAt: time.Now(),
	}
	conn, err := pgx.Connect(context.Background(), "postgres://novi:novi@localhost:5432/proddx_db?sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to database: %s", err.Error())
	}
	companyStorage := &CompanyDatabase{Conn: conn}
	if err = companyStorage.Save(cm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	productStorage := &ProductDatabase{Conn: conn}
	if err = productStorage.Save(pm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	reviewStorage := &ReviewDatabase{Conn: conn}
	if err = reviewStorage.Save(rm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if _, err = reviewStorage.Find(reviewID); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if err = reviewStorage.Delete(reviewID); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if err = productStorage.Delete(productID); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if err = companyStorage.Delete(companyID); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
}

func TestReviewDatabaseList(t *testing.T) {
	companyID := uuid.NewV4().String()
	cm := &CompanyModel{
		ID:            uuid.FromStringOrNil(companyID),
		CompanyUserID: uuid.NewV4().String(),
		CompanyName:   "Company One",
		Email:         "company@domain.com",
		Logo:          "https://proddx.com/company-one/logo.png",
		CreatedAt:     time.Now(),
	}
	productID := uuid.NewV4().String()
	pm := &ProductModel{
		ID:          uuid.FromStringOrNil(productID),
		CompanyID:   uuid.FromStringOrNil(companyID),
		ProductName: "Product One",
		FeedbackURL: "https://proddx.com/product-one/reviews",
		Rating:      0,
		CreatedAt:   time.Now(),
	}
	reviewID := uuid.NewV4().String()
	rm := &ReviewModel{
		ID:        uuid.FromStringOrNil(reviewID),
		CompanyID: uuid.FromStringOrNil(companyID),
		ProductID: uuid.FromStringOrNil(productID),
		Comment:   "Lorem ipsum dolor sit amet",
		Rating:    3,
		CreatedAt: time.Now(),
	}
	conn, err := pgx.Connect(context.Background(), "postgres://novi:novi@localhost:5432/proddx_db?sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to database: %s", err.Error())
	}
	companyStorage := &CompanyDatabase{Conn: conn}
	if err = companyStorage.Save(cm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	productStorage := &ProductDatabase{Conn: conn}
	if err = productStorage.Save(pm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	reviewStorage := &ReviewDatabase{Conn: conn}
	if err = reviewStorage.Save(rm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	records, err := reviewStorage.List()
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if len(records) != 1 {
		t.Errorf("Error: Wrong number of records returned - %d", len(records))
	}
	if err = reviewStorage.Delete(reviewID); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if err = productStorage.Delete(productID); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if err = companyStorage.Delete(companyID); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
}

func TestReviewDatabaseFind(t *testing.T) {
	companyID := uuid.NewV4().String()
	cm := &CompanyModel{
		ID:            uuid.FromStringOrNil(companyID),
		CompanyUserID: uuid.NewV4().String(),
		CompanyName:   "Company One",
		Email:         "company@domain.com",
		Logo:          "https://proddx.com/company-one/logo.png",
		CreatedAt:     time.Now(),
	}
	productID := uuid.NewV4().String()
	pm := &ProductModel{
		ID:          uuid.FromStringOrNil(productID),
		CompanyID:   uuid.FromStringOrNil(companyID),
		ProductName: "Product One",
		FeedbackURL: "https://proddx.com/product-one/reviews",
		Rating:      0,
		CreatedAt:   time.Now(),
	}
	reviewID := uuid.NewV4().String()
	rm := &ReviewModel{
		ID:        uuid.FromStringOrNil(reviewID),
		CompanyID: uuid.FromStringOrNil(companyID),
		ProductID: uuid.FromStringOrNil(productID),
		Comment:   "Lorem ipsum dolor sit amet",
		Rating:    3,
		CreatedAt: time.Now(),
	}
	conn, err := pgx.Connect(context.Background(), "postgres://novi:novi@localhost:5432/proddx_db?sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to database: %s", err.Error())
	}
	companyStorage := &CompanyDatabase{Conn: conn}
	if err = companyStorage.Save(cm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	productStorage := &ProductDatabase{Conn: conn}
	if err = productStorage.Save(pm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	reviewStorage := &ReviewDatabase{Conn: conn}
	if err = reviewStorage.Save(rm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	record, err := reviewStorage.Find(reviewID)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if record.ID != rm.ID {
		t.Errorf("Error: Record ID inconsistency: %s - %s", record.ID.String(), rm.ID.String())
	}
	if err = reviewStorage.Delete(reviewID); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if err = productStorage.Delete(productID); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if err = companyStorage.Delete(companyID); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
}

func TestReviewDatabaseDelete(t *testing.T) {
	companyID := uuid.NewV4().String()
	cm := &CompanyModel{
		ID:            uuid.FromStringOrNil(companyID),
		CompanyUserID: uuid.NewV4().String(),
		CompanyName:   "Company One",
		Email:         "company@domain.com",
		Logo:          "https://proddx.com/company-one/logo.png",
		CreatedAt:     time.Now(),
	}
	productID := uuid.NewV4().String()
	pm := &ProductModel{
		ID:          uuid.FromStringOrNil(productID),
		CompanyID:   uuid.FromStringOrNil(companyID),
		ProductName: "Product One",
		FeedbackURL: "https://proddx.com/product-one/reviews",
		Rating:      0,
		CreatedAt:   time.Now(),
	}
	reviewID := uuid.NewV4().String()
	rm := &ReviewModel{
		ID:        uuid.FromStringOrNil(reviewID),
		CompanyID: uuid.FromStringOrNil(companyID),
		ProductID: uuid.FromStringOrNil(productID),
		Comment:   "Lorem ipsum dolor sit amet",
		Rating:    3,
		CreatedAt: time.Now(),
	}
	conn, err := pgx.Connect(context.Background(), "postgres://novi:novi@localhost:5432/proddx_db?sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to database: %s", err.Error())
	}
	companyStorage := &CompanyDatabase{Conn: conn}
	if err = companyStorage.Save(cm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	productStorage := &ProductDatabase{Conn: conn}
	if err = productStorage.Save(pm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	reviewStorage := &ReviewDatabase{Conn: conn}
	if err = reviewStorage.Save(rm); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if err = reviewStorage.Delete(reviewID); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if _, err = reviewStorage.Find(reviewID); err == nil {
		t.Errorf("Error: Record was not deleted")
	}
	if err = productStorage.Delete(productID); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if err = companyStorage.Delete(companyID); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
}
