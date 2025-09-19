# ResourceLevelBillingHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingHistory** | Pointer to [**[]ResourceLevelBillingHistoryResources**](ResourceLevelBillingHistoryResources.md) |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**Pagination** | Pointer to [**PaginationData**](PaginationData.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewResourceLevelBillingHistory

`func NewResourceLevelBillingHistory() *ResourceLevelBillingHistory`

NewResourceLevelBillingHistory instantiates a new ResourceLevelBillingHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceLevelBillingHistoryWithDefaults

`func NewResourceLevelBillingHistoryWithDefaults() *ResourceLevelBillingHistory`

NewResourceLevelBillingHistoryWithDefaults instantiates a new ResourceLevelBillingHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingHistory

`func (o *ResourceLevelBillingHistory) GetBillingHistory() []ResourceLevelBillingHistoryResources`

GetBillingHistory returns the BillingHistory field if non-nil, zero value otherwise.

### GetBillingHistoryOk

`func (o *ResourceLevelBillingHistory) GetBillingHistoryOk() (*[]ResourceLevelBillingHistoryResources, bool)`

GetBillingHistoryOk returns a tuple with the BillingHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingHistory

`func (o *ResourceLevelBillingHistory) SetBillingHistory(v []ResourceLevelBillingHistoryResources)`

SetBillingHistory sets BillingHistory field to given value.

### HasBillingHistory

`func (o *ResourceLevelBillingHistory) HasBillingHistory() bool`

HasBillingHistory returns a boolean if a field has been set.

### GetOrgId

`func (o *ResourceLevelBillingHistory) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ResourceLevelBillingHistory) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ResourceLevelBillingHistory) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ResourceLevelBillingHistory) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPagination

`func (o *ResourceLevelBillingHistory) GetPagination() PaginationData`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ResourceLevelBillingHistory) GetPaginationOk() (*PaginationData, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ResourceLevelBillingHistory) SetPagination(v PaginationData)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ResourceLevelBillingHistory) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetTotalCount

`func (o *ResourceLevelBillingHistory) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ResourceLevelBillingHistory) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ResourceLevelBillingHistory) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ResourceLevelBillingHistory) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


