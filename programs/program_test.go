package programs

import (
	"io/fs"
	"os/exec"
	"path/filepath"
	"testing"
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
		t.Log(string(output))
	}
}

func TestPrograms(t *testing.T) {
	tt := getFiles()

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			checkFile(t, test.file)
		})
	}

}
func getFiles() []scriptTest {
	files := []scriptTest{}
	filepath.WalkDir("files", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, scriptTest{name: d.Name(), file: path})
		}

		return nil
	})

	return files

}
