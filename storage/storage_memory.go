package storage

import "errors"

type companyMemoryStore struct {
	companies []CompanyModel
}

func (cms *companyMemoryStore) save(model *CompanyModel) error {
	for index, record := range cms.companies {
		if record.ID == model.ID {
			cms.companies[index] = *model
			return nil
		}
	}
	cms.companies = append(cms.companies, *model)
	return nil
}

func (cms companyMemoryStore) list() ([]CompanyModel, error) {
	if len(cms.companies) == 0 {
		return cms.companies, errors.New("No companies found")
	}
	return cms.companies, nil
}

func (cms companyMemoryStore) find(id string) (*CompanyModel, error) {
	for _, record := range cms.companies {
		if record.ID.String() == id {
			return &record, nil
		}
	}
	return nil, errors.New("Company not found")
}

func (cms *companyMemoryStore) delete(id string) error {
	for index, record := range cms.companies {
		if record.ID.String() == id {
			cms.companies[index] = cms.companies[len(cms.companies)-1]
			cms.companies = cms.companies[:len(cms.companies)-1]
			return nil
		}
	}
	return errors.New("Company not found")
}

type productMemoryStore struct {
	products []ProductModel
}

func (pms *productMemoryStore) save(model *ProductModel) error {
	for index, record := range pms.products {
		if record.ID == model.ID {
			pms.products[index] = *model
			return nil
		}
	}
	pms.products = append(pms.products, *model)
	return nil
}

func (pms productMemoryStore) list() ([]ProductModel, error) {
	if len(pms.products) == 0 {
		return pms.products, errors.New("No products found")
	}
	return pms.products, nil
}

func (pms productMemoryStore) find(id string) (*ProductModel, error) {
	for _, record := range pms.products {
		if record.ID.String() == id {
			return &record, nil
		}
	}
	return nil, errors.New("Product not found")
}

func (pms *productMemoryStore) delete(id string) error {
	for index, record := range pms.products {
		if record.ID.String() == id {
			pms.products[index] = pms.products[len(pms.products)-1]
			pms.products = pms.products[:len(pms.products)-1]
			return nil
		}
	}
	return errors.New("Product not found")
}

type reviewMemoryStore struct {
	reviews []ReviewModel
}

func (rms *reviewMemoryStore) save(model *ReviewModel) error {
	for index, record := range rms.reviews {
		if record.ID == model.ID {
			rms.reviews[index] = *model
			return nil
		}
	}
	rms.reviews = append(rms.reviews, *model)
	return nil
}

func (rms reviewMemoryStore) list() ([]ReviewModel, error) {
	if len(rms.reviews) == 0 {
		return rms.reviews, errors.New("No reviews found")
	}
	return rms.reviews, nil
}

func (rms reviewMemoryStore) find(id string) (*ReviewModel, error) {
	for _, record := range rms.reviews {
		if record.ID.String() == id {
			return &record, nil
		}
	}
	return nil, errors.New("Review not found")
}

func (rms *reviewMemoryStore) delete(id string) error {
	for index, record := range rms.reviews {
		if record.ID.String() == id {
			rms.reviews[index] = rms.reviews[len(rms.reviews)-1]
			rms.reviews = rms.reviews[:len(rms.reviews)-1]
			return nil
		}
	}
	return errors.New("Review not found")
}
