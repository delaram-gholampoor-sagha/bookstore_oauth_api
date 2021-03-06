package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/Delaram-Gholampoor-Sagha/bookstore_oauth_api/src/utils/crypto_utils"
	"github.com/Delaram-Gholampoor-Sagha/bookstore_oauth_api/src/utils/errors"
)

// this entire api is going to provide access_token to outside world

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grandTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() *errors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		break

	case grandTypeClientCredentials:
		break

	default:
		return errors.NewBadRequessrError("invalid grant_type parameter")
	}

	//TODO: Validate parameters for each grant_type
	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	// whats the application the userid is having  for creating this access token ?
	ClientId int64 `json:"client_id"`
	Expires  int64 `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)

	if at.AccessToken == "" {
		return errors.NewBadRequessrError("invalid access token id")
	}

	if at.UserId <= 0 {
		return errors.NewBadRequessrError("invalid user id")
	}

	if at.ClientId <= 0 {
		return errors.NewBadRequessrError("invalid client id")
	}

	if at.Expires <= 0 {
		return errors.NewBadRequessrError("invalid expiration time")
	}

	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}
func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
