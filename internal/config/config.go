package config

import "os"

func Env() string {
	return os.Getenv("ENV")
}

func IsLocal() bool {
	return Env() == "local" || Env() == ""
}
