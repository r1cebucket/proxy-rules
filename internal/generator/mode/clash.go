package mode

import (
	"fmt"
	"os"
	"proxy-rules/internal/rule"

	"gopkg.in/yaml.v2"
)

type Clash struct{}

func (m Clash) GenRules(ruleSet rule.RuleSet, outDir string) error {
	type rulesConf struct {
		Payload []string `yaml:"payload"`
	}

	confReject := rulesConf{
		Payload: []string{},
	}
	confProxy := rulesConf{
		Payload: []string{},
	}
	confDirect := rulesConf{
		Payload: []string{},
	}

	// domain_suffix
	for _, domain := range ruleSet.Reject.DomainSuffix {
		rule := fmt.Sprintf("+.%s", domain)
		confReject.Payload = append(confReject.Payload, rule)
	}
	for _, domain := range ruleSet.Proxy.DomainSuffix {
		rule := fmt.Sprintf("+.%s", domain)
		confProxy.Payload = append(confProxy.Payload, rule)
	}
	for _, domain := range ruleSet.Direct.DomainSuffix {
		rule := fmt.Sprintf("+.%s", domain)
		confDirect.Payload = append(confDirect.Payload, rule)
	}
	// domain_keyword
	for _, domain := range ruleSet.Reject.DomainKeyword {
		rule := fmt.Sprintf("*%s*", domain)
		confReject.Payload = append(confReject.Payload, rule)
	}
	for _, domain := range ruleSet.Proxy.DomainKeyword {
		rule := fmt.Sprintf("*%s*", domain)
		confProxy.Payload = append(confProxy.Payload, rule)
	}
	for _, domain := range ruleSet.Direct.DomainKeyword {
		rule := fmt.Sprintf("*%s*", domain)
		confDirect.Payload = append(confDirect.Payload, rule)
	}

	rulesReject, err := os.Create(outDir + "/clash_reject.yaml")
	if err != nil {
		return err
	}
	rulesProxy, err := os.Create(outDir + "/clash_proxy.yaml")
	if err != nil {
		return err
	}
	rulesDirect, err := os.Create(outDir + "/clash_direct.yaml")
	if err != nil {
		return err
	}
	defer rulesReject.Close()
	defer rulesProxy.Close()
	defer rulesDirect.Close()

	dataReject, _ := yaml.Marshal(confReject)
	dataProxy, _ := yaml.Marshal(confProxy)
	dataDirect, _ := yaml.Marshal(confDirect)

	rulesReject.Write(dataReject)
	rulesProxy.Write(dataProxy)
	rulesDirect.Write(dataDirect)

	return nil
}

// for _, domain := range ruleSet.Reject.DomainSuffix {
// 	rule := fmt.Sprintf("  - '+.%s'\n", domain)
// 	rulesReject.Write([]byte(rule))
// }
// for _, domain := range ruleSet.Proxy.DomainSuffix {
// 	rule := fmt.Sprintf("  - '+.%s'\n", domain)
// 	rulesProxy.Write([]byte(rule))
// }
// for _, domain := range ruleSet.Direct.DomainSuffix {
// 	rule := fmt.Sprintf("  - '+.%s'\n", domain)
// 	rulesDirect.Write([]byte(rule))
// }
