package main

type ResponseBodyForResrvation struct {
	Status          bool   `json:"status"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	ReservationID   string `json:"reservationID"`
}

type ResponseBodyWithID struct {
	Status          bool   `json:"status"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	UserID          string `json:"userID"`
}

type ResponseBody struct {
	Status          bool   `json:"status"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

// type UpdateReservation struct {
// 	ServiceCode  string `json:"serviceCode"`
// 	ServicePrice string `json:"servicePrice"`
// 	ServiceType  string `json:"ServiceType"`
// 	Description  string `json:"description"`
// 	HomePage     string `json:"homePage"`
// }

type ResetPasswordInfo struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type UserInformation struct {
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Gender         string `json:"gender"`
	Birthdate      string `json:"birthdate"`
	FoodPreference string `json:"foodPreference"`
}

type RequestBody struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
	OTP         string `json:"otp"`
	Language    string `json:"language"`
}

type UserInfo struct {
	UserID      string `json:"userID"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Gender      string `json:"gender"`
	Birthdate   string `json:"birthdate"`
	PhoneNumber string `json:"phoneNumber"`
}
