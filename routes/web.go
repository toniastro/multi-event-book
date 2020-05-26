package routes

import (
	"github.com/gorilla/mux"
	"github.com/iamt-chadwick/multi-event-book/controllers"
	"github.com/ichtrojan/thoth"
	"html/template"
	"log"
	"net/http"
)

func Init() *mux.Router {

	route := mux.NewRouter().StrictSlash(true)
	// route.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./public"))))
	route.PathPrefix("/pdf/{code}.pdf").Handler(http.StripPrefix("/pdf/", http.FileServer(http.Dir("./storage/pdf/"))))
	route.NotFoundHandler = http.HandlerFunc(notFound)
	route.HandleFunc("/api/events/payment", controllers.Detail).Methods("POST")
	route.HandleFunc("/api/events/payment/verify", controllers.Verify).Methods("POST")
	route.HandleFunc("/api/events", controllers.GetAllEvents).Methods("GET")
	route.HandleFunc("/api/events/{code}", controllers.GetOneEvent).Methods("GET")
	route.HandleFunc("/api/events/type/{id}", controllers.GetTicket).Methods("GET")
	route.PathPrefix("/").Handler(http.FileServer(http.Dir("./build")))
	return route
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	logger, _ := thoth.Init("log")

	view, err := template.ParseFiles("views/errors/404.html")

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	}

	_ = view.Execute(w, nil)
}
