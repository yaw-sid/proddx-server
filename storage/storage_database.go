package storage

import (
	"context"
	"strconv"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	uuid "github.com/satori/go.uuid"
)

type UserDatabase struct {
	Pool *pgxpool.Pool
}

func (ub UserDatabase) Save(model *UserModel) error {
	_, err := ub.Pool.Exec(context.Background(), "insert into users(id, email, user_password, created_at) values($1, $2, $3, $4)",
		model.ID, model.Email, model.UserPassword, model.CreatedAt)
	return err
}

func (ub UserDatabase) Find(id string) (*UserModel, error) {
	row := ub.Pool.QueryRow(context.Background(), "select * from users where id=$1 or email=$2", uuid.FromStringOrNil(id), id)
	var model UserModel
	err := row.Scan(&model.ID, &model.Email, &model.UserPassword, &model.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (ub UserDatabase) Delete(id string) error {
	_, err := ub.Pool.Exec(context.Background(), "delete from users where id=$1", uuid.FromStringOrNil(id))
	return err
}

type CompanyDatabase struct {
	Pool *pgxpool.Pool
}

func (cb CompanyDatabase) Save(model *CompanyModel) error {
	var initRow CompanyModel
	row := cb.Pool.QueryRow(context.Background(), "select created_at from companies where id=$1", model.ID)
	err := row.Scan(&initRow.CreatedAt)
	if err == pgx.ErrNoRows {
		_, err := cb.Pool.Exec(context.Background(), "insert into companies(id, company_user_id, company_name, email, logo, created_at) values($1, $2, $3, $4, $5, $6)",
			model.ID, model.CompanyUserID, model.CompanyName, model.Email, model.Logo, model.CreatedAt)
		return err
	} else if err != nil {
		return err
	}
	_, err = cb.Pool.Exec(context.Background(), "update companies set company_name=$2, email=$3, logo=$4 where id=$1", model.ID, model.CompanyName, model.Email, model.Logo)
	if err != nil {
		return err
	}
	if &model.CreatedAt == nil {
		model.CreatedAt = initRow.CreatedAt
	}
	return nil
}

func (cb CompanyDatabase) List() ([]CompanyModel, error) {
	rows, err := cb.Pool.Query(context.Background(), "select * from companies")
	if err != nil {
		return nil, err
	}
	var models []CompanyModel
	for rows.Next() {
		var model CompanyModel
		err = rows.Scan(&model.ID, &model.CompanyUserID, &model.CompanyName, &model.Email, &model.Logo, &model.CreatedAt)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}
	return models, nil
}

func (cb CompanyDatabase) Find(id string) (*CompanyModel, error) {
	row := cb.Pool.QueryRow(context.Background(), "select * from companies where id=$1 or company_user_id=$2", uuid.FromStringOrNil(id), id)
	var model CompanyModel
	err := row.Scan(&model.ID, &model.CompanyUserID, &model.CompanyName, &model.Email, &model.Logo, &model.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (cb CompanyDatabase) Delete(id string) error {
	_, err := cb.Pool.Exec(context.Background(), "delete from companies where id=$1", uuid.FromStringOrNil(id))
	return err
}

type ProductDatabase struct {
	Pool *pgxpool.Pool
}

func (cb ProductDatabase) Save(model *ProductModel) error {
	var initRow ProductModel
	row := cb.Pool.QueryRow(context.Background(), "select rating, created_at from products where id=$1", model.ID)
	err := row.Scan(&initRow.Rating, &initRow.CreatedAt)
	if err == pgx.ErrNoRows {
		_, err := cb.Pool.Exec(context.Background(), "insert into products(id, company_id, product_name, feedback_url, rating, created_at) values($1, $2, $3, $4, $5, $6)",
			model.ID, model.CompanyID, model.ProductName, model.FeedbackURL, model.Rating, model.CreatedAt)
		return err
	} else if err != nil {
		return err
	}
	_, err = cb.Pool.Exec(context.Background(), "update products set product_name=$2, feedback_url=$3, rating=$4 where id=$1", model.ID, model.ProductName, model.FeedbackURL, model.Rating)
	if err != nil {
		return err
	}
	if model.Rating == 0 {
		model.Rating = initRow.Rating
	}
	if &model.CreatedAt == nil {
		model.CreatedAt = initRow.CreatedAt
	}
	return nil
}

func (cb ProductDatabase) List(companyID string) ([]ProductModel, error) {
	stmt := "select * from products"
	var args []interface{}
	if companyID != "" {
		stmt = stmt + " where company_id=$1"
		args = append(args, companyID)
	}
	rows, err := cb.Pool.Query(context.Background(), stmt, args...)
	if err != nil {
		return nil, err
	}
	var models []ProductModel
	for rows.Next() {
		var model ProductModel
		err = rows.Scan(&model.ID, &model.CompanyID, &model.ProductName, &model.FeedbackURL, &model.Rating, &model.CreatedAt)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}
	return models, nil
}

func (cb ProductDatabase) Find(id string) (*ProductModel, error) {
	row := cb.Pool.QueryRow(context.Background(), "select * from products where id=$1", uuid.FromStringOrNil(id))
	var model ProductModel
	err := row.Scan(&model.ID, &model.CompanyID, &model.ProductName, &model.FeedbackURL, &model.Rating, &model.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (cb ProductDatabase) Delete(id string) error {
	_, err := cb.Pool.Exec(context.Background(), "delete from products where id=$1", uuid.FromStringOrNil(id))
	return err
}

type ReviewDatabase struct {
	Pool *pgxpool.Pool
}

func (cb ReviewDatabase) Save(model *ReviewModel) error {
	var initRow ReviewModel
	row := cb.Pool.QueryRow(context.Background(), "select created_at from reviews where id=$1", model.ID)
	err := row.Scan(&initRow.CreatedAt)
	if err == pgx.ErrNoRows {
		_, err := cb.Pool.Exec(context.Background(), "insert into reviews(id, company_id, product_id, comment, rating, created_at) values($1, $2, $3, $4, $5, $6)",
			model.ID, model.CompanyID, model.ProductID, model.Comment, model.Rating, model.CreatedAt)
		return err
	} else if err != nil {
		return err
	}
	_, err = cb.Pool.Exec(context.Background(), "update reviews set comment=$2, rating=$3 where id=$1", model.ID, model.Comment, model.Rating)
	if err != nil {
		return err
	}
	if &model.CreatedAt == nil {
		model.CreatedAt = initRow.CreatedAt
	}
	return nil
}

func (cb ReviewDatabase) List(companyID string, productID string) ([]ReviewModel, error) {
	stmt := "select * from reviews where id is not null"
	var args []interface{}
	if companyID != "" {
		stmt = stmt + " and company_id=$" + strconv.Itoa(len(args)+1)
		args = append(args, companyID)
	}
	if productID != "" {
		stmt = stmt + " and product_id=$" + strconv.Itoa(len(args)+1)
		args = append(args, productID)
	}
	rows, err := cb.Pool.Query(context.Background(), stmt, args...)
	if err != nil {
		return nil, err
	}
	var models []ReviewModel
	for rows.Next() {
		var model ReviewModel
		err = rows.Scan(&model.ID, &model.CompanyID, &model.ProductID, &model.Comment, &model.Rating, &model.CreatedAt)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}
	return models, nil
}

func (cb ReviewDatabase) Find(id string) (*ReviewModel, error) {
	row := cb.Pool.QueryRow(context.Background(), "select * from reviews where id=$1", uuid.FromStringOrNil(id))
	var model ReviewModel
	err := row.Scan(&model.ID, &model.CompanyID, &model.ProductID, &model.Comment, &model.Rating, &model.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (cb ReviewDatabase) Delete(id string) error {
	_, err := cb.Pool.Exec(context.Background(), "delete from reviews where id=$1", uuid.FromStringOrNil(id))
	return err
}
