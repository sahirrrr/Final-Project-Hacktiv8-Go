package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	UserID   uint      `json:"user_id"`
	Title    string    `gorm:"not null" json:"title" form:"title" valid:"required~Your Title for photo is required"`
	Caption  string    `json:"caption" form:"caption"`
	PhotoUrl string    `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Your Photo for is required to post"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comments,omitempty"`
	User     *User     `json:"user,omitempty"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
