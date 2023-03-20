package bootstrap

import (
	"github.com/zenthangplus/call-billing-example/src/api/properties"
	"github.com/zenthangplus/call-billing-example/src/core/config"
)

func NewCallConfig(props *properties.CallProperties) *config.CallConfig {
	return &config.CallConfig{
		MaxUsernameLength: props.MaxUsernameLength,
	}
}

func NewBillingConfig(props *properties.BillingProperties) *config.BillingConfig {
	return &config.BillingConfig{
		BlockTime:     props.BlockTime,
		PricePerBlock: props.PricePerBlock,
	}
}
