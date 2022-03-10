package storage

type User interface {
	Save(*UserModel) error
	Find(string) (*UserModel, error)
	Delete(string) error
}

type Company interface {
	Save(*CompanyModel) error
	List() ([]CompanyModel, error)
	Find(string) (*CompanyModel, error)
	Delete(string) error
}

type Product interface {
	Save(*ProductModel) error
	List(companyID string) ([]ProductModel, error)
	Find(string) (*ProductModel, error)
	Delete(string) error
}

type Review interface {
	Save(*ReviewModel) error
	List(companyID string, productID string) ([]ReviewModel, error)
	Find(string) (*ReviewModel, error)
	Delete(string) error
}
