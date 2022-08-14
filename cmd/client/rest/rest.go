package rest

import (
	"github.com/fajarardiyanto/afaik-svc-client-news/internal/config"
	"github.com/fajarardiyanto/afaik-svc-client-news/pkg/adapter/router"
	"github.com/spf13/cobra"
)

var RESTCmd = &cobra.Command{
	Use:   "rest",
	Short: "Start the REST Server",
	RunE:  RestRun,
}

func RestRun(cmd *cobra.Command, args []string) error {
	config.Init()

	router.Router()

	return nil
}
