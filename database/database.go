package database

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	dbConnections    map[string]*gorm.DB
	dbConnectionName string
)

func Init() {
	dbConfigurations := map[string]Db{
		strings.ToUpper(dbConnectionName): &dbMySQL{
			db: db{
				Host: os.Getenv("DB_HOST"),
				User: os.Getenv("DB_USER"),
				Pass: os.Getenv("DB_PASS"),
				Port: os.Getenv("DB_PORT"),
				Name: os.Getenv("DB_NAME"),
			},
			//SslMode: os.Getenv("DB_SSLMODE"),
			//Tz:      os.Getenv("DB_TZ"),
			Charset:   os.Getenv("DB_Charset"),
			ParseTime: os.Getenv("DB_ParseTime"),
			Loc:       os.Getenv("DB_LOC"),
		},
		strings.ToUpper(os.Getenv("DB_NAME_MIGRATION")): &dbMySQL{
			db: db{
				Host: os.Getenv("DB_HOST"),
				User: os.Getenv("DB_USER"),
				Pass: os.Getenv("DB_PASS"),
				Port: os.Getenv("DB_PORT"),
				Name: os.Getenv("DB_NAME_MIGRATION"),
			},
			//SslMode: os.Getenv("DB_SSLMODE"),
			//Tz:      os.Getenv("DB_TZ"),
			Charset:   os.Getenv("DB_Charset"),
			ParseTime: os.Getenv("DB_ParseTime"),
			Loc:       os.Getenv("DB_LOC"),
		},
	}

	dbConnections = make(map[string]*gorm.DB)
	for k, v := range dbConfigurations {
		db, err := v.Init()
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to database %s", k))
		}
		dbConnections[k] = db
		logrus.Info(fmt.Sprintf("Successfully connected to database %s", k))
	}
}

func Connection(name string) (*gorm.DB, error) {
	if dbConnections[strings.ToUpper(name)] == nil {
		return nil, errors.New("Connection is undefined")
	}
	return dbConnections[strings.ToUpper(name)], nil
}
