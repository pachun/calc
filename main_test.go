package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/pjbgf/go-test/should"
)

func TestMain(t *testing.T) {
	assertThat := func(
		assumption string,
		command string,
		expectedOutput string,
	) {
		should := should.New(t)
		tmpfile, _ := ioutil.TempFile("", "calc-fake-stdout.*")
		defer os.Remove(tmpfile.Name())

		os.Stdout = tmpfile
		os.Args = strings.Split(command, " ")

		main()

		output, _ := ioutil.ReadFile(tmpfile.Name())
		actualOutput := string(output)

		should.BeEqual(expectedOutput, actualOutput, assumption)
	}

	assertThat(
		"should sum 1+1 and return 2",
		"calc 1 + 1",
		"sum total: 2\n",
	)
}
