package cmd

import (
	"log/slog"

	"github.com/dusted-go/logging/prettylog"
	"github.com/fr12k/cloudsql-exporter/pkg/version"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cloudsql-exporter",
	Short: "This is tool to export/import data from/to Cloud SQL instances.",
	Long:  `This is tool to export/import data from/to Cloud SQL instances.`,

	PersistentPreRun: setupLogging,
	SilenceUsage:     true,
	SilenceErrors:    true,
}

func init() {
	AddRequiredPersistentFlagShort(RootCmd, "bucket", "b", "The GCP bucket name to export/import data to.")
	AddRequiredPersistentFlagShort(RootCmd, "project", "p", "The GCP project name that contains the Cloud SQL instance.")
	AddRequiredPersistentFlagShort(RootCmd, "instance", "i", "The GCP Cloud SQL instance name to export/import data from.")
	RootCmd.PersistentFlags().String("user", "", "The Cloud SQL user to connect to the database.")
	RootCmd.PersistentFlags().Bool("debug", false, "Enable debug logging including source code lines as well. (default false)")

	RootCmd.Version = version.BuildVersion
}

func setupLogging(ccmd *cobra.Command, _ []string) {
	logOpts := slog.HandlerOptions{
		Level:       slog.LevelInfo,
		AddSource:   false,
		ReplaceAttr: nil,
	}
	if GetBool(ccmd, "debug") {
		logOpts.Level = slog.LevelDebug
		logOpts.AddSource = true
	}
	prettyHandler := prettylog.NewHandler(&logOpts)
	logger := slog.New(prettyHandler)
	slog.SetDefault(logger)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return RootCmd.Execute()
}
