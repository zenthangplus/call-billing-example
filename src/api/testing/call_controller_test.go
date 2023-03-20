package testing

import (
	"github.com/stretchr/testify/require"
	"github.com/zenthangplus/call-billing-example/src/adapter/repository/mysql/model"
	"github.com/zenthangplus/call-billing-example/src/core/enum"
	"gitlab.com/golibs-starter/golib-test"
	baseEx "gitlab.com/golibs-starter/golib/exception"
	"net/http"
	"testing"
	"time"
)

func TestCallController_WhenEndCallRequest_GivenInvalidBody_ShouldBadRequest(t *testing.T) {
	golibtest.NewRestAssured(t).
		When().
		Put("/v1/mobile/xxx/call").
		Body(`not-json-body`).
		Then().
		Status(http.StatusBadRequest).
		Body("meta.code", baseEx.BadRequest.Code()).
		Body("meta.message", baseEx.BadRequest.Message())
}

func TestCallController_WhenEndCallRequest_GivenMissingUsername_ShouldBadRequest(t *testing.T) {
	golibtest.NewRestAssured(t).
		When().
		Put("/v1/mobile//call").
		Body(`{"call_duration": 1000}`).
		Then().
		Status(http.StatusBadRequest).
		Body("meta.code", enum.ErrMissingUsername.Code()).
		Body("meta.message", enum.ErrMissingUsername.Message())
}

func TestCallController_WhenEndCallRequest_GivenInvalidUsername_ShouldBadRequest(t *testing.T) {
	golibtest.NewRestAssured(t).
		When().
		Put("/v1/mobile/22222222222222222222222222222222222222222222222222222222/call").
		Body(`{"call_duration": 1000}`).
		Then().
		Status(http.StatusBadRequest).
		Body("meta.code", enum.ErrInvalidUsername.Code()).
		Body("meta.message", enum.ErrInvalidUsername.Message())
}

func TestCallController_WhenEndCallRequest_GivenInvalidDuration_ShouldBadRequest(t *testing.T) {
	golibtest.NewRestAssured(t).
		When().
		Put("/v1/mobile/xxx/call").
		Body(`{"call_duration": 0}`).
		Then().
		Status(http.StatusBadRequest).
		Body("meta.code", enum.ErrInvalidDuration.Code()).
		Body("meta.message", enum.ErrInvalidDuration.Message())
}

func TestCallController_WhenEndCallRequest_GivenValidRequest_ShouldSuccess(t *testing.T) {
	golibtest.NewRestAssured(t).
		When().
		Put("/v1/mobile/user1/call").
		Body(`{"call_duration": 2000}`).
		Then().
		Status(http.StatusOK)
	callModel := model.Call{Username: "user1", Duration: 2000}
	tx := db.Last(&callModel)
	require.NoError(t, tx.Error)
	require.Greater(t, callModel.Id, 0)
	require.InDelta(t, time.Now().Unix(), callModel.CreatedAt.Unix(), 1)
	require.InDelta(t, time.Now().Unix(), callModel.UpdatedAt.Unix(), 1)
}
