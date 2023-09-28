package mysql_utils

import (
	"book/utils/errors"
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("no record matching with given id")
		}
		return errors.NewInternalServerError("error parsing db response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("email already exists")
	}
	return errors.NewInternalServerError("error processing request")
}
