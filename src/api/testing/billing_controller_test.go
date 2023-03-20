package testing

import (
	"github.com/stretchr/testify/require"
	"github.com/zenthangplus/call-billing-example/src/adapter/repository/mysql/model"
	"github.com/zenthangplus/call-billing-example/src/core/enum"
	golibdataTestUtil "gitlab.com/golibs-starter/golib-data/testutil"
	"gitlab.com/golibs-starter/golib-test"
	"net/http"
	"testing"
)

func TestBillingController_WhenGetBilling_GivenMissingUsername_ShouldBadRequest(t *testing.T) {
	golibtest.NewRestAssured(t).
		When().
		Get("/v1/mobile//billing").
		Then().
		Status(http.StatusBadRequest).
		Body("meta.code", enum.ErrMissingUsername.Code()).
		Body("meta.message", enum.ErrMissingUsername.Message())
}

func TestBillingController_WhenGetBilling_GivenValidRequestAndNoBillingFound_ShouldSuccess(t *testing.T) {
	golibdataTestUtil.TruncateTables("billings")
	golibtest.NewRestAssured(t).
		When().
		Get("/v1/mobile/user2/billing").
		Then().
		Status(http.StatusNotFound).
		Body("meta.code", enum.ErrBillingNotFound.Code()).
		Body("meta.message", enum.ErrBillingNotFound.Message())
}

func TestBillingController_WhenGetBilling_GivenValidRequestAndBillingFound_ShouldSuccess(t *testing.T) {
	golibdataTestUtil.TruncateTables("billings")
	bill1 := model.Billing{Username: "user1", CallDuration: 3000, CallCount: 1}
	bill2 := model.Billing{Username: "user2", CallDuration: 70000, CallCount: 3}
	db.Create(&bill1)
	db.Create(&bill2)
	golibtest.NewRestAssured(t).
		When().
		Get("/v1/mobile/user2/billing").
		Then().
		Status(http.StatusOK).
		Body("data.id", bill2.Id).
		Body("data.call_count", bill2.CallCount).
		Body("data.block_count", 3).
		BodyCb("data.price", func(value interface{}) {
			require.InDelta(t, 0.6, value, 10e-9)
		})
}
