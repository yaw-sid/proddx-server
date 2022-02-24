package router

import (
	"api.proddx.com/storage"
	uuid "github.com/satori/go.uuid"
)

func companyToStorage(c *company) *storage.CompanyModel {
	return &storage.CompanyModel{
		ID:        uuid.FromStringOrNil(c.ID),
		UserID:    c.UserID,
		Name:      c.Name,
		Email:     c.Email,
		Logo:      c.Logo,
		CreatedAt: c.CreatedAt,
	}
}

func companyFromStorage(model *storage.CompanyModel) *company {
	return &company{
		ID:        model.ID.String(),
		UserID:    model.UserID,
		Name:      model.Name,
		Email:     model.Email,
		Logo:      model.Logo,
		CreatedAt: model.CreatedAt,
	}
}

func companyFromTransport(req *companyRequest) *company {
	return &company{
		UserID: req.UserID,
		Name:   req.Name,
		Email:  req.Email,
		Logo:   req.Logo,
	}
}

func productToStorage(p *product) *storage.ProductModel {
	return &storage.ProductModel{
		ID:          uuid.FromStringOrNil(p.ID),
		CompanyID:   uuid.FromStringOrNil(p.CompanyID),
		Name:        p.Name,
		FeedbackURL: p.FeedbackURL,
		Rating:      p.Rating,
		CreatedAt:   p.CreatedAt,
	}
}

func productFromStorage(model *storage.ProductModel) *product {
	return &product{
		ID:          model.ID.String(),
		CompanyID:   model.CompanyID.String(),
		Name:        model.Name,
		FeedbackURL: model.FeedbackURL,
		Rating:      model.Rating,
		CreatedAt:   model.CreatedAt,
	}
}

func productFromTransport(req *productRequest) *product {
	return &product{
		CompanyID:   req.CompanyID,
		Name:        req.Name,
		FeedbackURL: req.FeedbackURL,
	}
}

func reviewToStorage(r *review) *storage.ReviewModel {
	return &storage.ReviewModel{
		ID:        uuid.FromStringOrNil(r.ID),
		CompanyID: uuid.FromStringOrNil(r.CompanyID),
		ProductID: uuid.FromStringOrNil(r.ProductID),
		Comment:   r.Comment,
		Rating:    r.Rating,
		CreatedAt: r.CreatedAt,
	}
}

func reviewFromStorage(model *storage.ReviewModel) *review {
	return &review{
		ID:        model.ID.String(),
		CompanyID: model.CompanyID.String(),
		ProductID: model.ProductID.String(),
		Comment:   model.Comment,
		Rating:    model.Rating,
		CreatedAt: model.CreatedAt,
	}
}

func reviewFromTransport(req *reviewRequest) *review {
	return &review{
		CompanyID: req.CompanyID,
		ProductID: req.ProductID,
		Comment:   req.Comment,
		Rating:    req.Rating,
	}
}
