package config

import (
	"errors"
	"strings"
)

var (
	ErrInvalidProfileString = errors.New("invalid profile string")
)

type Profile int8

func (p *Profile) String() string {
	switch *p {
	case ProfileLocal:
		return "local"
	case ProfileDev:
		return "dev"
	case ProfileProd:
		return "prod"
	default:
		return "unknown"
	}
}

func (p *Profile) FromString(s string) error {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "local":
		*p = ProfileLocal
	case "dev", "development":
		*p = ProfileDev
	case "prod", "production":
		*p = ProfileProd
	default:
		return ErrInvalidProfileString
	}

	return nil
}

const (
	ProfileLocal Profile = iota
	ProfileDev
	ProfileProd
)
