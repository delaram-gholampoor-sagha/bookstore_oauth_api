package db

import (
	"github.com/Delaram-Gholampoor-Sagha/bookstore_oauth_api/src/domain/access_token"
	"github.com/Delaram-Gholampoor-Sagha/bookstore_oauth_api/src/utils/errors"
)

func NewRepository() DBRepository {
	return &dbRepository{} 
}

type dbRepository struct {
}

type DBRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

func (r *dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewIntervalServerError("database connection not implemented yet")
}
