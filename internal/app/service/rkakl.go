package service

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/dto"
	"codeid-boiler/internal/app/model"
	"codeid-boiler/internal/app/repository"
	"codeid-boiler/internal/factory"
	"codeid-boiler/pkg/util/numeric"
	res "codeid-boiler/pkg/util/response"
	"codeid-boiler/pkg/util/trxmanager"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type RkaklCode int

const (
	Undefined RkaklCode = iota
	Program
	Kegiatan
	Output
	OutputLocation
	Suboutput
	Komponen
	Subkomponent
	Akun
	AkunLocation
	Rkaklitem
)

func (r RkaklCode) String() string {
	switch r {
	case Program:
		return "Program"
	case Kegiatan:
		return "Kegiatan"
	case Output:
		return "Output"
	case OutputLocation:
		return "Outputlocation"
	case Suboutput:
		return "Suboutput"
	case Komponen:
		return "Komponen"
	case Subkomponent:
		return "Subkomponent"
	case Akun:
		return "Akun"
	case AkunLocation:
		return "Akunlocation"
	case Rkaklitem:
		return "Rkaklitem"

	default:
		return "Undefined"
	}

}

type RkaklService interface {
	Save(*abstraction.Context, *dto.RkaklSaveRequest, *multipart.FileHeader) (*dto.RkaklResponse, error)
	Update(*abstraction.Context, *dto.RkaklUpdateRequest, *multipart.FileHeader) (*dto.RkaklResponse, error)
	Delete(*abstraction.Context, *dto.RkaklDeleteRequest) (*dto.RkaklDeleteResponse, error)
	Get(*abstraction.Context, *dto.RkaklGetRequest) (*dto.RkaklGetInfoResponse, error)
	GetByID(*abstraction.Context, *dto.RkaklGetByIDRequest) (*dto.RkaklResponse, error)
	GetByID2(*abstraction.Context, *dto.RkaklGetByIDRequest) (*dto.RkaklResponse, error)
}

type rkaklService struct {
	Repository                        repository.Rkakl
	RkaklFileRepository               repository.RkaklFile
	ProgramRepository                 repository.Program
	RkaklProgRepository               repository.RkaklProg
	KegiatanRepository                repository.Kegiatan
	ProgKegiatanRepository            repository.ProgKegiatan
	OutputRepository                  repository.Output
	KegiatanOutputRepository          repository.KegiatanOutput
	KegiatanOutputLocationRepository  repository.KegiatanOutputLocation
	SubOutputRepository               repository.SubOutput
	KomponenRepository                repository.Komponen
	SubKomponenRepository             repository.SubKomponen
	AkunRepository                    repository.Akun
	SubKomponenAkunRepository         repository.SubKomponenAkun
	SubKomponenAkunLocationRepository repository.SubKomponenAkunLocation
	RkaklItemRepository               repository.RkaklItem
	Db                                *gorm.DB
}

func NewRkaklService(f *factory.Factory) *rkaklService {
	repository := f.RkaklRepository
	rkaklFileRepository := f.RkaklFileRepository
	programRepository := f.ProgramRepository
	rkaklProgRepository := f.RkaklProgRepository
	kegiatanRepository := f.KegiatanRepository
	progKegiatanRepository := f.ProgKegiatanRepository
	outputRepository := f.OutputRepository
	kegiatanOutputRepository := f.KegiatanOutputRepository
	kegiatanOutputLocationRepository := f.KegiatanOutputLocationRepository
	subOutputRepository := f.SubOutputRepository
	komponenRepository := f.KomponenRepository
	subKomponenRepository := f.SubKomponenRepository
	akunRepository := f.AkunRepository
	subKomponenAkunRepository := f.SubKomponenAkunRepository
	subKomponenAkunLocationRepository := f.SubKomponenAkunLocationRepository
	rkaklItemRepository := f.RkaklItemRepository

	db := f.Db
	return &rkaklService{repository, rkaklFileRepository, programRepository,
		rkaklProgRepository, kegiatanRepository, progKegiatanRepository,
		outputRepository, kegiatanOutputRepository,
		kegiatanOutputLocationRepository, subOutputRepository,
		komponenRepository, subKomponenRepository, akunRepository,
		subKomponenAkunRepository, subKomponenAkunLocationRepository,
		rkaklItemRepository, db}

}

func (s *rkaklService) Save(ctx *abstraction.Context, payload *dto.RkaklSaveRequest, file *multipart.FileHeader) (*dto.RkaklResponse, error) {

	var result *dto.RkaklResponse
	var rkaklProg *model.RkaklProg
	var progKegiatan *model.ProgKegiatan
	var kegiatanOutput *model.KegiatanOutput
	var kegiatanOuputLocation *model.KegiatanOutputLocation
	var subOutput *model.SubOutput
	var komponen *model.Komponen
	var subKomponen *model.SubKomponen
	var akun *model.Akun
	var subKomponenAkun *model.SubKomponenAkun
	var subKomponenAkunLocation *model.SubKomponenAkunLocation

	//var data *model.ThnAng
	var rkaklCode RkaklCode = Undefined
	var code string
	var deskripsi string
	var volume string
	var harga decimal.Decimal
	var biaya decimal.Decimal
	var sdcp string

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		src, err := file.Open()
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UploadFileSrcError, err)
		}
		defer src.Close()

		rkakl, err := s.Repository.Create(ctx, &model.Rkakl{
			Context:     ctx,
			RkaklEntity: payload.RkaklEntity,
		})

		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			} else if strings.Contains(strings.ToLower(err.Error()), "foreign key") {
				return res.ErrorBuilder(&res.ErrorConstant.ForeignKey, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		fileName := fmt.Sprintf("rkakl%d%s", rkakl.ID, filepath.Ext(file.Filename))

		//pathApp, _ := os.Getwd()

		// Destination
		destinationPath := path.Join("upload", fileName)

		dst, err := os.Create(destinationPath)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UploadFileCreateError, err)
			//return res.ErrorResponse(err).Send(c)
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			//fmt.Println("no copy")
			return res.ErrorBuilder(&res.ErrorConstant.UploadFileDestError, err)
		}

		payload.FilePath = destinationPath

		f, err := excelize.OpenFile("./" + payload.FilePath)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.OpenFileErr, err)
		}
		defer func() {
			// Close the spreadsheet.
			if err := f.Close(); err != nil {
				fmt.Println(err)
			}
		}()

		rows, err := f.GetRows(payload.Sheetname)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.SheetFileXLSXSErr, err)
		}

		isValid, initRowHeader := InitFileHeader(rows)
		if isValid == false {
			//return res.ErrorBuilder(&res.ErrorConstant.SheetFileXLSXSErr, err)
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Header file yang valid adalah Kode, Deskripsi, Volume, Harga Satuan, Jumlah Biaya, SD/CP", "Invalid file header")
		}

		//exitLoopRow:
		for i, row := range rows {

			//cek initial data
			if i < initRowHeader+1 {
				continue
			}

			for j, colCell := range row {

				/*
					switch rkaklCode {
					case Program, Kegiatan, Output:
						if j > 1 {
							break forColCell
						}
					}
				*/

				colCell = strings.TrimSpace(colCell)

				//remove single quote
				switch true {
				case strings.Index(colCell, "'") == 0:
					colCell = strings.Replace(colCell, "'", "", -1)
					break
				case strings.Index(colCell, "`") == 0:
					colCell = strings.Replace(colCell, "`", "", -1)
				}

				/*
					if colCell == "2338" {
						break exitLoopRow
					}
				*/

				switch j {
				case 0: //Kode
					switch len(colCell) {
					case 0: //empty
						if rkaklCode == OutputLocation || rkaklCode == AkunLocation {
							if IsItem(row) == false {
								return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
									fmt.Sprintf("Baris:'%d' tidak valid, Karena Column: Kode,Volume,Harga Satuan,Jumlah Biaya,SD/CP tidak kosong", i+1), "Invalid value")
							}
						}
						break
					case 1: //kdsubkomponent
						if numeric.IsNumeric(colCell) == true {
							return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
								fmt.Sprintf("Baris:'%d' Column:'Kode', [%s] bukan string, "+
									"Berikut contoh format data valid A", i+1, colCell), "Invalid value")
						}
						rkaklCode = Subkomponent
						code = strings.TrimSpace(colCell)
						break
					case 3: //kdkomponent
						if numeric.IsNumeric(colCell) == false {
							return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
								fmt.Sprintf("Baris:'%d' Column:'Kode', [%s] bukan string, "+
									"Berikut contoh format data valid 052", i+1, colCell), "Invalid value")
						}
						rkaklCode = Komponen
						code = strings.TrimSpace(colCell)
						break
					case 4: //kdkegiatan
						if numeric.IsNumeric(colCell) == false {
							return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
								fmt.Sprintf("Baris:'%d' Column:'Kode', [%s] bukan numeric, "+
									"Berikut contoh format data valid 2337", i+1, colCell), "Invalid value")
						}
						rkaklCode = Kegiatan
						code = strings.TrimSpace(colCell)
						break
					case 6: //kdakun
						if numeric.IsNumeric(colCell) == false {
							return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
								fmt.Sprintf("Baris:'%d' Column:'Kode', [%s] bukan string, "+
									"Berikut contoh format data valid 521211", i+1, colCell), "Invalid value")
						}
						rkaklCode = Akun
						code = strings.TrimSpace(colCell)
						break
					case 8: //kdoutput
						if strings.Index(colCell, ".") != 4 {
							return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
								fmt.Sprintf("Baris:'%d' Column:'Kode', [%s] tidak valid, "+
									"Berikut contoh format data valid 2337.BDC", i+1, colCell), "Invalid value")
						}
						rkaklCode = Output
						code = strings.TrimSpace(strings.Split(colCell, ".")[1])
						break
					case 9: //kdprog
						if strings.Index(colCell, ".") != 3 || strings.LastIndex(colCell, ".") != 6 {
							return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
								fmt.Sprintf("Baris:'%d' Column:'Kode', [%s] tidak valid, "+
									"Berikut contoh format data valid 032.03.HB", i+1, colCell), "Invalid value")
						}
						rkaklCode = Program
						code = strings.TrimSpace(strings.Split(colCell, ".")[2])
						break
					case 12: //kdsuboutput
						if strings.Index(colCell, ".") != 4 || strings.LastIndex(colCell, ".") != 8 {
							return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
								fmt.Sprintf("Baris:'%d' Column:'Kode', [%s] tidak valid, "+
									"Berikut contoh format data valid 2337.BDC.001", i+1, colCell), "Invalid value")
						}
						rkaklCode = Suboutput
						code = strings.TrimSpace(strings.Split(colCell, ".")[2])
						break
					default:
						return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
							fmt.Sprintf("Baris:'%d' Column:'Kode', [%s] tidak valid, "+
								"Panjang data yang valid adalah 0, 1, 3, 4, 6, 8, 9, 12", i+1, colCell), "Invalid value")
					}

				case 1: // deskripsi
					deskripsi = colCell
					break
				case 2: // volume
					volume = colCell
					break
				case 3: //harga satuan
					if colCell == "" {
						colCell = "0"
					}
					harga, err = decimal.NewFromString(colCell)
					if err != nil {
						return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
							"Error convert harga satuan", err.Error())
					}
					break
				case 4: //jumlah biaya
					if colCell == "" {
						colCell = "0"
					}
					biaya, err = decimal.NewFromString(colCell)
					if err != nil {
						return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
							"Error convert biaya", err.Error())
					}
					break
				case 5: //SC/CP
					sdcp = colCell
				default:
					break

				} //end switch j

			} //end for j

			switch rkaklCode {
			case Program:
				program, err := s.ProgramRepository.FirstOrCreate(ctx, &model.Program{Context: ctx,
					ProgramEntity: model.ProgramEntity{Code: code, Name: deskripsi}})
				if err != nil {
					return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
						"Error program", err.Error())
				}

				rkaklProg, err = s.RkaklProgRepository.Upsert(ctx, &model.RkaklProg{Context: ctx,
					RkaklProgEntity: model.RkaklProgEntity{RkaklID: int(rkakl.ID), ProgramID: int(program.ID),
						Biaya: biaya}})
				if err != nil {
					return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
						"Error rkaklprog", err.Error())
				}
				break

			case Kegiatan:
				kegiatan, err := s.KegiatanRepository.FirstOrCreate(ctx, &model.Kegiatan{Context: ctx,
					KegiatanEntity: model.KegiatanEntity{Code: code, Name: deskripsi}})
				if err != nil {
					return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
						"Error kegiatan", err.Error())
				}

				progKegiatan, err = s.ProgKegiatanRepository.Upsert(ctx, &model.ProgKegiatan{Context: ctx,
					ProgKegiatanEntity: model.ProgKegiatanEntity{RkaklProgID: int(rkaklProg.ID),
						KegiatanID: int(kegiatan.ID), Biaya: biaya}})
				if err != nil {
					return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
						"Error progkegiatan", err.Error())
				}

				break
			case Output:

				output, err := s.OutputRepository.FirstOrCreate(ctx, &model.Output{Context: ctx,
					OutputEntity: model.OutputEntity{Code: code, Name: deskripsi}})
				if err != nil {
					return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
						"Error ouput", err.Error())
				}
				kegiatanOutput, err = s.KegiatanOutputRepository.Upsert(ctx, &model.KegiatanOutput{Context: ctx,
					KegiatanOutputEntity: model.KegiatanOutputEntity{ProgKegiatanID: int(progKegiatan.ID),
						OutputID: int(output.ID), Volume: volume, Biaya: biaya}})
				if err != nil {
					return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
						"Error kegiatan output", err.Error())

				}
				rkaklCode = OutputLocation
				break
			case OutputLocation:
				kegiatanOuputLocation, err = s.KegiatanOutputLocationRepository.Upsert(ctx, &model.KegiatanOutputLocation{Context: ctx,
					KegiatanOutputLocationEntity: model.KegiatanOutputLocationEntity{KegiatanOutputID: int(kegiatanOutput.ID),
						Name: deskripsi}})
				if err != nil {
					return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
						"Error kegiatan output location", err.Error())
				}
			case Suboutput:
				subOutput, err = s.SubOutputRepository.Upsert(ctx, &model.SubOutput{Context: ctx,
					SubOutputEntity: model.SubOutputEntity{KegiatanOutputLocationID: int(kegiatanOuputLocation.ID),
						Code: code, Name: deskripsi, Volume: volume, Biaya: biaya}})
				if err != nil {
					return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
						"Error sub output", err.Error())
				}
			case Komponen:
				komponen, err = s.KomponenRepository.Upsert(ctx, &model.Komponen{Context: ctx,
					KomponenEntity: model.KomponenEntity{SubOutputID: int(subOutput.ID),
						Code: code, Name: deskripsi, Volume: volume, Biaya: biaya}})
				if err != nil {
					return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
						"Error komponen", err.Error())
				}
			case Subkomponent:
				subKomponen, err = s.SubKomponenRepository.Upsert(ctx, &model.SubKomponen{Context: ctx,
					SubKomponenEntity: model.SubKomponenEntity{KomponenID: int(komponen.ID),
						Code: code, Name: deskripsi, Biaya: biaya}})
				if err != nil {
					return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
						"Error Sub komponen", err.Error())
				}
			case Akun:
				akun, err = s.AkunRepository.FirstOrCreate(ctx, &model.Akun{Context: ctx,
					AkunEntity: model.AkunEntity{Code: code, Name: deskripsi}})
				if err != nil {
					return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
						"Error Akun", err.Error())
				}

				subKomponenAkun, err = s.SubKomponenAkunRepository.Upsert(ctx, &model.SubKomponenAkun{Context: ctx,
					SubKomponenAkunEntity: model.SubKomponenAkunEntity{SubKomponenID: int(subKomponen.ID),
						AkunID: int(akun.ID), Name: deskripsi, Biaya: biaya, Sdcp: sdcp}})
				if err != nil {
					return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
						"Error Sub komponen akun", err.Error())
				}
				rkaklCode = AkunLocation
			case AkunLocation:
				subKomponenAkunLocation, err = s.SubKomponenAkunLocationRepository.Upsert(ctx, &model.SubKomponenAkunLocation{Context: ctx,
					SubKomponenAkunLocationEntity: model.SubKomponenAkunLocationEntity{SubKomponenAkunID: int(subKomponenAkun.ID),
						Name: deskripsi}})
				if err != nil {
					return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
						"Error sub komponen akun location", err.Error())
				}
				rkaklCode = Rkaklitem
			case Rkaklitem:
				_, err = s.RkaklItemRepository.Upsert(ctx, &model.RkaklItem{Context: ctx,
					RkaklItemEntity: model.RkaklItemEntity{SubKomponenAkunID: int(subKomponenAkun.ID),
						Name: deskripsi, Volume: volume, Harga: harga, Biaya: biaya}})
				if err != nil {
					return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
						"Error rkakl item", err.Error())
				}
				rkaklCode = Rkaklitem

			} //end switch rkaklcode

		} //end for i

		//fmt.Println(rows)
		//time.Sleep(time.Minute * 5)

		rkaklFile, err := s.RkaklFileRepository.Create(ctx, &model.RkaklFile{
			Entity:          abstraction.Entity{ID: abstraction.ID{ID: rkakl.ID}},
			RkaklFileEntity: model.RkaklFileEntity{Filepath: strings.TrimSpace(strings.Split(payload.FilePath, "/")[1])},
		})

		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			} else if strings.Contains(strings.ToLower(err.Error()), "foreign key") {
				return res.ErrorBuilder(&res.ErrorConstant.ForeignKey, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.RkaklResponse{
			ID:          abstraction.ID{ID: rkakl.ID},
			RkaklEntity: rkakl.RkaklEntity,
			Filepath:    &rkaklFile.Filepath,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func InitFileHeader(rows [][]string) (bool, int) {

	//validasi header
	initFileHeader := 0
	isHeadervalid := false
	isHeaderKode := false
	isHeaderDeskripsi := false
	isHeaderVolume := false
	isHeaderHarga := false
	isHeaderJumlahBiaya := false
	isHeaderSDCP := false
	for i, row := range rows {
		for _, colCell := range row {
			if len(colCell) > 0 {
				if strings.Contains(strings.ToLower(colCell), "kode") {
					isHeaderKode = true
				}
				if strings.Contains(strings.ToLower(colCell), "deskripsi") {
					isHeaderDeskripsi = true
				}
				if strings.Contains(strings.ToLower(colCell), "volume") {
					isHeaderVolume = true
				}
				if strings.Contains(strings.ToLower(colCell), "harga") {
					isHeaderHarga = true
				}
				if strings.Contains(strings.ToLower(colCell), "jumlah") {
					isHeaderJumlahBiaya = true
				}
				if strings.Contains(strings.ToLower(colCell), "sd") || strings.Contains(strings.ToLower(colCell), "cp") {
					isHeaderSDCP = true
				}
			}

		}
		if isHeaderKode == true && isHeaderDeskripsi == true && isHeaderVolume == true &&
			isHeaderHarga == true && isHeaderJumlahBiaya == true && isHeaderSDCP == true {
			isHeadervalid = true
			initFileHeader = i
			break
		}

	}
	return isHeadervalid, initFileHeader
}

func IsItem(row []string) bool {
	isEmptyCode := false
	isEmptyDeskripsi := true
	isEmptyVolum := false
	isEmptyHarga := false
	isEmptyBiaya := false
	for i, colCell := range row {

		//remove single quote
		switch true {
		case strings.Index(colCell, "'") == 0:
			colCell = strings.Replace(colCell, "'", "", -1)
			break
		case strings.Index(colCell, "`") == 0:
			colCell = strings.Replace(colCell, "`", "", -1)
			break
		}
		colCell = strings.TrimSpace(colCell)

		switch i {
		case 0: //Code
			if colCell == "" {
				isEmptyCode = true
			}
		case 1: //Deskripsi
			if colCell != "" {
				isEmptyDeskripsi = false
			}
		case 2: //volume
			if colCell == "" {
				isEmptyVolum = true
			}
		case 3: //harga
			if colCell == "" {
				isEmptyHarga = true
			}
		case 4: //Biaya

			if colCell == "" || colCell == "0" {
				//fmt.Println("masuk sini")
				isEmptyBiaya = true
			}
			//case 5: //SD/CP
			//	if colCell == "" {
			//		isEmptySDCP = true
			//	}
		}
	} //end loop row
	if isEmptyCode == true && isEmptyVolum == true && isEmptyHarga == true && isEmptyBiaya == true &&
		isEmptyDeskripsi == false {
		return true
	}
	return false
}

func (s *rkaklService) Update(ctx *abstraction.Context, payload *dto.RkaklUpdateRequest, file *multipart.FileHeader) (*dto.RkaklResponse, error) {

	var result *dto.RkaklResponse
	var oldFile, fileLocation string
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		rkakl, err := s.Repository.FindByID(ctx, &model.Rkakl{Context: ctx, EntityInc: abstraction.EntityInc{
			IDInc: abstraction.IDInc{ID: payload.ID.ID}}})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		rkakl, err = s.Repository.Update(ctx, &model.Rkakl{Context: ctx,
			EntityInc:   abstraction.EntityInc{IDInc: abstraction.IDInc{ID: rkakl.ID}},
			RkaklEntity: payload.RkaklEntity,
		})
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			} else if strings.Contains(strings.ToLower(err.Error()), "foreign key") {
				return res.ErrorBuilder(&res.ErrorConstant.ForeignKey, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		rkaklFile, err := s.RkaklFileRepository.FindByID(ctx, &model.RkaklFile{Context: ctx,
			Entity: abstraction.Entity{ID: abstraction.ID{ID: rkakl.ID}}})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		if file != nil {

			src, err := file.Open()
			if err != nil {
				return res.ErrorBuilder(&res.ErrorConstant.UploadFileSrcError, err)
			}
			defer src.Close()

			fileName := fmt.Sprintf("rkakl%d%s", rkakl.ID, filepath.Ext(file.Filename))

			// Destination
			destinationPath := path.Join("upload", fileName)

			dst, err := os.Create(destinationPath)
			if err != nil {
				return res.ErrorBuilder(&res.ErrorConstant.UploadFileCreateError, err)
				//return res.ErrorResponse(err).Send(c)
			}
			defer dst.Close()

			// Copy
			if _, err = io.Copy(dst, src); err != nil {
				//fmt.Println("no copy")
				return res.ErrorBuilder(&res.ErrorConstant.UploadFileDestError, err)
			}

			oldFile = path.Join("upload", rkaklFile.Filepath)

			fileLocation = destinationPath
			//payload.ImagePath = destinationPath
		} else {
			//fmt.Println("No file")
			fileLocation = rkaklFile.Filepath
		}

		rkaklFile, err = s.RkaklFileRepository.Update(ctx, &model.RkaklFile{
			Entity:          abstraction.Entity{ID: abstraction.ID{ID: rkakl.ID}},
			RkaklFileEntity: model.RkaklFileEntity{Filepath: strings.TrimSpace(strings.Split(fileLocation, "/")[1])},
		})

		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			} else if strings.Contains(strings.ToLower(err.Error()), "foreign key") {
				return res.ErrorBuilder(&res.ErrorConstant.ForeignKey, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		//tidak diperlukan karena nama file nya sama
		//if len(oldFile) > 0 {
		//	_ = os.Remove(oldFile)
		//}

		result = &dto.RkaklResponse{
			ID:          abstraction.ID{ID: rkakl.ID},
			RkaklEntity: rkakl.RkaklEntity,
			Filepath:    &rkaklFile.Filepath,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *rkaklService) Delete(ctx *abstraction.Context, payload *dto.RkaklDeleteRequest) (*dto.RkaklDeleteResponse, error) {
	var result *dto.RkaklDeleteResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		rkakl, err := s.Repository.FindByID(ctx, &model.Rkakl{Context: ctx, EntityInc: abstraction.EntityInc{
			IDInc: abstraction.IDInc{ID: payload.ID.ID}}})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		rkaklFile, err := s.RkaklFileRepository.FindByID(ctx, &model.RkaklFile{Context: ctx, Entity: abstraction.Entity{
			ID: abstraction.ID{ID: rkakl.ID}}})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		fmt.Println(path.Join("upload", rkaklFile.Filepath))

		if len(rkaklFile.Filepath) > 0 {

			_ = os.Remove(path.Join("upload", rkaklFile.Filepath))
		}

		_, err = s.Repository.Delete(ctx, &model.Rkakl{Context: ctx,
			EntityInc: abstraction.EntityInc{IDInc: abstraction.IDInc{ID: rkakl.ID}},
		})
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.RkaklDeleteResponse{ID: payload.ID}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *rkaklService) Get(ctx *abstraction.Context, payload *dto.RkaklGetRequest) (*dto.RkaklGetInfoResponse, error) {
	var result *dto.RkaklGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		rkakls, info, err := s.Repository.Find(ctx, &payload.RkaklFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		result = &dto.RkaklGetInfoResponse{
			Datas:          &rkakls,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *rkaklService) GetByID(ctx *abstraction.Context, payload *dto.RkaklGetByIDRequest) (*dto.RkaklResponse, error) {
	var result *dto.RkaklResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		rkakl, err := s.Repository.FindByID(ctx, &model.Rkakl{
			Context:   ctx,
			EntityInc: abstraction.EntityInc{IDInc: abstraction.IDInc{ID: payload.ID.ID}},
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		rkaklFile, err := s.RkaklFileRepository.FindByID(ctx, &model.RkaklFile{
			Context: ctx,
			Entity:  abstraction.Entity{ID: abstraction.ID{ID: payload.ID.ID}},
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.RkaklResponse{
			ID:          abstraction.ID{ID: rkakl.ID},
			RkaklEntity: rkakl.RkaklEntity,
			Filepath:    &rkaklFile.Filepath,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *rkaklService) GetByID2(ctx *abstraction.Context, payload *dto.RkaklGetByIDRequest) (*dto.RkaklResponse, error) {
	var result *dto.RkaklResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		rkakl, err := s.Repository.FindThnAngSatkerByID(ctx, &model.Rkakl{
			Context:   ctx,
			EntityInc: abstraction.EntityInc{IDInc: abstraction.IDInc{ID: payload.ID.ID}},
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				//result = &dto.RkaklResponse{}
				return nil
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		//if rkakl.ID != 0 && rkakl.SatkerID != 0 {

		rkaklFile, err := s.RkaklFileRepository.FindByID(ctx, &model.RkaklFile{
			Context: ctx,
			Entity:  abstraction.Entity{ID: abstraction.ID{ID: payload.ID.ID}},
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.RkaklResponse{
			ID:          abstraction.ID{ID: rkakl.ID},
			RkaklEntity: rkakl.RkaklEntity,
			ThnAngYear:  &rkakl.ThnAng.Year,
			SatkerName:  &rkakl.Satker.Name,
			Filepath:    &rkaklFile.Filepath,
		}
		//} else {
		//result = &dto.RkaklResponse{}
		//}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
