package abstractions

import "gorm.io/gorm"

type Repository struct {
	DBConnection *gorm.DB
}
