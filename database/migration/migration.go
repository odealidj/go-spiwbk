package migration

import (
	"codeid-boiler/database"
	//"codeid-boiler/internal/model"
	"codeid-boiler/internal/app/auth/model"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Migration interface {
	DropMigrate()
	AutoMigrate()
	SetDb(*gorm.DB)
}

type migration struct {
	Db            *gorm.DB
	DbModels      *[]interface{}
	IsAutoMigrate bool
}

func Init() {
	mgConfigurations := map[string]Migration{
		strings.ToUpper(os.Getenv("DB_NAME_MIGRATION")): &migration{
			DbModels: &[]interface{}{
				//&model.UserEntityModel{},
				//&model.SampleEntityModel{},
				//&model.SampleChildEntityModel{},
				//&model.SampleGrandChildEntityModel{},
				&model.UserApp{},
				&model.LoginApp{},
			},
			IsAutoMigrate: true,
		},
	}

	for k, v := range mgConfigurations {
		dbConnection, err := database.Connection(k)
		if err != nil {
			logrus.Error(fmt.Sprintf("Failed to run migration, database not found %s", k))
		} else {
			v.SetDb(dbConnection)
			v.DropMigrate()
			v.AutoMigrate()
			logrus.Info(fmt.Sprintf("Successfully run migration for database %s", k))
		}
	}

}

func (m *migration) AutoMigrate() {
	if m.IsAutoMigrate {
		m.Db.AutoMigrate(*m.DbModels...)
	}
}

func (m *migration) DropMigrate() {
	if m.IsAutoMigrate {
		m.Db.Migrator().DropTable(*m.DbModels...)
	}
}

func (m *migration) SetDb(db *gorm.DB) {
	m.Db = db
}
