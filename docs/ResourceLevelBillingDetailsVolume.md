# ResourceLevelBillingDetailsVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingHistory** | Pointer to [**[]ResourceLevelBillingVolumeDetailsResources**](ResourceLevelBillingVolumeDetailsResources.md) |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewResourceLevelBillingDetailsVolume

`func NewResourceLevelBillingDetailsVolume() *ResourceLevelBillingDetailsVolume`

NewResourceLevelBillingDetailsVolume instantiates a new ResourceLevelBillingDetailsVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceLevelBillingDetailsVolumeWithDefaults

`func NewResourceLevelBillingDetailsVolumeWithDefaults() *ResourceLevelBillingDetailsVolume`

NewResourceLevelBillingDetailsVolumeWithDefaults instantiates a new ResourceLevelBillingDetailsVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingHistory

`func (o *ResourceLevelBillingDetailsVolume) GetBillingHistory() []ResourceLevelBillingVolumeDetailsResources`

GetBillingHistory returns the BillingHistory field if non-nil, zero value otherwise.

### GetBillingHistoryOk

`func (o *ResourceLevelBillingDetailsVolume) GetBillingHistoryOk() (*[]ResourceLevelBillingVolumeDetailsResources, bool)`

GetBillingHistoryOk returns a tuple with the BillingHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingHistory

`func (o *ResourceLevelBillingDetailsVolume) SetBillingHistory(v []ResourceLevelBillingVolumeDetailsResources)`

SetBillingHistory sets BillingHistory field to given value.

### HasBillingHistory

`func (o *ResourceLevelBillingDetailsVolume) HasBillingHistory() bool`

HasBillingHistory returns a boolean if a field has been set.

### GetOrgId

`func (o *ResourceLevelBillingDetailsVolume) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ResourceLevelBillingDetailsVolume) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ResourceLevelBillingDetailsVolume) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ResourceLevelBillingDetailsVolume) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetTotalCount

`func (o *ResourceLevelBillingDetailsVolume) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ResourceLevelBillingDetailsVolume) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ResourceLevelBillingDetailsVolume) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ResourceLevelBillingDetailsVolume) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


