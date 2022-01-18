package sfp_logic

import (
    "errors"

	"github.com/vs-uulm/ztsfc_http_sfp_logic/internal/app/policies"
	md "github.com/vs-uulm/ztsfc_http_sfp_logic/internal/app/metadata"
    logger "github.com/vs-uulm/ztsfc_http_logger"
)

func TransformSFCintoSFP(sysLogger *logger.Logger, cpm *md.Cp_metadata) error {
	if cpm == nil {
		return errors.New("sfp_logic: TransformSFCintoSFP(): passed metadata pointer is nil")
	}

	if len(cpm.SFC) == 0 {
        sysLogger.Infof("sfp_logic: TransformSFCintoSFP(): empty SFC passed. please take a look. could be an error")
		cpm.SFP = []md.SF{}
		return nil
	}

	// reintroduced ip translation, but now provide it together wih SF name.
	for _, sfName := range cpm.SFC {
		switch sfName {
		case "dpi":
			sf := md.SF{Name: "dpi", URL: policies.Policies.SfPool["dpi"].InstanceURLs[0]}
			cpm.SFP = append(cpm.SFP, sf)
		case "logger":
            sysLogger.Debugf("logger has been choosen: %s", policies.Policies.SfPool["logger"].InstanceURLs[0])
			sf := md.SF{Name: "logger", URL: policies.Policies.SfPool["logger"].InstanceURLs[0]}
			cpm.SFP = append(cpm.SFP, sf)
		default:
            sysLogger.Infof("sfp_logic: TransformSFCintoSFP(): there is an empty entry in the SFC slice. this could indicate an error")
			sf := md.SF{Name: "", URL: ""}
			cpm.SFP = append(cpm.SFP, sf)
		}
	}
    return nil
}
