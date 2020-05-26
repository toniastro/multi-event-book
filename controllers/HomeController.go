package controllers

import (
	"encoding/json"
	"multi-event-booking/models"
	u "multi-event-booking/utils"
	"net/http"
	"time"
)

var results []string


func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

var Detail = func(w http.ResponseWriter, r *http.Request) {

	detail := &models.Users{}
	err := json.NewDecoder(r.Body).Decode(detail) 
	if err != nil {
		u.Respond(w, u.Message(http.StatusInternalServerError, "Invalid request"))
		return
	}
	detail.Create(w)
}
var Verify = func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	detail := &models.Payload{}
	err := json.NewDecoder(r.Body).Decode(detail) 
	if err != nil {
		u.Respond(w, u.Message(http.StatusBadRequest, "Invalid request"))
		return
	}

	detail.Confirm(w)
	//u.Respond(w, resp)
}

