package models


import (
	"github.com/jinzhu/gorm"
	u "github.com/iamt-chadwick/multi-event-book/utils"
	"net/http"
	"time"
)

type Events struct{
	ID          uint `json:"id" gorm:"primary_key"`
	Name        string      `gorm:"type:varchar(100)" json:"name"`
	Address     string      `gorm:"type:varchar(200)" json:"address"`
	Code        string      `gorm:"not null;unique_index;type:varchar(6)" json:"code"`
	Image       string      `gorm:"not null;type:varchar(255)" json:"image"`
	Details     string      `json:"details"`
	Event_Kinds []EventKind `json:"event_kinds"`
	Active      bool        `gorm:"default:true" json:"active"`
	Date        time.Time
	CreatedAt   time.Time    `gorm:"not null;default:CURRENT_TIMESTAMP" json:"-" `
	UpdatedAt   time.Time    `gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"-" `
	DeletedAt   *time.Time   `sql:"index" json:"-"`
}

type EventKind struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	Name       string    `gorm:"type:varchar(100)" json:"name"`
	Price      float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	Active     bool      `gorm:"not null" json:"active"`
	EventsID   uint      `gorm:"association_foreignkey:event_kinds"`
	CreatedAt time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"-" `
	UpdatedAt time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"-"`
	DeletedAt *time.Time `sql:"index"json:"-"`
}


func AllEvents() map[string] interface{} {
	var result int64
	var allEvents []Events
	errors :=  GetDB().Preload("Event_Kinds").Find(&allEvents).Count(&result).Error
	if errors != nil && errors != gorm.ErrRecordNotFound {
		return u.Message(http.StatusNotFound, "Something went wrong. Please retry")

	}
	if result < 1 {
		return u.Message(http.StatusNotFound, "No event here man.")
	}
	mapD := allEvents
	response := map[string] interface{} {"events": mapD,"total_events": result}
	return response

}

