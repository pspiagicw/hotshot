package tests

import "testing"

func TestTheTest(t *testing.T) {
	if true != true {
		t.Fatalf("Testing framework failed!")
	}
}
