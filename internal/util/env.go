package util

import (
	"log"
	"os"
	"strconv"
)

func GetEnvStr(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatal("Environment variable ", key, " doesn't exist")
	}
	return v
}

func CheckEnvStr(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Println("Environment variable %s doesn't exist", key)
	}
	return v
}

func GetEnvInt(key string) int {
	s := GetEnvStr(key)
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func GetEnvBool(key string) bool {
	s := GetEnvStr(key)
	v, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
