package cmd

import (
  "golang.org/x/net/context"
  "gopkg.in/olivere/elastic.v5"
  "github.com/spf13/cobra"
)

var client *elastic.Client

var ctx context.Context

var RootCmd = &cobra.Command {
  Use: "es-jmdict",
  Short: "",
  PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
    ctx = context.Background()
    var err error
    client, err = elastic.NewSimpleClient()
    return err
  },
}

