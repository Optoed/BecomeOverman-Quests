package repositories

import (
	"errors"
	"github.com/jmoiron/sqlx"
)

type BaseRepository struct {
	db *sqlx.DB
}

func NewBaseRepository(db *sqlx.DB) *BaseRepository {
	return &BaseRepository{db: db}
}

func (r *BaseRepository) CheckConnection() error {
	if err := r.db.Ping(); err != nil {
		return errors.New("failed to connect to the database")
	}
	return nil
}
