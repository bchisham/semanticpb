package commands

import (
	"context"
	"github.com/spf13/cobra"
)

var rdfCmd = &cobra.Command{
	Use:   "rdf",
	Short: "extract annotations as rdf triples",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		RunApp(NewRdfService)
	},
}

func init() {
	rootCmd.AddCommand(rdfCmd)
}

type rdfService struct{}

func NewRdfService() Service {
	return &rdfService{}
}

func (rs *rdfService) Start(ctx context.Context) error {
	return nil
}

func (rs *rdfService) Stop(ctx context.Context) error {
	return nil
}
