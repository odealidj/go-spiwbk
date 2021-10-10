package factory

import (
	"codeid-boiler/database"
	"codeid-boiler/internal/app/auth/repository"

	"os"
	"strings"

	"gorm.io/gorm"
)

type Factory struct {
	Db               *gorm.DB
	AuthRepository    repository.Auth
}

func NewFactory() *Factory {
	f := &Factory{}
	f.SetupDb()
	f.SetupRepository()

	return f
}

func (f *Factory) SetupDb() {
	db, err := database.Connection(strings.ToUpper(os.Getenv("DB_NAME_MIGRATION")))
	if err != nil {
		panic("Failed setup db, connection is undefined")
	}
	f.Db = db
}

func (f *Factory) SetupRepository() {
	if f.Db == nil {
		panic("Failed setup repository, db is undefined")
	}

	f.AuthRepository = repository.NewAuth(f.Db)
}
