package model

import (
	"code-boiler/internal/abstractions"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	//model standart design
	abstractions.Model

	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Password     string `json:"password" gorm:"-"`
	Status       string `json:"status"`
	IsActive     bool   `json:"is_active"`

	//relations definitions
	Samples []Sample `json:"samples" gorm:"foreignKey:UserId"`
}

func (m *User) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = time.Now()
	if m.Password != "" {
		m.hashPassword()
	}
	return
}

func (m *User) BeforeCreate(tx *gorm.DB) (err error) {
	m.hashPassword()
	return
}

func (m *User) hashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	m.PasswordHash = string(bytes)
}

func (m *User) GenerateToken() (string, error) {
	var (
		jwtKey = os.Getenv("JWT_KEY")
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": m.Email,
		"name":  m.Name,
		"phone": m.Phone,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	return tokenString, err
}
