package config

import (
	"github.com/jmoiron/sqlx"
	"goswagger/constants"
)

// ConnectDBSqlx ConnectDB opens a connection to the database
func ConnectDBSqlx() *sqlx.DB {
	db, err := sqlx.Open(constants.DBTYPE, constants.DBUSERNAME+":"+constants.DBPASSWORD+"@/"+constants.DBNAME)
	if err != nil {
		panic(err.Error())
	}

	return db
}