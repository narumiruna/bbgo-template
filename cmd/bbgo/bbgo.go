package main

import (
	"github.com/c9s/bbgo/pkg/bbgo"
	"github.com/c9s/bbgo/pkg/cmd"
	_ "github.com/narumiruna/bbgo-template"
)

func init() {
	bbgo.SetWrapperBinary()
}

func main() {
	cmd.Execute()
}
