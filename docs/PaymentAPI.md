# \PaymentAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
<<<<<<< HEAD
[**GetDetails**](PaymentAPI.md#GetDetails) | **Get** /billing/payment/payment-details | GET: View payment details
[**GetPaymentReceipt2**](PaymentAPI.md#GetPaymentReceipt2) | **Get** /billing/payment/receipt/{payment_id} | Retrieve Payment Receipt
[**PostPayment**](PaymentAPI.md#PostPayment) | **Post** /billing/payment/payment-initiate | POST: Initiate payment



## GetDetails

> PaymentDetailsResponse GetDetails(ctx).Execute()

GET: View payment details



=======
[**GetViewPaymentDetails**](PaymentAPI.md#GetViewPaymentDetails) | **Get** /billing/payment/payment-details | GET: View payment details
[**PostInitiatePayment**](PaymentAPI.md#PostInitiatePayment) | **Post** /billing/payment/payment-initiate | POST: Initiate payment



## GetViewPaymentDetails

> PaymentDetailsResponse GetViewPaymentDetails(ctx).Execute()

GET: View payment details

>>>>>>> main
### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/NexGenCloud/hyperstack-sdk-go/hyperstack"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
<<<<<<< HEAD
	resp, r, err := apiClient.PaymentAPI.GetDetails(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.GetDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDetails`: PaymentDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.GetDetails`: %v\n", resp)
=======
	resp, r, err := apiClient.PaymentAPI.GetViewPaymentDetails(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.GetViewPaymentDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetViewPaymentDetails`: PaymentDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.GetViewPaymentDetails`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiGetDetailsRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiGetViewPaymentDetailsRequest struct via the builder pattern
>>>>>>> main


### Return type

[**PaymentDetailsResponse**](PaymentDetailsResponse.md)

### Authorization

<<<<<<< HEAD
[apiKey](../README.md#apiKey)
=======
[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)
>>>>>>> main

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


<<<<<<< HEAD
## GetPaymentReceipt2

> GetPaymentReceipt2(ctx, paymentId).Execute()

Retrieve Payment Receipt



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/NexGenCloud/hyperstack-sdk-go/hyperstack"
)

func main() {
	paymentId := "paymentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PaymentAPI.GetPaymentReceipt2(context.Background(), paymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.GetPaymentReceipt2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentReceipt2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPayment

> PaymentInitiateResponse PostPayment(ctx).Payload(payload).Execute()

POST: Initiate payment



=======
## PostInitiatePayment

> PaymentInitiateResponse PostInitiatePayment(ctx).Payload(payload).Execute()

POST: Initiate payment

>>>>>>> main
### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/NexGenCloud/hyperstack-sdk-go/hyperstack"
)

func main() {
	payload := *openapiclient.NewPaymentInitiatePayload() // PaymentInitiatePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
<<<<<<< HEAD
	resp, r, err := apiClient.PaymentAPI.PostPayment(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PostPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPayment`: PaymentInitiateResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PostPayment`: %v\n", resp)
=======
	resp, r, err := apiClient.PaymentAPI.PostInitiatePayment(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PostInitiatePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostInitiatePayment`: PaymentInitiateResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PostInitiatePayment`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters



### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiPostPaymentRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiPostInitiatePaymentRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**PaymentInitiatePayload**](PaymentInitiatePayload.md) |  | 

### Return type

[**PaymentInitiateResponse**](PaymentInitiateResponse.md)

### Authorization

<<<<<<< HEAD
[apiKey](../README.md#apiKey)
=======
[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)
>>>>>>> main

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

