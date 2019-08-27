package common

import (
	"context"
	"testing"

	"github.com/antihax/optional"
	openapi "github.com/openservicebrokerapi/osb-checker/autogenerated/go-client"
	. "github.com/openservicebrokerapi/osb-checker/config"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUpdateInstance(
	t *testing.T,
	instanceID string,
	req *openapi.ServiceInstanceUpdateRequest,
	async bool,
) {
	Convey("UPDATE - request syntax", t, func() {

		Convey("should return 412 PreconditionFailed if missing X-Broker-API-Version header", func() {
			_, resp, err := cli.ServiceInstancesApi.ServiceInstanceUpdate(
				authCtx, "", instanceID, openapi.ServiceInstanceUpdateRequest{},
				&openapi.ServiceInstanceUpdateOpts{AcceptsIncomplete: optional.NewBool(async)})

			So(err, ShouldNotBeNil)
			So(resp.StatusCode, ShouldEqual, 412)
		})

		if CONF.Authentication.AuthType != TypeNoauth {
			Convey("should return 401 Unauthorized if missing Authorization header", func() {
				_, resp, err := cli.ServiceInstancesApi.ServiceInstanceUpdate(
					context.Background(), CONF.APIVersion, instanceID, openapi.ServiceInstanceUpdateRequest{},
					&openapi.ServiceInstanceUpdateOpts{AcceptsIncomplete: optional.NewBool(async)})

				So(err, ShouldNotBeNil)
				So(resp.StatusCode, ShouldEqual, 401)
			})
		}

		if async {
			Convey("should return 422 UnprocessableEntity if missing accepts_incomplete", func() {
				tempBody := openapi.ServiceInstanceUpdateRequest{}
				deepCopy(req, &tempBody)
				_, resp, err := cli.ServiceInstancesApi.ServiceInstanceUpdate(
					authCtx, CONF.APIVersion, instanceID, tempBody,
					&openapi.ServiceInstanceUpdateOpts{AcceptsIncomplete: optional.NewBool(false)})

				So(err, ShouldNotBeNil)
				So(resp.StatusCode, ShouldEqual, 422)
			})
		}

		Convey("should reject if missing service_id", func() {
			tempBody := openapi.ServiceInstanceUpdateRequest{}
			deepCopy(req, &tempBody)
			tempBody.ServiceId = ""
			_, resp, err := cli.ServiceInstancesApi.ServiceInstanceUpdate(
				authCtx, CONF.APIVersion, instanceID, tempBody,
				&openapi.ServiceInstanceUpdateOpts{AcceptsIncomplete: optional.NewBool(async)})

			So(err, ShouldNotBeNil)
			So(resp.StatusCode, ShouldEqual, 400)
		})

		Convey("should reject if service_id is invalid", func() {
			tempBody := openapi.ServiceInstanceUpdateRequest{}
			deepCopy(req, &tempBody)
			tempBody.ServiceId = "xxxx-xxxx-xxxx-xxxx"
			_, resp, err := cli.ServiceInstancesApi.ServiceInstanceUpdate(
				authCtx, CONF.APIVersion, instanceID, tempBody,
				&openapi.ServiceInstanceUpdateOpts{AcceptsIncomplete: optional.NewBool(async)})

			So(err, ShouldNotBeNil)
			So(resp.StatusCode, ShouldEqual, 400)
		})

		Convey("should reject if parameters are not following schema", func() {
			tempBody := openapi.ServiceInstanceUpdateRequest{}
			deepCopy(req, &tempBody)
			tempBody.Parameters = map[string]interface{}{
				"can-not": "be-good",
			}
			if err := testCatalogSchema(&SchemaOpts{
				ServiceID:  tempBody.ServiceId,
				PlanID:     tempBody.PlanId,
				Parameters: tempBody.Parameters,
				SchemaType: TypeServiceInstance,
				Action:     ActionUpdate,
			}); err == nil {
				return
			}
			_, resp, err := cli.ServiceInstancesApi.ServiceInstanceUpdate(
				authCtx, CONF.APIVersion, instanceID, tempBody,
				&openapi.ServiceInstanceUpdateOpts{AcceptsIncomplete: optional.NewBool(async)})

			So(err, ShouldNotBeNil)
			So(resp.StatusCode, ShouldEqual, 400)
		})
	})

	Convey("UPDATE - new", t, func() {
		Convey("should accept a valid update request", func() {
			tempBody := openapi.ServiceInstanceUpdateRequest{}
			deepCopy(req, &tempBody)
			_, resp, err := cli.ServiceInstancesApi.ServiceInstanceUpdate(
				authCtx, CONF.APIVersion, instanceID, tempBody,
				&openapi.ServiceInstanceUpdateOpts{AcceptsIncomplete: optional.NewBool(async)})

			So(err, ShouldBeNil)
			if async {
				So(resp.StatusCode, ShouldEqual, 202)
			} else {
				So(resp.StatusCode, ShouldEqual, 200)
			}
		})
	})

	if async {
		Convey("UPDATE - poll", t, func(c C) {
			testPollInstanceLastOperation(instanceID)

			So(pollInstanceLastOperationStatus(instanceID, "update"), ShouldBeNil)
		})
	}
}