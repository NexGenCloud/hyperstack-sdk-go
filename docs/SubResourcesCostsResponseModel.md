# SubResourcesCostsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingHistory** | Pointer to [**[]SubResourcesGraphBillingHistoryFields**](SubResourcesGraphBillingHistoryFields.md) |  | [optional] 
**Granularity** | Pointer to **int32** |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewSubResourcesCostsResponseModel

`func NewSubResourcesCostsResponseModel() *SubResourcesCostsResponseModel`

NewSubResourcesCostsResponseModel instantiates a new SubResourcesCostsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubResourcesCostsResponseModelWithDefaults

`func NewSubResourcesCostsResponseModelWithDefaults() *SubResourcesCostsResponseModel`

NewSubResourcesCostsResponseModelWithDefaults instantiates a new SubResourcesCostsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingHistory

`func (o *SubResourcesCostsResponseModel) GetBillingHistory() []SubResourcesGraphBillingHistoryFields`

GetBillingHistory returns the BillingHistory field if non-nil, zero value otherwise.

### GetBillingHistoryOk

`func (o *SubResourcesCostsResponseModel) GetBillingHistoryOk() (*[]SubResourcesGraphBillingHistoryFields, bool)`

GetBillingHistoryOk returns a tuple with the BillingHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingHistory

`func (o *SubResourcesCostsResponseModel) SetBillingHistory(v []SubResourcesGraphBillingHistoryFields)`

SetBillingHistory sets BillingHistory field to given value.

### HasBillingHistory

`func (o *SubResourcesCostsResponseModel) HasBillingHistory() bool`

HasBillingHistory returns a boolean if a field has been set.

### GetGranularity

`func (o *SubResourcesCostsResponseModel) GetGranularity() int32`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *SubResourcesCostsResponseModel) GetGranularityOk() (*int32, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *SubResourcesCostsResponseModel) SetGranularity(v int32)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *SubResourcesCostsResponseModel) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetOrgId

`func (o *SubResourcesCostsResponseModel) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SubResourcesCostsResponseModel) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SubResourcesCostsResponseModel) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *SubResourcesCostsResponseModel) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetTotalCount

`func (o *SubResourcesCostsResponseModel) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SubResourcesCostsResponseModel) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SubResourcesCostsResponseModel) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *SubResourcesCostsResponseModel) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


