package router

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
	CompanyID string  `json:"company_id,omitempty"`
	ProductID string  `json:"product_id,omitempty"`
	Comment   string  `json:"comment,omitempty"`
	Rating    float64 `json:"rating,omitempty"`
}

type company struct {
	ID        string `json:"id,omitempty"`
	UserID    string `json:"user_id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Logo      string `json:"logo,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type product struct {
	ID          string  `json:"id,omitempty"`
	CompanyID   string  `json:"company_id,omitempty"`
	Name        string  `json:"name,omitempty"`
	FeedbackUrl string  `json:"feedback_url,omitempty"`
	Rating      float64 `json:"rating,omitempty"`
	CreatedAt   string  `json:"created_at,omitempty"`
}

type review struct {
	ID        string  `json:"id,omitempty"`
	CompanyID string  `json:"company_id,omitempty"`
	ProductID string  `json:"product_id,omitempty"`
	Comment   string  `json:"comment,omitempty"`
	Rating    float64 `json:"rating,omitempty"`
	CreatedAt string  `json:"created_at,omitempty"`
}
