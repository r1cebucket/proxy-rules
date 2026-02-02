package mode

import (
	"encoding/json"
	"os"
	"proxy-rules/internal/rule"
)

type SingBox struct{}

func (m SingBox) GenRules(ruleSet rule.RuleSet, outDir string) error {
	type rulesConf struct {
		Rules   []map[string]interface{} `json:"rules"`
		Version int64                    `json:"version"`
	}

	confReject := rulesConf{
		Rules: []map[string]interface{}{
			{"domain_suffix": ruleSet.Reject.DomainSuffix},
		},
		Version: 1,
	}
	confProxy := rulesConf{
		Rules: []map[string]interface{}{
			{"domain_suffix": ruleSet.Proxy.DomainSuffix},
		},
		Version: 1,
	}
	confDirect := rulesConf{
		Rules: []map[string]interface{}{
			{"domain_suffix": ruleSet.Direct.DomainSuffix},
		},
		Version: 1,
	}

	rulesReject, err := os.Create(outDir + "/sing-box_reject.json")
	if err != nil {
		return err
	}
	rulesProxy, err := os.Create(outDir + "/sing-box_proxy.json")
	if err != nil {
		return err
	}
	rulesDirect, err := os.Create(outDir + "/sing-box_direct.json")
	if err != nil {
		return err
	}
	defer rulesReject.Close()
	defer rulesProxy.Close()
	defer rulesDirect.Close()

	dataReject, _ := json.Marshal(confReject)
	dataProxy, _ := json.Marshal(confProxy)
	dataDirect, _ := json.Marshal(confDirect)

	rulesReject.Write(dataReject)
	rulesProxy.Write(dataProxy)
	rulesDirect.Write(dataDirect)

	return nil
}

// confReject, err := os.Create("./rules/sing-box_reject.conf")
// if err != nil {
// 	fmt.Println(err.Error())
// 	return
// }
// confProxy, err := os.Create("./rules/sing-box_proxy.conf")
// if err != nil {
// 	fmt.Println(err.Error())
// 	return
// }
// confDirect, err := os.Create("./rules/sing-box_direct.conf")
// if err != nil {
// 	fmt.Println(err.Error())
// 	return
// }
