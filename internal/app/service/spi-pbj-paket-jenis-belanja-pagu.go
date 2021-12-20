package service

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/dto"
	"codeid-boiler/internal/app/model"
	"codeid-boiler/internal/app/repository"
	"codeid-boiler/internal/factory"
	res "codeid-boiler/pkg/util/response"
	"codeid-boiler/pkg/util/trxmanager"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type SpiPbjPaketJenisBelanjaPaguService interface {
	Save(*abstraction.Context, *dto.SpiPbjPaketJenisBelanjaPaguSaveRequest) (*dto.SpiAngResponse, error)
	//Upsert(*abstraction.Context, *dto.SpiPbjPaketJenisBelanjaPaguUpsertRequest) ([]dto.SpiPbjRekapitulasiResponse, error)
	GetSpiPbjPaketJenisBelanjaPaguByThnAngIDAndSatkerID(*abstraction.Context, *dto.SpiPbjPaketJenisBelanjaPaguGetRequest) (
		*dto.SpiPbjPaketJenisBelanjaPaguGetInfoResponse, error)
}

type spiPbjPaketJenisBelanjaPaguService struct {
	SpiAngRepository                      repository.SpiAng
	SpiPbjPaketRepository                 repository.SpiPbjPaket
	SpiPbjPaketJenisBelanjaPaguRepository repository.SpiPbjPaketJenisBelanjaPagu
	GroupPackageValueRepository           repository.GroupPackageValue
	KomponenRepository                    repository.Komponen
	SubKomponenAkunRepository             repository.SubKomponenAkun
	Db                                    *gorm.DB
}

func NewSpiPbjPaketJenisBelanjaPaguService(f *factory.Factory) *spiPbjPaketJenisBelanjaPaguService {
	spiAngRepository := f.SpiAngRepository
	spiPbjPaketRepository := f.SpiPbjPaketRepository
	spiPbjPaketJenisBelanjaPaguRepository := f.SpiPbjPaketJenisBelanjaPaguRepository
	groupPackageValueRepository := f.GroupPackageValueRepository
	komponenRepository := f.KomponenRepository
	subKomponenAkunRepository := f.SubKomponenAkunRepository
	db := f.Db
	return &spiPbjPaketJenisBelanjaPaguService{spiAngRepository, spiPbjPaketRepository,
		spiPbjPaketJenisBelanjaPaguRepository, groupPackageValueRepository,
		komponenRepository, subKomponenAkunRepository, db}

}

func (s *spiPbjPaketJenisBelanjaPaguService) Save(ctx *abstraction.Context, payload *dto.SpiPbjPaketJenisBelanjaPaguSaveRequest) (
	*dto.SpiAngResponse, error) {

	var isSave bool = false
	var result *dto.SpiAngResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		isSave = true
		spiAng, err := s.SpiAngRepository.Create(ctx, &model.SpiAng{Context: ctx, SpiAngEntity: model.SpiAngEntity{
			ThnAngID: uint16(payload.ThnAngID), SatkerID: uint16(payload.SatkerID),
		}})
		if err != nil {
			//if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			//	return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
			//		"Duplicate spi ang", "Invalid spi ang")
			//}
			isSave = false
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid spi ang", "Invalid spi ang")
		}

		//Find All GroupPackageValue
		SortByGroupPackageValue := "id"
		SortGroupPackageValue := "asc"
		groupPackageValues, _, err := s.GroupPackageValueRepository.Find(ctx,
			&model.GroupPackageValueFilter{}, &abstraction.Pagination{
				SortBy: &SortByGroupPackageValue, Sort: &SortGroupPackageValue,
			})
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid group package", err.Error())
		}

		if len(groupPackageValues) == 0 {
			isSave = false
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid group package", "GroupPackageValue tidak ditemukan")
		}

		//Find Komponen by satkerID dan thnAngID
		komponens, err := s.KomponenRepository.FindByThnAngIDAndSatkerID(ctx, &dto.KomponenFindByThnAngIDAndSatkerIDRequest{
			ThnAngID: payload.ThnAngID, SatkerID: payload.SatkerID,
		})
		if err != nil {
			isSave = false
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid komponen", err.Error())
		}
		if len(*komponens) == 0 {
			isSave = false
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid komponen", "Komponen tidak ditemukan")
		}

		tempGroupPackageValue := 1
		tempJenisBelanjaPaguID := 1

		for _, komponen := range *komponens {

			//valiasi group package value
			for _, groupPackageValue := range groupPackageValues {
				if komponen.Biaya.GreaterThan(groupPackageValue.MinValue) && komponen.Biaya.LessThanOrEqual(groupPackageValue.MaxValue) {
					tempGroupPackageValue = int(groupPackageValue.ID.ID)
					break
				} else {
					tempGroupPackageValue = int(groupPackageValue.ID.ID)
					break
				}
			}

			spiPbjPaket, err := s.SpiPbjPaketRepository.Create(ctx, &model.SpiPbjPaket{Context: ctx,
				SpiPbjPaketEntity: model.SpiPbjPaketEntity{SpiAngID: int(spiAng.ID),
					GroupPackageValueID: tempGroupPackageValue,
					KomponenID:          int(komponen.ID)},
			})
			if err != nil {
				isSave = false
				return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
					"Invalid spi-pbj-paket", err.Error())
			}

			subKomponenAkuns, err := s.SubKomponenAkunRepository.FindByKomponenID(ctx, &dto.SubKomponenAkunFindByKomponenIDRequest{
				KomponenID: int(komponen.ID),
			})
			if err != nil {
				isSave = false
				return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
					"Invalid Sub Komponen", err.Error())
			}
			if len(*subKomponenAkuns) == 0 {
				isSave = false
				return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
					"Invalid Sub komponen akun", "Sub Komponen Akun tidak ditemukan")
			}

			for _, subKomponenAkun := range *subKomponenAkuns {

				//validasi jenis belanja pagu
				if strings.Index(strings.TrimSpace(subKomponenAkun.AkunCode), "52") == 0 {
					if strings.Index(strings.TrimSpace(subKomponenAkun.AkunCode), "523") == 0 {
						tempJenisBelanjaPaguID = 3 //Sosial
					} else {
						tempJenisBelanjaPaguID = 1 //Barang
					}
				} else if strings.Index(strings.TrimSpace(subKomponenAkun.AkunCode), "53") == 0 {
					tempJenisBelanjaPaguID = 2 //Modal
				} else {
					tempJenisBelanjaPaguID = 4
				}

				_, err := s.SpiPbjPaketJenisBelanjaPaguRepository.Upsert(ctx, &model.SpiPbjPaketJenisBelanjaPagu{
					ID: abstraction.ID{ID: spiPbjPaket.ID},
					SpiPbjPaketJenisBelanjaPaguEntity: model.SpiPbjPaketJenisBelanjaPaguEntity{
						SpiPbjPaketID: int(spiPbjPaket.ID), JenisBelanjaPaguID: tempJenisBelanjaPaguID,
						SubKomponenAkunID: int(subKomponenAkun.ID.ID),
					},
				})
				if err != nil {
					isSave = false
					return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
						"Invalid SpiPbjPaketJenisBelanjaPagu", err.Error())
				}

			} //end loop sub komponen akun

			//fmt.Println(spiPbjPaket)

		} // end loop komponens

		if isSave == true {

			result = &dto.SpiAngResponse{
				ID:           abstraction.ID{ID: spiAng.ID},
				SpiAngEntity: spiAng.SpiAngEntity,
			}
		}

		//time.Sleep(time.Minute * 5)

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

/*
func (s *spiPbjPaketJenisBelanjaPaguService) Upsert(ctx *abstraction.Context, payload *dto.SpiPbjRekapitulasiUpsertRequest) ([]dto.SpiPbjRekapitulasiResponse, error) {

	var result []dto.SpiPbjRekapitulasiResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		spiAng, err := s.SpiAngRepository.Create(ctx, &model.SpiAng{Context: ctx, SpiAngEntity: model.SpiAngEntity{
			ThnAngID: uint16(payload.ThnAngID), SatkerID: uint16(payload.SatkerID),
		}})
		if err != nil {
			//if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			//	return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
			//		"Duplicate spi ang", "Invalid spi ang")
			//}

			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid spi ang", "Invalid spi ang")
		}

		payload.SpiPbjRekapitulasiEntity.SpiAngID = int(spiAng.ID)
		spiPbjRekapitulasi, err := s.SpiPbjRekapitulasiRepository.Upsert(ctx, &model.SpiPbjRekapitulasi{Context: ctx,
			IDInc:                    abstraction.IDInc{ID: payload.ID.ID},
			SpiPbjRekapitulasiEntity: payload.SpiPbjRekapitulasiEntity})
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid upsert spi pbj rekapitulasi", err.Error())
		}

		result = append(result, dto.SpiPbjRekapitulasiResponse{
			ID:                       abstraction.ID{ID: spiPbjRekapitulasi.ID},
			SpiPbjRekapitulasiEntity: spiPbjRekapitulasi.SpiPbjRekapitulasiEntity,
			SatkerID:                 payload.SatkerID,
			ThnAngID:                 payload.ThnAngID,
		})

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}
*/

func (s *spiPbjPaketJenisBelanjaPaguService) GetSpiPbjPaketJenisBelanjaPaguByThnAngIDAndSatkerID(ctx *abstraction.Context,
	payload *dto.SpiPbjPaketJenisBelanjaPaguGetRequest) (
	*dto.SpiPbjPaketJenisBelanjaPaguGetInfoResponse, error) {

	var result *dto.SpiPbjPaketJenisBelanjaPaguGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		groupPackageValues, info, err := s.GroupPackageValueRepository.FindExistSpiPbjPaket(ctx,
			&model.GroupPackageValueFilter{
				GroupPackageValueEntityFilter: model.GroupPackageValueEntityFilter{ThnAngID: payload.ThnAngID,
					SatkerID: payload.SatkerID}}, &abstraction.Pagination{Page: payload.Page, PageSize: payload.PageSize})
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		if len(groupPackageValues) == 0 {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid group package", "GroupPackageValue tidak ditemukan")
		}

		alfabetCounter := [4]string{"A", "B", "C", "D"}
		subTotalBarang := decimal.NewFromInt(0)
		subTotalModal := decimal.NewFromInt(0)
		subTotalSosial := decimal.NewFromInt(0)
		totalBarang := decimal.NewFromInt(0)
		totalModal := decimal.NewFromInt(0)
		totalSosial := decimal.NewFromInt(0)

		spiPbjPaketJenisBelanjaPaguGetReformatResponse := &[]dto.SpiPbjPaketJenisBelanjaPaguGetReformatResponse{}

		for i, groupPackageValue := range groupPackageValues {
			//fmt.Println(groupPackageValue.ID.ID)
			//*payload.GroupPackageValueID = int(groupPackageValue.ID.ID)
			//fmt.Println("test")
			groupPackageValueID := int(groupPackageValue.ID.ID)
			spiPbjPaketJenisBelanjaPagues, _, err := s.SpiPbjPaketJenisBelanjaPaguRepository.FindspiPbjPaketJenisBelanjaPaguByThnAngIDAndSatkerID(ctx,
				&model.SpiPbjPaketJenisBelanjaPaguFilter{
					SpiPbjPaketJenisBelanjaPaguEntityFilter: model.SpiPbjPaketJenisBelanjaPaguEntityFilter{
						ThnAngID: payload.ThnAngID, SatkerID: payload.SatkerID,
						GroupPackageValueID: &groupPackageValueID,
					},
				}, &payload.Pagination)
			if err != nil {
				return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
					"Invalid SpiPbjPaketJenisBelanjaPagu", err.Error())
			}

			subTotalBarang = decimal.NewFromInt(0)
			subTotalModal = decimal.NewFromInt(0)
			subTotalSosial = decimal.NewFromInt(0)
			for j, data := range spiPbjPaketJenisBelanjaPagues {
				spiPbjPaketJenisBelanjaPagues[j].Row = j + 1
				subTotalBarang = subTotalBarang.Add(data.Barang)
				subTotalModal = subTotalModal.Add(data.Modal)
				subTotalSosial = subTotalSosial.Add(data.Sosial)
				totalBarang = totalBarang.Add(data.Barang)
				totalModal = totalModal.Add(data.Modal)
				totalSosial = totalSosial.Add(data.Sosial)
			}

			*spiPbjPaketJenisBelanjaPaguGetReformatResponse = append(*spiPbjPaketJenisBelanjaPaguGetReformatResponse,
				dto.SpiPbjPaketJenisBelanjaPaguGetReformatResponse{
					Row: alfabetCounter[i], PaketName: groupPackageValue.Name,
					Barang: nil, Modal: nil, Sosial: nil,
				})

			for _, data := range spiPbjPaketJenisBelanjaPagues {
				*spiPbjPaketJenisBelanjaPaguGetReformatResponse = append(*spiPbjPaketJenisBelanjaPaguGetReformatResponse,
					dto.SpiPbjPaketJenisBelanjaPaguGetReformatResponse{
						Row: strconv.Itoa(data.Row), SpiAngID: &data.SpiAngID, ThnAngID: &data.ThnAngID,
						SatkerID: &data.SatkerID, PaketName: data.PaketName, Barang: &data.Barang,
						Modal: &data.Modal, Sosial: &data.Sosial, Lainnya: &data.Lainnya,
						GroupPackageValueID: &data.GroupPackageValueID, MethodPbj: data.MethodPbj,
						Rencana1: &data.Rencana1, Rencana2: &data.Rencana2, Rencana3: &data.Rencana3,
						Rencana4: &data.Rencana4, Rencana5: &data.Rencana5, Rencana6: &data.Rencana6,
						Rencana7: &data.Rencana7, Rencana8: &data.Rencana8, Rencana9: &data.Rencana9,
						Rencana10: &data.Rencana10, Rencana11: &data.Rencana11, Rencana12: &data.Rencana12,
						Realisasi1: &data.Realisasi1, Realisasi2: &data.Realisasi2, Realisasi3: &data.Realisasi3,
						Realisasi4: &data.Realisasi4, Realisasi5: &data.Realisasi5, Realisasi6: &data.Realisasi6,
						Realisasi7: &data.Realisasi7, Realisasi8: &data.Realisasi8, Realisasi9: &data.Realisasi9,
						Realisasi10: &data.Realisasi10, Realisasi11: &data.Realisasi11, Realisasi12: &data.Realisasi12,
					})
			}

			*spiPbjPaketJenisBelanjaPaguGetReformatResponse = append(*spiPbjPaketJenisBelanjaPaguGetReformatResponse,
				dto.SpiPbjPaketJenisBelanjaPaguGetReformatResponse{
					PaketName: "SubTotal",
					Barang:    &subTotalBarang, Modal: &subTotalModal, Sosial: &subTotalSosial,
				})

		} // end for

		*spiPbjPaketJenisBelanjaPaguGetReformatResponse = append(*spiPbjPaketJenisBelanjaPaguGetReformatResponse,
			dto.SpiPbjPaketJenisBelanjaPaguGetReformatResponse{
				PaketName: "Total",
				Barang:    &totalBarang, Modal: &totalModal, Sosial: &totalSosial,
			})

		/*
			if len(dealers) <= 0 {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, errors.New("Data not found"))
			}
		*/
		result = &dto.SpiPbjPaketJenisBelanjaPaguGetInfoResponse{
			Datas:          spiPbjPaketJenisBelanjaPaguGetReformatResponse,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
