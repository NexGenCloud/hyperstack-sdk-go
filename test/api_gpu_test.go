/*
Infrahub-API

Testing GpuAPIService

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

func Test_hyperstack_GpuAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GpuAPIService ListGpus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GpuAPI.ListGpus(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
