package mysql

import (
	"github.com/zenthangplus/call-billing-example/src/core/enum"
	"gorm.io/gorm"
)

type base struct {
	db *gorm.DB
}

func (b base) handleError(err error) error {
	if err == gorm.ErrRecordNotFound {
		return enum.ErrResourceNotFound
	}
	return err
}
