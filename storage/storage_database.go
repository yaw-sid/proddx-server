package storage

import (
	"context"

	"github.com/jackc/pgx/v4"
	uuid "github.com/satori/go.uuid"
)

type UserDatabase struct {
	Conn *pgx.Conn
}

func (ub UserDatabase) Save(model *UserModel) error {
	_, err := ub.Conn.Exec(context.Background(), "insert into users(id, email, user_password, created_at) values($1, $2, $3, $4)",
		model.ID, model.Email, model.UserPassword, model.CreatedAt)
	return err
}

func (ub UserDatabase) Find(id string) (*UserModel, error) {
	row := ub.Conn.QueryRow(context.Background(), "select * from users where id=$1 or email=$2", uuid.FromStringOrNil(id), id)
	var model UserModel
	err := row.Scan(&model.ID, &model.Email, &model.UserPassword, &model.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (ub UserDatabase) Delete(id string) error {
	_, err := ub.Conn.Exec(context.Background(), "delete from users where id=$1", uuid.FromStringOrNil(id))
	return err
}

type CompanyDatabase struct {
	Conn *pgx.Conn
}

func (cb CompanyDatabase) Save(model *CompanyModel) error {
	row := cb.Conn.QueryRow(context.Background(), "select * from companies where id=$1", model.ID)
	if row.Scan() == pgx.ErrNoRows {
		_, err := cb.Conn.Exec(context.Background(), "insert into companies(id, company_user_id, company_name, email, logo, created_at) values($1, $2, $3, $4, $5, $6)",
			model.ID, model.CompanyUserID, model.CompanyName, model.Email, model.Logo, model.CreatedAt)
		return err
	}
	_, err := cb.Conn.Exec(context.Background(), "update companies set company_name=$2, email=$3, logo=$4 where id=$1", model.ID, model.CompanyName, model.Email, model.Logo)
	return err
}

func (cb CompanyDatabase) List() ([]CompanyModel, error) {
	rows, err := cb.Conn.Query(context.Background(), "select * from companies")
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
	row := cb.Conn.QueryRow(context.Background(), "select * from companies where id=$1 or company_user_id=$2", uuid.FromStringOrNil(id), id)
	var model CompanyModel
	err := row.Scan(&model.ID, &model.CompanyUserID, &model.CompanyName, &model.Email, &model.Logo, &model.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (cb CompanyDatabase) Delete(id string) error {
	_, err := cb.Conn.Exec(context.Background(), "delete from companies where id=$1", uuid.FromStringOrNil(id))
	return err
}

type ProductDatabase struct {
	Conn *pgx.Conn
}

func (cb ProductDatabase) Save(model *ProductModel) error {
	row := cb.Conn.QueryRow(context.Background(), "select * from products where id=$1", model.ID)
	if row.Scan() == pgx.ErrNoRows {
		_, err := cb.Conn.Exec(context.Background(), "insert into products(id, company_id, product_name, feedback_url, rating, created_at) values($1, $2, $3, $4, $5, $6)",
			model.ID, model.CompanyID, model.ProductName, model.FeedbackURL, model.Rating, model.CreatedAt)
		return err
	}
	_, err := cb.Conn.Exec(context.Background(), "update products set product_name=$2, feedback_url=$3, rating=$4 where id=$1", model.ID, model.ProductName, model.FeedbackURL, model.Rating)
	return err
}

func (cb ProductDatabase) List() ([]ProductModel, error) {
	rows, err := cb.Conn.Query(context.Background(), "select * from products")
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
	row := cb.Conn.QueryRow(context.Background(), "select * from products where id=$1", uuid.FromStringOrNil(id))
	var model ProductModel
	err := row.Scan(&model.ID, &model.CompanyID, &model.ProductName, &model.FeedbackURL, &model.Rating, &model.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (cb ProductDatabase) Delete(id string) error {
	_, err := cb.Conn.Exec(context.Background(), "delete from products where id=$1", uuid.FromStringOrNil(id))
	return err
}

type ReviewDatabase struct {
	Conn *pgx.Conn
}

func (cb ReviewDatabase) Save(model *ReviewModel) error {
	row := cb.Conn.QueryRow(context.Background(), "select * from reviews where id=$1", model.ID)
	if row.Scan() == pgx.ErrNoRows {
		_, err := cb.Conn.Exec(context.Background(), "insert into reviews(id, company_id, product_id, comment, rating, created_at) values($1, $2, $3, $4, $5, $6)",
			model.ID, model.CompanyID, model.ProductID, model.Comment, model.Rating, model.CreatedAt)
		return err
	}
	_, err := cb.Conn.Exec(context.Background(), "update reviews set comment=$2, rating=$3 where id=$1", model.ID, model.Comment, model.Rating)
	return err
}

func (cb ReviewDatabase) List() ([]ReviewModel, error) {
	rows, err := cb.Conn.Query(context.Background(), "select * from reviews")
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
	row := cb.Conn.QueryRow(context.Background(), "select * from reviews where id=$1", uuid.FromStringOrNil(id))
	var model ReviewModel
	err := row.Scan(&model.ID, &model.CompanyID, &model.ProductID, &model.Comment, &model.Rating, &model.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (cb ReviewDatabase) Delete(id string) error {
	_, err := cb.Conn.Exec(context.Background(), "delete from reviews where id=$1", uuid.FromStringOrNil(id))
	return err
}
