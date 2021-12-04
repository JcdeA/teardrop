package api

import (
	"io/ioutil"

	"github.com/fosshostorg/teardrop/models"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

func ReadConfig(path string) (*models.Config, error) {
	var c models.Config

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed reading server config file: %s", path)
	}

	if err := yaml.UnmarshalStrict(bytes, &c); err != nil {
		return nil, errors.Wrap(err, "failed parsing configuration file")
	}

	return &c, nil
}
