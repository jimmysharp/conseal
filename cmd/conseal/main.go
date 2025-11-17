package main

import (
	"os"

	"github.com/jimmysharp/conseal"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	config, err := conseal.ParseConfig(".conseal.yml")
	if err != nil {
		os.Exit(1)
	}

	a := conseal.NewAnalyzer(config)

	singlechecker.Main(a)
}
