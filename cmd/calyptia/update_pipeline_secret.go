package main

import (
	"fmt"

	"github.com/spf13/cobra"

	cloud "github.com/calyptia/api/types"
	"github.com/calyptia/cli/cmd/calyptia/utils"
)

func newCmdUpdatePipelineSecret(config *utils.Config) *cobra.Command {
	return &cobra.Command{
		Use:               "pipeline_secret ID VALUE",
		Short:             "Update a pipeline secret value",
		Args:              cobra.ExactArgs(2),
		ValidArgsFunction: config.CompleteSecretIDs,
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: update secret by its key. The key is unique per pipeline.
			secretID, value := args[0], args[1]
			err := config.Cloud.UpdatePipelineSecret(config.Ctx, secretID, cloud.UpdatePipelineSecret{
				Value: ptrBytes([]byte(value)),
			})
			if err != nil {
				return fmt.Errorf("could not update pipeline secret: %w", err)
			}

			return nil
		},
	}
}
