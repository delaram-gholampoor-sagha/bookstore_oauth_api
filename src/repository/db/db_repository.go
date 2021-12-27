package db

import (
	"github.com/Delaram-Gholampoor-Sagha/bookstore_oauth_api/src/clients/cassandra"
	"github.com/Delaram-Gholampoor-Sagha/bookstore_oauth_api/src/domain/access_token"
	"github.com/Delaram-Gholampoor-Sagha/bookstore_oauth_api/src/utils/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token , user_id , client_id , expires 	FROM access_tokens WHERE access_token = ? ;"
	queryCreateAccessToken = "INSER INTO access_token(access_token , user_id , client_id , expires) VALUES ( ? , ? , ? , ?) ;"
	queryUpdateExpires     = "UPDATE access_token SET expires = ? WHERE access_token = ? ;"
)

func NewRepository() DBRepository {
	return &dbRepository{}
}

type dbRepository struct {
}

type DBRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {

	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found with the given id")
		}
		return nil, errors.NewIntervalServerError(err.Error())
	}

	return &result, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {

	if err := cassandra.GetSession().Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.ClientId, at.Expires).Exec(); err != nil {
		return errors.NewIntervalServerError(err.Error())
	}

	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {

	if err := cassandra.GetSession().Query(queryUpdateExpires, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.NewIntervalServerError(err.Error())
	}

	return nil
}
