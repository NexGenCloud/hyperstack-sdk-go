/*
Infrahub-API

Testing PaymentAPIService

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

func Test_hyperstack_PaymentAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PaymentAPIService GetViewPaymentDetails", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PaymentAPI.GetViewPaymentDetails(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentAPIService PostInitiatePayment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PaymentAPI.PostInitiatePayment(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
