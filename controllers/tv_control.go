package controllers

import (
	"github.com/EtoNeJa00/Rest_Api_TV/models"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
	"net/http"
	"database/sql"
	DB "github.com/EtoNeJa00/Rest_Api_TV/database"
)

type TVHandler struct {
	Db *sql.DB
}
func (h *TVHandler) GetTVs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	TVs, err := DB.GetAllTV(h.Db)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(TVs)
}
func (h *TVHandler) GetTV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}	
	
	tv, err := DB.GetTVbyID(h.Db, id)
	if (err!= nil)||(tv.Id==0) {
		w.WriteHeader(http.StatusNotFound)
		return	
	}	

	json.NewEncoder(w).Encode(tv)
}

func (h *TVHandler) AddTV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var tv models.TVModel 
	var tempTV models.JsonTVDec
	err := json.NewDecoder(r.Body).Decode(&tempTV)
	if (err!= nil){
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	err = tv.Parse(tempTV)
	if (err!= nil){
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = DB.AddTV(h.Db, tv)
	if (err!= nil){
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tv)
}

func (h *TVHandler) DeleteTV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err!= nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	tv, err := DB.GetTVbyID(h.Db, id)
	if (err!= nil)||(tv.Id==0) {
		w.WriteHeader(http.StatusNotFound)
		return
	}	
	
	err = DB.DeleteTV (h.Db, id)
	if (err!= nil){
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tv)
}
