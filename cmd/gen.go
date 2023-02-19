package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	MODES := []string{
		"clash",
		"quan x",
		"matsuri",
	}

	// TODO
	//

	var domainProxy, domainDirect []string

	// PROXY
	PROXY := "amp-api-edge.apps.apple.com push.apple.com inappcheck.itunes.apple.com app-measurement.com nexoncdn.co.kr nexon.com nexon.io "
	// Microsoft
	{
		PROXY += "windows.com windows.net office.com microsoft.com live.com "
		PROXY += "contentsync.onenote.com hierarchyapi.onenote.com www.onenote.com "
		PROXY += "bing.com "
	}
	// Apple
	{
		PROXY += "app.adjust.com "
	}
	// Crusaders Quest
	{
		PROXY += "cq.hangame.com "
		PROXY += "nhn.com gslb-gamebase.nhncloudservice.com toast.com "
	}

	// DIRECT
	DIRECT := ""
	// Microsoft
	{
		DIRECT += "microsoftonline.com sharepoint.com office.net onenote.com "
	}
	// Crusaders Quest
	{
		DIRECT += "nhnst.com "
		DIRECT += "cq-pvp.hangame.com " // 匹配队友，对手（需要直连）
		DIRECT += "cq-cha.hangame.com " // 进入匹配服务器（需要直连）
		DIRECT += "toastoven.net "
		// DIRECT += "cru.cdn.toastoven.net adam.gslb.toastoven.net api-iaptacc.gslb.toastoven.net "
		DIRECT += "unity3d.com "
	}

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
		conf, err := os.Create("./rules/quan_x.conf")
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
	case "clash":
		confProxy, err := os.Create("./rules/clash_proxy.conf")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		confDirect, err := os.Create("./rules/clash_direct.conf")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer confProxy.Close()
		defer confDirect.Close()

		confProxy.Write([]byte("payload:\n"))
		confDirect.Write([]byte("payload:\n"))

		for _, domain := range domainProxy[:len(domainProxy)-1] {
			rule = fmt.Sprintf("  - '+.%s'\n", domain)
			confProxy.Write([]byte(rule))
		}
		for _, domain := range domainDirect[:len(domainDirect)-1] {
			rule = fmt.Sprintf("  - '+.%s'\n", domain)
			confDirect.Write([]byte(rule))
		}

		conf, err := os.Create("./rules/clash.conf")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer conf.Close()

		for _, domain := range domainProxy[:len(domainProxy)-1] {
			rule = fmt.Sprintf("  - DOMAIN-SUFFIX,%s,PROXY\n", domain)
			conf.Write([]byte(rule))
		}
		for _, domain := range domainDirect[:len(domainDirect)-1] {
			rule = fmt.Sprintf("  - DOMAIN-SUFFIX,%s,DIRECT\n", domain)
			conf.Write([]byte(rule))
		}
	default:
		fmt.Println("no such mode:", MODE)
	}
}
