package utils

import "os"

func MustGetEnv(name string) string {
	param := os.Getenv(name)
	if param == "" {
		panic("can not found env " + name)
	}
	return param
}
