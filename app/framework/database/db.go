package database

import (
	"database/sql"
	"fmt"
	config "minotauro/app/configuracao"

	_ "github.com/lib/pq"
)

var (
	host      string
	user      string
	password  string
	dbname    string
	port      int64
	pgsqlConn string
)

func Conn() (*sql.DB, error) {

	host = config.EnvLocalHostDb()
	user = config.EnvUsrDb()
	password = config.EnvSqlPassword()
	dbname = config.EnvLocalDb()
	port = 5432

	if config.EnvDev() {
		pgsqlConn = fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable", host, port, user, password, dbname)
	}

	db, err := sql.Open("postgres", pgsqlConn)

	if err != nil {
		panic(err)
	}

	if err != nil {
		return nil, err
	}
	return db, nil
}
