package storage

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type CompanyModel struct {
	ID            uuid.UUID
	CompanyUserID string
	CompanyName   string
	Email         string
	Logo          string
	CreatedAt     time.Time
}

type ProductModel struct {
	ID          uuid.UUID
	CompanyID   uuid.UUID
	ProductName string
	FeedbackURL string
	Rating      uint
	CreatedAt   time.Time
}

type ReviewModel struct {
	ID        uuid.UUID
	CompanyID uuid.UUID
	ProductID uuid.UUID
	Comment   string
	Rating    uint
	CreatedAt time.Time
}
