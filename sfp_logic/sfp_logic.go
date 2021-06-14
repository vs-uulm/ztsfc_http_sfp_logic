package sfp_logic

import (
    md "local.com/leobrada/ztsfc_http_sfpLogic/metadata"
    "strings"
    "fmt"
)

//func TransformSFCintoSFP(cpm *md.Cp_metadata) {
//    cpm.SFP = ""
//}


func TransformSFCintoSFP(cpm *md.Cp_metadata) {
    if cpm == nil {
        fmt.Printf("cpm is nil")
        return
    }

    if len(cpm.SFC) == 0 {
        cpm.SFP = ""
        return
    }

    sfc := strings.Split(cpm.SFC, ",")
    sfp := ""
    for _, sf := range sfc {
        switch sf {
        case "dpi":
            sfp += ",https://10.5.0.54:8888"
        case "logger":
            sfp += ",https://10.5.0.50:8889"
        default:
            sfp += ""
        }
    }
    sfp = strings.TrimLeft(sfp, ",")
    cpm.SFP = sfp
}
