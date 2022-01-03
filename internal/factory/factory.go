package factory

import (
	"codeid-boiler/database"
	"codeid-boiler/internal/app/repository"
	"codeid-boiler/internal/app/repository/spi-pbj"
	"codeid-boiler/internal/app/repository/wbk"

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
	JenisRekapitulasiRepository           spi_pbj.JenisRekapitulasi
	SpiPbjRekapitulasiRepository          spi_pbj.SpiPbjRekapitulasi
	BulanRepository                       repository.Bulan
	GroupPackageValueRepository           spi_pbj.GroupPackageValue
	JenisBelanjaPaguRepository            spi_pbj.JenisBelanjaPagu
	MethodApbjRepository                  spi_pbj.MethodApbj
	SpiPbjPaketRepository                 spi_pbj.SpiPbjPaket
	SpiPbjPaketJenisBelanjaPaguRepository spi_pbj.SpiPbjPaketJenisBelanjaPagu
	SpiBmnRepository                      repository.SpiBmn
	WbkProgramRankerRepository            wbk.WbkProgramRanker
	WbkKomponenRepository                 wbk.WbkKomponen
	WbkProgramRepository                  wbk.WbkProgram
	WbkProgramTujuanRepository            wbk.WbkProgramTujuan
	WbkProgramTargetRepository            wbk.WbkProgramTarget
	WbkSubProgramRankerRepository         wbk.WbkSubProgramRanker
	FrekuensiRankerRepository             wbk.FrekuensiRanker
	WbkSubProgramUraianRepository         wbk.WbkSubProgramUraian
	WbkSubProgramUraianBulanRepository    wbk.WbkSubProgramUraianBulan
	WbkSubProgramRankerBulanRepository    wbk.WbkSubProgramRankerBulan
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
	f.JenisRekapitulasiRepository = spi_pbj.NewJenisRekapitulasi(f.Db)
	f.SpiPbjRekapitulasiRepository = spi_pbj.NewSpiPbjRekapitulasi(f.Db)
	f.BulanRepository = repository.NewBulan(f.Db)
	f.GroupPackageValueRepository = spi_pbj.NewGroupPackageValue(f.Db)
	f.JenisBelanjaPaguRepository = spi_pbj.NewJenisBelanjaPagu(f.Db)
	f.MethodApbjRepository = spi_pbj.NewMethodApbj(f.Db)
	f.SpiPbjPaketRepository = spi_pbj.NewSpiPbjPaket(f.Db)
	f.SpiPbjPaketJenisBelanjaPaguRepository = spi_pbj.NewSpiPbjPaketJenisBelanjaPagu(f.Db)
	f.SpiBmnRepository = repository.NewSpiBmn(f.Db)
	f.WbkProgramRankerRepository = wbk.NewWbkProgramRanker(f.Db)
	f.WbkKomponenRepository = wbk.NewWbkKomponen(f.Db)
	f.WbkProgramRepository = wbk.NewWbkProgram(f.Db)
	f.WbkProgramTujuanRepository = wbk.NewWbkProgramTujuan(f.Db)
	f.WbkProgramTargetRepository = wbk.NewWbkProgramTarget(f.Db)
	f.WbkSubProgramRankerRepository = wbk.NewWbkSubProgramRanker(f.Db)

	f.FrekuensiRankerRepository = wbk.NewFrekuensiRanker(f.Db)
	f.WbkSubProgramUraianRepository = wbk.NewWbkSubProgramUraian(f.Db)
	f.WbkSubProgramUraianBulanRepository = wbk.NewWbkSubProgramUraianBulan(f.Db)
	f.WbkSubProgramRankerBulanRepository = wbk.NewWbkSubProgramRankerBulan(f.Db)
}
