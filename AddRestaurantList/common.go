package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/motaz/codeutils"
)

func GetConfigurationParameter(param, defaultValue string) string {

	value := codeutils.GetConfigValue("config.ini", param)
	if value == "" {
		value = defaultValue
	}
	return value
}

func readStringFromConfig(key string, defaultValue string) string {
	value := codeutils.GetConfigValue("config.ini", key)
	if value == "" {
		return defaultValue
	}
	return value
}

func writeLog(logtext ...string) {
	var logstring = ""
	for _, log := range logtext {
		logstring += log + " "
	}

	codeutils.WriteToLog(logstring, "log/info")
	mode := readStringFromConfig("mode", "")
	if mode == "dev" {
		fmt.Println(logtext)
	}
}

func writeLogProxy(logtext ...string) {
	var logstring = ""
	for _, log := range logtext {
		logstring += log + " "
	}

	codeutils.WriteToLog(logstring, "log/proxy")
	mode := readStringFromConfig("mode", "")
	if mode == "dev" {
		fmt.Println(logtext)
	}
}

func writeErrorLog(logtext ...string) {
	var logstring = ""
	for _, log := range logtext {
		logstring += log + " "
	}

	codeutils.WriteToLog(logstring, "log/error")
	fmt.Println(logtext)
}
func convertStringToInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		//writeErrorLog("Error Converting String at convertStringToInt ", err.Error())
		return -404
	}
	return i
}

func randomInt(min int, max int) int {

	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min+1) + min
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	s := ""
	for i := 0; i < len; i++ {
		s += fmt.Sprintf("%d", randomInt(0, 9))
	}
	return s
}

func floatToString(f float64) string {
	return fmt.Sprintf("%.0f", f*100000)
}

func parseUnixToDate(unix int64) {
	tm := time.Unix(unix, 0)
	fmt.Println(tm)
}

func RandStringBytes(n int) string {
	const letterBytes = "0123456789"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

