package cmd

import (
  "fmt"
  "os"
  "errors"

  "github.com/nec286/es-jmdict/jmdict"
  "github.com/spf13/cobra"
  "github.com/nec286/xmlstream"
)

var cmd = &cobra.Command {
  Use: "import /path/to/jmdict.xml",
  Short: "import jmdict into elasticsearch",
  RunE: func(cmd *cobra.Command, args []string) error {
    if len(args) < 1 {
      return errors.New("no path supplied")
    }

    file, err := os.Open(args[0])
    if err != nil { return err }
    defer file.Close()

    scanner := xmlstream.NewScanner(file, new(jmdict.Entry))
    scanner.Decoder().Entity = jmdict.Entity
    count := 0
    for scanner.Scan() {
      tag := scanner.Element()
      switch el := tag.(type) {
      case *jmdict.Entry:
        indexEntry(*el)
        count++
        fmt.Printf("\r%d of about 170,000", count)
      }
    }

    return scanner.Err()
  },
}

func indexEntry(entry jmdict.Entry) {
  _, err := client.Index().
    Index("jmdict").
    Type("entry").
    Id(entry.EntSeq).
    BodyJson(entry).
    Do(ctx)
  if err != nil { panic(err) }
}

func init() {
  RootCmd.AddCommand(cmd)
}
