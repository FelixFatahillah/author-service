package models

import (
	"author-service/internal/shared"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Author struct {
	ID string `json:"id" gorm:"type:uuid;primaryKey;not null"`
	shared.BaseModel
	FirstName   string  `json:"first_name,omitempty" gorm:"type:varchar(255);not null"`
	LastName    *string `json:"last_name,omitempty" gorm:"type:varchar(255)"`
	PhoneNumber *string `json:"phone_number,omitempty" gorm:"type:varchar(50);unique"`
	Email       string  `json:"email,omitempty" gorm:"type:varchar(255);unique"`
	Biography   *string `json:"biography" gorm:"type:text"`
	Education   *string `json:"education" gorm:"type:varchar(255)"`
}

func (u *Author) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String() // Generate UUID
	}
	return
}
