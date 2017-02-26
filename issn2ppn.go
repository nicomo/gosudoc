package gosudoc

import (
	"encoding/xml"
	"errors"
)

// Issn2ppnData: used to unmarshal the XML returned by Sudoc
// for the issn2ppn web service
type Issn2ppnData struct {
	Err   string `xml:"error"`
	Pairs []struct {
		Issn          string   `xml:"issn,omitempty"`
		PPNs          []string `xml:"result>ppn"`
		PPNsNoHolding []string `xml:"resultNoHolding>ppn"`
	} `xml:"query"`
}

// Issn2ppn takes 1 or more ISSNs
// and will return the corresponding sudoc record IDs
// with an additional information : does any library have holdings for this title? Y/N
func Issn2ppn(input []string) (map[string][]string, error) {

	result := make(map[string][]string)

	// construct the url
	baseURL := "http://www.sudoc.fr/services/issn2ppn"
	getURL, err := getQryString(baseURL, input)
	if err != nil {
		return result, err
	}

	// call Sudoc & put the response into a []byte
	b, err := callSudoc(getURL)

	// unmarshall
	var parsedResp Issn2ppnData
	if xml.Unmarshal(b, &parsedResp); err != nil {
		return result, err
	}

	if parsedResp.Err != "" {
		ErrNoResult := errors.New("No result")
		return result, ErrNoResult
	}

	for _, v := range parsedResp.Pairs {
		PResults := []string{}
		if len(v.PPNs) > 0 {
			for _, value := range v.PPNs {
				PResults = append(PResults, value)
			}
		}
		if len(v.PPNsNoHolding) > 0 {
			for _, value := range v.PPNsNoHolding {
				PResults = append(PResults, value)
			}
		}
		result[v.Issn] = PResults
	}

	return result, nil
}
