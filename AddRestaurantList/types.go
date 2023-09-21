package main

type Service struct {
	ServiceCode  string `json:"serviceCode"`
	ServicePrice string `json:"servicePrice"`
	ServiceType  string `json:"ServiceType"`
	Description  string `json:"description"`
	HomePage     string `json:"homePage"`
}

type ResponseBody struct {
	Status          bool   `json:"status"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

type AccessDetails struct {
	UserId uint64 `json:"UserId"`
	Exp    int64  `json:"Exp"`
}

type UpdateService struct {
	ServiceCode  string `json:"serviceCode"`
	ServicePrice string `json:"servicePrice"`
	ServiceType  string `json:"ServiceType"`
	Description  string `json:"description"`
	HomePage     string `json:"homePage"`
}

type ServiceAfterUpdate struct {
	ServiceCode  string `json:"serviceCode"`
	ServicePrice string `json:"servicePrice"`
	ServiceType  string `json:"ServiceType"`
	Description  string `json:"description"`
	HomePage     string `json:"homePage"`
}
