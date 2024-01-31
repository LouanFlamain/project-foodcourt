package stores

import (
	"database/sql"
	//"foodcourt/app/model"
)

type FileStore struct {
	*sql.DB
}

func NewFileStore(db *sql.DB) *FileStore {
	return &FileStore{
		db,
	}
}

func (r *FileStore) GetImage(id int) (string, error) {
	//var item model.FileImageItem

	return "", nil
}