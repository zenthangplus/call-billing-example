package properties

import (
	"gitlab.com/golibs-starter/golib/config"
	"time"
)

type BillingProperties struct {
	BlockTime     time.Duration `validate:"required"`
	PricePerBlock float64
}

func NewBillingProperties(loader config.Loader) (*BillingProperties, error) {
	props := BillingProperties{}
	err := loader.Bind(&props)
	return &props, err
}

func (o BillingProperties) Prefix() string {
	return "app.billing"
}
