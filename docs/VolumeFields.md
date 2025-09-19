# VolumeFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to [**[]AttachmentsFieldsForVolume**](AttachmentsFieldsForVolume.md) |  | [optional] 
**Bootable** | Pointer to **bool** |  | [optional] 
**CallbackUrl** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to [**EnvironmentFieldsForVolume**](EnvironmentFieldsForVolume.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**ImageId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OsImage** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**VolumeType** | Pointer to **string** |  | [optional] 

## Methods

### NewVolumeFields

`func NewVolumeFields() *VolumeFields`

NewVolumeFields instantiates a new VolumeFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeFieldsWithDefaults

`func NewVolumeFieldsWithDefaults() *VolumeFields`

NewVolumeFieldsWithDefaults instantiates a new VolumeFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *VolumeFields) GetAttachments() []AttachmentsFieldsForVolume`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *VolumeFields) GetAttachmentsOk() (*[]AttachmentsFieldsForVolume, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *VolumeFields) SetAttachments(v []AttachmentsFieldsForVolume)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *VolumeFields) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetBootable

`func (o *VolumeFields) GetBootable() bool`

GetBootable returns the Bootable field if non-nil, zero value otherwise.

### GetBootableOk

`func (o *VolumeFields) GetBootableOk() (*bool, bool)`

GetBootableOk returns a tuple with the Bootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootable

`func (o *VolumeFields) SetBootable(v bool)`

SetBootable sets Bootable field to given value.

### HasBootable

`func (o *VolumeFields) HasBootable() bool`

HasBootable returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *VolumeFields) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *VolumeFields) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *VolumeFields) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *VolumeFields) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VolumeFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VolumeFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VolumeFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VolumeFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *VolumeFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VolumeFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VolumeFields) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VolumeFields) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *VolumeFields) GetEnvironment() EnvironmentFieldsForVolume`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *VolumeFields) GetEnvironmentOk() (*EnvironmentFieldsForVolume, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *VolumeFields) SetEnvironment(v EnvironmentFieldsForVolume)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *VolumeFields) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *VolumeFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VolumeFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageId

`func (o *VolumeFields) GetImageId() int32`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *VolumeFields) GetImageIdOk() (*int32, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *VolumeFields) SetImageId(v int32)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *VolumeFields) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetName

`func (o *VolumeFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumeFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsImage

`func (o *VolumeFields) GetOsImage() string`

GetOsImage returns the OsImage field if non-nil, zero value otherwise.

### GetOsImageOk

`func (o *VolumeFields) GetOsImageOk() (*string, bool)`

GetOsImageOk returns a tuple with the OsImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsImage

`func (o *VolumeFields) SetOsImage(v string)`

SetOsImage sets OsImage field to given value.

### HasOsImage

`func (o *VolumeFields) HasOsImage() bool`

HasOsImage returns a boolean if a field has been set.

### GetSize

`func (o *VolumeFields) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VolumeFields) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VolumeFields) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *VolumeFields) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *VolumeFields) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VolumeFields) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VolumeFields) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VolumeFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VolumeFields) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VolumeFields) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VolumeFields) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VolumeFields) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVolumeType

`func (o *VolumeFields) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *VolumeFields) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *VolumeFields) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *VolumeFields) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


