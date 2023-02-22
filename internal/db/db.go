package db

import "database/sql"

type Database struct {
	Client *sql.DB
}
