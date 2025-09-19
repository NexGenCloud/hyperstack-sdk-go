# VolumeAttachmentFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Device** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Protected** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Volume** | Pointer to [**VolumeFieldsForInstance**](VolumeFieldsForInstance.md) |  | [optional] 

## Methods

### NewVolumeAttachmentFields

`func NewVolumeAttachmentFields() *VolumeAttachmentFields`

NewVolumeAttachmentFields instantiates a new VolumeAttachmentFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeAttachmentFieldsWithDefaults

`func NewVolumeAttachmentFieldsWithDefaults() *VolumeAttachmentFields`

NewVolumeAttachmentFieldsWithDefaults instantiates a new VolumeAttachmentFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *VolumeAttachmentFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VolumeAttachmentFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VolumeAttachmentFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VolumeAttachmentFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDevice

`func (o *VolumeAttachmentFields) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *VolumeAttachmentFields) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *VolumeAttachmentFields) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *VolumeAttachmentFields) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetId

`func (o *VolumeAttachmentFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeAttachmentFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeAttachmentFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VolumeAttachmentFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProtected

`func (o *VolumeAttachmentFields) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *VolumeAttachmentFields) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *VolumeAttachmentFields) SetProtected(v bool)`

SetProtected sets Protected field to given value.

### HasProtected

`func (o *VolumeAttachmentFields) HasProtected() bool`

HasProtected returns a boolean if a field has been set.

### GetStatus

`func (o *VolumeAttachmentFields) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VolumeAttachmentFields) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VolumeAttachmentFields) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VolumeAttachmentFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVolume

`func (o *VolumeAttachmentFields) GetVolume() VolumeFieldsForInstance`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *VolumeAttachmentFields) GetVolumeOk() (*VolumeFieldsForInstance, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *VolumeAttachmentFields) SetVolume(v VolumeFieldsForInstance)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *VolumeAttachmentFields) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


