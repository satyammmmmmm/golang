package services

import (
	"src/api/config"
	"src/api/domain/github"
	"src/api/domain/provider/github_provider"
	"src/api/domain/repositories"
	"src/api/utils/errors"
	"strings"
)

type repoService struct {
}

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}

}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repo name")
	}
	request := github.CreateRepoRequest{
		Name:        input.name,
		Description: input.Description,
		Private:     false,
	}
	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}
	result := repositories.CreateRepoResponse{
		Id:    response.id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil

}
func (s *repoService) CreateRepos(input []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {

}
