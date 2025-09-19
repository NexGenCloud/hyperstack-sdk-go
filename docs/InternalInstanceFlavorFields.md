# InternalInstanceFlavorFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Disk** | Pointer to **int32** |  | [optional] 
**Gpu** | Pointer to **string** |  | [optional] 
**GpuCount** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Ram** | Pointer to **int32** |  | [optional] 
**RegionName** | Pointer to **string** |  | [optional] 

## Methods

### NewInternalInstanceFlavorFields

`func NewInternalInstanceFlavorFields() *InternalInstanceFlavorFields`

NewInternalInstanceFlavorFields instantiates a new InternalInstanceFlavorFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalInstanceFlavorFieldsWithDefaults

`func NewInternalInstanceFlavorFieldsWithDefaults() *InternalInstanceFlavorFields`

NewInternalInstanceFlavorFieldsWithDefaults instantiates a new InternalInstanceFlavorFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *InternalInstanceFlavorFields) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *InternalInstanceFlavorFields) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *InternalInstanceFlavorFields) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *InternalInstanceFlavorFields) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InternalInstanceFlavorFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InternalInstanceFlavorFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InternalInstanceFlavorFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InternalInstanceFlavorFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisk

`func (o *InternalInstanceFlavorFields) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *InternalInstanceFlavorFields) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *InternalInstanceFlavorFields) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *InternalInstanceFlavorFields) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetGpu

`func (o *InternalInstanceFlavorFields) GetGpu() string`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *InternalInstanceFlavorFields) GetGpuOk() (*string, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *InternalInstanceFlavorFields) SetGpu(v string)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *InternalInstanceFlavorFields) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetGpuCount

`func (o *InternalInstanceFlavorFields) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *InternalInstanceFlavorFields) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *InternalInstanceFlavorFields) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *InternalInstanceFlavorFields) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetId

`func (o *InternalInstanceFlavorFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalInstanceFlavorFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalInstanceFlavorFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InternalInstanceFlavorFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InternalInstanceFlavorFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternalInstanceFlavorFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternalInstanceFlavorFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InternalInstanceFlavorFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRam

`func (o *InternalInstanceFlavorFields) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *InternalInstanceFlavorFields) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *InternalInstanceFlavorFields) SetRam(v int32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *InternalInstanceFlavorFields) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetRegionName

`func (o *InternalInstanceFlavorFields) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *InternalInstanceFlavorFields) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *InternalInstanceFlavorFields) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *InternalInstanceFlavorFields) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


