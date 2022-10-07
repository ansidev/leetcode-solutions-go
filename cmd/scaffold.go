package cmd

import (
	"fmt"
	"github.com/ansidev/leetcode/utils"
	"github.com/urfave/cli/v2"
)

func GenerateLeetCodeProblemFilesCommand(c *cli.Context) error {
	sl := c.Args().Slice()

	problem := utils.NormalizeProblemInput(sl[1:])

	if len(problem) == 0 {
		fmt.Printf("Invalid input string")
	} else {
		utils.GenerateProblemFiles(problem)
	}

	return nil
}
