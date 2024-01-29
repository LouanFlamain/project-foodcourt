package model

import "database/sql"

type Store struct {
	*sql.DB
}
