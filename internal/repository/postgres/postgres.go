package postgres

import (
	"fmt"

	"github.com/Gunga-D/taskL0/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func CreateKernel(config *config.DBConfig) (*sqlx.DB, error) {
	connectionInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.User, config.DbName, config.Password, config.SSLMode)
	kernel, err := sqlx.Connect("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}

	return kernel, nil
}
