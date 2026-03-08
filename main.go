// OpenCode Auto-comment
// File: /Users/f.l/Coding/geoip/main.go
// Purpose: Auto-comment header added by OpenCode. Review and refine.

package main

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "geoip",
	Short: "geoip is a convenient tool to merge, convert and lookup IP & CIDR from various formats of geoip data.",
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
