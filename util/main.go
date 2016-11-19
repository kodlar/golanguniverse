package util

import "fmt"

func RunGuest(ipPort string) {
	fmt.Println("Running guest", ipPort)
}

func RunHost(ipPort string) {
	fmt.Println("Running host", ipPort)
}

func GetLocalNetworkIp() string {
	return "192.168.1.0"
}
