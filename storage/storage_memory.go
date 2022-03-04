package storage

import "errors"

type UserMemoryStore struct {
	users []UserModel
}

func (ums *UserMemoryStore) Save(model *UserModel) error {
	for _, record := range ums.users {
		if record.ID == model.ID {
			return errors.New("User already exists")
		}
	}
	ums.users = append(ums.users, *model)
	return nil
}

func (ums UserMemoryStore) Find(id string) (*UserModel, error) {
	for _, record := range ums.users {
		if record.ID.String() == id || record.Email == id {
			return &record, nil
		}
	}
	return nil, errors.New("User not found")
}

func (ums *UserMemoryStore) Delete(id string) error {
	for index, record := range ums.users {
		if record.ID.String() == id {
			ums.users[index] = ums.users[len(ums.users)-1]
			ums.users = ums.users[:len(ums.users)-1]
			return nil
		}
	}
	return errors.New("User not found")
}

type CompanyMemoryStore struct {
	companies []CompanyModel
}

func (cms *CompanyMemoryStore) Save(model *CompanyModel) error {
	for index, record := range cms.companies {
		if record.ID == model.ID {
			if &model.CompanyName != nil {
				cms.companies[index].CompanyName = model.CompanyName
			}
			if &model.Email != nil {
				cms.companies[index].Email = model.Email
			}
			if &model.Logo != nil {
				cms.companies[index].Logo = model.Logo
			}
			model = &cms.companies[index]
			return nil
		}
	}
	cms.companies = append(cms.companies, *model)
	return nil
}

func (cms CompanyMemoryStore) List() ([]CompanyModel, error) {
	if len(cms.companies) == 0 {
		return cms.companies, errors.New("No companies found")
	}
	return cms.companies, nil
}

func (cms CompanyMemoryStore) Find(id string) (*CompanyModel, error) {
	for _, record := range cms.companies {
		if record.ID.String() == id {
			return &record, nil
		}
	}
	return nil, errors.New("Company not found")
}

func (cms *CompanyMemoryStore) Delete(id string) error {
	for index, record := range cms.companies {
		if record.ID.String() == id {
			cms.companies[index] = cms.companies[len(cms.companies)-1]
			cms.companies = cms.companies[:len(cms.companies)-1]
			return nil
		}
	}
	return errors.New("Company not found")
}

type ProductMemoryStore struct {
	products []ProductModel
}

func (pms *ProductMemoryStore) Save(model *ProductModel) error {
	for index, record := range pms.products {
		if record.ID == model.ID {
			if &model.ProductName != nil {
				pms.products[index].ProductName = model.ProductName
			}
			if &model.Rating != nil {
				pms.products[index].Rating = model.Rating
			}
			model = &pms.products[index]
			return nil
		}
	}
	pms.products = append(pms.products, *model)
	return nil
}

func (pms ProductMemoryStore) List() ([]ProductModel, error) {
	if len(pms.products) == 0 {
		return pms.products, errors.New("No products found")
	}
	return pms.products, nil
}

func (pms ProductMemoryStore) Find(id string) (*ProductModel, error) {
	for _, record := range pms.products {
		if record.ID.String() == id {
			return &record, nil
		}
	}
	return nil, errors.New("Product not found")
}

func (pms *ProductMemoryStore) Delete(id string) error {
	for index, record := range pms.products {
		if record.ID.String() == id {
			pms.products[index] = pms.products[len(pms.products)-1]
			pms.products = pms.products[:len(pms.products)-1]
			return nil
		}
	}
	return errors.New("Product not found")
}

type ReviewMemoryStore struct {
	reviews []ReviewModel
}

func (rms *ReviewMemoryStore) Save(model *ReviewModel) error {
	for index, record := range rms.reviews {
		if record.ID == model.ID {
			if &model.Comment != nil {
				rms.reviews[index].Comment = model.Comment
			}
			if &model.Rating != nil {
				rms.reviews[index].Rating = model.Rating
			}
			model = &rms.reviews[index]
			return nil
		}
	}
	rms.reviews = append(rms.reviews, *model)
	return nil
}

func (rms ReviewMemoryStore) List() ([]ReviewModel, error) {
	if len(rms.reviews) == 0 {
		return rms.reviews, errors.New("No reviews found")
	}
	return rms.reviews, nil
}

func (rms ReviewMemoryStore) Find(id string) (*ReviewModel, error) {
	for _, record := range rms.reviews {
		if record.ID.String() == id {
			return &record, nil
		}
	}
	return nil, errors.New("Review not found")
}

func (rms *ReviewMemoryStore) Delete(id string) error {
	for index, record := range rms.reviews {
		if record.ID.String() == id {
			rms.reviews[index] = rms.reviews[len(rms.reviews)-1]
			rms.reviews = rms.reviews[:len(rms.reviews)-1]
			return nil
		}
	}
	return errors.New("Review not found")
}
