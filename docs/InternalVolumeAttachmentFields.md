# InternalVolumeAttachmentFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Device** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Volume** | Pointer to [**InternalVolumeFields**](InternalVolumeFields.md) |  | [optional] 

## Methods

### NewInternalVolumeAttachmentFields

`func NewInternalVolumeAttachmentFields() *InternalVolumeAttachmentFields`

NewInternalVolumeAttachmentFields instantiates a new InternalVolumeAttachmentFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalVolumeAttachmentFieldsWithDefaults

`func NewInternalVolumeAttachmentFieldsWithDefaults() *InternalVolumeAttachmentFields`

NewInternalVolumeAttachmentFieldsWithDefaults instantiates a new InternalVolumeAttachmentFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *InternalVolumeAttachmentFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InternalVolumeAttachmentFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InternalVolumeAttachmentFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InternalVolumeAttachmentFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDevice

`func (o *InternalVolumeAttachmentFields) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InternalVolumeAttachmentFields) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InternalVolumeAttachmentFields) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InternalVolumeAttachmentFields) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetStatus

`func (o *InternalVolumeAttachmentFields) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalVolumeAttachmentFields) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalVolumeAttachmentFields) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InternalVolumeAttachmentFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVolume

`func (o *InternalVolumeAttachmentFields) GetVolume() InternalVolumeFields`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *InternalVolumeAttachmentFields) GetVolumeOk() (*InternalVolumeFields, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *InternalVolumeAttachmentFields) SetVolume(v InternalVolumeFields)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *InternalVolumeAttachmentFields) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


