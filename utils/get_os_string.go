package utils

import "runtime"

func GetOSStr() string {
	osStr := runtime.GOOS

	if osStr == "darwin" {
		osStr = "osx"
	}

	return osStr
}