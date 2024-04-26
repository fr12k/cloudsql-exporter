package main

import (
	"log/slog"
	"os"

	"github.com/fr12k/cloudsql-exporter/cmd"
	_ "github.com/fr12k/cloudsql-exporter/cmd/backup"
	_ "github.com/fr12k/cloudsql-exporter/cmd/restore"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		slog.Error("error executing command", "error", err)
		os.Exit(1)
	}
}
