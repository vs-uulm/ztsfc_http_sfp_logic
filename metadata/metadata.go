package metadata

import (
//    "fmt"
    "net/http"
)

type Cp_metadata struct {
    //Auth_decision bool
    //User string
    //Pw_authenticated bool
    //Cert_authenticated bool
    //Resource string
    //Action string
    //Device string
    //RequestToday int
    //FailedToday int
    //Location string
    SFC string
    SFP string
}

func (cpm *Cp_metadata) ExtractMetadata(req *http.Request) {
    //cpm.User = req.Header.Get("user")
    //cpm.Pw_authenticated, _ = strconv.ParseBool(req.Header.Get("pwAuthenticated"))
    //cpm.Cert_authenticated, _ = strconv.ParseBool(req.Header.Get("certAuthenticated"))
    //cpm.Resource = req.Header.Get("resource")
    //cpm.Action = req.Header.Get("action")
    //cpm.Device = req.Header.Get("device")
    //cpm.RequestToday, _ = strconv.Atoi(req.Header.Get("requestToday"))
    //cpm.FailedToday, _ = strconv.Atoi(req.Header.Get("failedToday"))
    //cpm.Location = req.Header.Get("location")
    cpm.SFC = req.Header.Get("sfc")
//    fmt.Printf("%s\n", cpm.SFC)
}
