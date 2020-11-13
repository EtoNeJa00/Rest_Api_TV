package database

import (
    "fmt"
    "database/sql"
	"github.com/EtoNeJa00/Rest_Api_TV/models"
	
)

func GetAllTV(db *sql.DB) ([]models.TVModel, error){
    rows, err := db.Query("SELECT * FROM tv")
    if err!=nil {
        return nil, err
    }
    TVs := []models.TVModel{}  
	for rows.Next(){
 
        tv := models.TVModel{}
        err := rows.Scan(&tv.Id,&tv.Brand, &tv.Year, &tv.Manufacturer, &tv.Model)
        if err != nil{
            fmt.Println(err)
            continue
        }
        TVs = append(TVs, tv)
   }
   return TVs, err
}

func GetTVbyID (db *sql.DB, id int) (models.TVModel, error){
   
    row, err := db.Query("SELECT * FROM tv WHERE id=$1;", id)
    tv := models.TVModel{}
    if err != nil{
        return tv, err
    }    
    row.Next()
    err = row.Scan(&tv.Id,&tv.Brand, &tv.Year, &tv.Manufacturer, &tv.Model)
    if err != nil{
        return tv, err
    }
    return tv, nil
}

func AddTV (db *sql.DB, tv models.TVModel) error{
    _, err := db.Exec("INSERT INTO tv (brand, manufacturer, model, year) VALUES ($1,$2,$3,$4)",tv.Brand,tv.Manufacturer,tv.Model,tv.Year)
    return err
}

func DeleteTV (db *sql.DB, id int) error{
    res,  err := db.Exec("DELETE FROM tv WHERE id=$1;", id)
    fmt.Println(res,"hhh")
    return err
}
