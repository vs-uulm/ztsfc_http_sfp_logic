package metadata

import (
	"net/http"
)

type Cp_metadata struct {
	SFC []string
	SFP []string
}

func (cpm *Cp_metadata) ExtractMetadata(req *http.Request) {

	// @author:marie
	// Retreive parameters from query instead from custom headers
	// (SFC is transmitted as a list of several SFs)
	cpm.SFC = req.URL.Query()["sf"]
	// cpm.SFC = req.Header.Get("sfc")
}
