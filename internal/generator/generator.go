package generator

import (
	"encoding/json"
	"fmt"
	"os"
	"proxy-rules/internal/rule"
)

type Generator interface {
	GenRules(ruleSet rule.RuleSet, outDir string) error
}

// TODO: migrate pending
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
