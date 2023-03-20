package testing

import (
	"github.com/stretchr/testify/require"
	"github.com/zenthangplus/call-billing-example/src/adapter/repository/mysql/model"
	golibdataTestUtil "gitlab.com/golibs-starter/golib-data/testutil"
	golibtest "gitlab.com/golibs-starter/golib-test"
	"net/http"
	"testing"
	"time"
)

func TestBillingAggregationListener_GivenValidRequest_WhenBillNotExists_ShouldCreateNewBill(t *testing.T) {
	golibdataTestUtil.TruncateTables("billings")
	golibtest.NewRestAssured(t).
		When().
		Put("/v1/mobile/user1/call").
		Body(`{"call_duration": 2000}`).
		Then().
		Status(http.StatusOK)
	time.Sleep(50 * time.Millisecond)
	actualBill := model.Billing{Username: "user1"}
	tx := db.First(&actualBill)
	require.NoError(t, tx.Error)
	require.Greater(t, actualBill.Id, 0)
	require.EqualValues(t, 2000, actualBill.CallDuration)
	require.EqualValues(t, 1, actualBill.CallCount)
}

func TestBillingAggregationListener_GivenValidRequest_WhenBillExists_ShouldUpdateBill(t *testing.T) {
	golibdataTestUtil.TruncateTables("billings")
	bill := model.Billing{Username: "user1", CallDuration: 5000, CallCount: 1}
	db.Create(&bill)
	golibtest.NewRestAssured(t).
		When().
		Put("/v1/mobile/user1/call").
		Body(`{"call_duration": 2000}`).
		Then().
		Status(http.StatusOK)
	time.Sleep(50 * time.Millisecond)
	actualBill := model.Billing{Username: "user1"}
	tx := db.First(&actualBill)
	require.NoError(t, tx.Error)
	require.Greater(t, actualBill.Id, 0)
	require.EqualValues(t, 7000, actualBill.CallDuration)
	require.EqualValues(t, 2, actualBill.CallCount)
}
