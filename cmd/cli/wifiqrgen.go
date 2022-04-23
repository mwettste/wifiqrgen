package main

import (
	"flag"
	"fmt"

	"github.com/skip2/go-qrcode"
)

func main() {
	var ssid string
	var password string
	var isHidden bool
	var encryption string

	flag.StringVar(&ssid, "ssid", "", "Defines the SSID of the Wifi (required)")
	flag.StringVar(&password, "pass", "", "Defines the password of the Wifi (required)")
	flag.StringVar(&encryption, "encryption", "WPA", "Encryption type of the Wifi [WPA|WEP] (default: WPA)")
	flag.BoolVar(&isHidden, "hidden", false, "Set this if the SSID is hidden")
	flag.Parse()

	if ssid == "" || password == "" {
		fmt.Println("Missing parameters")
		flag.PrintDefaults()
		return
	}

	if encryption != "WEP" && encryption != "WPA" {
		fmt.Println("Invalid encryption type")
		flag.PrintDefaults()
		return
	}

	hiddenIdentifier := ""
	if isHidden {
		hiddenIdentifier = ";H:true"
	}
	qrcodeText := fmt.Sprintf("WIFI:T:%s;S:%s;P:%s%s", encryption, ssid, password, hiddenIdentifier)

	fmt.Println("Generating code...")

	qc, err := qrcode.New(qrcodeText, qrcode.Medium)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Writing file...")
	qc.WriteFile(256, "wificode.png")
}
