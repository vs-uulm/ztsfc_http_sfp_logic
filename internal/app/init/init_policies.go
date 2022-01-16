package init

import (
    "errors"
    "fmt"
    "net/url"

    "github.com/vs-uulm/ztsfc_http_sfp_logic/internal/app/policies"
)

func InitPolicyParams() error {
    if policies.Policies.SfPool == nil {
        return errors.New("init: InitResourcesParams(): no sf_pool defined")
    }

    for sfName, sf := range policies.Policies.SfPool {
        if sf == nil {
            return errors.New("init: InitPolicyParams(): sf '" + sfName + "' is empty")
        }

        if len(sf.InstanceURLs) == 0 {
            return errors.New("init: InitPolicyParams(): for sf '" + sfName + "' no instances are defined")
        }

        for _, instanceUrl := range sf.InstanceURLs {
            if _, err := url.Parse(instanceUrl); err != nil {
                return fmt.Errorf("init: InitPolicyParams(): for %s given url %s is not a vald url: %v", sfName, instanceUrl, err)
            }
        }
    }

    return nil
}
