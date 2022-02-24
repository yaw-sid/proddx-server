package storage

type Company interface {
	Save(*CompanyModel) error
	List() ([]CompanyModel, error)
	Find(string) (*CompanyModel, error)
	Delete(string) error
}

type Product interface {
	Save(*ProductModel) error
	List() ([]ProductModel, error)
	Find(string) (*ProductModel, error)
	Delete(string) error
}

type Review interface {
	Save(*ReviewModel) error
	List() ([]ReviewModel, error)
	Find(string) (*ReviewModel, error)
	Delete(string) error
}
