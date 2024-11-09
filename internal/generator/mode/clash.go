package mode

import (
	"fmt"
	"os"
	"proxy-rules/internal/rule"
)

type Clash struct{}

func (m Clash) GenRules(ruleSet rule.RuleSet, outDir string) error {
	rulesReject, err := os.Create(outDir + "/clash_reject.conf")
	if err != nil {
		return err
	}
	rulesProxy, err := os.Create(outDir + "/clash_proxy.conf")
	if err != nil {
		return err
	}
	rulesDirect, err := os.Create(outDir + "/clash_direct.conf")
	if err != nil {
		return err
	}
	defer rulesReject.Close()
	defer rulesProxy.Close()
	defer rulesDirect.Close()

	rulesReject.Write([]byte("payload:\n"))
	rulesProxy.Write([]byte("payload:\n"))
	rulesDirect.Write([]byte("payload:\n"))

	for _, domain := range ruleSet.Reject.DomainSuffix {
		rule := fmt.Sprintf("  - '+.%s'\n", domain)
		rulesReject.Write([]byte(rule))
	}
	for _, domain := range ruleSet.Proxy.DomainSuffix {
		rule := fmt.Sprintf("  - '+.%s'\n", domain)
		rulesProxy.Write([]byte(rule))
	}
	for _, domain := range ruleSet.Direct.DomainSuffix {
		rule := fmt.Sprintf("  - '+.%s'\n", domain)
		rulesDirect.Write([]byte(rule))
	}

	return nil
}
