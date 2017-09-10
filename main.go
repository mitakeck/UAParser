package main

import "os"

// Version : ヴァージョン表示用文字列
const Version string = "v0.0.1"

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
