package main

// import (
// 	"encoding/json"
// 	"log"
// )

type ResponseBody struct {
	Status          bool   `json:"status"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

type UserInfo struct {
	UserID      string `json:"userID"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Gender      string `json:"gender"`
	Birthdate   string `json:"birthdate"`
	PhoneNumber string `json:"phoneNumber"`
}

type Fruit struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

// func main() {
// 	a := Fruit{Name: "Apple", Quantity: 1}
// 	o := Fruit{Name: "Orange", Quantity: 2}

// 	var fs []Fruit
// 	fs = append(fs, a)
// 	fs = append(fs, o)
// 	log.Println(fs)

// 	j, _ := json.Marshal(fs)
// 	log.Println(string(j))

// 	j, _ = json.MarshalIndent(fs, "", "  ")
// 	log.Println(string(j))
// }
