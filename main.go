package main

import (
	"bufio"
	"os"

	"github.com/KIYOMORIDESU/gotestui/collector"
	"github.com/KIYOMORIDESU/gotestui/view"
)

func main() {
	stdin := os.Stdin
	stdinScanner := bufio.NewScanner(stdin)
	tes, results, _ := collector.ReadLogStdout(stdinScanner)
	app := view.CreateApplication(tes, results)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
