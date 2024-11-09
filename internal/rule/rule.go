package rule

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type RuleSet struct {
	Reject Rules `toml:"reject"`
	Proxy  Rules `toml:"proxy"`
	Direct Rules `toml:"direct"`
}

type Rules struct {
	DomainSuffix []string `toml:"domain_suffix"`
}

func ReadConf(path string) (RuleSet, error) {
	var filePaths []string

	s, err := os.Stat(path)
	if err != nil {
		return RuleSet{}, err
	}

	if s.IsDir() {
		ents, err := os.ReadDir(path)
		if err != nil {
			return RuleSet{}, err
		}
		for _, ent := range ents {
			if ent.IsDir() {
				continue // only read 1st layer
			}
			filePaths = append(filePaths, fmt.Sprintf("%s/%s", path, ent.Name()))
		}
	}

	// read files and unmarshal
	ruleSet := RuleSet{}
	for _, filePath := range filePaths {
		data, err := os.ReadFile(filePath)
		if err != nil {
			return RuleSet{}, err
		}
		if err := toml.Unmarshal(data, &ruleSet); err != nil {
			return RuleSet{}, err
		}
	}

	return ruleSet, nil
}
