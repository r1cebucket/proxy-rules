package main

import (
	"fmt"
	"log"
	"os"
	"proxy-rules/internal/generator"
	"proxy-rules/internal/rule"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

const (
	ruleConfPath = "./rules"
	ruleOutDir   = "./data/rules"
)

var rootCmd = &cobra.Command{
	Use:   "proxy-rules",
	Short: "Generate proxy rules for various proxy tools",
	RunE:  run,
}

func init() {
	modesAllowed := getAllowedModes()
	rootCmd.Flags().StringP("modes", "m", strings.Join(modesAllowed, ","),
		"modes which rules need to be generated, separated by comma. Available modes: "+strings.Join(modesAllowed, ","))
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	if err := os.MkdirAll(ruleOutDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}

	ruleSet, err := rule.ReadConf(ruleConfPath)
	if err != nil {
		return fmt.Errorf("read rule conf failed: %w", err)
	}

	modesStr, err := cmd.Flags().GetString("modes")
	if err != nil {
		return fmt.Errorf("failed to get modes flag: %w", err)
	}

	var modes []string
	if len(modesStr) > 0 {
		modes = strings.Split(modesStr, ",")
	} else {
		modes = getAllowedModes()
	}

	for _, mode := range modes {
		mode = strings.TrimSpace(mode)
		if mode == "" {
			continue
		}
		if gen, ok := generator.ModeGenerator[mode]; ok {
			if err := gen.GenRules(ruleSet, ruleOutDir); err != nil {
				return fmt.Errorf("gen rules for %s failed: %w", mode, err)
			}
		}
	}

	return nil
}

func getAllowedModes() []string {
	var modesAllowed []string
	for mode := range generator.MODES_ALLOWED {
		if generator.MODES_ALLOWED[mode] {
			modesAllowed = append(modesAllowed, mode)
		}
	}
	sort.Strings(modesAllowed)
	return modesAllowed
}
