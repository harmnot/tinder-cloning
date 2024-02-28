package driver

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"log"
	"os"
	"sync"
	"time"
)

func init() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

var (
	sqlDBOnce sync.Once
	sqlDB     *sql.DB
)

func NewSQL(config *GlobalConf) *sql.DB {
	sqlDBOnce.Do(func() {

		connString := fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=disable",
			config.DB.User,
			config.DB.Password,
			config.DB.Host,
			config.DB.Port,
			config.DB.Name,
		)

		db, err := sql.Open(config.DB.Driver, connString)
		if err != nil {
			log.Fatal(err)
		}

		db.SetConnMaxLifetime(time.Duration(config.DB.ConMaxLifetime) * time.Second)
		db.SetMaxIdleConns(config.DB.ConMaxIdle)
		db.SetMaxOpenConns(config.DB.ConMaxOpen)

		sqlDB = db
	})

	return sqlDB
}

func loadEnvironment(filename string) error {
	var err error
	if filename != "" {
		err = godotenv.Overload(filename)
	} else {
		err = godotenv.Load()
		// handle if .env file does not exist, this is OK
		if os.IsNotExist(err) {
			return nil
		}
	}
	return err
}

func LoadGlobal(filename string) (*GlobalConf, error) {
	if err := loadEnvironment(filename); err != nil {
		return nil, err
	}

	config := new(GlobalConf)
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}
