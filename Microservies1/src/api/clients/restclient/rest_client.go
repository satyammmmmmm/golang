package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

//	func Something() {
//		Post("https://api.github.com/user/repos", `{"name":"golang-tutorial"}`)
//	}
var (
	enabledMocks = false
	mocks        = make(map[string]*Mock)
)

type Mock struct {
	Url        string
	HttpMethod string
	response   *http.Response
	Err        error
}

func GetMockId(httpMethod string, url string) string {
	return fmt.Sprintf("%s %s", httpMethod, url)
}

func StartMock() {
	enabledMocks = true

}
func FlushMocks() {
	mocks = make(map[string]*Mock)
}

func StopMock() {
	enabledMocks = false
}

func AddMock(mock Mock) {
	mocks[GetMockId(mock.HttpMethod, mock.Url)] = &mock
}

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	if enabledMocks {
		//return local mocks
		mock := mocks[GetMockId(http.MethodPost, url)]
		if mock == nil {
			return nil, errors.New("no mockup found")
		}
		return mock.response, mock.Err

	}
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers
	return client.Do(request)
}
