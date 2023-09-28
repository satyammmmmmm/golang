package users

import (
	"book/datasources/mysql/users_db"
	"book/logger"
	"book/utils/errors"
	"book/utils/errors/date_utils"
	"book/utils/errors/mysql_utils"
	"strings"
)

const (
	queryUpdateUser             = "UPDATE users SET first_name=?,last_name=?,email=? WHERE id=?;"
	queryInsertUser             = "INSERT INTO users(first_name,last_name,email,date_created,status,password) VALUES(?,?,?,?,?,?);"
	queryGetUser                = "SELECT id,first_name,last_name,email,date_created,status from users where id = ?;"
	queryDeleteUser             = "DELETE FROM users where id=?;"
	queryFindUSerByStatus       = "SELECT id,first_name,last_name,email,date_created FROM users where status=?;"
	queryFindByEmailAndPassword = "SELECT id,first_name,last_name,email,date_created,status FROM users where email=? AND password=? AND status=? "
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statemt", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		//return errors.NewInternalServerError(getErr.Error())
		logger.Error("error when trying to  get user by id", getErr)
		return errors.NewInternalServerError("database error")
	}

	return nil

}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare user statemt for save", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)

	if saveErr != nil {
		logger.Error("error when trying to save user ", saveErr)
		return errors.NewInternalServerError("database error")
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after user creation ", err)
		return errors.NewInternalServerError("database error")
	}
	user.Id = userId
	return nil
}
func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to get delete user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id)
	if err != nil {
		logger.Error("error when trying to delete user ", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUSerByStatus)
	if err != nil {
		logger.Error("error when trying to execute find user statement ", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)

	if err != nil {
		logger.Error("error when trying to find user ", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when trying to find next user ", err)
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, user)

	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError("no user matching this status")
	}

	return results, nil

}
func (user *User) FindByEmailAndPassword() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare get user by email and passwordstatemt", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		//return errors.NewInternalServerError(getErr.Error())
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return errors.NewNotFoundError("no user found/invalid user crdentials")
		}
		logger.Error("error when trying to  get user by email and password", getErr)
		return errors.NewInternalServerError("database error")
	}

	return nil

}
