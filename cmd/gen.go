package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var MODES []string
	var domainProxy []string
	var domainDirect []string

	MODES = []string{
		"clash",
		"quan x",
		"matsuri",
	}

	// PROXY
	PROXY := "amp-api-edge.apps.apple.com push.apple.com inappcheck.itunes.apple.com app-measurement.com nexoncdn.co.kr nexon.com nexon.io "
	// Crusaders Quest
	PROXY += "nhn.com adam.gslb.toastoven.net "
	domainProxy = strings.Split(PROXY, " ")

	// DIRECT
	DIRECT := ""
	// Microsoft
	DIRECT += "microsoft.com microsoftonline.com sharepoint.com office.com office.net "
	// Crusaders Quest
	DIRECT += "cq.hangame.com cq-pvp.hangame.com cq-cha.hangame.com gslb-gamebase.nhncloudservice.com " // ios
	DIRECT += "nhnst.com "
	domainDirect = strings.Split(DIRECT, " ")

	for _, MODE := range MODES {
		SaveConfig(domainProxy, domainDirect, MODE)
	}
}

func SaveConfig(domainProxy, domainDirect []string, MODE string) {
	var rule string
	switch MODE {
	case "quan x":
		conf, err := os.Create("./rules/quanx.conf")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer conf.Close()

		for _, domain := range domainProxy[:len(domainProxy)-1] {
			rule = fmt.Sprintf("HOST-SUFFIX,%s,PROXY\n", domain)
			conf.Write([]byte(rule))
		}
		for _, domain := range domainDirect[:len(domainDirect)-1] {
			rule = fmt.Sprintf("HOST-SUFFIX,%s,DIRECT\n", domain)
			conf.Write([]byte(rule))
		}
	case "matsuri":
		confProxy, err := os.Create("./rules/matsuri_proxy.conf")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		confDirect, err := os.Create("./rules/matsuri_direct.conf")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer confProxy.Close()
		defer confDirect.Close()

		for _, domain := range domainProxy[:len(domainProxy)-1] {
			rule = fmt.Sprintf("domain:%s\n", domain)
			confProxy.Write([]byte(rule))
		}
		confDirect.Write([]byte("geosite:cn\n"))
		for _, domain := range domainDirect[:len(domainDirect)-1] {
			rule = fmt.Sprintf("domain:%s\n", domain)
			confDirect.Write([]byte(rule))
		}
	default:
		fmt.Println("no such mode:", MODE)
	}
}
