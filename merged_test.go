package gosudoc

import (
	"reflect"
	"testing"
)

var mergedTests = []struct {
	Input    []string
	Expected map[string]string
}{
	{[]string{"071860576"}, map[string]string{"071860576": "007527446"}},
	{[]string{"071860576", "04770280X"}, map[string]string{"071860576": "007527446", "04770280X": "096729856"}},
	{[]string{"071860576", "04770XXXX"}, map[string]string{"071860576": "007527446"}},
	{[]string{"123456789"}, map[string]string{}},
	{[]string{}, map[string]string{}},
}

func TestGetMerged(t *testing.T) {
	for _, test := range mergedTests {
		actual, _ := GetMerged(test.Input)
		if reflect.DeepEqual(test.Expected, actual) {
			t.Logf("PASS: got %v", test.Expected)
		} else {
			t.Fatalf("FAIL for %s: expected %v, actual result was %v", test.Input, test.Expected, actual)
		}
	}
}

// TODO: test with a mock server rather than a live http.Get
