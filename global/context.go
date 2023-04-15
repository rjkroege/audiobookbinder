package global

import (
	"database/sql"
)

type Context struct {
	Debug bool
	Dbname string
	Db *sql.DB
}
