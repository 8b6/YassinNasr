package main

import "github.com/motaz/codeutils"

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
func GetConfigurationParameter(param, defaultValue string) string {

	value := codeutils.GetConfigValue("remvesubscriber.ini", param)
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
