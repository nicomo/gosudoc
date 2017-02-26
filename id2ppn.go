package gosudoc

import (
	"encoding/xml"
	"errors"
)

var ID2 = []string{
	"frbn2ppn",
	"ocn2ppn",
	"dnb2ppn",
	"ucatb2ppn",
	"frcairninfo2ppn",
	"springerln2ppn",
	"isbn2ppn",
	"ean2ppn",
}

// ID2ppnData: used to unmarshal the XML returned by Sudoc
// will populate the appropriate field given a service,
// e.g. Ocn for the ocn2ppn web service,
// and ignore the other fields
type ID2ppnData struct {
	Err   string `xml:"error"`
	Pairs []struct {
		Frbn        string   `xml:"frbn,omitempty"`
		Ocn         string   `xml:"ocn,omitempty"`
		Dnb         string   `xml:"dnb,omitempty"`
		Ucatb       string   `xml:"ucatb,omitempty"`
		FrCairnInfo string   `xml:"frcairninfo,omitempty"`
		SpringerLN  string   `xml:"springerln,omitempty"`
		Isbn        string   `xml:"isbn,omitempty"`
		Ean         string   `xml:"ean,omitempty"`
		PPNs        []string `xml:"result>ppn"`
	} `xml:"query"`
}

// ID2ppn is used for several related web services:
// given one or more IDs such as isbns, OCLC ocn, etc.
// and given the appropriate service to try (isbn, ocn, etc)
// will return the corresponding sudoc record IDs
func ID2ppn(input []string, id2 string) (map[string][]string, error) {

	result := make(map[string][]string)

	// check id2 is a known web service
	if !containsID2(id2) {
		ErrID2 := errors.New("unknown web service")
		return result, ErrID2
	}

	// construct the url
	baseURL := "http://www.sudoc.fr/services/" + id2
	getURL, err := getQryString(baseURL, input)
	if err != nil {
		return result, err
	}

	// call Sudoc & put the response into a []byte
	b, err := callSudoc(getURL)

	// unmarshall
	var parsedResp ID2ppnData
	if xml.Unmarshal(b, &parsedResp); err != nil {
		return result, err
	}

	if parsedResp.Err != "" {
		ErrNoResult := errors.New("No result")
		return result, ErrNoResult
	}

	for _, v := range parsedResp.Pairs {
		if v.Frbn != "" {
			result[v.Frbn] = v.PPNs
			continue
		}
		if v.Ocn != "" {
			result[v.Ocn] = v.PPNs
			continue
		}
		if v.Dnb != "" {
			result[v.Dnb] = v.PPNs
			continue
		}
		if v.Ucatb != "" {
			result[v.Ucatb] = v.PPNs
			continue
		}
		if v.FrCairnInfo != "" {
			result[v.FrCairnInfo] = v.PPNs
			continue
		}
		if v.SpringerLN != "" {
			result[v.SpringerLN] = v.PPNs
			continue
		}
		if v.Isbn != "" {
			result[v.Isbn] = v.PPNs
			continue
		}
		if v.Ean != "" {
			result[v.Ean] = v.PPNs
			continue
		}

	}

	return result, nil
}

func containsID2(id2 string) bool {
	for _, v := range ID2 {
		if v == id2 {
			return true
		}
	}
	return false
}
