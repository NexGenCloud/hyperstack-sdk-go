# FlavorFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Disk** | Pointer to **int32** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Ephemeral** | Pointer to **int32** |  | [optional] 
**Features** | Pointer to **map[string]interface{}** |  | [optional] 
**Gpu** | Pointer to **string** |  | [optional] 
**GpuCount** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Labels** | Pointer to [**[]LableResonse**](LableResonse.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Ram** | Pointer to **float32** |  | [optional] 
**RegionName** | Pointer to **string** |  | [optional] 
**StockAvailable** | Pointer to **bool** |  | [optional] 

## Methods

### NewFlavorFields

`func NewFlavorFields() *FlavorFields`

NewFlavorFields instantiates a new FlavorFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlavorFieldsWithDefaults

`func NewFlavorFieldsWithDefaults() *FlavorFields`

NewFlavorFieldsWithDefaults instantiates a new FlavorFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *FlavorFields) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *FlavorFields) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *FlavorFields) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *FlavorFields) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FlavorFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FlavorFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FlavorFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FlavorFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisk

`func (o *FlavorFields) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *FlavorFields) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *FlavorFields) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *FlavorFields) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetDisplayName

`func (o *FlavorFields) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FlavorFields) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FlavorFields) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FlavorFields) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEphemeral

`func (o *FlavorFields) GetEphemeral() int32`

GetEphemeral returns the Ephemeral field if non-nil, zero value otherwise.

### GetEphemeralOk

`func (o *FlavorFields) GetEphemeralOk() (*int32, bool)`

GetEphemeralOk returns a tuple with the Ephemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeral

`func (o *FlavorFields) SetEphemeral(v int32)`

SetEphemeral sets Ephemeral field to given value.

### HasEphemeral

`func (o *FlavorFields) HasEphemeral() bool`

HasEphemeral returns a boolean if a field has been set.

### GetFeatures

`func (o *FlavorFields) GetFeatures() map[string]interface{}`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *FlavorFields) GetFeaturesOk() (*map[string]interface{}, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *FlavorFields) SetFeatures(v map[string]interface{})`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *FlavorFields) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetGpu

`func (o *FlavorFields) GetGpu() string`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *FlavorFields) GetGpuOk() (*string, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *FlavorFields) SetGpu(v string)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *FlavorFields) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetGpuCount

`func (o *FlavorFields) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *FlavorFields) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *FlavorFields) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *FlavorFields) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetId

`func (o *FlavorFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlavorFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlavorFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FlavorFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabels

`func (o *FlavorFields) GetLabels() []LableResonse`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *FlavorFields) GetLabelsOk() (*[]LableResonse, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *FlavorFields) SetLabels(v []LableResonse)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *FlavorFields) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *FlavorFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlavorFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlavorFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlavorFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRam

`func (o *FlavorFields) GetRam() float32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *FlavorFields) GetRamOk() (*float32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *FlavorFields) SetRam(v float32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *FlavorFields) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetRegionName

`func (o *FlavorFields) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *FlavorFields) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *FlavorFields) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *FlavorFields) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetStockAvailable

`func (o *FlavorFields) GetStockAvailable() bool`

GetStockAvailable returns the StockAvailable field if non-nil, zero value otherwise.

### GetStockAvailableOk

`func (o *FlavorFields) GetStockAvailableOk() (*bool, bool)`

GetStockAvailableOk returns a tuple with the StockAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockAvailable

`func (o *FlavorFields) SetStockAvailable(v bool)`

SetStockAvailable sets StockAvailable field to given value.

### HasStockAvailable

`func (o *FlavorFields) HasStockAvailable() bool`

HasStockAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


