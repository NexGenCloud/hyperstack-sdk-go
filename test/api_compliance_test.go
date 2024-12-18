/*
Infrahub-API

Testing ComplianceAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package hyperstack

import (
	"context"
	"testing"

	openapiclient "github.com/NexGenCloud/hyperstack-sdk-go/hyperstack"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_hyperstack_ComplianceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ComplianceAPIService CreateCompliance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ComplianceAPI.CreateCompliance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ComplianceAPIService DeleteACompliance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var gpuModel string

		resp, httpRes, err := apiClient.ComplianceAPI.DeleteACompliance(context.Background(), gpuModel).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ComplianceAPIService RetrieveGpuCompliance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ComplianceAPI.RetrieveGpuCompliance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ComplianceAPIService UpdateACompliance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ComplianceAPI.UpdateACompliance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
