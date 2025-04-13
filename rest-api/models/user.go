package models

import (
	"github.com/pkg/errors"
	"github.com/wfernandez/rest-api/db"
	"github.com/wfernandez/rest-api/utils"
)

func init() {
	RegisterModel(&User{})
}

type User struct {
	ID       int64  `gorm:"primaryKey"`
	Email 	 string `binding:"required" gorm:"not null;unique"`
	Password string `binding:"required" gorm:"not null"`

  Registrations []Registration `gorm:"foreignKey:UserID" json:"registrations,omitempty"`  
  Events        []Event        `gorm:"many2many:registrations;joinForeignKey:UserID;joinReferences:EventID" json:"events,omitempty"`	
}

func (u *User) Save() error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return errors.Wrap(err, "failed to hash password")
	}
	u.Password = hashedPassword

	if err := db.GetInstance().DB.Create(u).Error; err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}

func (u *User) Authenticate() error {
	var user User
	if err := db.GetInstance().DB.Where("email = ?", u.Email).First(&user).Error; err != nil {
		return errors.Wrap(err, "invalid credentials")
	}

	if !utils.CheckPasswordHash(u.Password, user.Password) {
		return errors.New("invalid credentials")
	}

	u.ID = user.ID
	return nil
}
