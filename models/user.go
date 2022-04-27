package models

import (
	"final-project/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Email       string        `gorm:"not null;uniqueIndex" json:"email,omitempty" form:"email" valid:"required~Your email is required"`
	Username    string        `gorm:"not null;uniqueIndex" json:"username,omitempty" form:"username" valid:"required~Your username is required"`
	Password    string        `gorm:"not null" json:"password,omitempty" form:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have minimum length of 6 characters"`
	Age         int           `gorm:"not null" json:"age,omitempty" form:"age" valid:"required~Your age is required"`
	Photos      []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photos,omitempty"`
	Comments    []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comments,omitempty"`
	SocialMedia []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"social_media,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
