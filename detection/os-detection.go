package detection

import "runtime"

func OsDetection() string {
	var osType string
	if runtime.GOOS == "windows" {
		osType = "Windows"
	} else if runtime.GOOS == "linux" {
		osType = "Linux"
	} else {
		osType = "Mac"
	}

	return osType
}
