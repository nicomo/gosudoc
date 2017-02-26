package gosudoc

import (
	"encoding/xml"
	"errors"
)

// MergedData is used to parse xml response
type MergedData struct {
	Err   string `xml:"error"`
	Pairs []struct {
		Queried string `xml:"ppn"`
		Merged  string `xml:"result>ppn"`
	} `xml:"query"`
}

// GetMerged takes any number of Sudoc ID numbers (PPN)
// and matches them with the ID of the merged marc record
// it echoes the IDs you input, then outputs the returned PPNs
// Sudoc responds with an error if there's nothing to return
func GetMerged(input []string) (map[string]string, error) {
	result := make(map[string]string)

	// construct the url
	getURL, err := getQryString("http://www.sudoc.fr/services/merged", input)
	if err != nil {
		return result, err
	}

	// call Sudoc & put the response into a []byte
	b, err := callSudoc(getURL)

	// unmarshall the xml response into a struct
	var parsedResp MergedData
	if xml.Unmarshal(b, &parsedResp); err != nil {
		return result, err
	}

	if parsedResp.Err != "" {
		ErrNoResult := errors.New("No result")
		return result, ErrNoResult
	}

	for _, v := range parsedResp.Pairs {
		result[v.Queried] = v.Merged
	}

	return result, nil
}
