package factory

import (
	"codeid-boiler/database"
	"codeid-boiler/internal/app/repository"

	"os"
	"strings"

	"gorm.io/gorm"
)

type Factory struct {
	Db                         *gorm.DB
	LoginAppRepository         repository.LoginApp
	UserAppRepository          repository.UserApp
	ThnAngRepository           repository.ThnAng
	SatkerRepository           repository.Satker
	SpiSdmRepository           repository.SpiSdm
	JenisSdmRepository         repository.JenisSdm
	JenisCertificateRepository repository.JenisCertificate
	PegawaiRepository          repository.Pegawai
	SpiSdmItemRepository       repository.SpiSdmItem
	SpiSdmFileRepository       repository.SpiSdmFile
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

	f.LoginAppRepository = repository.NewLoginApp(f.Db)
	f.UserAppRepository = repository.NewUserApp(f.Db)
	f.ThnAngRepository = repository.NewThnAng(f.Db)
	f.SatkerRepository = repository.NewSatker(f.Db)
	f.SpiSdmRepository = repository.NewSpiSdm(f.Db)
	f.JenisSdmRepository = repository.NewJenisSdm(f.Db)
	f.JenisCertificateRepository = repository.NewJenisCertificate(f.Db)
	f.PegawaiRepository = repository.NewPegawai(f.Db)
	f.SpiSdmItemRepository = repository.NewSpiSdmItem(f.Db)
	f.SpiSdmFileRepository = repository.NewSpiSdmFile(f.Db)
}
