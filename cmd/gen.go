package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var MODES []string
	var domainProxy, domainDirect []string

	MODES = []string{
		"clash",
		"quan x",
		"matsuri",
	}

	// TODO
	//

	// PROXY
	PROXY := "amp-api-edge.apps.apple.com push.apple.com inappcheck.itunes.apple.com app-measurement.com nexoncdn.co.kr nexon.com nexon.io "
	// Microsoft
	{
		PROXY += "windows.com windows.net office.com microsoft.com "
		PROXY += "contentsync.onenote.com hierarchyapi.onenote.com www.onenote.com "
	}
	// Apple
	PROXY += "app.adjust.com "
	// Crusaders Quest
	PROXY += "hangame.com "
	// PROXY += "cq.hangame.com cq-pvp.hangame.com cq-cha.hangame.com "
	PROXY += "nhn.com gslb-gamebase.nhncloudservice.com toast.com "

	// DIRECT
	DIRECT := ""
	// Microsoft
	DIRECT += "microsoftonline.com sharepoint.com office.net live.com onenote.com "
	// Crusaders Quest
	DIRECT += "nhnst.com "
	DIRECT += "toastoven.net "
	// DIRECT += "cru.cdn.toastoven.net adam.gslb.toastoven.net api-iaptacc.gslb.toastoven.net "
	DIRECT += "unity3d.com "
	domainProxy = strings.Split(PROXY, " ")
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
