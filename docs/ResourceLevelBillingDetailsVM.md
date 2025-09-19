# ResourceLevelBillingDetailsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingHistory** | Pointer to [**[]ResourceLevelBillingVMDetailsResources**](ResourceLevelBillingVMDetailsResources.md) |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewResourceLevelBillingDetailsVM

`func NewResourceLevelBillingDetailsVM() *ResourceLevelBillingDetailsVM`

NewResourceLevelBillingDetailsVM instantiates a new ResourceLevelBillingDetailsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceLevelBillingDetailsVMWithDefaults

`func NewResourceLevelBillingDetailsVMWithDefaults() *ResourceLevelBillingDetailsVM`

NewResourceLevelBillingDetailsVMWithDefaults instantiates a new ResourceLevelBillingDetailsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingHistory

`func (o *ResourceLevelBillingDetailsVM) GetBillingHistory() []ResourceLevelBillingVMDetailsResources`

GetBillingHistory returns the BillingHistory field if non-nil, zero value otherwise.

### GetBillingHistoryOk

`func (o *ResourceLevelBillingDetailsVM) GetBillingHistoryOk() (*[]ResourceLevelBillingVMDetailsResources, bool)`

GetBillingHistoryOk returns a tuple with the BillingHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingHistory

`func (o *ResourceLevelBillingDetailsVM) SetBillingHistory(v []ResourceLevelBillingVMDetailsResources)`

SetBillingHistory sets BillingHistory field to given value.

### HasBillingHistory

`func (o *ResourceLevelBillingDetailsVM) HasBillingHistory() bool`

HasBillingHistory returns a boolean if a field has been set.

### GetOrgId

`func (o *ResourceLevelBillingDetailsVM) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ResourceLevelBillingDetailsVM) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ResourceLevelBillingDetailsVM) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ResourceLevelBillingDetailsVM) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetTotalCount

`func (o *ResourceLevelBillingDetailsVM) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ResourceLevelBillingDetailsVM) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ResourceLevelBillingDetailsVM) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ResourceLevelBillingDetailsVM) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


