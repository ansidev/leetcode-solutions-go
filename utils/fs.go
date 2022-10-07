package utils

import (
	"fmt"
	"os"
	"text/template"
)

func GenerateProblemFiles(problemName string) {
	// variables
	vars := make(map[string]interface{})
	vars["ProblemPackageName"] = "problem"
	problemDir := "problem"

	// parse the template
	problemTmpl, _ := template.ParseFiles("templates/problem.tmpl")
	problemTestTmpl, _ := template.ParseFiles("templates/problem_test.tmpl")

	// create new files
	problemFile, _ := os.Create(fmt.Sprintf("%s/%s.go", problemDir, problemName))
	problemTestFile, _ := os.Create(fmt.Sprintf("%s/%s_test.go", problemDir, problemName))
	defer problemFile.Close()
	defer problemTestFile.Close()

	// apply the template to the vars map and write the result to file.
	problemTmpl.Execute(problemFile, vars)
	fmt.Printf("Generated file %s.go\n", problemName)
	problemTestTmpl.Execute(problemTestFile, vars)
	fmt.Printf("Generated file %s_test.go\n", problemName)
}
