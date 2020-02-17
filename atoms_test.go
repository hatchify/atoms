package atoms

import (
	"os"
	"testing"
)

const (
	testErrInvalidValueFmt = "Invalid value, expected %v and received %v"
	testErrInvalidTypeFmt  = "Invalid type, expected %T and received %T"
	testErrInvalidSwapFmt  = "Swapped successfully when should have failed"
)

var testsink int

func TestMain(m *testing.M) {
	sc := m.Run()
	os.Exit(sc)
}
