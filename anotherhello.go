package main

import (
	"fmt"
	"os"
	"time"
	"github.com/kodlar/anotherhello/util"
)

func main() {
	listenPort := "8080"
	args := os.Args

	if len(args) > 1 {
		
		hostIP := args[1]
		fmt.Println("oyun merkezine hoşgeldiniz")
		
		//fmt.Println(args[1])
		
		util.RunGuest(hostIP)
	} else {
		fmt.Println("saatimiz", time.Now())
		var listenIP string
		listenIP = util.GetLocalNetworkIp()
		util.RunHost(listenIP + ":" + listenPort)
	}

}
