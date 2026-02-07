package mode

import (
	"fmt"
	"os"
	"proxy-rules/internal/rule"
)

type Loon struct{}

func (m Loon) GenRules(ruleSet rule.RuleSet, outDir string) error {
	rulesReject, err := os.Create(outDir + "/loon_reject.conf")
	if err != nil {
		return err
	}
	rulesProxy, err := os.Create(outDir + "/loon_proxy.conf")
	if err != nil {
		return err
	}
	rulesDirect, err := os.Create(outDir + "/loon_direct.conf")
	if err != nil {
		return err
	}
	defer rulesReject.Close()
	defer rulesProxy.Close()
	defer rulesDirect.Close()

	// domain_suffix
	for _, domain := range ruleSet.Reject.DomainSuffix {
		rule := fmt.Sprintf("DOMAIN-SUFFIX,%s\n", domain)
		rulesReject.Write([]byte(rule))
	}
	for _, domain := range ruleSet.Proxy.DomainSuffix {
		rule := fmt.Sprintf("DOMAIN-SUFFIX,%s\n", domain)
		rulesProxy.Write([]byte(rule))
	}
	for _, domain := range ruleSet.Direct.DomainSuffix {
		rule := fmt.Sprintf("DOMAIN-SUFFIX,%s\n", domain)
		rulesDirect.Write([]byte(rule))
	}
	// domain_keyword
	for _, domain := range ruleSet.Reject.DomainKeyword {
		rule := fmt.Sprintf("DOMAIN-KEYWORD,%s\n", domain)
		rulesReject.Write([]byte(rule))
	}
	for _, domain := range ruleSet.Proxy.DomainKeyword {
		rule := fmt.Sprintf("DOMAIN-KEYWORD,%s\n", domain)
		rulesProxy.Write([]byte(rule))
	}
	for _, domain := range ruleSet.Direct.DomainKeyword {
		rule := fmt.Sprintf("DOMAIN-KEYWORD,%s\n", domain)
		rulesDirect.Write([]byte(rule))
	}

	return nil
}
