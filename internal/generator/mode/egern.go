package mode

import (
	"os"
	"proxy-rules/internal/rule"

	"gopkg.in/yaml.v2"
)

type Egern struct{}

func (m Egern) GenRules(ruleSet rule.RuleSet, outDir string) error {
	type rulesConf struct {
		NoResolve       bool     `yaml:"no_resolve"`
		DomainSuffixSet []string `yaml:"domain_suffix_set"`
	}

	confReject := rulesConf{
		NoResolve:       true,
		DomainSuffixSet: ruleSet.Reject.DomainSuffix,
	}
	confProxy := rulesConf{
		NoResolve:       true,
		DomainSuffixSet: ruleSet.Proxy.DomainSuffix,
	}
	confDirect := rulesConf{
		NoResolve:       true,
		DomainSuffixSet: ruleSet.Direct.DomainSuffix,
	}

	rulesReject, err := os.Create(outDir + "/egern_reject.yaml")
	if err != nil {
		return err
	}
	rulesProxy, err := os.Create(outDir + "/egern_proxy.yaml")
	if err != nil {
		return err
	}
	rulesDirect, err := os.Create(outDir + "/egern_direct.yaml")
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

// no_resolve: true
// domain_set:
// - www.google.com
// domain_suffix_set:
// - apple.com
