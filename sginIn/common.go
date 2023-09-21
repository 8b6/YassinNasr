package main

import "github.com/motaz/codeutils"

func GetConfigurationParameter(param, defaultValue string) string {

	value := codeutils.GetConfigValue("config.ini", param)
	if value == "" {
		value = defaultValue
	}
	return value
}

//writeLog Wrapper Function for 'WriteToLog' for taking multiple strings and log File Name
func writeLog(logtext ...string) {
	var logstring = ""
	for _, log := range logtext {
		logstring += log + " "
	}

	codeutils.WriteToLog(logstring, "log/info")
	mode := readStringFromConfig("mode", "")
	if mode == "dev" {
		//fmt.Println(logtext)
	}
}

//readStringFromConfig read parameter from 'config.ini' as string
func readStringFromConfig(key string, defaultValue string) string {
	value := codeutils.GetConfigValue("config.ini", key)
	if value == "" {
		return defaultValue
	}
	return value
}
