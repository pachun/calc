package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/pjbgf/go-test/should"
)

func TestAddition(t *testing.T) {
	app := Run(main, "3 + 3")

	should.New(t).BeEqual(
		"6\n",
		app.output(),
		"Running 'calc 3 + 3' outputs '6'",
	)

	app.cleanup()
}

func TestMultiplication(t *testing.T) {
	app := Run(main, "3 x 3")

	should.New(t).BeEqual(
		"9\n",
		app.output(),
		"Running 'calc 3 x 3' outputs '9'",
	)

	app.cleanup()
}

type commandLineApp struct {
	realStandardOutput          *os.File
	temporaryStandardOutputFile *os.File
}

func Run(main func(), commandLineArguments string) commandLineApp {
	fakeAppName := "main"
	fullRunCommand := fmt.Sprintf("%s %s", fakeAppName, commandLineArguments)
	os.Args = strings.Split(fullRunCommand, " ")

	realStandardOutput := os.Stdout
	temporaryStandardOutputFile, _ := os.CreateTemp("", "")
	os.Stdout = temporaryStandardOutputFile

	main()

	return commandLineApp{
		realStandardOutput,
		temporaryStandardOutputFile,
	}
}

func (commandLineApp commandLineApp) output() string {
	output, _ := os.ReadFile(commandLineApp.temporaryStandardOutputFile.Name())
	return string(output)
}

func (commandLineApp commandLineApp) cleanup() {
	os.Stdout = commandLineApp.realStandardOutput
	os.Remove(commandLineApp.temporaryStandardOutputFile.Name())
}
