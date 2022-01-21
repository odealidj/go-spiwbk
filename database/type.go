package database

import (
	"fmt"

	"gorm.io/driver/mysql"
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

type dbMySQL struct {
	db
	Charset   string
	ParseTime string
	Loc       string
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

	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=%s", "djptkkps_heroku", "dUZCa8m5QFKVTMy", "49.128.186.146", "djptkkps_heroku", c.ParseTime)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
		Logger:                                   logger.Default.LogMode(logger.Info),
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              true,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
