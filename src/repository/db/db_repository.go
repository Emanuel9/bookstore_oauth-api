package db

import (
	"github.com/Emanuel9/bookstore_oauth-api/src/domain/access_token"
	"github.com/Emanuel9/bookstore_oauth-api/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestError) {
	return nil, errors.NewInternalServerError("database connection not implemeneted yet")
}
