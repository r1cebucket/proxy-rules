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
		"surge",
	}

	// TODO
	//

	var domainProxy, domainDirect, domainReject []string

	// PROXY
	PROXY := "amp-api-edge.apps.apple.com push.apple.com inappcheck.itunes.apple.com nexoncdn.co.kr nexon.com nexon.io "
	// Microsoft
	{
		PROXY += "bing.com "
		PROXY += "windows.com windows.net office.com microsoft.com live.com "
		PROXY += "onenote.com contentsync.onenote.com hierarchyapi.onenote.com "
		PROXY += "microsoftonline.com office.net " // 原来是direct
		PROXY += "sharepoint.com 1drv.com "        // onedrive
		PROXY += "googleapis.cn "
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
	// cn
	{
		DIRECT += "baidu.com qq.com zhihu.com dcarstatic.com byteimg.com 163.com "
	}
	// alist
	{
		DIRECT += "alist.r1cebucket.top "
		DIRECT += "alist-proxy.r1cebucket.top "
	}
	// Microsoft
	{
		// DIRECT += ""
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

	// REJECT
	REJECT := ""
	//
	{
		REJECT += "app-measurement.com appsflyer.com "

	}
	// ad
	{
		REJECT += "amazon-adsystem.com doubleclick.net rubiconproject.com adservice.google.com "
	}

	domainProxy = strings.Split(PROXY, " ")
	domainDirect = strings.Split(DIRECT, " ")
	domainReject = strings.Split(REJECT, " ")

	for _, MODE := range MODES {
		SaveConfig(domainReject, domainProxy, domainDirect, MODE)
	}
}

func SaveConfig(domainReject, domainProxy, domainDirect []string, MODE string) {
	var rule string
	switch MODE {
	case "quan x":
		conf, err := os.Create("./rules/quan_x.conf")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer conf.Close()

		for _, domain := range domainReject[:len(domainReject)-1] {
			rule = fmt.Sprintf("HOST-SUFFIX,%s,REJECT\n", domain)
			conf.Write([]byte(rule))
		}
		for _, domain := range domainProxy[:len(domainProxy)-1] {
			rule = fmt.Sprintf("HOST-SUFFIX,%s,PROXY\n", domain)
			conf.Write([]byte(rule))
		}
		for _, domain := range domainDirect[:len(domainDirect)-1] {
			rule = fmt.Sprintf("HOST-SUFFIX,%s,DIRECT\n", domain)
			conf.Write([]byte(rule))
		}
	case "matsuri":
		confReject, err := os.Create("./rules/matsuri_reject.conf")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
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
		defer confReject.Close()
		defer confProxy.Close()
		defer confDirect.Close()

		for _, domain := range domainReject[:len(domainReject)-1] {
			rule = fmt.Sprintf("domain:%s\n", domain)
			confReject.Write([]byte(rule))
		}
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
		confReject, err := os.Create("./rules/clash_reject.conf")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
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
		defer confReject.Close()
		defer confProxy.Close()
		defer confDirect.Close()

		confReject.Write([]byte("payload:\n"))
		confProxy.Write([]byte("payload:\n"))
		confDirect.Write([]byte("payload:\n"))

		for _, domain := range domainReject[:len(domainReject)-1] {
			rule = fmt.Sprintf("  - '+.%s'\n", domain)
			confReject.Write([]byte(rule))
		}
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

		for _, domain := range domainReject[:len(domainReject)-1] {
			rule = fmt.Sprintf("  - DOMAIN-SUFFIX,%s,REJECT\n", domain)
			conf.Write([]byte(rule))
		}
		for _, domain := range domainProxy[:len(domainProxy)-1] {
			rule = fmt.Sprintf("  - DOMAIN-SUFFIX,%s,PROXY\n", domain)
			conf.Write([]byte(rule))
		}
		for _, domain := range domainDirect[:len(domainDirect)-1] {
			rule = fmt.Sprintf("  - DOMAIN-SUFFIX,%s,DIRECT\n", domain)
			conf.Write([]byte(rule))
		}
	case "surge":
		confReject, err := os.Create("./rules/surge_reject.conf")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		confProxy, err := os.Create("./rules/surge_proxy.conf")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		confDirect, err := os.Create("./rules/surge_direct.conf")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer confReject.Close()
		defer confProxy.Close()
		defer confDirect.Close()

		for _, domain := range domainReject[:len(domainReject)-1] {
			rule = fmt.Sprintf("DOMAIN-SUFFIX,%s\n", domain)
			confReject.Write([]byte(rule))
		}
		for _, domain := range domainProxy[:len(domainProxy)-1] {
			rule = fmt.Sprintf("DOMAIN-SUFFIX,%s\n", domain)
			confProxy.Write([]byte(rule))
		}
		for _, domain := range domainDirect[:len(domainDirect)-1] {
			rule = fmt.Sprintf("DOMAIN-SUFFIX,%s\n", domain)
			confDirect.Write([]byte(rule))
		}
	default:
		fmt.Println("no such mode:", MODE)
	}
}
