package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "we were not expecting user with id 0")
	assert.NotNil(t, err, "we were expectng error when userid is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	// if user != nil {
	// 	t.Error("we were not expecting user with id 0")
	// }
	// if err == nil {
	// 	t.Error("we were expectng error when userid is 0")
	// }
	// if err.StatusCode != http.StatusNotFound {
	// 	t.Error("we are expecting 404 when user not found")
	// }
}
func TestGetUserUserFound(t *testing.T) {
	user, err := GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)

}
