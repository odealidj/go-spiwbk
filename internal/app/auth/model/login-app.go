package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginAppEntity struct {
	Username     string `json:"username" validate:"required" gorm:"index:idx_login_app_name,unique"`
	Passwordhash string `json:"-"`
	Password     string `json:"password" validate:"required" gorm:"-"`
}

type LoginApp struct {

	abstraction.Entity 

	LoginAppEntity

	//Relation
	UserApp UserApp `json:"user_app" gorm:"foreignKey:ID"`

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
	
}

func (m *LoginApp) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateToday()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY
	m.hashPassword()
	return
}

func (m *LoginApp) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateToday()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}

func (m *LoginApp) hashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	m.Passwordhash = string(bytes)
}

func (m *LoginApp) GenerateToken() (string, error) {
	var (
		jwtKey = os.Getenv("JWT_KEY")
		//jwtKey = constant.JWT_KEY
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":          m.ID,
		"username":    m.Username,
		"exp":         time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	return tokenString, err
}