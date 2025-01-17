package cmd

import (
	"github.com/spf13/cobra"

	"github.com/calyptia/cli/cmd/agent"
	cnfg "github.com/calyptia/cli/cmd/config"
	"github.com/calyptia/cli/cmd/coreinstance"
	"github.com/calyptia/cli/cmd/endpoint"
	"github.com/calyptia/cli/cmd/environment"
	"github.com/calyptia/cli/cmd/fleet"
	"github.com/calyptia/cli/cmd/ingestcheck"
	"github.com/calyptia/cli/cmd/pipeline"
	"github.com/calyptia/cli/cmd/tracesession"
	cfg "github.com/calyptia/cli/config"
)

func newCmdDelete(config *cfg.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete core instances, pipelines, etc.",
	}

	cmd.AddCommand(
		agent.NewCmdDeleteAgents(config),
		agent.NewCmdDeleteAgent(config),
		fleet.NewCmdDeleteFleet(config),
		fleet.NewCmdDeleteFleetFile(config),
		pipeline.NewCmdDeletePipeline(config),
		pipeline.NewCmdDeletePipelines(config),
		endpoint.NewCmdDeleteEndpoint(config),
		pipeline.NewCmdDeletePipelineFile(config),
		pipeline.NewCmdDeletePipelineClusterObject(config),
		coreinstance.NewCmdDeleteCoreInstance(config, nil),
		coreinstance.NewCmdDeleteCoreInstanceFile(config),
		coreinstance.NewCmdDeleteCoreInstanceSecret(config),
		coreinstance.NewCmdDeleteCoreInstances(config),
		environment.NewCmdDeleteEnvironment(config),
		tracesession.NewCmdDeleteTraceSession(config),
		cnfg.NewCmdDeleteConfigSection(config),
		ingestcheck.NewCmdDeleteIngestCheck(config),
	)

	return cmd
}
