package main

import (
	"fmt"
	"os"

	"github.com/buenocode/caas/pkg/pdf"
	"github.com/buenocode/caas/pkg/screenshot"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "caas",
		Short: "caas",
	}

	cmd.AddCommand(screenshot.Cmd)
	cmd.AddCommand(pdf.Cmd)

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
