package main

import (
	"math"
	"os"
	"testing"
)

//TestAbs comment needed
func TestAbs(t *testing.T) {
	// TestAbs comment
	got := math.Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}

//TestMain comment needed
func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func main() {
	TestAbs(&testing.T{})

}
