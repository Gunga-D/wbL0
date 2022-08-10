package postgres

import (
	"fmt"

	"github.com/Gunga-D/taskL0/internal/config"
	"github.com/jmoiron/sqlx"
)

func CreateKernel(config *config.DBConfig) (*sqlx.DB, error) {
	kernel, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		"localhost", "5432", config.User, config.Database, config.Password, "disable"))
	if err != nil {
		return nil, err
	}

	err = kernel.Ping()
	if err != nil {
		return nil, err
	}

	return kernel, nil
}
