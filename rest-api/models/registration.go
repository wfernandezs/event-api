package models

import "time"

func init() {
	RegisterModel(&Registration{})
}

func (Registration) TableName() string {
	return "registrations"
}

type Registration struct {
	ID 		 int64 `gorm:"primaryKey"`
	EventID	 int64 `gorm:"not null"`
	UserID	 int64 `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Event  *Event `gorm:"foreignKey:EventID" json:"event.omitempty"`
	User   *User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
}