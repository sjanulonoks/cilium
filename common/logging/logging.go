package logging

type Logger interface {
	// ValidateOpts validates the keys in the provided map to ensure that they are supported by the specified logger.
	ValidateOpts(map[string]string) error
}

