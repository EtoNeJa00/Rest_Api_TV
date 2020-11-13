package configuraton

import (
	"encoding/json"
	"fmt"
	"os"

)

type Configuration struct {
	Version  string   `json:"Version"`
	DataBase DataBase `json:"DataBase"`
}
type DataBase struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
	IP       string `json:"ip"`
	Port     string `json:"port"`
}


func GetConfig() *Configuration {
	file, err := os.Open("configuraton/configuration.json")
	if err != nil{fmt.Print(err)}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err1 := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err1)
	}
	return &configuration
}
