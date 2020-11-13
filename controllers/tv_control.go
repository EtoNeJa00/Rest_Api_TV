package controllers

import (
	"fmt"
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
	TVs := DB.GetAllTV(h.Db)
	json.NewEncoder(w).Encode(TVs)
}
func (h *TVHandler) GetTV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	tv := DB.GetTVbyID(h.Db, id)
	if (err!= nil)||(id==0) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
   json.NewEncoder(w).Encode(tv)
}
func (h *TVHandler) AddTV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tv models.TVModel 
	var tempTV models.JsonTVDec
	errj := json.NewDecoder(r.Body).Decode(&tempTV)
	tv.Parse(tempTV)
	err :=	DB.AddTV(h.Db, tv)
	fmt.Println(err)
	if (err!=nil) || (errj!=nil){
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errj)
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(tv)
}
func (h *TVHandler) DeleteTV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err!= nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	DB.DeleteTV (h.Db, id)
}
