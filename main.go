package main

import (
	"log"
	"os"

	"github.com/ansidev/leetcode/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "lc",
		Usage: "LeetCode Helper Command",
		UsageText: `Generate new problem: lc {problem_number}, ex.: lc 1234

Test a specific problem: lc test {problem_number}, ex: lc test 1234`,
		Action: cmd.LeetCodeCommandHandler,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
