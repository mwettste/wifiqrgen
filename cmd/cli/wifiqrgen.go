package main

import (
	"flag"
	"fmt"

	"github.com/skip2/go-qrcode"
)

func main() {
	var ssid string
	var password string

	flag.StringVar(&ssid, "ssid", "", "Defines the SSID of the Wifi.")
	flag.StringVar(&password, "pass", "", "Defines the password of the Wifi.")
	flag.Parse()

	qrcodeText := fmt.Sprintf("WIFI:T:WPA;S:%s;P:%s", ssid, password)

	fmt.Println("Generating code...")

	qc, err := qrcode.New(qrcodeText, qrcode.Medium)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Writing file...")
	qc.WriteFile(256, "wificode.png")
}
