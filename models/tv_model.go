package models

import "time"
//import "Rest_Api_TV"

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
func (t *TVModel) Parse(j JsonTVDec){
	t.Brand = j.Brand
	t.Manufacturer = j.Manufacturer
	t.Model = j.Model
	t.Year, _ = time.Parse("2006-01-02", j.Year+"-01-01")
}