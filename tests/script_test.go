package tests

import (
	"os/exec"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

type scriptTest struct {
	name string
	file string
}

func checkFile(t *testing.T, filename string) {
	cmd := exec.Command("../hotshot", filename)
	output, err := cmd.Output()
	if err != nil {
		t.Errorf("Error running interpreter! %v", err)
	}
	snaps.MatchSnapshot(t, string(output))
}

func TestScripts(t *testing.T) {
	t.Skip()
	tt := []scriptTest{
		{
			name: "Hello World",
			file: "files/hello-world.ht",
		},
		{
			name: "Data Types",
			file: "files/data-types.ht",
		},
		{
			name: "Arithmetic Types",
			file: "files/arithmetic.ht",
		},
		{
			name: "Booleans",
			file: "files/booleans.ht",
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			checkFile(t, test.file)
		})
	}

}
