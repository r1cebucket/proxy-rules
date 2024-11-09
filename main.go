package main

import (
	"context"
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

var (
	globalCtx context.Context
	modes     []string
)

func main() {
	parseInputFlags()

	ruleSet, err := rule.ReadConf(ruleConfPath)
	if err != nil {
		log.Printf("read rule conf failed: %+v", err)
	}

	for _, mode := range modes {
		if gen, ok := generator.ModeGenerator[mode]; ok {
			if err := gen.GenRules(ruleSet, ruleOutDir); err != nil {
				log.Printf("gen rules for %s failed: %+v", mode, err)
			}
		}
	}
}

func parseInputFlags() {
	var (
		modesStr     string
		modesAllowed []string
	)
	for mode, ok := range generator.MODES_ALLOWED {
		if ok && generator.MODES_ALLOWED[mode] {
			modesAllowed = append(modesAllowed, mode)
		}
	}
	sort.Strings(modesAllowed)

	var rootCmd = &cobra.Command{
		Use:   "main",
		Short: "modes allowed: " + strings.Join(modesAllowed, ","),
		PreRun: func(cmd *cobra.Command, args []string) {
			// --help or -h
			helpFlag, _ := cmd.Flags().GetBool("help")
			if helpFlag {
				cmd.Help()
				os.Exit(0)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			modes = strings.Split(modesStr, ",")
		},
	}

	rootCmd.Flags().StringVarP(&modesStr, "modes", "m", strings.Join(modesAllowed, ","), "modes which rules need to be generated, sep with comma")

	if err := rootCmd.Execute(); err != nil {
		log.Println("parse input failed:", err.Error())
		return
	}
}
