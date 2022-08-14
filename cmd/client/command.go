package main

import (
	"github.com/fajarardiyanto/afaik-svc-client-news/cmd/client/rest"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "afaik-service",
	Short: "AFAIK Service Backend",
}

func init() {
	rootCmd.AddCommand(rest.RESTCmd)
}

func Run(args []string) error {
	rootCmd.SetArgs(args)
	return rootCmd.Execute()
}
