package github_provider

import (
	"src/domain/github"
	"testing"

	"github.com/stretchr/testify/assert"
)




func TestMain(m *testing.M){
	restclient.StartMock()
	os.Exit(m.Run())

}



func TestGetAuthorizatonHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}

func TestCreatRepoErrorRestClient(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMock(restclient.Mock{
		Url: "https://api.github.com/user/repos"
		HttpMethod: http.MethodPost, 
		Err : erros.New("invaliad response")
	})
	restclient.StopMock()
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t,"invvalid resclient response",err.Message)
}
func TestCreatRepoInvalidResponseBody(t *testing.T) {
	restclient.FlushMocks()
	invalidCloser,_:=os.Open("-asf3")
	restclient.AddMock(restclient.Mock{
		Url: "https://api.github.com/user/repos"
		HttpMethod: http.MethodPost, 
		Response: &http.Response{
			StatusCode:http.StatusCreated,
			Body:invalidCloser,

		},
		
	})
	restclient.StopMock()
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t,http.StatusInternalServerError,err.StatusCode)
	assert.EqualValues(t,"invvalid resclient response",err.Message)
}

func TestCreatRepoInvalidErrorInterface(t *testing.T) {
	restclient.FlushMocks()
	invalidCloser,_:=os.Open("-asf3")
	restclient.AddMock(restclient.Mock{
		Url: "https://api.github.com/user/repos"
		HttpMethod: http.MethodPost, 
		Response: &http.Response{
			StatusCode:http.StatusUnauthorized,
			Body:ioutil.NopCloser(strings.NewReader(`{"message":1}`)),

		},
		
	})
	restclient.StopMock()
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t,http.StatusInternalServerError,err.StatusCode)
	assert.EqualValues(t,"invvalid resclient response",err.Message)
}


