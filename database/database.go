package database

import (
	"code-boiler/internal/model"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbEnv struct {
	DB_HOST    string
	DB_USER    string
	DB_PASS    string
	DB_PORT    string
	DB_NAME    string
	DB_SSLMODE string
	DB_TZ      string
	models     []interface{}
}

var (
	DBConnections map[string]*gorm.DB
)

func Connect() (map[string]*gorm.DB, error) {
	DBConnections = make(map[string]*gorm.DB)
	databases := map[string]DbEnv{
		// "SYS_AUTH": {
		// 	DB_HOST:    os.Getenv("DB_HOST_SYS_AUTH"),
		// 	DB_USER:    os.Getenv("DB_USER_SYS_AUTH"),
		// 	DB_PASS:    os.Getenv("DB_PASS_SYS_AUTH"),
		// 	DB_PORT:    os.Getenv("DB_PORT_SYS_AUTH"),
		// 	DB_NAME:    os.Getenv("DB_NAME_SYS_AUTH"),
		// 	DB_SSLMODE: os.Getenv("DB_SSLMODE_SYS_AUTH"),
		// 	DB_TZ:      os.Getenv("DB_TZ_SYS_AUTH"),
		// 	models:     []interface{}{
		// 		// &model.User{},
		// 	},
		// },
		"SAMPLE2": {
			DB_HOST:    os.Getenv("DB_HOST_SAMPLE"),
			DB_USER:    os.Getenv("DB_USER_SAMPLE"),
			DB_PASS:    os.Getenv("DB_PASS_SAMPLE"),
			DB_PORT:    os.Getenv("DB_PORT_SAMPLE"),
			DB_NAME:    os.Getenv("DB_NAME_SAMPLE"),
			DB_SSLMODE: os.Getenv("DB_SSLMODE_SAMPLE"),
			DB_TZ:      os.Getenv("DB_TZ_SAMPLE"),
			models: []interface{}{
				&model.Sample{},
				&model.User{},
			},
		},
	}
	for key, item := range databases {
		db, err := getDBConnection(item.DB_HOST, item.DB_USER, item.DB_PASS, item.DB_NAME, item.DB_PORT, item.DB_SSLMODE, item.DB_TZ)
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to database %s", key))
		}
		DBConnections[key] = db
		if item.models != nil {
			db.AutoMigrate(item.models...)
		}
		logrus.Info(fmt.Sprintf("Successfully connected to database %s", item.DB_NAME))
	}
	return DBConnections, nil
}

func getDBConnection(dbHost, dbUser, dbPass, dbName, dbPort, sslMode, tz string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", dbHost, dbUser, dbPass, dbName, dbPort, sslMode, tz)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate()
	return db, nil
}

func DBManager(name string) (*gorm.DB, error) {
	if DBConnections[strings.ToUpper(name)] == nil {
		return nil, errors.New("connection is undefined")
	}
	return DBConnections[strings.ToUpper(name)], nil
}
