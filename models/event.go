package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/wfernandez/rest-api/db"
)

func init() {
	RegisterModel(&Event{})
}

func (Event) TableName() string {
    return "events"
}


type Event struct {	
	ID					int64 `gorm:"primaryKey" json:"id" example:"1"`
	Name				string `binding:"required" gorm:"size:255;not null" json:"name" example:"Tech Conference 2023"`
	Description			string `binding:"required" gorm:"not null" json:"description" example:"A conference about the latest technologies"`
	Location			string `binding:"required" gorm:"not null" json:"location" example:"San Francisco Convention Center"`
	DateTime			time.Time `binding:"required" gorm:"not null" json:"dateTime" example:"2023-12-15T18:00:00Z"`
	UserID			int64 `json:"userId" example:"1"`
	Registrations 	[]Registration `gorm:"foreignKey:EventID" json:"registrations,omitempty"`
	Attendees []User `gorm:"many2many:registrations;foreignKey:ID;joinForeignKey:EventID;References:ID;joinReferences:UserID" json:"attendees,omitempty"`
}

func (e *Event) Save() error {
	return db.GetInstance().DB.Create(e).Error
}

func GetAllEvents()	([]Event, error) {
	var events []Event
	result := db.GetInstance().DB.Find(&events)
	fmt.Println(events)
	if result.Error != nil {
		return nil, result.Error
	}
	return events, nil
}

func GetEvent(eventId int64) (*Event, error) {
	var event Event
	result := db.GetInstance().DB.First(&event, eventId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &event, nil
}

func (event Event) Update() error {
	return db.GetInstance().DB.Save(&event).Error
}

func (event Event) Delete() error {
  return db.GetInstance().DB.Delete(&Event{}, event.ID).Error
}

func (e Event) Register(userId int64) error {
    var count int64
    if err := db.GetInstance().DB.Model(&Registration{}).
        Where("event_id = ? AND user_id = ?", e.ID, userId).
        Count(&count).Error; err != nil {
        return err
    }
    
    if count > 0 {
        return errors.New("user already registered for this event")
    }
    
    registration := Registration{
        EventID: e.ID,
        UserID: userId,
    }
    
    return db.GetInstance().DB.Create(&registration).Error
}

func (e Event) CancelRegistration(userId int64) error {
    result := db.GetInstance().DB.Where("event_id = ? AND user_id = ?", e.ID, userId).
        Delete(&Registration{})
    
    if result.Error != nil {
        return result.Error
    }
    
    if result.RowsAffected == 0 {
        return errors.New("user is not registered for this event")
    }
    
    return nil
}

