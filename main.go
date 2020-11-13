package main

import (
	"github.com/gorilla/mux"
	conf "github.com/EtoNeJa00/Rest_Api_TV/configuraton"
	_ "github.com/lib/pq"
	DB "github.com/EtoNeJa00/Rest_Api_TV/database"
	"net/http"
	c "github.com/EtoNeJa00/Rest_Api_TV/controllers"
)



func main() {
	config := conf.GetConfig()
	dbConfig := config.DataBase
	db:= DB.CreateDB(dbConfig)
	
	 tvHandler := c.TVHandler{Db : db}

	r := mux.NewRouter()
	r.HandleFunc("/api/tvs", tvHandler.GetTVs).Methods("GET")
	r.HandleFunc("/api/tvs", tvHandler.AddTV).Methods("POST")
	r.HandleFunc("/api/tvs/{id}", tvHandler.GetTV).Methods("GET")
	r.HandleFunc("/api/tvs/{id}", tvHandler.DeleteTV).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}



