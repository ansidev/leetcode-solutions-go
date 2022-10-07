# LeetCode Starter Go

## Introduction

This is a starter project for solving LeetCode problem using Go.

## How to use?

1. Clone this repository.
2. Update LICENSE (ex, change the author name).
3. Run `make prepare` to install necessary packages.
4. Run `task lc -- {leetcode_problem_number}` to generate code for solving new LeetCode problem using Go.
5. Run `task test` to test all of your code.
6. Update your progress in [PROBLEM.md](/PROBLEMS.md).

## Command line

The command line was built with [urfave/cli](https://github.com/urfave/cli).
```
NAME:
   lc - Generate code to solve a specific LeetCode problem

USAGE:
   Run lc {problem_number}, ex.: lc 1234

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## Contact

Le Minh Tri [@ansidev](https://ansidev.xyz/about).

## License

This source code is available under the [MIT License](/LICENSE).
