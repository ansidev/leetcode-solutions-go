package cmd

import (
	"fmt"
	"os/exec"

	"github.com/ansidev/leetcode/utils"
	"github.com/urfave/cli/v2"
)

func LeetCodeCommandHandler(c *cli.Context) error {
	args := c.Args().Slice()

	if len(args) == 0 {
		fmt.Print("Invalid input arguments. Showing app usage.\n\n")
		return cli.ShowAppHelp(c)
	}

	var problemName string

	if args[0] == "test" {
		problemName = utils.NormalizeProblemInput(args[1:])
		problemFile := fmt.Sprintf("%s/%s.go", utils.ProblemDir, problemName)
		problemTestFile := fmt.Sprintf("%s/%s_test.go", utils.ProblemDir, problemName)

		cmd := exec.Command("go", "test", "-v", problemFile, problemTestFile)

		stdout, err := cmd.Output()

		if err != nil {
			return err
		}

		fmt.Print(string(stdout))

		return nil
	}

	problemName = utils.NormalizeProblemInput(args[0:])

	if len(problemName) == 0 {
		fmt.Printf("Invalid input string: %s\n", args)
	} else {
		utils.GenerateProblemFiles(problemName)
	}

	return nil
}
