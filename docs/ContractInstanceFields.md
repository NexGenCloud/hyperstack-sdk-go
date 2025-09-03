# ContractInstanceFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterFields**](ClusterFields.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**FlavorName** | Pointer to **string** |  | [optional] 
**GpuCount** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TerminationTime** | Pointer to **time.Time** |  | [optional] 
**TotalUsageTime** | Pointer to **int32** |  | [optional] 

## Methods

### NewContractInstanceFields

`func NewContractInstanceFields() *ContractInstanceFields`

NewContractInstanceFields instantiates a new ContractInstanceFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractInstanceFieldsWithDefaults

`func NewContractInstanceFieldsWithDefaults() *ContractInstanceFields`

NewContractInstanceFieldsWithDefaults instantiates a new ContractInstanceFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *ContractInstanceFields) GetCluster() ClusterFields`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ContractInstanceFields) GetClusterOk() (*ClusterFields, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ContractInstanceFields) SetCluster(v ClusterFields)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ContractInstanceFields) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ContractInstanceFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ContractInstanceFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ContractInstanceFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ContractInstanceFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFlavorName

`func (o *ContractInstanceFields) GetFlavorName() string`

GetFlavorName returns the FlavorName field if non-nil, zero value otherwise.

### GetFlavorNameOk

`func (o *ContractInstanceFields) GetFlavorNameOk() (*string, bool)`

GetFlavorNameOk returns a tuple with the FlavorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorName

`func (o *ContractInstanceFields) SetFlavorName(v string)`

SetFlavorName sets FlavorName field to given value.

### HasFlavorName

`func (o *ContractInstanceFields) HasFlavorName() bool`

HasFlavorName returns a boolean if a field has been set.

### GetGpuCount

`func (o *ContractInstanceFields) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *ContractInstanceFields) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *ContractInstanceFields) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *ContractInstanceFields) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetId

`func (o *ContractInstanceFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContractInstanceFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContractInstanceFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ContractInstanceFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ContractInstanceFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContractInstanceFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContractInstanceFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContractInstanceFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ContractInstanceFields) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContractInstanceFields) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContractInstanceFields) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ContractInstanceFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTerminationTime

`func (o *ContractInstanceFields) GetTerminationTime() time.Time`

GetTerminationTime returns the TerminationTime field if non-nil, zero value otherwise.

### GetTerminationTimeOk

`func (o *ContractInstanceFields) GetTerminationTimeOk() (*time.Time, bool)`

GetTerminationTimeOk returns a tuple with the TerminationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationTime

`func (o *ContractInstanceFields) SetTerminationTime(v time.Time)`

SetTerminationTime sets TerminationTime field to given value.

### HasTerminationTime

`func (o *ContractInstanceFields) HasTerminationTime() bool`

HasTerminationTime returns a boolean if a field has been set.

### GetTotalUsageTime

`func (o *ContractInstanceFields) GetTotalUsageTime() int32`

GetTotalUsageTime returns the TotalUsageTime field if non-nil, zero value otherwise.

### GetTotalUsageTimeOk

`func (o *ContractInstanceFields) GetTotalUsageTimeOk() (*int32, bool)`

GetTotalUsageTimeOk returns a tuple with the TotalUsageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageTime

`func (o *ContractInstanceFields) SetTotalUsageTime(v int32)`

SetTotalUsageTime sets TotalUsageTime field to given value.

### HasTotalUsageTime

`func (o *ContractInstanceFields) HasTotalUsageTime() bool`

HasTotalUsageTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


