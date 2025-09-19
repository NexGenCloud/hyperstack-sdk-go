# SnapshotRetrieveFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Creation timestamp | 
**CustomImage** | Pointer to **string** |  | [optional] 
**Description** | **string** | Description of the snapshot | 
**HasFloatingIp** | Pointer to **bool** | Indicates if the VM had a floating IP assigned | [optional] 
**Id** | **int32** | Snapshot ID | 
**IsImage** | **bool** | Indicates if the snapshot is an image | 
**Labels** | Pointer to **string** |  | [optional] 
**Name** | **string** | Snapshot name | 
**Region** | Pointer to **string** |  | [optional] 
**RegionId** | **int32** | Region where the snapshot will be available | 
**Size** | **int32** | Size in GB of the snapshot | 
**Status** | **string** | Status of the snapshot | 
**UpdatedAt** | **time.Time** | Last update timestamp | 
**VmEnvironment** | Pointer to **string** |  | [optional] 
**VmFlavor** | Pointer to **string** |  | [optional] 
**VmId** | **int32** | ID of the VM from which the snapshot is created | 
**VmImage** | Pointer to **string** |  | [optional] 
**VmKeypair** | Pointer to **string** |  | [optional] 
**VmName** | Pointer to **string** |  | [optional] 
**VmStatus** | Pointer to **string** |  | [optional] 
**VolumeId** | Pointer to **string** |  | [optional] 
**VolumeName** | Pointer to **string** |  | [optional] 

## Methods

### NewSnapshotRetrieveFields

`func NewSnapshotRetrieveFields(createdAt time.Time, description string, id int32, isImage bool, name string, regionId int32, size int32, status string, updatedAt time.Time, vmId int32, ) *SnapshotRetrieveFields`

NewSnapshotRetrieveFields instantiates a new SnapshotRetrieveFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotRetrieveFieldsWithDefaults

`func NewSnapshotRetrieveFieldsWithDefaults() *SnapshotRetrieveFields`

NewSnapshotRetrieveFieldsWithDefaults instantiates a new SnapshotRetrieveFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SnapshotRetrieveFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SnapshotRetrieveFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SnapshotRetrieveFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCustomImage

`func (o *SnapshotRetrieveFields) GetCustomImage() string`

GetCustomImage returns the CustomImage field if non-nil, zero value otherwise.

### GetCustomImageOk

`func (o *SnapshotRetrieveFields) GetCustomImageOk() (*string, bool)`

GetCustomImageOk returns a tuple with the CustomImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomImage

`func (o *SnapshotRetrieveFields) SetCustomImage(v string)`

SetCustomImage sets CustomImage field to given value.

### HasCustomImage

`func (o *SnapshotRetrieveFields) HasCustomImage() bool`

HasCustomImage returns a boolean if a field has been set.

### GetDescription

`func (o *SnapshotRetrieveFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SnapshotRetrieveFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SnapshotRetrieveFields) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHasFloatingIp

`func (o *SnapshotRetrieveFields) GetHasFloatingIp() bool`

GetHasFloatingIp returns the HasFloatingIp field if non-nil, zero value otherwise.

### GetHasFloatingIpOk

`func (o *SnapshotRetrieveFields) GetHasFloatingIpOk() (*bool, bool)`

GetHasFloatingIpOk returns a tuple with the HasFloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFloatingIp

`func (o *SnapshotRetrieveFields) SetHasFloatingIp(v bool)`

SetHasFloatingIp sets HasFloatingIp field to given value.

### HasHasFloatingIp

`func (o *SnapshotRetrieveFields) HasHasFloatingIp() bool`

HasHasFloatingIp returns a boolean if a field has been set.

### GetId

`func (o *SnapshotRetrieveFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotRetrieveFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotRetrieveFields) SetId(v int32)`

SetId sets Id field to given value.


### GetIsImage

`func (o *SnapshotRetrieveFields) GetIsImage() bool`

GetIsImage returns the IsImage field if non-nil, zero value otherwise.

### GetIsImageOk

`func (o *SnapshotRetrieveFields) GetIsImageOk() (*bool, bool)`

GetIsImageOk returns a tuple with the IsImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImage

`func (o *SnapshotRetrieveFields) SetIsImage(v bool)`

SetIsImage sets IsImage field to given value.


### GetLabels

`func (o *SnapshotRetrieveFields) GetLabels() string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SnapshotRetrieveFields) GetLabelsOk() (*string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SnapshotRetrieveFields) SetLabels(v string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SnapshotRetrieveFields) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *SnapshotRetrieveFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnapshotRetrieveFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnapshotRetrieveFields) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *SnapshotRetrieveFields) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SnapshotRetrieveFields) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SnapshotRetrieveFields) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *SnapshotRetrieveFields) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRegionId

`func (o *SnapshotRetrieveFields) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *SnapshotRetrieveFields) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *SnapshotRetrieveFields) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.


### GetSize

`func (o *SnapshotRetrieveFields) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SnapshotRetrieveFields) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SnapshotRetrieveFields) SetSize(v int32)`

SetSize sets Size field to given value.


### GetStatus

`func (o *SnapshotRetrieveFields) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SnapshotRetrieveFields) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SnapshotRetrieveFields) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *SnapshotRetrieveFields) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SnapshotRetrieveFields) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SnapshotRetrieveFields) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVmEnvironment

`func (o *SnapshotRetrieveFields) GetVmEnvironment() string`

GetVmEnvironment returns the VmEnvironment field if non-nil, zero value otherwise.

### GetVmEnvironmentOk

`func (o *SnapshotRetrieveFields) GetVmEnvironmentOk() (*string, bool)`

GetVmEnvironmentOk returns a tuple with the VmEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmEnvironment

`func (o *SnapshotRetrieveFields) SetVmEnvironment(v string)`

SetVmEnvironment sets VmEnvironment field to given value.

### HasVmEnvironment

`func (o *SnapshotRetrieveFields) HasVmEnvironment() bool`

HasVmEnvironment returns a boolean if a field has been set.

### GetVmFlavor

`func (o *SnapshotRetrieveFields) GetVmFlavor() string`

GetVmFlavor returns the VmFlavor field if non-nil, zero value otherwise.

### GetVmFlavorOk

`func (o *SnapshotRetrieveFields) GetVmFlavorOk() (*string, bool)`

GetVmFlavorOk returns a tuple with the VmFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmFlavor

`func (o *SnapshotRetrieveFields) SetVmFlavor(v string)`

SetVmFlavor sets VmFlavor field to given value.

### HasVmFlavor

`func (o *SnapshotRetrieveFields) HasVmFlavor() bool`

HasVmFlavor returns a boolean if a field has been set.

### GetVmId

`func (o *SnapshotRetrieveFields) GetVmId() int32`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *SnapshotRetrieveFields) GetVmIdOk() (*int32, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *SnapshotRetrieveFields) SetVmId(v int32)`

SetVmId sets VmId field to given value.


### GetVmImage

`func (o *SnapshotRetrieveFields) GetVmImage() string`

GetVmImage returns the VmImage field if non-nil, zero value otherwise.

### GetVmImageOk

`func (o *SnapshotRetrieveFields) GetVmImageOk() (*string, bool)`

GetVmImageOk returns a tuple with the VmImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmImage

`func (o *SnapshotRetrieveFields) SetVmImage(v string)`

SetVmImage sets VmImage field to given value.

### HasVmImage

`func (o *SnapshotRetrieveFields) HasVmImage() bool`

HasVmImage returns a boolean if a field has been set.

### GetVmKeypair

`func (o *SnapshotRetrieveFields) GetVmKeypair() string`

GetVmKeypair returns the VmKeypair field if non-nil, zero value otherwise.

### GetVmKeypairOk

`func (o *SnapshotRetrieveFields) GetVmKeypairOk() (*string, bool)`

GetVmKeypairOk returns a tuple with the VmKeypair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmKeypair

`func (o *SnapshotRetrieveFields) SetVmKeypair(v string)`

SetVmKeypair sets VmKeypair field to given value.

### HasVmKeypair

`func (o *SnapshotRetrieveFields) HasVmKeypair() bool`

HasVmKeypair returns a boolean if a field has been set.

### GetVmName

`func (o *SnapshotRetrieveFields) GetVmName() string`

GetVmName returns the VmName field if non-nil, zero value otherwise.

### GetVmNameOk

`func (o *SnapshotRetrieveFields) GetVmNameOk() (*string, bool)`

GetVmNameOk returns a tuple with the VmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmName

`func (o *SnapshotRetrieveFields) SetVmName(v string)`

SetVmName sets VmName field to given value.

### HasVmName

`func (o *SnapshotRetrieveFields) HasVmName() bool`

HasVmName returns a boolean if a field has been set.

### GetVmStatus

`func (o *SnapshotRetrieveFields) GetVmStatus() string`

GetVmStatus returns the VmStatus field if non-nil, zero value otherwise.

### GetVmStatusOk

`func (o *SnapshotRetrieveFields) GetVmStatusOk() (*string, bool)`

GetVmStatusOk returns a tuple with the VmStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmStatus

`func (o *SnapshotRetrieveFields) SetVmStatus(v string)`

SetVmStatus sets VmStatus field to given value.

### HasVmStatus

`func (o *SnapshotRetrieveFields) HasVmStatus() bool`

HasVmStatus returns a boolean if a field has been set.

### GetVolumeId

`func (o *SnapshotRetrieveFields) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *SnapshotRetrieveFields) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *SnapshotRetrieveFields) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *SnapshotRetrieveFields) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### GetVolumeName

`func (o *SnapshotRetrieveFields) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *SnapshotRetrieveFields) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *SnapshotRetrieveFields) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *SnapshotRetrieveFields) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


