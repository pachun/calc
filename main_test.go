package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
	app := Run(main, "3 + 3")

	assert.Equal(t, app.standardOutput(), "6\n")

	app.cleanup()
}

func TestMultiplication(t *testing.T) {
	app := Run(main, "3 x 3")

	assert.Equal(t, app.standardOutput(), "9\n")

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

func (commandLineApp commandLineApp) standardOutput() string {
	output, _ := os.ReadFile(commandLineApp.temporaryStandardOutputFile.Name())
	return string(output)
}

func (commandLineApp commandLineApp) cleanup() {
	os.Stdout = commandLineApp.realStandardOutput
	os.Remove(commandLineApp.temporaryStandardOutputFile.Name())
}
