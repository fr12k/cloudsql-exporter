package cmd

import (
	"os"

	"github.com/fr12k/cloudsql-exporter/pkg/version"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cloudsql-exporter",
	Short: "This is tool to export/import data from/to Cloud SQL instances.",
	Long: `This is tool to export/import data from/to Cloud SQL instances.`,
}

func init() {
	AddRequiredPersistentFlagShort(RootCmd, "bucket", "b", "GCS bucket name")
	AddRequiredPersistentFlagShort(RootCmd, "project", "p", "GCP project name")
	AddRequiredPersistentFlagShort(RootCmd, "instance", "i", "Cloud SQL instance name")
	AddRequiredPersistentFlag(RootCmd, "user", "Cloud SQL user")

	RootCmd.Version = version.BuildVersion
}

func AddRequiredPersistentFlagShort(ccmd *cobra.Command ,name, shorthand, usage string) {
	ccmd.PersistentFlags().StringP(name, shorthand, "", usage)
	err := ccmd.MarkPersistentFlagRequired(name)
	if err != nil {
		panic(err)
	}
}

func AddRequiredFlag(ccmd *cobra.Command, ref *string, name, usage string) {
	ccmd.Flags().StringVar(ref, name, "", usage)
	err := ccmd.MarkFlagRequired(name)
	if err != nil {
		panic(err)
	}
}

func AddRequiredPersistentFlag(ccmd *cobra.Command ,name, usage string) {
	ccmd.PersistentFlags().String(name, "", usage)
	err := ccmd.MarkPersistentFlagRequired(name)
	if err != nil {
		panic(err)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}