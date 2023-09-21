package main

type RequestBody struct {
	Msisdn      string `json:"msisdn"`
	ServiceCode string `json:"serviceCode"`
}

type Subscriber struct {
	Msisdn      string `json:"msisdn"`
	ServiceCode string `json:"serviceCode"`
}

type UserInformation struct {
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Gender         string `json:"gender"`
	birthdate      string `json:"lastname"`
	FoodPreference string `json:"foodPreference"`
}
type UserRegistration struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
	UserType    string `json:"userType"`
}
