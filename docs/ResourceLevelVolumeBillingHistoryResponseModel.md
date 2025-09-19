# ResourceLevelVolumeBillingHistoryResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingHistoryVolume** | Pointer to [**ResourceLevelBillingHistory**](ResourceLevelBillingHistory.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewResourceLevelVolumeBillingHistoryResponseModel

`func NewResourceLevelVolumeBillingHistoryResponseModel() *ResourceLevelVolumeBillingHistoryResponseModel`

NewResourceLevelVolumeBillingHistoryResponseModel instantiates a new ResourceLevelVolumeBillingHistoryResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceLevelVolumeBillingHistoryResponseModelWithDefaults

`func NewResourceLevelVolumeBillingHistoryResponseModelWithDefaults() *ResourceLevelVolumeBillingHistoryResponseModel`

NewResourceLevelVolumeBillingHistoryResponseModelWithDefaults instantiates a new ResourceLevelVolumeBillingHistoryResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingHistoryVolume

`func (o *ResourceLevelVolumeBillingHistoryResponseModel) GetBillingHistoryVolume() ResourceLevelBillingHistory`

GetBillingHistoryVolume returns the BillingHistoryVolume field if non-nil, zero value otherwise.

### GetBillingHistoryVolumeOk

`func (o *ResourceLevelVolumeBillingHistoryResponseModel) GetBillingHistoryVolumeOk() (*ResourceLevelBillingHistory, bool)`

GetBillingHistoryVolumeOk returns a tuple with the BillingHistoryVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingHistoryVolume

`func (o *ResourceLevelVolumeBillingHistoryResponseModel) SetBillingHistoryVolume(v ResourceLevelBillingHistory)`

SetBillingHistoryVolume sets BillingHistoryVolume field to given value.

### HasBillingHistoryVolume

`func (o *ResourceLevelVolumeBillingHistoryResponseModel) HasBillingHistoryVolume() bool`

HasBillingHistoryVolume returns a boolean if a field has been set.

### GetMessage

`func (o *ResourceLevelVolumeBillingHistoryResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResourceLevelVolumeBillingHistoryResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResourceLevelVolumeBillingHistoryResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResourceLevelVolumeBillingHistoryResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ResourceLevelVolumeBillingHistoryResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceLevelVolumeBillingHistoryResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceLevelVolumeBillingHistoryResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResourceLevelVolumeBillingHistoryResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


