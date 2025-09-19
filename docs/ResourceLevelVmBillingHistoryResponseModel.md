# ResourceLevelVmBillingHistoryResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingHistoryVm** | Pointer to [**ResourceLevelBillingHistory**](ResourceLevelBillingHistory.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewResourceLevelVmBillingHistoryResponseModel

`func NewResourceLevelVmBillingHistoryResponseModel() *ResourceLevelVmBillingHistoryResponseModel`

NewResourceLevelVmBillingHistoryResponseModel instantiates a new ResourceLevelVmBillingHistoryResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceLevelVmBillingHistoryResponseModelWithDefaults

`func NewResourceLevelVmBillingHistoryResponseModelWithDefaults() *ResourceLevelVmBillingHistoryResponseModel`

NewResourceLevelVmBillingHistoryResponseModelWithDefaults instantiates a new ResourceLevelVmBillingHistoryResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingHistoryVm

`func (o *ResourceLevelVmBillingHistoryResponseModel) GetBillingHistoryVm() ResourceLevelBillingHistory`

GetBillingHistoryVm returns the BillingHistoryVm field if non-nil, zero value otherwise.

### GetBillingHistoryVmOk

`func (o *ResourceLevelVmBillingHistoryResponseModel) GetBillingHistoryVmOk() (*ResourceLevelBillingHistory, bool)`

GetBillingHistoryVmOk returns a tuple with the BillingHistoryVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingHistoryVm

`func (o *ResourceLevelVmBillingHistoryResponseModel) SetBillingHistoryVm(v ResourceLevelBillingHistory)`

SetBillingHistoryVm sets BillingHistoryVm field to given value.

### HasBillingHistoryVm

`func (o *ResourceLevelVmBillingHistoryResponseModel) HasBillingHistoryVm() bool`

HasBillingHistoryVm returns a boolean if a field has been set.

### GetMessage

`func (o *ResourceLevelVmBillingHistoryResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResourceLevelVmBillingHistoryResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResourceLevelVmBillingHistoryResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResourceLevelVmBillingHistoryResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ResourceLevelVmBillingHistoryResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceLevelVmBillingHistoryResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceLevelVmBillingHistoryResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResourceLevelVmBillingHistoryResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


