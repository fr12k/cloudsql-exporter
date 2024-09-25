package cmd

import (
	"github.com/spf13/cobra"
)

func AddRequiredPersistentFlagShort(ccmd *cobra.Command, name, shorthand, usage string) {
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

func GetString(ccmd *cobra.Command, name string) string {
	str, err := ccmd.Flags().GetString(name)
	if err != nil {
		panic(err)
	}
	return str
}

func GetBool(ccmd *cobra.Command, name string) bool {
	b, err := ccmd.Flags().GetBool(name)
	if err != nil {
		panic(err)
	}
	return b
}
