package access_token

import (
	"time"
)

const (
	expirationTime = 24
)

// this entire api is going to provide access_token to outside world

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	// whats the application the userid is having  for creating this access token ?
	ClientId int64 `json:"client_id"`
	Expires  int64 `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
