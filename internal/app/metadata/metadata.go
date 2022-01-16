package metadata

import (
	"net/http"
)

type Cp_metadata struct {
	SFC []string `json:"sfc"`
	SFP []SF     `json:"sfp"`
}
type SF struct {
	Name    string `json:"name"`
	URL string `json:"url"`
}

func (cpm *Cp_metadata) ExtractMetadata(req *http.Request) error {

	// Retreive parameters from query instead from custom headers
	// (SFC is transmitted as a list of several SFs)
    // TODO: Error Handling???
	cpm.SFC = req.URL.Query()["sfc"]
    return nil
	// cpm.SFC = req.Header.Get("sfc")
}
