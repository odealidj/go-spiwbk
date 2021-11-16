package migration

import (
	"codeid-boiler/internal/app/model"
	"fmt"
)

func InitData() {
	thnang := []model.ThnAng{
		model.ThnAng{
			ThnAngEntity: model.ThnAngEntity{Year: "2020"},
		},
		model.ThnAng{
			ThnAngEntity: model.ThnAngEntity{Year: "2021"},
		},
		model.ThnAng{
			ThnAngEntity: model.ThnAngEntity{Year: "2022"},
		},
	}

	fmt.Println(thnang)

}
