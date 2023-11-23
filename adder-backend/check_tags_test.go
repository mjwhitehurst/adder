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
		t.Fatalf(`Test 1 - checkStringsInFile("%s %s %s") - err="%v", want nil`, fileName, string1, string2, err)
	}

	string3 := "checkString3"
	string4 := "checkString4"

	err = checkStringsInFile(fileName, string3, string4)

	if err == nil {
		t.Fatalf(`Test 2 - checkStringsInFile(%s %s %s) - err=nil, want file not exist`, fileName, string3, string4)
	}
}
