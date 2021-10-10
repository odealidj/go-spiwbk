package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Db interface {
	Init() (*gorm.DB, error)
}

type db struct {
	Host string
	User string
	Pass string
	Port string
	Name string
}

type dbPostgreSQL struct {
	db
	SslMode string
	Tz      string
}

type dbMySQL struct {
	db
	Charset   string
	ParseTime string
	Loc       string
}

func (c *dbPostgreSQL) Init() (*gorm.DB, error) {
	//heroku sslmode=require
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", c.Host, c.User, c.Pass, c.Name, c.Port, c.SslMode, c.Tz)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
		Logger:                                   logger.Default.LogMode(logger.Info),
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (c *dbMySQL) Init() (*gorm.DB, error) {
	//Heroku
	//benar
	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", c.User, c.Pass, c.Host, c.Name)
	//mysql://b34b5824ee06f6:923f1590@us-cdbr-east-04.cleardb.com/heroku_08b64deef14c329?reconnect=true
	//MyLocal
	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=%s", c.User, c.Pass, c.Host, c.Name, c.Charset, c.ParseTime, c.Loc)
	//Heroku sukses
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=%s", c.User, c.Pass, c.Host, c.Name, c.ParseTime)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
		Logger:                                   logger.Default.LogMode(logger.Info),
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
