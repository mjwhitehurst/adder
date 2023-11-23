package main

import (
	"testing"
)

func TestCheckStringsInFile(t *testing.T) {
	fileName := "test_files/stringfile.c"
	string1 := "checkString1"
	string2 := "checkString2"

	err := checkStringsInFile(fileName, string1, string2)

	if err != nil {
		t.Fatalf(`checkStringsInFile("%s %s %s") = %v, want nil`, fileName, string1, string2, err)
	}
}
