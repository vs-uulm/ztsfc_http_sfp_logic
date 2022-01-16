package policies

import (
// for later...
)

var (
    Policies PoliciesT
)

type PoliciesT struct {
    SfPool map[string]*SfT `yaml:"sf_pool"`
}

type SfT struct {
    InstanceURLs []string `yaml:"instance_urls"`
}
