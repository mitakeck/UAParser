package main

import (
	"flag"
	"fmt"
	"io"

	"github.com/mssola/user_agent"
)

const (
	// ExistCodeOK 正常終了コード
	ExistCodeOK = iota
	// ExistCodeParseFlagError 異常終了コード
	ExistCodeParseFlagError
)

// CLI 入力先と出力先
type CLI struct {
	outStream, errStream io.Writer
}

// Run 実行
func (c *CLI) Run(args []string) int {
	var version bool
	flags := flag.NewFlagSet("uaparser", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.BoolVar(&version, "version", false, "Print version information and quit")

	if err := flags.Parse(args[1:]); err != nil {
		return ExistCodeParseFlagError
	}

	// バージョン表示
	if version {
		fmt.Fprintf(c.errStream, "uaparser - version %s\n", Version)
		return ExistCodeOK
	}

	// user agent
	ua := user_agent.New(flags.Arg(0))

	browser, _ := ua.Browser()
	fmt.Printf("\"%s\", \"%s\"\n", ua.OS(), browser)

	return ExistCodeOK
}
