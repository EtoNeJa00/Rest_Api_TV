package models

import (
	"errors"
	"strconv"
	"time"
)

type TVModel struct{
	Id int 				`json:"ID"`
	Brand string 		`json:"Brand"`
	Manufacturer string `json:"Manufacturer"`
	Model string		`json:"Model"`
	Year time.Time		`json:"Year"`
}

type JsonTVDec struct{
	Brand string 		`json:"Brand"`
	Manufacturer string `json:"Manufacturer"`
	Model string		`json:"Model"`
	Year string			`json:"Year"`
}

func (t *TVModel) Parse(j JsonTVDec) error{
	year, err := strconv.Atoi(j.Year)
	if err != nil{
		return err
	}

	if (len(j.Manufacturer) >= 3)&&(len(j.Model) >= 2)&&(year >= 2010){
		t.Brand = j.Brand
		t.Manufacturer = j.Manufacturer
		t.Model = j.Model
		t.Year, _ = time.Parse("2006-01-02", j.Year+"-01-01")
		return nil
	}else{ 
		return errors.New("Invalid data")
	}	
}