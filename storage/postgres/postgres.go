package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Mubinabd/RestaurantService/config"
	"github.com/Mubinabd/RestaurantService/storage"
)

type Storage struct {
	db          *sql.DB
	OrderS storage.OrdersI
}

func ConnectDB() (*Storage, error) {
	cfg := config.Load()
	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	orderS := NewOrderStorage(db)
	return &Storage{
		db:    db,
		OrderS: orderS,
	}, nil
}
func (s *Storage) Orders() storage.OrdersI {
	if s. OrderS== nil {
		s.OrderS = NewOrderStorage(s.db)
	}
	return s.OrderS
}
