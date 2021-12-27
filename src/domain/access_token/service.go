package access_token

import (
	"strings"

	"github.com/Delaram-Gholampoor-Sagha/bookstore_oauth_api/src/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	// look for this access token in a given database
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(AccessTokenId string) (*AccessToken, *errors.RestErr) {
	accessTokenId := strings.TrimSpace(AccessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewIntervalServerError("invalid access token id ")
	}
	AccessToken, err := s.repository.GetById(AccessTokenId)
	if err != nil {
		return nil, err
	}
	return AccessToken, nil
}

func (s *service) Create(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
