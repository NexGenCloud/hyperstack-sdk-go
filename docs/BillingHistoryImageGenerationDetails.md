# BillingHistoryImageGenerationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingHistory** | Pointer to [**[]BillingHistory**](BillingHistory.md) |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewBillingHistoryImageGenerationDetails

`func NewBillingHistoryImageGenerationDetails() *BillingHistoryImageGenerationDetails`

NewBillingHistoryImageGenerationDetails instantiates a new BillingHistoryImageGenerationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingHistoryImageGenerationDetailsWithDefaults

`func NewBillingHistoryImageGenerationDetailsWithDefaults() *BillingHistoryImageGenerationDetails`

NewBillingHistoryImageGenerationDetailsWithDefaults instantiates a new BillingHistoryImageGenerationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingHistory

`func (o *BillingHistoryImageGenerationDetails) GetBillingHistory() []BillingHistory`

GetBillingHistory returns the BillingHistory field if non-nil, zero value otherwise.

### GetBillingHistoryOk

`func (o *BillingHistoryImageGenerationDetails) GetBillingHistoryOk() (*[]BillingHistory, bool)`

GetBillingHistoryOk returns a tuple with the BillingHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingHistory

`func (o *BillingHistoryImageGenerationDetails) SetBillingHistory(v []BillingHistory)`

SetBillingHistory sets BillingHistory field to given value.

### HasBillingHistory

`func (o *BillingHistoryImageGenerationDetails) HasBillingHistory() bool`

HasBillingHistory returns a boolean if a field has been set.

### GetOrgId

`func (o *BillingHistoryImageGenerationDetails) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *BillingHistoryImageGenerationDetails) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *BillingHistoryImageGenerationDetails) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *BillingHistoryImageGenerationDetails) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPagination

`func (o *BillingHistoryImageGenerationDetails) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *BillingHistoryImageGenerationDetails) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *BillingHistoryImageGenerationDetails) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *BillingHistoryImageGenerationDetails) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetTotalCount

`func (o *BillingHistoryImageGenerationDetails) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *BillingHistoryImageGenerationDetails) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *BillingHistoryImageGenerationDetails) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *BillingHistoryImageGenerationDetails) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


