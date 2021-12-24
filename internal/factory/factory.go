package factory

import (
	"codeid-boiler/database"
	"codeid-boiler/internal/app/repository"

	"gorm.io/gorm"
	"os"
)

type Factory struct {
	Db                                    *gorm.DB
	LoginAppRepository                    repository.LoginApp
	UserAppRepository                     repository.UserApp
	ThnAngRepository                      repository.ThnAng
	SatkerRepository                      repository.Satker
	SpiSdmRepository                      repository.SpiSdm
	JenisSdmRepository                    repository.JenisSdm
	JenisCertificateRepository            repository.JenisCertificate
	PegawaiRepository                     repository.Pegawai
	SpiSdmItemRepository                  repository.SpiSdmItem
	SpiSdmFileRepository                  repository.SpiSdmFile
	RkaklRepository                       repository.Rkakl
	RkaklFileRepository                   repository.RkaklFile
	ProgramRepository                     repository.Program
	RkaklProgRepository                   repository.RkaklProg
	KegiatanRepository                    repository.Kegiatan
	ProgKegiatanRepository                repository.ProgKegiatan
	OutputRepository                      repository.Output
	KegiatanOutputRepository              repository.KegiatanOutput
	KegiatanOutputLocationRepository      repository.KegiatanOutputLocation
	SubOutputRepository                   repository.SubOutput
	KomponenRepository                    repository.Komponen
	SubKomponenRepository                 repository.SubKomponen
	AkunRepository                        repository.Akun
	SubKomponenAkunRepository             repository.SubKomponenAkun
	SubKomponenAkunLocationRepository     repository.SubKomponenAkunLocation
	RkaklItemRepository                   repository.RkaklItem
	SpiAngRepository                      repository.SpiAng
	SpiAngItemRepository                  repository.SpiAngItem
	SpiAngKesesuaianRepository            repository.SpiAngKesesuaian
	JenisKesesuaianRepository             repository.JenisKesesuaian
	JenisPengendaliRepository             repository.JenisPengendali
	JenisRekapitulasiRepository           repository.JenisRekapitulasi
	SpiPbjRekapitulasiRepository          repository.SpiPbjRekapitulasi
	BulanRepository                       repository.Bulan
	GroupPackageValueRepository           repository.GroupPackageValue
	JenisBelanjaPaguRepository            repository.JenisBelanjaPagu
	MethodApbjRepository                  repository.MethodApbj
	SpiPbjPaketRepository                 repository.SpiPbjPaket
	SpiPbjPaketJenisBelanjaPaguRepository repository.SpiPbjPaketJenisBelanjaPagu
	SpiBmnRepository                      repository.SpiBmn
	WbkProgramRankerRepository            repository.WbkProgramRanker
	WbkKomponenRepository                 repository.WbkKomponen
	WbkProgramRepository                  repository.WbkProgram
}

func NewFactory() *Factory {
	f := &Factory{}
	f.SetupDb()
	f.SetupRepository()

	return f
}

func (f *Factory) SetupDb() {
	db, err := database.Connection(os.Getenv("DB_CONN_NAME_DJPT_SPIWBK"))
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
	f.RkaklRepository = repository.NewRkakl(f.Db)
	f.RkaklFileRepository = repository.NewRkaklFile(f.Db)
	f.ProgramRepository = repository.NewProgram(f.Db)
	f.RkaklProgRepository = repository.NewRkaklProg(f.Db)
	f.KegiatanRepository = repository.NewKegiatan(f.Db)
	f.ProgKegiatanRepository = repository.NewProgKegiatan(f.Db)
	f.OutputRepository = repository.NewOutput(f.Db)
	f.KegiatanOutputRepository = repository.NewKegiatanOutput(f.Db)
	f.KegiatanOutputLocationRepository = repository.NewKegiatanOuputLocation(f.Db)
	f.SubOutputRepository = repository.NewSubOutput(f.Db)
	f.KomponenRepository = repository.NewKomponen(f.Db)
	f.SubKomponenRepository = repository.NewSubKomponen(f.Db)
	f.AkunRepository = repository.NewAkun(f.Db)
	f.SubKomponenAkunRepository = repository.NewSubKomponenAkun(f.Db)
	f.SubKomponenAkunLocationRepository = repository.NewSubKomponenAkunLocation(f.Db)
	f.RkaklItemRepository = repository.NewRkaklItem(f.Db)
	f.SpiAngRepository = repository.NewSpiAng(f.Db)
	f.SpiAngItemRepository = repository.NewSpiAngItem(f.Db)
	f.SpiAngKesesuaianRepository = repository.NewSpiAngKesesuaian(f.Db)
	f.JenisKesesuaianRepository = repository.NewJenisKesesuaian(f.Db)
	f.JenisPengendaliRepository = repository.NewJenisPengendali(f.Db)
	f.JenisRekapitulasiRepository = repository.NewJenisRekapitulasi(f.Db)
	f.SpiPbjRekapitulasiRepository = repository.NewSpiPbjRekapitulasi(f.Db)
	f.BulanRepository = repository.NewBulan(f.Db)
	f.GroupPackageValueRepository = repository.NewGroupPackageValue(f.Db)
	f.JenisBelanjaPaguRepository = repository.NewJenisBelanjaPagu(f.Db)
	f.MethodApbjRepository = repository.NewMethodApbj(f.Db)
	f.SpiPbjPaketRepository = repository.NewSpiPbjPaket(f.Db)
	f.SpiPbjPaketJenisBelanjaPaguRepository = repository.NewSpiPbjPaketJenisBelanjaPagu(f.Db)
	f.SpiBmnRepository = repository.NewSpiBmn(f.Db)
	f.WbkProgramRankerRepository = repository.NewWbkProgramRanker(f.Db)
	f.WbkKomponenRepository = repository.NewWbkKomponen(f.Db)
	f.WbkProgramRepository = repository.NewWbkProgram(f.Db)

}
