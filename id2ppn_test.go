package gosudoc

import (
	"fmt"
	"reflect"
	"testing"
)

type resultTest struct {
	ID2      string
	Input    []string
	Expected map[string][]string
}

func TestID2ppn(t *testing.T) {

	var ID2Tests = []resultTest{
		{
			ID2:      "frbn2ppn",
			Input:    []string{"00000004X"},
			Expected: map[string][]string{"00000004X": {"000000027"}},
		},
		{
			ID2:      "frbn2ppn",
			Input:    []string{"00000051X"},
			Expected: map[string][]string{"00000051X": {"000000167", "02635778X"}},
		},
		{
			ID2:      "frbn2ppn",
			Input:    []string{"00000004X", "00000051X"},
			Expected: map[string][]string{"00000004X": {"000000027"}, "00000051X": {"000000167", "02635778X"}},
		},
		{
			ID2:      "frbn2ppn",
			Input:    []string{"123456789"},
			Expected: map[string][]string{},
		},
		{
			ID2:      "frbn2ppn",
			Input:    []string{},
			Expected: map[string][]string{},
		},
		{
			ID2:      "ocn2ppn",
			Input:    []string{"10002646"},
			Expected: map[string][]string{"10002646": {"025439944"}},
		},
		{
			ID2:      "ocn2ppn",
			Input:    []string{"10002646", "10006281"},
			Expected: map[string][]string{"10002646": {"025439944"}, "10006281": {"025441027"}},
		},
		{
			ID2:      "ocn2ppn",
			Input:    []string{},
			Expected: map[string][]string{},
		},
		{
			ID2:      "ocn2ppn",
			Input:    []string{"123456789"},
			Expected: map[string][]string{},
		},
		{
			ID2:      "dnb2ppn",
			Input:    []string{"015130681"},
			Expected: map[string][]string{"015130681": {"155841408"}},
		},
		{
			ID2:      "dnb2ppn",
			Input:    []string{"015130681", "100024797X"},
			Expected: map[string][]string{"015130681": {"155841408"}, "100024797X": {"149750382"}},
		},
		{
			ID2:      "dnb2ppn",
			Input:    []string{},
			Expected: map[string][]string{},
		},
		{
			ID2:      "dnb2ppn",
			Input:    []string{"123456789"},
			Expected: map[string][]string{},
		},
		{
			ID2:      "ucatb2ppn",
			Input:    []string{"10011536"},
			Expected: map[string][]string{"10011536": {"155976931"}},
		},
		{
			ID2:      "ucatb2ppn",
			Input:    []string{"10011536", "10037641"},
			Expected: map[string][]string{"10011536": {"155976931"}, "10037641": {"154733253"}},
		},
		{
			ID2:      "ucatb2ppn",
			Input:    []string{},
			Expected: map[string][]string{},
		},
		{
			ID2:      "ucatb2ppn",
			Input:    []string{"123456789"},
			Expected: map[string][]string{},
		},
		{
			ID2:      "frcairninfo2ppn",
			Input:    []string{"DBU_ADRIE_2008_01"},
			Expected: map[string][]string{"DBU_ADRIE_2008_01": {"151452512"}},
		},
		{
			ID2:      "frcairninfo2ppn",
			Input:    []string{"DBU_ADRIE_2008_01", "DBU_ALTET_2002_01"},
			Expected: map[string][]string{"DBU_ADRIE_2008_01": {"151452512"}, "DBU_ALTET_2002_01": {"151452547"}},
		},
		{
			ID2:      "frcairninfo2ppn",
			Input:    []string{},
			Expected: map[string][]string{},
		},
		{
			ID2:      "frcairninfo2ppn",
			Input:    []string{"123456789"},
			Expected: map[string][]string{},
		},
		{
			ID2:      "springerln2ppn",
			Input:    []string{"978-0-387-96842-1"},
			Expected: map[string][]string{"978-0-387-96842-1": {"155216880"}},
		},
		{
			ID2:      "springerln2ppn",
			Input:    []string{"978-0-387-96842-1", "978-1-85233-291-4"},
			Expected: map[string][]string{"978-0-387-96842-1": {"155216880"}, "978-1-85233-291-4": {"155222163"}},
		},
		{
			ID2:      "springerln2ppn",
			Input:    []string{},
			Expected: map[string][]string{},
		},
		{
			ID2:      "springerln2ppn",
			Input:    []string{"123456789"},
			Expected: map[string][]string{},
		},
		{
			ID2:      "isbn2ppn",
			Input:    []string{"978-2-7073-1326-3"},
			Expected: map[string][]string{"978-2-7073-1326-3": {"001584944", "099012111", "110704053", "156942208", "190212306"}},
		},
		{
			ID2:      "isbn2ppn",
			Input:    []string{},
			Expected: map[string][]string{},
		},
		{
			ID2:      "isbn2ppn",
			Input:    []string{"123456789"},
			Expected: map[string][]string{},
		},
		{
			ID2:      "ean2ppn",
			Input:    []string{"9782204015899", "9782222024507"},
			Expected: map[string][]string{"9782204015899": {"00034088X", "000343064"}, "9782222024507": {"000272388", "000272396"}},
		},
		{
			ID2:      "ean2ppn",
			Input:    []string{},
			Expected: map[string][]string{},
		},
		{
			ID2:      "ean2ppn",
			Input:    []string{"123456789"},
			Expected: map[string][]string{},
		},
	}

	for _, test := range ID2Tests {
		actual, _ := ID2ppn(test.Input, test.ID2)
		if reflect.DeepEqual(test.Expected, actual) {
			t.Logf("PASS: got %v", test.Expected)
		} else {
			t.Fatalf("FAIL for %s: expected %v, actual result was %v", test.Input, test.Expected, actual)
		}
	}
}

func ExampleID2ppn() {
	myIDService := "frbn2ppn"
	myInput := []string{"00000051X"}

	result, _ := ID2ppn(myInput, myIDService)
	for k, v := range result {
		fmt.Println(k)
		for _, value := range v {
			fmt.Println(value)
		}
	}

	// Output:
	// 00000051X
	// 000000167
	// 02635778X

}
