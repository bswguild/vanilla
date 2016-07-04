package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type FacebookAuthInfo struct {
	AppId      string
	SecretCode string
  Locale string
}

type AppConfig struct {
	FbAuth FacebookAuthInfo
}

var appConfig AppConfig

func LoadAppConfig() {
	b, err := ioutil.ReadFile("appconfig.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(b, &appConfig)
	if err != nil {
		log.Fatal(err)
	} else {
    log.Println("ok")
  }
}

func GetFbAuthInfo() FacebookAuthInfo {
	return appConfig.FbAuth
}
