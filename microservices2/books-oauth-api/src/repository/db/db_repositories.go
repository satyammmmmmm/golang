package db

import (
	"oauth/clients/cassandra"
	"oauth/domain/access_token"
	"oauth/utils/errors"

	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token,user_id,client_id,expires FROM access_token WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_token(access_token,user_id,client_id,expires) VALUES (?,?,?,?);"
	queryUpdateAccessToken = "UPDATE access_token SET expires=? WHERE access_token=?; "
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}
type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer session.Close()
	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found with given id ")
		}

		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}

func NewRepository() DbRepository {
	return &dbRepository{}
}
func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer session.Close()
	if err := session.Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.ClientId, at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer session.Close()
	if err := session.Query(queryUpdateAccessToken, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
