# VolumeAttachmentFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Device** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Volume** | Pointer to [**VolumeFieldsforInstance**](VolumeFieldsforInstance.md) |  | [optional] 

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

`func (o *VolumeAttachmentFields) GetVolume() VolumeFieldsforInstance`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *VolumeAttachmentFields) GetVolumeOk() (*VolumeFieldsforInstance, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *VolumeAttachmentFields) SetVolume(v VolumeFieldsforInstance)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *VolumeAttachmentFields) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


