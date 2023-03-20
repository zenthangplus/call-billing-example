package properties

import (
	"gitlab.com/golibs-starter/golib/config"
)

type CallProperties struct {
	MaxUsernameLength int `validate:"required" default:"32"`
}

func NewCallProperties(loader config.Loader) (*CallProperties, error) {
	props := CallProperties{}
	err := loader.Bind(&props)
	return &props, err
}

func (o CallProperties) Prefix() string {
	return "app.call"
}
