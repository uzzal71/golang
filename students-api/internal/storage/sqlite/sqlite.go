package sqlite

import "database/sql"

type Sqlite struct {
	Db *sql.DB
}
