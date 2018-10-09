package dao

import "database/sql"

// Dao database connection
type Dao struct {
	Driver     string
	Dsn        string
	Connection *sql.DB
}

// daoBuffer dao buffer
type daoBuffer struct {
	tx   *sql.Tx
	stmt *sql.Stmt
	res  sql.Result
}
