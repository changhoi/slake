package config

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	errKeyNotExist = errors.New("not exist key")
)

func StringFromEnv(key string) (string, error) {
	v, ok := os.LookupEnv(key)
	if !ok {
		return "", errors.Join(errKeyNotExist, errors.New(key))
	}

	return v, nil
}

func BytesFromEnv(key string) ([]byte, error) {
	raw, err := StringFromEnv(key)
	if err != nil {
		return nil, err
	}

	return []byte(raw), nil
}

func StringSliceFromEnv(key string) ([]string, error) {
	raw, err := StringFromEnv(key)
	if err != nil {
		return nil, err
	}

	ret := strings.Split(raw, ",")

	for i, v := range ret {
		ret[i] = strings.TrimSpace(v)
	}

	return ret, nil
}

func IntFromEnv(key string) (int, error) {
	raw, err := StringFromEnv(key)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(raw)
}

func BoolFromEnv(key string) (bool, error) {
	raw, err := StringFromEnv(key)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(raw)
}

func ProfileFromEnv(key string) (Profile, error) {
	raw, err := StringFromEnv(key)
	if err != nil {
		return 0, err
	}

	var e Profile
	if err := e.FromString(raw); err != nil {
		return 0, err
	}

	return e, nil
}

func DurationFromEnv(key string) (time.Duration, error) {
	raw, err := StringFromEnv(key)
	if err != nil {
		return 0, err
	}

	return time.ParseDuration(raw)
}

func Int64FromEnv(key string) (int64, error) {
	raw, err := StringFromEnv(key)
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(raw, 10, 64)
}

func Int32FromEnv(key string) (int32, error) {
	raw, err := StringFromEnv(key)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseInt(raw, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(v), nil
}

type Getter[T any] func(key string) (T, error)

func Or[T any](fn Getter[T], key string, fallback T) T {
	v, err := fn(key)
	if err != nil {
		return fallback
	}

	return v
}

func Must[T any](fn Getter[T], key string) T {
	v, err := fn(key)
	if err != nil {
		panic(err)
	}

	return v
}
