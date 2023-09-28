package users_db

import (
	"database/sql"
	"fmt"
	"log"
	_ "os"

	_ "github.com/go-sql-driver/mysql"
)

// const (
// 	mysql_users_username = "mysql_users_username"
// 	mysql_users_password = "mysql_users_password"
// 	mysql_users_host     = "mysql_users_host"
// 	mysql_users_schema   = "mysql_users_schema"
// )

var (
	Client *sql.DB
	// username = os.Getenv(mysql_users_username)
	// password = os.Getenv(mysql_users_password)
	// host     = os.Getenv(mysql_users_host)
	// schema   = os.Getenv(mysql_users_schema)
)

func init() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", "root", "root", "localhost:3306", "user_db")
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	//mysql.SetLogger()
	log.Println("database successfully configure")

}
