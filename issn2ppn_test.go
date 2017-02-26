package gosudoc

import (
	"fmt"
	"reflect"
	"testing"
)

type issnTest struct {
	Input    []string
	Expected map[string][]string
}

func TestIssn2ppn(t *testing.T) {

	var Issn2Tests = []issnTest{
		{
			Input:    []string{"02672472"},
			Expected: map[string][]string{"02672472": {"013630687"}},
		},
		{
			Input:    []string{"02672472", "0296-2454"},
			Expected: map[string][]string{"02672472": {"013630687"}, "0296-2454": {"001020617"}},
		},
	}

	for _, test := range Issn2Tests {
		actual, _ := Issn2ppn(test.Input)
		if reflect.DeepEqual(test.Expected, actual) {
			t.Logf("PASS: got %v", test.Expected)
		} else {
			t.Fatalf("FAIL for %s: expected %v, actual result was %v", test.Input, test.Expected, actual)
		}
	}
}

func ExampleIssn2ppn() {
	myInput := []string{"02672472"}

	result, _ := Issn2ppn(myInput)
	for k, v := range result {
		fmt.Println(k)
		for _, value := range v {
			fmt.Println(value)
		}
	}

	// Output:
	// 02672472
	// 013630687

}
