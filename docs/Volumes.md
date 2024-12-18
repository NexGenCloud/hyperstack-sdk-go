# Volumes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Volume** | Pointer to [**[]VolumeFields**](VolumeFields.md) |  | [optional] 

## Methods

### NewVolumes

`func NewVolumes() *Volumes`

NewVolumes instantiates a new Volumes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumesWithDefaults

`func NewVolumesWithDefaults() *Volumes`

NewVolumesWithDefaults instantiates a new Volumes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *Volumes) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Volumes) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Volumes) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Volumes) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *Volumes) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Volumes) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Volumes) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Volumes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVolume

`func (o *Volumes) GetVolume() []VolumeFields`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *Volumes) GetVolumeOk() (*[]VolumeFields, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *Volumes) SetVolume(v []VolumeFields)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *Volumes) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


