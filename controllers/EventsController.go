package controllers

import (
	_ "encoding/json"
	"github.com/gorilla/mux"
	"github.com/ichtrojan/thoth"
	"github.com/jinzhu/gorm"
	"multi-event-booking/models"
	u "multi-event-booking/utils"
	"net/http"
	"time"
	_ "time"
)

type Ticket struct {
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	TicketActive bool      `json:"ticket_active"`
	Code         string    `json:"code"`
	EventName    string    `json:"event_name"`
	EventDate    time.Time `json:"event_date"`
	Address      string    `json:"address"`
	Active       bool      `json:"active"`
}
type Sellout struct{
	Active 	  bool       `json:"active"`
}

var db_call = models.GetDB()

var logger, _ = thoth.Init("log")

func GetAllEvents(w http.ResponseWriter, r *http.Request) {

	var events []models.Events

	err := db_call.Preload("Event_Kinds").Where("active = ?", 1).Find(&events).Error

	if err != nil{
		logger.Log(err)
	}

	u.Response(w, http.StatusOK,"", events )
}

 func GetOneEvent(w http.ResponseWriter, r *http.Request) {

	eventCode := mux.Vars(r)["code"]

	temp := models.Events{}

	errors := db_call.Where("code = ?", eventCode).Preload("Event_Kinds").First(&temp).Error

	if errors != nil || errors == gorm.ErrRecordNotFound {

		logger.Log(errors)

		u.Errors(w, http.StatusBadRequest, "This record does not exist",nil)

		return
	}

	u.Response(w,http.StatusOK, "", temp)
}

func GetTicket(w http.ResponseWriter, r *http.Request){
	eventKindId := mux.Vars(r)["id"]

	ticketKind := models.EventKind{}

	events := models.Events{}

    errors := db_call.Where("id = ?", eventKindId).First(&ticketKind).Error

	errors = db_call.Where("id = ?", ticketKind.EventsID).First(&events).Error

	if errors != nil || errors == gorm.ErrRecordNotFound {

		logger.Log(errors)

		u.Errors(w, http.StatusBadRequest, "This record does not exist",nil)

		return
	}
	if !ticketKind.Active {

		u.Response(w,http.StatusOK, "The ticket selected is sold out.", Sellout{Active: false})

		return
	}

	data:= Ticket{
		Name:         ticketKind.Name,
		Price:        ticketKind.Price,
		TicketActive: ticketKind.Active,
		Code:         events.Code,
		EventName:    events.Name,
		EventDate:    events.Date,
		Address:      events.Address,
		Active:       events.Active,
	}

	u.Response(w,http.StatusOK, "", data)
}

