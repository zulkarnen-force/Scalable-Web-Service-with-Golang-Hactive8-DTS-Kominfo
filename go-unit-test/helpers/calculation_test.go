package helpers

import (
	"fmt"
	"testing"
)
func TestFailSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 40 {
		t.Fail()
	}

	fmt.Println("TestFailSum eksekusi terhenti")
}


func TestSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		t.FailNow()
	}

	fmt.Println("TestFailSum eksekusi terhenti")
}