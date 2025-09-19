# FlavorItemGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flavors** | Pointer to [**[]FlavorFields**](FlavorFields.md) |  | [optional] 
**Gpu** | Pointer to **string** |  | [optional] 
**RegionName** | Pointer to **string** |  | [optional] 

## Methods

### NewFlavorItemGetResponse

`func NewFlavorItemGetResponse() *FlavorItemGetResponse`

NewFlavorItemGetResponse instantiates a new FlavorItemGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlavorItemGetResponseWithDefaults

`func NewFlavorItemGetResponseWithDefaults() *FlavorItemGetResponse`

NewFlavorItemGetResponseWithDefaults instantiates a new FlavorItemGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlavors

`func (o *FlavorItemGetResponse) GetFlavors() []FlavorFields`

GetFlavors returns the Flavors field if non-nil, zero value otherwise.

### GetFlavorsOk

`func (o *FlavorItemGetResponse) GetFlavorsOk() (*[]FlavorFields, bool)`

GetFlavorsOk returns a tuple with the Flavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavors

`func (o *FlavorItemGetResponse) SetFlavors(v []FlavorFields)`

SetFlavors sets Flavors field to given value.

### HasFlavors

`func (o *FlavorItemGetResponse) HasFlavors() bool`

HasFlavors returns a boolean if a field has been set.

### GetGpu

`func (o *FlavorItemGetResponse) GetGpu() string`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *FlavorItemGetResponse) GetGpuOk() (*string, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *FlavorItemGetResponse) SetGpu(v string)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *FlavorItemGetResponse) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetRegionName

`func (o *FlavorItemGetResponse) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *FlavorItemGetResponse) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *FlavorItemGetResponse) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *FlavorItemGetResponse) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


