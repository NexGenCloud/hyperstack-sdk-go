# Contract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocatedGpuCountGraph** | Pointer to [**[]AllocatedGPUCountGraph**](AllocatedGPUCountGraph.md) |  | [optional] 
**Granularity** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TotalGpuAllocation** | Pointer to **int32** |  | [optional] 

## Methods

### NewContract

`func NewContract() *Contract`

NewContract instantiates a new Contract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractWithDefaults

`func NewContractWithDefaults() *Contract`

NewContractWithDefaults instantiates a new Contract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocatedGpuCountGraph

`func (o *Contract) GetAllocatedGpuCountGraph() []AllocatedGPUCountGraph`

GetAllocatedGpuCountGraph returns the AllocatedGpuCountGraph field if non-nil, zero value otherwise.

### GetAllocatedGpuCountGraphOk

`func (o *Contract) GetAllocatedGpuCountGraphOk() (*[]AllocatedGPUCountGraph, bool)`

GetAllocatedGpuCountGraphOk returns a tuple with the AllocatedGpuCountGraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedGpuCountGraph

`func (o *Contract) SetAllocatedGpuCountGraph(v []AllocatedGPUCountGraph)`

SetAllocatedGpuCountGraph sets AllocatedGpuCountGraph field to given value.

### HasAllocatedGpuCountGraph

`func (o *Contract) HasAllocatedGpuCountGraph() bool`

HasAllocatedGpuCountGraph returns a boolean if a field has been set.

### GetGranularity

`func (o *Contract) GetGranularity() int32`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *Contract) GetGranularityOk() (*int32, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *Contract) SetGranularity(v int32)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *Contract) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetId

`func (o *Contract) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Contract) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Contract) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Contract) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *Contract) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Contract) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Contract) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Contract) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetStatus

`func (o *Contract) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Contract) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Contract) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Contract) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalGpuAllocation

`func (o *Contract) GetTotalGpuAllocation() int32`

GetTotalGpuAllocation returns the TotalGpuAllocation field if non-nil, zero value otherwise.

### GetTotalGpuAllocationOk

`func (o *Contract) GetTotalGpuAllocationOk() (*int32, bool)`

GetTotalGpuAllocationOk returns a tuple with the TotalGpuAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGpuAllocation

`func (o *Contract) SetTotalGpuAllocation(v int32)`

SetTotalGpuAllocation sets TotalGpuAllocation field to given value.

### HasTotalGpuAllocation

`func (o *Contract) HasTotalGpuAllocation() bool`

HasTotalGpuAllocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


