package rest

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("about tp strat test cases")
	rest.StartMockupServer()
	os.Exit(m.Run())

}

func TestLoginUSerTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "localhost:8080/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"satyamgupta40@gmail.com","password":"micro1"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "password")
	assert.Nil(t, user)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)

}
func TestLoginUSerInvalidInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "localhost:8080/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"satyamgupta40@gmail.com","password":"micro1"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "password")
	assert.Nil(t, user)
	assert.EqualValues(t, http.StatusNotFound, err.Status)

}
func TestLoginUSerInvalidLoginCredential(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "localhost:8080/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"satyamgupta40@gmail.com","password":"micro1"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "password")
	assert.Nil(t, user)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
}
func TestLoginUSerInvalidUSrREsponse(t *testing.T) {

}

func TestLoginUSerNoError(t *testing.T) {

}
