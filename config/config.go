package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

var Envs map[string]string

func GetD(key string, def string) string {
	if val, ok := Envs[key]; ok {
		return val
	}
	return def
}

func GetR(key string) string {
	if val, ok := Envs[key]; ok {
		return val
	}
	panic(fmt.Sprintf("Environment Variable %q was not set!", key))
}

func init() {
	var err error
	Envs, err = godotenv.Read()
	if err != nil {
		log.Fatalln(err)
	}
}
