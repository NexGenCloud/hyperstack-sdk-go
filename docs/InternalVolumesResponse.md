# InternalVolumesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Volumes** | Pointer to [**[]InternalVolumeFields**](InternalVolumeFields.md) |  | [optional] 

## Methods

### NewInternalVolumesResponse

`func NewInternalVolumesResponse() *InternalVolumesResponse`

NewInternalVolumesResponse instantiates a new InternalVolumesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalVolumesResponseWithDefaults

`func NewInternalVolumesResponseWithDefaults() *InternalVolumesResponse`

NewInternalVolumesResponseWithDefaults instantiates a new InternalVolumesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *InternalVolumesResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InternalVolumesResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InternalVolumesResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InternalVolumesResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *InternalVolumesResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalVolumesResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalVolumesResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InternalVolumesResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVolumes

`func (o *InternalVolumesResponse) GetVolumes() []InternalVolumeFields`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *InternalVolumesResponse) GetVolumesOk() (*[]InternalVolumeFields, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *InternalVolumesResponse) SetVolumes(v []InternalVolumeFields)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *InternalVolumesResponse) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


