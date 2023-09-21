package main

type RequestBody struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
	OTP         string `json:"otp"`
	Language    string `json:"language"`
}

type ResponseBody struct {
	Status          bool   `json:"status"`
	ResponseCode    string `json:"responseCode"`
	IsSend          bool   `json:"isSent"`
	ResponseMessage string `json:"responseMessage"`
}

type SMSRequest struct {
	Password      string `json:"password"`
	Shortcodetext string `json:"shortcodetext"`
	Tomdn         string `json:"tomdn"`
	Messagetext   string `json:"messagetext"`
	User          string `json:"user"`
	OperatorID    int    `json:"operatorid"`
}
type ValidResponseBody struct {
	Status          bool   `json:"status"`
	ResponseCode    string `json:"responseCode"`
	Ticket          string `json:"ticket"`
	ResponseMessage string `json:"responseMessage"`
}

type AccessDetails struct {
	UserId uint64 `json:"UserId"`
	Exp    int64  `json:"Exp"`
}
