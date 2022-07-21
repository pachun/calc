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

func TestUnsupportedOperations(t *testing.T) {
	app := Run(main, "3 - 3")

	assert.Equal(t, app.standardOutput(), "Only addition (+) and multiplication (x) are supported.\n")

	app.cleanup()
}

func TestNonNumericOperands(t *testing.T) {
	app1 := Run(main, "Four + 3")

	assert.Equal(t, app1.standardOutput(), "Operands must be numeric.\n")

	app1.cleanup()

	app2 := Run(main, "3 + six")

	assert.Equal(t, app2.standardOutput(), "Operands must be numeric.\n")

	app2.cleanup()
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
