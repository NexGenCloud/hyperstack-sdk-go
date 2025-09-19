# WorkloadBillingHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingHistoryFineTuning** | Pointer to [**BillingHistoryFineTuning**](BillingHistoryFineTuning.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewWorkloadBillingHistoryResponse

`func NewWorkloadBillingHistoryResponse() *WorkloadBillingHistoryResponse`

NewWorkloadBillingHistoryResponse instantiates a new WorkloadBillingHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadBillingHistoryResponseWithDefaults

`func NewWorkloadBillingHistoryResponseWithDefaults() *WorkloadBillingHistoryResponse`

NewWorkloadBillingHistoryResponseWithDefaults instantiates a new WorkloadBillingHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingHistoryFineTuning

`func (o *WorkloadBillingHistoryResponse) GetBillingHistoryFineTuning() BillingHistoryFineTuning`

GetBillingHistoryFineTuning returns the BillingHistoryFineTuning field if non-nil, zero value otherwise.

### GetBillingHistoryFineTuningOk

`func (o *WorkloadBillingHistoryResponse) GetBillingHistoryFineTuningOk() (*BillingHistoryFineTuning, bool)`

GetBillingHistoryFineTuningOk returns a tuple with the BillingHistoryFineTuning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingHistoryFineTuning

`func (o *WorkloadBillingHistoryResponse) SetBillingHistoryFineTuning(v BillingHistoryFineTuning)`

SetBillingHistoryFineTuning sets BillingHistoryFineTuning field to given value.

### HasBillingHistoryFineTuning

`func (o *WorkloadBillingHistoryResponse) HasBillingHistoryFineTuning() bool`

HasBillingHistoryFineTuning returns a boolean if a field has been set.

### GetMessage

`func (o *WorkloadBillingHistoryResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WorkloadBillingHistoryResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WorkloadBillingHistoryResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WorkloadBillingHistoryResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSuccess

`func (o *WorkloadBillingHistoryResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *WorkloadBillingHistoryResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *WorkloadBillingHistoryResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *WorkloadBillingHistoryResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


