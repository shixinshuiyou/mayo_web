package appconf

import (
	"os"
)

type App struct {
	Name       string
	LogLevel   int
	MainHost   string
	MainPort   int64
	StatusHost string
	StatusPort int64
}

func GetEnv() string {
	result := os.Getenv("GO_ENV")
	if result == "" {
		return "dev"
	}
	return result
}

func IsProduct() bool {
	return GetEnv() == "prod"
}
func IsOnline() bool {
	return GetEnv() == "online"
}
func IsDev() bool {
	result := GetEnv() == "dev"
	if result {
		return true
	}
	if !IsProduct() && !IsTest() && !IsOnline() {
		return true
	}
	return false
}

func IsTest() bool {
	return GetEnv() == "test"
}

func GetGoPath() string {
	return os.Getenv("GOPATH")
}
