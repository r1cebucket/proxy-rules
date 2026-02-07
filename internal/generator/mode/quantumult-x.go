package mode

import (
	"fmt"
	"os"
	"proxy-rules/internal/rule"
)

type QuanX struct{}

func (m QuanX) GenRules(ruleSet rule.RuleSet, outDir string) error {
	rulesReject, err := os.Create(outDir + "/quantumult_x_reject.conf")
	if err != nil {
		return err
	}
	rulesProxy, err := os.Create(outDir + "/quantumult_x_proxy.conf")
	if err != nil {
		return err
	}
	rulesDirect, err := os.Create(outDir + "/quantumult_x_direct.conf")
	if err != nil {
		return err
	}
	defer rulesReject.Close()
	defer rulesProxy.Close()
	defer rulesDirect.Close()

	// domain_suffix
	for _, domain := range ruleSet.Reject.DomainSuffix {
		rule := fmt.Sprintf("host-suffix,%s\n", domain)
		rulesReject.Write([]byte(rule))
	}
	for _, domain := range ruleSet.Proxy.DomainSuffix {
		rule := fmt.Sprintf("host-suffix,%s\n", domain)
		rulesProxy.Write([]byte(rule))
	}
	for _, domain := range ruleSet.Direct.DomainSuffix {
		rule := fmt.Sprintf("host-suffix,%s\n", domain)
		rulesDirect.Write([]byte(rule))
	}
	// domain_keyword
	for _, domain := range ruleSet.Reject.DomainKeyword {
		rule := fmt.Sprintf("host-keyword,%s\n", domain)
		rulesReject.Write([]byte(rule))
	}
	for _, domain := range ruleSet.Proxy.DomainKeyword {
		rule := fmt.Sprintf("host-keyword,%s\n", domain)
		rulesProxy.Write([]byte(rule))
	}
	for _, domain := range ruleSet.Direct.DomainKeyword {
		rule := fmt.Sprintf("host-keyword,%s\n", domain)
		rulesDirect.Write([]byte(rule))
	}

	return nil
}
