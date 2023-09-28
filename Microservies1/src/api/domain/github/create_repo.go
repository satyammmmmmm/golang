package github

import (
	"src/api/utils/errors"
)

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}
type CreateReposResponse struct {
	StatusCode int                        `json:"status"`
	Results    []CreateRepositoriesResult `json:"results"`
	Response   CreateRepoResponse         `json:"result"`
	Error      errors.ApiError            `json:"error"`
}
type CreateRepositoriesResult struct {
	Response CreateRepoResponse `json:"repo"`
	Error    errors.ApiError    `json:"error"`
}

type CreateRepoResponse struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	FullName string    `json:"full_name"`
	Owner    RepoOwner `json:"owner"`
}

type RepoOwner struct {
	Id      int64  `json:"id"`
	Login   string `json:"login"`
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
}
type RepoPermissions struct {
	IsAdmin bool `json:"admin"`
	HasPull bool `json:"push"`
	Hsspush bool `json:"pull"`
}
