package sfp_logic

import (
	"fmt"

	md "github.com/vs-uulm/ztsfc_http_sfp_logic/internal/app/metadata"
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
		cpm.SFP = []md.SF{}
		return
	}

	// @author:marie
	// reintroduced ip translation, but now provide it together wih SF name.
	for _, sfName := range cpm.SFC {
		switch sfName {
		case "dpi":
			sf := md.SF{Name: "dpi", Address: "https://10.5.0.54:8888"}
			cpm.SFP = append(cpm.SFP, sf)
		case "logger":
			sf := md.SF{Name: "logger", Address: "https://10.5.0.50:8889"}
			cpm.SFP = append(cpm.SFP, sf)
		default:
			sf := md.SF{Name: "", Address: ""}
			cpm.SFP = append(cpm.SFP, sf)
		}
	}
}
