package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/pjbgf/go-test/should"
)

func TestMain(t *testing.T) {
	tmpfile, _ := ioutil.TempFile("", "calc-fake-stdout.*")
	os.Stdout = tmpfile
	defer os.Remove(tmpfile.Name())
	os.Args = strings.Split("calc 1 + 1", " ")

	main()

	output, _ := ioutil.ReadFile(tmpfile.Name())
	should.New(t).BeEqual(
		"sum total: 2\n",
		string(output),
		"should sum 1+1 and return 2",
	)
}
