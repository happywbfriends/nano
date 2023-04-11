package config

import (
	"github.com/caarlos0/env/v6"
)

func ParseValues(values map[string]string, target interface{}) error {
	if err := env.Parse(target, env.Options{Environment: values}); err != nil {
		return err
	}
	return validate(target)
}

func ParseEnv(target interface{}) error {
	if err := env.Parse(target); err != nil {
		return err
	}
	return validate(target)
}
