// Package gosudoc is a set of client functions for the web services
// available on top of the Library Catalog of French University libraries
// named Sudoc.
// see http://documentation.abes.fr/sudoc/manuels/administration/aidewebservices/
package gosudoc

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// getQryString builds the query string from the base URL
// which is different for each web service
// and the list of IDs to pass : the construction is the same for every web service here
func getQryString(baseurl string, ids []string) (string, error) {
	if len(ids) == 0 {
		return "", errors.New("empty IDs string")
	}

	// build the GET string
	var idsString string
	for i := 0; i < len(ids); i++ {
		idsString += ids[i]
		if i != len(ids)-1 {
			idsString += ","
		}
	}
	URL := baseurl + "/" + idsString
	return URL, nil
}

// callSudoc performs the http GET
// retrieves the response and puts it in a slice of bytes
func callSudoc(getURL string) ([]byte, error) {
	// get the result from the url
	resp, err := http.Get(getURL)
	if err != nil {
		return []byte{}, err
	}

	// put the response into a []byte
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return []byte{}, err
	}

	return b, nil
}
