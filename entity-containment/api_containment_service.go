/*
 * Entity Containment API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package entitycontainment

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/rtk-logan-fillo/sentinelone-sdk-go/sentinelone"
)

// ContainmentApiService is a service that implements the logic for the ContainmentApiServicer
// This service should implement the business logic for every endpoint for the ContainmentApi API.
// Include any external packages or services that will be required by this service.
type ContainmentApiService struct {
	cli *sentinelone.APIClient
}

// NewContainmentApiService creates a default api service
func NewContainmentApiService() ContainmentApiServicer {
	return &ContainmentApiService{cli: sentinelone.NewAPIClient(getConfig())}
}

// BulkLookupContainmentStatus -
func (s *ContainmentApiService) BulkLookupContainmentStatus(ctx context.Context, organizationId string, vendor string, deviceIdList DeviceIdList) (ImplResponse, error) {
	// TODO - update BulkLookupContainmentStatus with the required logic for this service method.
	// Add api_containment_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	var res = ContainmentStatusResponse{}
	for _, deviceId := range deviceIdList.DeviceIds {
		agent := getAgent(s.cli, deviceId)
		status := agent.GetNetworkStatus()
		res.Count += 1
		res.Items = append(res.Items, ContainmentStatus{DeviceId: deviceId, Status: getStatus(status)})
	}
	return Response(200, res), nil
	// curl -X POST -H 'Content-Type: application/json' -d '{"device_ids":["1500187799060682232"]}' http://localhost:1234/containment/v1/organizations/a/devices/s1/containment-status

	//TODO: Uncomment the next line to return response Response(200, ContainmentStatusResponse{}) or use other options such as http.Ok ...
	//return Response(200, ContainmentStatusResponse{}), nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(403, {}) or use other options such as http.Ok ...
	//return Response(403, nil),nil

	//TODO: Uncomment the next line to return response Response(404, ErrorUnknown{}) or use other options such as http.Ok ...
	//return Response(404, ErrorUnknown{}), nil

	//TODO: Uncomment the next line to return response Response(0, GenericErrorResponse{}) or use other options such as http.Ok ...
	//return Response(0, GenericErrorResponse{}), nil

	// return Response(http.StatusNotImplemented, nil), errors.New("BulkLookupContainmentStatus method not implemented")
}

// HealthCheck - Health check
func (s *ContainmentApiService) HealthCheck(ctx context.Context) (ImplResponse, error) {
	// TODO - update HealthCheck with the required logic for this service method.
	// Add api_containment_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, HealthCheckResponse{}) or use other options such as http.Ok ...
	//return Response(200, HealthCheckResponse{}), nil

	//TODO: Uncomment the next line to return response Response(0, GenericErrorResponse{}) or use other options such as http.Ok ...
	//return Response(0, GenericErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("HealthCheck method not implemented")
}

// LookupIsContainmentActive -
func (s *ContainmentApiService) LookupIsContainmentActive(ctx context.Context, organizationId string, vendor string) (ImplResponse, error) {
	// TODO - update LookupIsContainmentActive with the required logic for this service method.
	// Add api_containment_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return Response(200, ContainmentActiveResponse{Active: true}), nil
	//TODO: Uncomment the next line to return response Response(200, ContainmentActiveResponse{}) or use other options such as http.Ok ...
	//return Response(200, ContainmentActiveResponse{}), nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(403, {}) or use other options such as http.Ok ...
	//return Response(403, nil),nil

	//TODO: Uncomment the next line to return response Response(404, ErrorUnknown{}) or use other options such as http.Ok ...
	//return Response(404, ErrorUnknown{}), nil

	//TODO: Uncomment the next line to return response Response(0, GenericErrorResponse{}) or use other options such as http.Ok ...
	//return Response(0, GenericErrorResponse{}), nil

	// return Response(http.StatusNotImplemented, nil), errors.New("LookupIsContainmentActive method not implemented")
}

// RequestContainment -
func (s *ContainmentApiService) RequestContainment(ctx context.Context, organizationId string, vendor string, deviceId string) (ImplResponse, error) {
	// TODO - update RequestContainment with the required logic for this service method.
	// Add api_containment_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	changeAgentNetworkStatus(s.cli, deviceId, "disconnected")
	//TODO: Uncomment the next line to return response Response(202, {}) or use other options such as http.Ok ...
	return Response(202, nil), nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(403, {}) or use other options such as http.Ok ...
	//return Response(403, nil),nil

	//TODO: Uncomment the next line to return response Response(404, ErrorUnknown{}) or use other options such as http.Ok ...
	//return Response(404, ErrorUnknown{}), nil

	//TODO: Uncomment the next line to return response Response(0, GenericErrorResponse{}) or use other options such as http.Ok ...
	//return Response(0, GenericErrorResponse{}), nil

	// return Response(http.StatusNotImplemented, nil), errors.New("RequestContainment method not implemented")
}

// RequestLiftContainment -
func (s *ContainmentApiService) RequestLiftContainment(ctx context.Context, organizationId string, vendor string, deviceId string) (ImplResponse, error) {
	// TODO - update RequestLiftContainment with the required logic for this service method.
	// Add api_containment_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	changeAgentNetworkStatus(s.cli, deviceId, "connected")
	//TODO: Uncomment the next line to return response Response(202, {}) or use other options such as http.Ok ...
	return Response(202, nil), nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(403, {}) or use other options such as http.Ok ...
	//return Response(403, nil),nil

	//TODO: Uncomment the next line to return response Response(404, ErrorUnknown{}) or use other options such as http.Ok ...
	//return Response(404, ErrorUnknown{}), nil

	//TODO: Uncomment the next line to return response Response(0, GenericErrorResponse{}) or use other options such as http.Ok ...
	//return Response(0, GenericErrorResponse{}), nil

	// return Response(http.StatusNotImplemented, nil), errors.New("RequestLiftContainment method not implemented")
}

func getStatus(status string) string {
	switch status {
	case "connected":
		return "normal"
	case "connecting":
		return "lift_containment_pending"
	case "disconnecting":
		return "containment_pending"
	case "disconnected":
		return "contained"
	default:
		return ""
	}
}

// gets client config, uses env var for api token
func getConfig() *sentinelone.Configuration {
	token := os.Getenv("SENTINELONE_API_TOKEN")
	if token == "" {
		panic("you must run \"export SENTINELONE_API_TOKEN=<A valid API token>\"")
	}
	cfg := sentinelone.NewConfiguration()
	cfg.Host = "usea1-partners.sentinelone.net"
	cfg.UserAgent = "Arctic-Wolf/1.0"
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", fmt.Sprintf("APIToken %s", token))
	return cfg
}

// get agent by Id
func getAgent(client *sentinelone.APIClient, name string) sentinelone.AgentsSchemasAgentViewSchemaMany200DataInner {
	res, _, err := client.AgentsApi.GetAgents(context.Background()).Ids([]string{name}).Execute()
	if err != nil {
		panic(err)
	}
	if len(res.Data) <= 0 {
		panic(fmt.Sprintf("could not find agent %s", name))
	}
	agent := res.Data[0]
	return agent
}

// changes agent status to desired status
func changeAgentNetworkStatus(client *sentinelone.APIClient, name string, desiredStatus string) {
	agent := getAgent(client, name)
	id := agent.GetId()
	fmt.Printf("current network status '%s', desired network status '%s' \n", agent.GetNetworkStatus(), desiredStatus)
	var err error
	switch desiredStatus {
	case "connected":
		filter := sentinelone.NewAgentsSchemasAgentsActionSchemaFilter()
		filter.SetIds([]string{id})
		params := sentinelone.NewAgentsSchemasAgentsActionSchema(*filter)
		_, _, err = client.AgentActionsApi.ConnectAgents(context.Background()).Body(*params).Execute()
	case "disconnected":
		filter := sentinelone.NewAgentsSchemasAgentsDangerousActionSchemaFilter()
		filter.SetIds([]string{id})
		params := sentinelone.NewAgentsSchemasAgentsDangerousActionSchema(*filter)
		_, _, err = client.AgentActionsApi.DisconnectAgents(context.Background()).Body(*params).Execute()
	}
	if err != nil {
		panic(err)
	}
}