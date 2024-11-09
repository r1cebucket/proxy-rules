package mode

import (
	"fmt"
	"os"
	"proxy-rules/internal/rule"
)

type ShadowRocket struct{}

func (m ShadowRocket) GenRules(ruleSet rule.RuleSet, outDir string) error {
	rules, err := os.Create(outDir + "/shadowrocket.conf")
	if err != nil {
		return err
	}
	defer rules.Close()
	rules.Write([]byte("[Rule]\n"))
	for _, domain := range ruleSet.Reject.DomainSuffix {
		rule := fmt.Sprintf("DOMAIN-SUFFIX,%s,REJECT\n", domain)
		rules.Write([]byte(rule))
	}
	for _, domain := range ruleSet.Proxy.DomainSuffix {
		rule := fmt.Sprintf("DOMAIN-SUFFIX,%s,PROXY\n", domain)
		rules.Write([]byte(rule))
	}
	for _, domain := range ruleSet.Direct.DomainSuffix {
		rule := fmt.Sprintf("DOMAIN-SUFFIX,%s,DIRECT\n", domain)
		rules.Write([]byte(rule))
	}
	return nil
}
