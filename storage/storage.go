package storage

type Company interface {
	save(*CompanyModel) error
	list() ([]CompanyModel, error)
	find(string) (*CompanyModel, error)
	delete(string) error
}

type Product interface {
	save(*ProductModel) error
	list() ([]ProductModel, error)
	find(string) (*ProductModel, error)
	delete(string) error
}

type Review interface {
	save(*ReviewModel) error
	list() ([]ReviewModel, error)
	find(string) (*ReviewModel, error)
	delete(string) error
}
