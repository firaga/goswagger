package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// ConnectDBSqlx ConnectDB opens a connection to the database
func ConnectDBSqlx() *sqlx.DB {
	dsn := "root:123456@tcp(127.0.0.1:3308)/company"
	// 也可以使用MustConnect连接不成功就panic
	db, err := sqlx.Connect("mysql", dsn)
	//db, err := sqlx.Open(constants.DBTYPE, constants.DBUSERNAME+":"+constants.DBPASSWORD+"@/"+constants.DBNAME)
	if err != nil {
		panic(err.Error())
	}

	return db
}
