package main

import (
	"math/rand"
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

func randomInt(min int, max int) int {

	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min+1) + min
}
