package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"primary_key; unique; type:uuid; column:id; default:uuid_generate_v4()"`
	Name        string    `gorm:"type:varchar(255); NOT NULL" json:"name"`
	Email       string    `gorm:"type:varchar(255)" json:"email"`
	PreferdName int       `gorm:"type:varchar(255) default:18" json:"prefered_name"`
	Address     string    `gorm:"type:varchar(255)" json:"address"`
	TimeRecords
}
