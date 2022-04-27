package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserID  uint   `json:"user_id"`
	PhotoID uint   `gorm:"not null" json:"photo_id" form:"photo_id" valid:"required~Your Photo ID for comments is required"`
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Your message for comments is required"`
	User    *User  `json:"user,omitempty"`
	Photo   *Photo `json:"photo,omitempty"`
}

func (p *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
