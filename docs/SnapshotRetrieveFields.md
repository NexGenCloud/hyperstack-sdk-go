# SnapshotRetrieveFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the snapshot | 
**HasFloatingIp** | Pointer to **bool** | Indicates if the VM had a floating IP assigned | [optional] 
**Id** | **int32** | Snapshot ID | 
**IsImage** | **bool** | Indicates if the snapshot is an image | 
**Name** | **string** | Snapshot name | 
**RegionId** | **int32** | Region where the snapshot will be available | 
**Size** | **int32** | Size in GB of the snapshot | 
**Status** | **string** | Status of the snapshot | 
**VmId** | **int32** | ID of the VM from which the snapshot is created | 

## Methods

### NewSnapshotRetrieveFields

`func NewSnapshotRetrieveFields(description string, id int32, isImage bool, name string, regionId int32, size int32, status string, vmId int32, ) *SnapshotRetrieveFields`

NewSnapshotRetrieveFields instantiates a new SnapshotRetrieveFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotRetrieveFieldsWithDefaults

`func NewSnapshotRetrieveFieldsWithDefaults() *SnapshotRetrieveFields`

NewSnapshotRetrieveFieldsWithDefaults instantiates a new SnapshotRetrieveFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


