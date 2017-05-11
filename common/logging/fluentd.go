package logging

import "fmt"

const (
	addr = "fluentd-address"
	tag = "tag"

)

// ValidateOpts validates the provided keys of l
func ValidateOpts(logOpts map[string]string) error {
	for k, _ := range logOpts {
		switch k {
		case addr:
		case tag:
		default :
			return fmt.Errorf("provided option %s is not supported as a Fluentd logging option", k)
		}
	}
	return nil
}
