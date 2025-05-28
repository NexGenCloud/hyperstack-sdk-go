# VolumesLastStatusChangeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**VolumeStatusList** | Pointer to [**[]VolumeStatusChangeFields**](VolumeStatusChangeFields.md) |  | [optional] 

## Methods

### NewVolumesLastStatusChangeResponse

`func NewVolumesLastStatusChangeResponse() *VolumesLastStatusChangeResponse`

NewVolumesLastStatusChangeResponse instantiates a new VolumesLastStatusChangeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumesLastStatusChangeResponseWithDefaults

`func NewVolumesLastStatusChangeResponseWithDefaults() *VolumesLastStatusChangeResponse`

NewVolumesLastStatusChangeResponseWithDefaults instantiates a new VolumesLastStatusChangeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *VolumesLastStatusChangeResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VolumesLastStatusChangeResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VolumesLastStatusChangeResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *VolumesLastStatusChangeResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *VolumesLastStatusChangeResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VolumesLastStatusChangeResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VolumesLastStatusChangeResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VolumesLastStatusChangeResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVolumeStatusList

`func (o *VolumesLastStatusChangeResponse) GetVolumeStatusList() []VolumeStatusChangeFields`

GetVolumeStatusList returns the VolumeStatusList field if non-nil, zero value otherwise.

### GetVolumeStatusListOk

`func (o *VolumesLastStatusChangeResponse) GetVolumeStatusListOk() (*[]VolumeStatusChangeFields, bool)`

GetVolumeStatusListOk returns a tuple with the VolumeStatusList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeStatusList

`func (o *VolumesLastStatusChangeResponse) SetVolumeStatusList(v []VolumeStatusChangeFields)`

SetVolumeStatusList sets VolumeStatusList field to given value.

### HasVolumeStatusList

`func (o *VolumesLastStatusChangeResponse) HasVolumeStatusList() bool`

HasVolumeStatusList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


