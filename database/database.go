package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
)

var (
	dbConnections    map[string]*gorm.DB
	dbConnectionName string
)

func Init() {
	dbConfigurations := map[string]Db{
		os.Getenv("DB_CONN_NAME_DJPT_SPIWBK"): &dbMySQL{
			db: db{
				Host: "49.128.186.146",  // os.Getenv("DB_HOST"),
				User: "djptkkps_heroku", //os.Getenv("DB_USER"),
				Pass: "dUZCa8m5QFKVTMy", // os.Getenv("DB_PASS"),
				Port: "3306",            //os.Getenv("DB_PORT"),
				Name: "djptkkps_heroku", //os.Getenv("DB_NAME"),
			},
			//SslMode: os.Getenv("DB_SSLMODE"),
			//Tz:      os.Getenv("DB_TZ"),
			Charset:   "utf8mb4", //os.Getenv("DB_Charset"),
			ParseTime: "True",    // os.Getenv("DB_ParseTime"),
			Loc:       "Local",   //os.Getenv("DB_LOC"),
		},

		/*
			os.Getenv("DB_CONN_NAME_DJPT_SPIWBK_MIG"): &dbMySQL{
				db: db{
					Host: os.Getenv("DB_HOST_MIG"),
					User: os.Getenv("DB_USER_MIG"),
					Pass: os.Getenv("DB_PASS_MIG"),
					Port: os.Getenv("DB_PORT_MIG"),
					Name: os.Getenv("DB_NAME_MIG"),
				},
				//SslMode: os.Getenv("DB_SSLMODE"),
				//Tz:      os.Getenv("DB_TZ"),
				Charset:   os.Getenv("DB_Charset"),
				ParseTime: os.Getenv("DB_ParseTime"),
				Loc:       os.Getenv("DB_LOC"),
			},
		*/
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
	if dbConnections[name] == nil {
		return nil, errors.New("Connection is undefined")
	}
	return dbConnections[name], nil
}

func Close() {
	var sqlDB *sql.DB
	var err error

	for key, conn := range dbConnections {
		if sqlDB, err = conn.DB(); err == nil {
			err = sqlDB.Close()
		}
		if err != nil {
			logrus.WithField("message", "failed to close database connection "+key).Error(err.Error())
		} else {
			logrus.Infof("Connection to %v closed", key)
		}
	}
}
