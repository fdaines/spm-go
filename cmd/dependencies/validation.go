package dependencies

import "errors"

func ValidateOutputFormat(outputFormat string) error {
	supportedOutputFormats := map[string]bool{"csv": true, "console": true, "json": true}
	if !supportedOutputFormats[outputFormat] {
		return errors.New("output format should be one of 'plain', 'console' or 'json'")
	}
	return nil
}