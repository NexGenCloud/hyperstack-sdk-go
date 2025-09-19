# SnapshotFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Creation timestamp | 
**Description** | **string** | Description of the snapshot | 
**HasFloatingIp** | Pointer to **bool** | Indicates if the VM had a floating IP assigned | [optional] 
**Id** | **int32** | Snapshot ID | 
**IsImage** | **bool** | Indicates if the snapshot is an image | 
**Labels** | Pointer to **[]string** | Labels associated with snapshot | [optional] 
**Name** | **string** | Snapshot name | 
**RegionId** | **int32** | Region where the snapshot will be available | 
**Size** | **int32** | Size in GB of the snapshot | 
**Status** | **string** | Status of the snapshot | 
**UpdatedAt** | **time.Time** | Last update timestamp | 
**VmId** | **int32** | ID of the VM from which the snapshot is created | 

## Methods

### NewSnapshotFields

`func NewSnapshotFields(createdAt time.Time, description string, id int32, isImage bool, name string, regionId int32, size int32, status string, updatedAt time.Time, vmId int32, ) *SnapshotFields`

NewSnapshotFields instantiates a new SnapshotFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotFieldsWithDefaults

`func NewSnapshotFieldsWithDefaults() *SnapshotFields`

NewSnapshotFieldsWithDefaults instantiates a new SnapshotFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SnapshotFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SnapshotFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SnapshotFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *SnapshotFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SnapshotFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SnapshotFields) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHasFloatingIp

`func (o *SnapshotFields) GetHasFloatingIp() bool`

GetHasFloatingIp returns the HasFloatingIp field if non-nil, zero value otherwise.

### GetHasFloatingIpOk

`func (o *SnapshotFields) GetHasFloatingIpOk() (*bool, bool)`

GetHasFloatingIpOk returns a tuple with the HasFloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFloatingIp

`func (o *SnapshotFields) SetHasFloatingIp(v bool)`

SetHasFloatingIp sets HasFloatingIp field to given value.

### HasHasFloatingIp

`func (o *SnapshotFields) HasHasFloatingIp() bool`

HasHasFloatingIp returns a boolean if a field has been set.

### GetId

`func (o *SnapshotFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotFields) SetId(v int32)`

SetId sets Id field to given value.


### GetIsImage

`func (o *SnapshotFields) GetIsImage() bool`

GetIsImage returns the IsImage field if non-nil, zero value otherwise.

### GetIsImageOk

`func (o *SnapshotFields) GetIsImageOk() (*bool, bool)`

GetIsImageOk returns a tuple with the IsImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImage

`func (o *SnapshotFields) SetIsImage(v bool)`

SetIsImage sets IsImage field to given value.


### GetLabels

`func (o *SnapshotFields) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SnapshotFields) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SnapshotFields) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SnapshotFields) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *SnapshotFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnapshotFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnapshotFields) SetName(v string)`

SetName sets Name field to given value.


### GetRegionId

`func (o *SnapshotFields) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *SnapshotFields) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *SnapshotFields) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.


### GetSize

`func (o *SnapshotFields) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SnapshotFields) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SnapshotFields) SetSize(v int32)`

SetSize sets Size field to given value.


### GetStatus

`func (o *SnapshotFields) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SnapshotFields) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SnapshotFields) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *SnapshotFields) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SnapshotFields) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SnapshotFields) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVmId

`func (o *SnapshotFields) GetVmId() int32`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *SnapshotFields) GetVmIdOk() (*int32, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *SnapshotFields) SetVmId(v int32)`

SetVmId sets VmId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


