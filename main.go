package main

import (
  "fmt"
  "os"

  "github.com/nec286/es-jmdict/cmd"
)

func main() {
  if err := cmd.RootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }
}

