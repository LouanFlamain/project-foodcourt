package structure

import "database/sql"

type Store struct{
	*sql.DB
}