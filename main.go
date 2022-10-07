package main

import (
	"github.com/ansidev/leetcode/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "lc",
		Usage: "Generate code to solve a specific LeetCode problem",
		UsageText: `Run lc {problem_number}, ex.: lc 1234`,
		Action: cmd.GenerateLeetCodeProblemFilesCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
