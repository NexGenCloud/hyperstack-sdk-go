# ResourceBillingEventsHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingEventsHistory** | Pointer to [**[]ResourceBillingEventsHistoryMetrics**](ResourceBillingEventsHistoryMetrics.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewResourceBillingEventsHistoryResponse

`func NewResourceBillingEventsHistoryResponse() *ResourceBillingEventsHistoryResponse`

NewResourceBillingEventsHistoryResponse instantiates a new ResourceBillingEventsHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceBillingEventsHistoryResponseWithDefaults

`func NewResourceBillingEventsHistoryResponseWithDefaults() *ResourceBillingEventsHistoryResponse`

NewResourceBillingEventsHistoryResponseWithDefaults instantiates a new ResourceBillingEventsHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingEventsHistory

`func (o *ResourceBillingEventsHistoryResponse) GetBillingEventsHistory() []ResourceBillingEventsHistoryMetrics`

GetBillingEventsHistory returns the BillingEventsHistory field if non-nil, zero value otherwise.

### GetBillingEventsHistoryOk

`func (o *ResourceBillingEventsHistoryResponse) GetBillingEventsHistoryOk() (*[]ResourceBillingEventsHistoryMetrics, bool)`

GetBillingEventsHistoryOk returns a tuple with the BillingEventsHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEventsHistory

`func (o *ResourceBillingEventsHistoryResponse) SetBillingEventsHistory(v []ResourceBillingEventsHistoryMetrics)`

SetBillingEventsHistory sets BillingEventsHistory field to given value.

### HasBillingEventsHistory

`func (o *ResourceBillingEventsHistoryResponse) HasBillingEventsHistory() bool`

HasBillingEventsHistory returns a boolean if a field has been set.

### GetMessage

`func (o *ResourceBillingEventsHistoryResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResourceBillingEventsHistoryResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResourceBillingEventsHistoryResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResourceBillingEventsHistoryResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ResourceBillingEventsHistoryResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceBillingEventsHistoryResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceBillingEventsHistoryResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResourceBillingEventsHistoryResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


