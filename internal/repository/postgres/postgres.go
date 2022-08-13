package postgres

import (
	"errors"
	"fmt"

	"github.com/Gunga-D/taskL0/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func CreateKernel(config *config.DBConfig) *sqlx.DB {
	connectionInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.User, config.DbName, config.Password, config.SSLMode)

	var kernel *sqlx.DB
	err := errors.New("Not connected to db")
	for err != nil {
		kernel, err = sqlx.Connect("postgres", connectionInfo)
		logrus.Info("Trying to connect to db...")
	}

	return kernel
}
