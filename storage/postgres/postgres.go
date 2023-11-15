package postgres

import (
	"database/sql"
	"fmt"

	"github.com/asadbekGo/market_system/config"
	"github.com/asadbekGo/market_system/storage"

	_ "github.com/lib/pq"
)

type Store struct {
	db       *sql.DB
	category storage.CategoryRepoI
	product  storage.ProductRepoI
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {

	connect := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	)

	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err)
	}

	return &Store{
		db: db,
	}, nil
}

func (s *Store) Category() storage.CategoryRepoI {

	if s.category == nil {
		s.category = NewCategoryRepo(s.db)
	}

	return s.category
}

func (s *Store) Product() storage.ProductRepoI {

	if s.product == nil {
		s.product = NewProductRepo(s.db)
	}

	return s.product
}
