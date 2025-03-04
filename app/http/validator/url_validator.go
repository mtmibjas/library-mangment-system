package validator

import (
	"errors"
	"library-mngmt/pkg/logger/zap"
	"net/url"
)

func ValidateURL(inputURL string) error {
	parsedURL, err := url.ParseRequestURI(inputURL)
	if err != nil {
		zap.Error("validator:", err)
		return errors.New("invalid URL format")
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		err = errors.New("URL scheme must be http or https")
		zap.Error("validator:", err)
		return err
	}

	return nil
}
