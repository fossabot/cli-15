package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

func newCmdGetPipelineStatusHistory(config *config) *cobra.Command {
	var format string
	var pipelineID string
	var last uint64
	cmd := &cobra.Command{
		Use:   "pipeline_status_history",
		Short: "Display latest status history from a pipeline",
		RunE: func(cmd *cobra.Command, args []string) error {
			ss, err := config.cloud.PipelineStatusHistory(config.ctx, pipelineID, last)
			if err != nil {
				return fmt.Errorf("could not fetch your pipeline status history: %w", err)
			}

			switch format {
			case "table":
				tw := table.NewWriter()
				tw.AppendHeader(table.Row{"ID", "Status", "Config ID", "Created at"})
				tw.SetStyle(table.StyleRounded)
				if w, _, err := term.GetSize(int(os.Stdout.Fd())); err == nil {
					tw.SetAllowedRowLength(w)
				}

				for _, s := range ss {
					tw.AppendRow(table.Row{s.ID, s.Status, s.Config.ID, s.CreatedAt})
				}
				fmt.Println(tw.Render())
			case "json":
				err := json.NewEncoder(os.Stdout).Encode(ss)
				if err != nil {
					return fmt.Errorf("could not json encode your pipeline status history: %w", err)
				}
			default:
				return fmt.Errorf("unknown output format %q", format)
			}
			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&format, "output-format", "o", "table", "Output format. Allowed: table, json")
	fs.StringVar(&pipelineID, "pipeline-id", "", "Parent pipeline ID")
	fs.Uint64VarP(&last, "last", "l", 0, "Last `N` pipeline status history entries. 0 means no limit")

	_ = cmd.RegisterFlagCompletionFunc("output-format", config.completeOutputFormat)
	// _ = cmd.RegisterFlagCompletionFunc("pipeline-id", nil) // TODO: complete pipelineID.

	_ = cmd.MarkFlagRequired("pipeline-id") // TODO: use default pipeline ID from config cmd.

	return cmd
}