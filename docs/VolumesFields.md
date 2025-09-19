# VolumesFields

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
**Size** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**VolumeType** | Pointer to **string** |  | [optional] 

## Methods

### NewVolumesFields

`func NewVolumesFields() *VolumesFields`

NewVolumesFields instantiates a new VolumesFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumesFieldsWithDefaults

`func NewVolumesFieldsWithDefaults() *VolumesFields`

NewVolumesFieldsWithDefaults instantiates a new VolumesFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *VolumesFields) GetAttachments() []AttachmentsFieldsForVolume`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *VolumesFields) GetAttachmentsOk() (*[]AttachmentsFieldsForVolume, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *VolumesFields) SetAttachments(v []AttachmentsFieldsForVolume)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *VolumesFields) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetBootable

`func (o *VolumesFields) GetBootable() bool`

GetBootable returns the Bootable field if non-nil, zero value otherwise.

### GetBootableOk

`func (o *VolumesFields) GetBootableOk() (*bool, bool)`

GetBootableOk returns a tuple with the Bootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootable

`func (o *VolumesFields) SetBootable(v bool)`

SetBootable sets Bootable field to given value.

### HasBootable

`func (o *VolumesFields) HasBootable() bool`

HasBootable returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *VolumesFields) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *VolumesFields) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *VolumesFields) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *VolumesFields) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VolumesFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VolumesFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VolumesFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VolumesFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *VolumesFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VolumesFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VolumesFields) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VolumesFields) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *VolumesFields) GetEnvironment() EnvironmentFieldsForVolume`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *VolumesFields) GetEnvironmentOk() (*EnvironmentFieldsForVolume, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *VolumesFields) SetEnvironment(v EnvironmentFieldsForVolume)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *VolumesFields) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *VolumesFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumesFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumesFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VolumesFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageId

`func (o *VolumesFields) GetImageId() int32`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *VolumesFields) GetImageIdOk() (*int32, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *VolumesFields) SetImageId(v int32)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *VolumesFields) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetName

`func (o *VolumesFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumesFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumesFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumesFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *VolumesFields) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VolumesFields) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VolumesFields) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *VolumesFields) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *VolumesFields) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VolumesFields) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VolumesFields) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VolumesFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VolumesFields) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VolumesFields) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VolumesFields) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VolumesFields) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVolumeType

`func (o *VolumesFields) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *VolumesFields) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *VolumesFields) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *VolumesFields) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


