# \CreditAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
<<<<<<< HEAD
[**GetCredit2**](CreditAPI.md#GetCredit2) | **Get** /billing/user-credit/credit | GET: View credit and threshold



## GetCredit2

> GetCreditAndThresholdInfoInResponse GetCredit2(ctx).Execute()

GET: View credit and threshold



=======
[**GetViewCreditAndThreshold**](CreditAPI.md#GetViewCreditAndThreshold) | **Get** /billing/user-credit/credit | GET: View credit and threshold



## GetViewCreditAndThreshold

> Getcreditandthresholdinfoinresponse GetViewCreditAndThreshold(ctx).Execute()

GET: View credit and threshold

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
	resp, r, err := apiClient.CreditAPI.GetCredit2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditAPI.GetCredit2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredit2`: GetCreditAndThresholdInfoInResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditAPI.GetCredit2`: %v\n", resp)
=======
	resp, r, err := apiClient.CreditAPI.GetViewCreditAndThreshold(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditAPI.GetViewCreditAndThreshold``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetViewCreditAndThreshold`: Getcreditandthresholdinfoinresponse
	fmt.Fprintf(os.Stdout, "Response from `CreditAPI.GetViewCreditAndThreshold`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiGetCredit2Request struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiGetViewCreditAndThresholdRequest struct via the builder pattern
>>>>>>> main


### Return type

<<<<<<< HEAD
[**GetCreditAndThresholdInfoInResponse**](GetCreditAndThresholdInfoInResponse.md)

### Authorization

[apiKey](../README.md#apiKey)
=======
[**Getcreditandthresholdinfoinresponse**](Getcreditandthresholdinfoinresponse.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)
>>>>>>> main

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

