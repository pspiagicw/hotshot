package tests

import (
	"os"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func TestMain(m *testing.M) {
	v := m.Run()

	snaps.Clean(m)

	os.Exit(v)
}

func TestTheTest(t *testing.T) {
	if true != true {
		t.Fatalf("Testing framework failed!")
	}
}
