package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConsatnt(t *testing.T) {
	// if expirationTime != 24 {
	// 	t.Error("expiration time should be 24 hour")
	// }
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hour")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken(AccessToken)
	if at.IsExpired() {
		t.Error("brand new access token should not be wxpired")

	}
	if at.AccessToken != "" {
		t.Error("new acces token should not have defined access token id")
	}
	if at.UserId != 0 {
		t.Error("new access toke should not have any associated user id")
	}
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	if !at.IsExpired() {
		t.Error("empty access token should be expired by default")
	}
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if !at.IsExpired() {
		t.Error("access token expiring 3 hours from now should not be wexpired")
	}
}
