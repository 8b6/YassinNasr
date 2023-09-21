package main

import (
	"strings"

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

func GetFileType(fileName string) string {
	s := strings.Split(fileName, ".")
	return s[1]
}
