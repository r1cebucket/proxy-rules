package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const (
	SHADOWROCKET_PREFIX = `
[General]
bypass-system = true
skip-proxy = 192.168.0.0/16, 10.0.0.0/8, 172.16.0.0/12, localhost, *.local, captive.apple.com
tun-excluded-routes = 10.0.0.0/8, 100.64.0.0/10, 127.0.0.0/8, 169.254.0.0/16, 172.16.0.0/12, 192.0.0.0/24, 192.0.2.0/24, 192.88.99.0/24, 192.168.0.0/16, 198.51.100.0/24, 203.0.113.0/24, 224.0.0.0/4, 255.255.255.255/32, 239.255.255.250/32
dns-server = system
fallback-dns-server = system
ipv6 = true
prefer-ipv6 = false
dns-fallback-system = false
dns-direct-system = false
icmp-auto-reply = true
always-reject-url-rewrite = false
private-ip-answer = true
# direct domain fail to resolve use proxy rule
dns-direct-fallback-proxy = true
# The fallback behavior when UDP traffic matches a policy that doesn't support the UDP relay. Possible values: DIRECT, REJECT.
udp-policy-not-supported-behaviour = REJECT

[Rule]
`
	SHADOWROCKET_SUFFIX = `
# LAN
IP-CIDR,192.168.0.0/16,DIRECT
IP-CIDR,10.0.0.0/8,DIRECT
IP-CIDR,172.16.0.0/12,DIRECT
IP-CIDR,127.0.0.0/8,DIRECT
# China
GEOIP,CN,DIRECT
# Final
FINAL,PROXY

[Host]
localhost = 127.0.0.1

[URL Rewrite]
^https?://(www.)?g.cn https://www.google.com 302
^https?://(www.)?google.cn https://www.google.com 302
`
)

func main() {
	MODES := []string{
		"sing-box",
		"clash",
		"quan x",
		"matsuri",
		"surge",
		"shadowrocket",
	}

	var domainProxy, domainDirect, domainReject []string

	// PROXY
	PROXY := "amp-api-edge.apps.apple.com push.apple.com inappcheck.itunes.apple.com nexoncdn.co.kr nexon.com nexon.io "
	// Google
	{
		PROXY += "googleapis.com "
	}
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
	// {
	// 	PROXY += "cq.hangame.com "
	// 	PROXY += "nhn.com gslb-gamebase.nhncloudservice.com toast.com "
	// }
	// LinkedIn
	{
		PROXY += "linkedin.com "
	}
	// Leetcode
	{
		PROXY += "leetcode.cn "
	}

	// DIRECT
	DIRECT := ""
	// cn
	{
		DIRECT += "baidu.com qq.com zhihu.com dcarstatic.com byteimg.com 163.com csdn.net qcloudimg.com tencent.com bilibili.com "
	}
	// Common
	{
		DIRECT += "aliyuncs.com "
	}
	// hosts
	{
		DIRECT += "megasrv.de "
		DIRECT += "432104.xyz "
	}
	// Microsoft
	{
		// DIRECT += ""
	}
	// Crusaders Quest
	// {
	// 	DIRECT += "nhnst.com "
	// 	DIRECT += "cq-pvp.hangame.com " // 匹配队友，对手（需要直连）
	// 	DIRECT += "cq-cha.hangame.com " // 进入匹配服务器（需要直连）
	// 	DIRECT += "toastoven.net "
	// 	// DIRECT += "cru.cdn.toastoven.net adam.gslb.toastoven.net api-iaptacc.gslb.toastoven.net "
	// 	DIRECT += "unity3d.com "
	// }
	// Steam
	{
		DIRECT += "akamaihd.net "
	}
	// WeTab
	{
		DIRECT += "wetab.link "
	}
	// Talkatone
	{
		DIRECT += "mobilefuse.com "
	}
	// Other
	{
		DIRECT += "nodeseek.com "
	}

	// REJECT
	REJECT := ""
	// analytic
	{
		REJECT += "app-measurement.com appsflyer.com google-analytics.com openinstall.io "

	}
	// ad
	{
		REJECT += "amazon-adsystem.com doubleclick.net rubiconproject.com adservice.google.com wwads.cn "
	}
	// Baidu
	{
		REJECT += "tieba-ares.cdn.bcebos.com "
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
	case "sing-box":
		type rulesConf struct {
			Rules   []map[string]interface{} `json:"rules"`
			Version int64                    `json:"version"`
		}

		rulesReject := rulesConf{
			Rules: []map[string]interface{}{
				{"domain_suffix": domainReject[:len(domainReject)-1]},
			},
			Version: 1,
		}
		rulesProxy := rulesConf{
			Rules: []map[string]interface{}{
				{"domain_suffix": domainProxy[:len(domainProxy)-1]},
			},
			Version: 1,
		}
		rulesDirect := rulesConf{
			Rules: []map[string]interface{}{
				{"domain_suffix": domainDirect[:len(domainDirect)-1]},
			},
			Version: 1,
		}

		confReject, err := os.Create("./rules/sing-box_reject.conf")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		confProxy, err := os.Create("./rules/sing-box_proxy.conf")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		confDirect, err := os.Create("./rules/sing-box_direct.conf")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		dataReject, _ := json.Marshal(rulesReject)
		dataProxy, _ := json.Marshal(rulesProxy)
		dataDirect, _ := json.Marshal(rulesDirect)

		confReject.Write(dataReject)
		confProxy.Write(dataProxy)
		confDirect.Write(dataDirect)

		defer confReject.Close()
		defer confProxy.Close()
		defer confDirect.Close()
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

		confFallback_HK_JP_SG, err := os.Create("./rules/quan_x_fallback_hk_jp_sg.conf")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer confFallback_HK_JP_SG.Close()
		for _, domain := range domainReject[:len(domainReject)-1] {
			rule = fmt.Sprintf("HOST-SUFFIX,%s,REJECT\n", domain)
			confFallback_HK_JP_SG.Write([]byte(rule))
		}
		for _, domain := range domainProxy[:len(domainProxy)-1] {
			rule = fmt.Sprintf("HOST-SUFFIX,%s,FALLBACK-HK&JP&SG\n", domain)
			confFallback_HK_JP_SG.Write([]byte(rule))
		}
		for _, domain := range domainDirect[:len(domainDirect)-1] {
			rule = fmt.Sprintf("HOST-SUFFIX,%s,DIRECT\n", domain)
			confFallback_HK_JP_SG.Write([]byte(rule))
		}

		confFallback_JP_SG, err := os.Create("./rules/quan_x_fallback_jp_sg.conf")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer confFallback_JP_SG.Close()
		for _, domain := range domainReject[:len(domainReject)-1] {
			rule = fmt.Sprintf("HOST-SUFFIX,%s,REJECT\n", domain)
			confFallback_JP_SG.Write([]byte(rule))
		}
		for _, domain := range domainProxy[:len(domainProxy)-1] {
			rule = fmt.Sprintf("HOST-SUFFIX,%s,FALLBACK-JP&SG\n", domain)
			confFallback_JP_SG.Write([]byte(rule))
		}
		for _, domain := range domainDirect[:len(domainDirect)-1] {
			rule = fmt.Sprintf("HOST-SUFFIX,%s,DIRECT\n", domain)
			confFallback_JP_SG.Write([]byte(rule))
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
	case "shadowrocket":
		conf, err := os.Create("./rules/shadowrocket.conf")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer conf.Close()

		conf.Write([]byte(SHADOWROCKET_PREFIX))

		for _, domain := range domainReject[:len(domainReject)-1] {
			rule = fmt.Sprintf("DOMAIN-SUFFIX,%s,REJECT\n", domain)
			conf.Write([]byte(rule))
		}
		for _, domain := range domainProxy[:len(domainProxy)-1] {
			rule = fmt.Sprintf("DOMAIN-SUFFIX,%s,PROXY\n", domain)
			conf.Write([]byte(rule))
		}
		for _, domain := range domainDirect[:len(domainDirect)-1] {
			rule = fmt.Sprintf("DOMAIN-SUFFIX,%s,DIRECT\n", domain)
			conf.Write([]byte(rule))
		}

		conf.Write([]byte(SHADOWROCKET_SUFFIX))
	default:
		fmt.Println("no such mode:", MODE)
	}
}
