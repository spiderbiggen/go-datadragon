package datadragon

import "errors"

var (
	ErrNoVersion = errors.New("no version")
	ErrNoLocale  = errors.New("no language")
)

type configValidator func(RequestConfig) error

type RequestConfig struct {
	Version string
	Locale  string
}

func RequireVersion() func(RequestConfig) error {
	return func(r RequestConfig) error {
		if r.Version == "" {
			return ErrNoVersion
		}
		return nil
	}
}

func RequireLocale() func(RequestConfig) error {
	return func(r RequestConfig) error {
		if r.Locale == "" {
			return ErrNoLocale
		}
		return nil
	}
}

func (d *DataDragon) mergeConfig(params []RequestConfig, validators ...configValidator) (*RequestConfig, error) {
	config := RequestConfig{
		Version: d.Version,
		Locale:  d.Locale,
	}
	if len(params) > 0 {
		param := params[0]
		if param.Version != "" {
			config.Version = param.Version
		}
		if param.Locale != "" {
			config.Locale = param.Locale
		}
	}

	for _, validator := range validators {
		if err := validator(config); err != nil {
			return nil, err
		}
	}

	return &config, nil
}
