/*
Infrahub-API

Testing CalculateAPIService

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

func Test_hyperstack_CalculateAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CalculateAPIService RetrieveBillingRateForResource", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var resourceType string
		var id int32

		resp, httpRes, err := apiClient.CalculateAPI.RetrieveBillingRateForResource(context.Background(), resourceType, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
