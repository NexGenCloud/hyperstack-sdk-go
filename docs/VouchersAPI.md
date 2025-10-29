# \VouchersAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RedeemAVoucher**](VouchersAPI.md#RedeemAVoucher) | **Post** /billing/billing/vouchers/redeem | Redeem a voucher with a voucher_code



## RedeemAVoucher

> VoucherRedeemResponseSchema RedeemAVoucher(ctx).Payload(payload).Execute()

Redeem a voucher with a voucher_code



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
	payload := *openapiclient.NewRedeemVoucherPayload("VoucherCode_example") // RedeemVoucherPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VouchersAPI.RedeemAVoucher(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VouchersAPI.RedeemAVoucher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RedeemAVoucher`: VoucherRedeemResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `VouchersAPI.RedeemAVoucher`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRedeemAVoucherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**RedeemVoucherPayload**](RedeemVoucherPayload.md) |  | 

### Return type

[**VoucherRedeemResponseSchema**](VoucherRedeemResponseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

