package router

import "time"

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registrationRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type companyRequest struct {
	UserID string `json:"user_id,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Logo   string `json:"logo,omitempty"`
}

type productRequest struct {
	CompanyID   string `json:"company_id,omitempty"`
	Name        string `json:"name,omitempty"`
	FeedbackURL string `json:"feedback_url,omitempty"`
}

type reviewRequest struct {
	CompanyID string `json:"company_id,omitempty"`
	ProductID string `json:"product_id,omitempty"`
	Comment   string `json:"comment,omitempty"`
	Rating    uint   `json:"rating,omitempty"`
}

type user struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type company struct {
	ID        string    `json:"id,omitempty"`
	UserID    string    `json:"user_id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Logo      string    `json:"logo,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type product struct {
	ID          string    `json:"id,omitempty"`
	CompanyID   string    `json:"company_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	FeedbackURL string    `json:"feedback_url,omitempty"`
	Rating      uint      `json:"rating,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

type review struct {
	ID        string    `json:"id,omitempty"`
	CompanyID string    `json:"company_id,omitempty"`
	ProductID string    `json:"product_id,omitempty"`
	Comment   string    `json:"comment,omitempty"`
	Rating    uint      `json:"rating,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
