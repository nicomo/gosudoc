package gosudoc

import (
	"encoding/xml"
	"errors"
)

// MultiwhereData is used to parse xml response for the multiwhere web service
type MultiwhereData struct {
	Err     string `xml:"error"`
	Results []struct {
		Queried   string    `xml:"ppn"`
		Libraries []Library `xml:"result>library"`
	} `xml:"query"`
}

// Library is used when we unmarshal the library info embedded in the returned XML
// i.e. library ID, name, lat and longitude.
type Library struct {
	RCR       string `xml:"rcr"`
	Shortname string `xml:"shortname"`
	Latitude  string `xml:"latitude"`
	Longitude string `xml:"longitude"`
}

// GetMultiwhere, given a record ID in the Sudoc catalog
// will return the list of libraries that hold this item
func GetMultiwhere(input []string) (map[string][]Library, error) {
	result := make(map[string][]Library)

	// construct the url
	baseURL := "http://www.sudoc.fr/services/multiwhere"
	getURL, err := getQryString(baseURL, input)
	if err != nil {
		return result, err
	}

	// call Sudoc & put the response into a []byte
	b, err := callSudoc(getURL)

	// unmarshall
	var parsedResp MultiwhereData
	if xml.Unmarshal(b, &parsedResp); err != nil {
		return result, err
	}

	if parsedResp.Err != "" {
		ErrNoResult := errors.New("No result")
		return result, ErrNoResult
	}

	for _, v := range parsedResp.Results {
		result[v.Queried] = v.Libraries
	}

	return result, nil
}
