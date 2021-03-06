package gosudoc

import (
	"fmt"
	"reflect"
	"testing"
)

var multiwhereTests = []struct {
	Input    []string
	Expected map[string][]Library
}{
	{
		Input: []string{"154923206"},

		Expected: map[string][]Library{"154923206": {
			{RCR: "751052105",
				Shortname: "PARIS-Bib. de la Sorbonne - BIS",
				Latitude:  "48.8517361",
				Longitude: "2.3484821",
			},
			{RCR: "751052116",
				Shortname: "PARIS-Bib. Sainte Geneviève",
				Latitude:  "48.8467139",
				Longitude: "2.3463854",
			},
		}},
	},
}

func TestGetMultiwhere(t *testing.T) {
	for _, test := range multiwhereTests {
		actual, _ := GetMultiwhere(test.Input)
		if reflect.DeepEqual(test.Expected, actual) {
			t.Logf("PASS: got %v", test.Expected)
		} else {
			t.Fatalf("FAIL for %s: expected %v, actual result was %v", test.Input, test.Expected, actual)
		}
	}
}

func ExampleGetMultiwhere() {
	myInput := []string{"154923206"}
	result, _ := GetMultiwhere(myInput)
	for k, v := range result {
		fmt.Println(k)
		for _, lib := range v {
			fmt.Printf("RCR: %v / Shortname: %v / Latitude: %v / Longitude: %v\n", lib.RCR, lib.Shortname, lib.Latitude, lib.Longitude)
		}
	}

	// Output:
	//154923206
	//RCR: 751052105 / Shortname: PARIS-Bib. de la Sorbonne - BIS / Latitude: 48.8517361 / Longitude: 2.3484821
	//RCR: 751052116 / Shortname: PARIS-Bib. Sainte Geneviève / Latitude: 48.8467139 / Longitude: 2.3463854

}
