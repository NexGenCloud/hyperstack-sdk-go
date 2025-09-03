# BillingMetricesFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**BillPerMinute** | Pointer to **float32** |  | [optional] 
**CreateTime** | Pointer to **time.Time** |  | [optional] 
**ExcludeBilling** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**ResourceId** | Pointer to **int32** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**TerminateTime** | Pointer to **time.Time** |  | [optional] 
**TotalBill** | Pointer to **float32** |  | [optional] 
**TotalUpTime** | Pointer to **float32** |  | [optional] 

## Methods

### NewBillingMetricesFields

`func NewBillingMetricesFields() *BillingMetricesFields`

NewBillingMetricesFields instantiates a new BillingMetricesFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingMetricesFieldsWithDefaults

`func NewBillingMetricesFieldsWithDefaults() *BillingMetricesFields`

NewBillingMetricesFieldsWithDefaults instantiates a new BillingMetricesFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *BillingMetricesFields) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BillingMetricesFields) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BillingMetricesFields) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *BillingMetricesFields) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBillPerMinute

`func (o *BillingMetricesFields) GetBillPerMinute() float32`

GetBillPerMinute returns the BillPerMinute field if non-nil, zero value otherwise.

### GetBillPerMinuteOk

`func (o *BillingMetricesFields) GetBillPerMinuteOk() (*float32, bool)`

GetBillPerMinuteOk returns a tuple with the BillPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillPerMinute

`func (o *BillingMetricesFields) SetBillPerMinute(v float32)`

SetBillPerMinute sets BillPerMinute field to given value.

### HasBillPerMinute

`func (o *BillingMetricesFields) HasBillPerMinute() bool`

HasBillPerMinute returns a boolean if a field has been set.

### GetCreateTime

`func (o *BillingMetricesFields) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *BillingMetricesFields) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *BillingMetricesFields) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *BillingMetricesFields) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetExcludeBilling

`func (o *BillingMetricesFields) GetExcludeBilling() bool`

GetExcludeBilling returns the ExcludeBilling field if non-nil, zero value otherwise.

### GetExcludeBillingOk

`func (o *BillingMetricesFields) GetExcludeBillingOk() (*bool, bool)`

GetExcludeBillingOk returns a tuple with the ExcludeBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeBilling

`func (o *BillingMetricesFields) SetExcludeBilling(v bool)`

SetExcludeBilling sets ExcludeBilling field to given value.

### HasExcludeBilling

`func (o *BillingMetricesFields) HasExcludeBilling() bool`

HasExcludeBilling returns a boolean if a field has been set.

### GetName

`func (o *BillingMetricesFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingMetricesFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingMetricesFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingMetricesFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *BillingMetricesFields) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *BillingMetricesFields) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *BillingMetricesFields) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *BillingMetricesFields) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetResourceId

`func (o *BillingMetricesFields) GetResourceId() int32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *BillingMetricesFields) GetResourceIdOk() (*int32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *BillingMetricesFields) SetResourceId(v int32)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *BillingMetricesFields) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *BillingMetricesFields) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BillingMetricesFields) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BillingMetricesFields) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BillingMetricesFields) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetTerminateTime

`func (o *BillingMetricesFields) GetTerminateTime() time.Time`

GetTerminateTime returns the TerminateTime field if non-nil, zero value otherwise.

### GetTerminateTimeOk

`func (o *BillingMetricesFields) GetTerminateTimeOk() (*time.Time, bool)`

GetTerminateTimeOk returns a tuple with the TerminateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminateTime

`func (o *BillingMetricesFields) SetTerminateTime(v time.Time)`

SetTerminateTime sets TerminateTime field to given value.

### HasTerminateTime

`func (o *BillingMetricesFields) HasTerminateTime() bool`

HasTerminateTime returns a boolean if a field has been set.

### GetTotalBill

`func (o *BillingMetricesFields) GetTotalBill() float32`

GetTotalBill returns the TotalBill field if non-nil, zero value otherwise.

### GetTotalBillOk

`func (o *BillingMetricesFields) GetTotalBillOk() (*float32, bool)`

GetTotalBillOk returns a tuple with the TotalBill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBill

`func (o *BillingMetricesFields) SetTotalBill(v float32)`

SetTotalBill sets TotalBill field to given value.

### HasTotalBill

`func (o *BillingMetricesFields) HasTotalBill() bool`

HasTotalBill returns a boolean if a field has been set.

### GetTotalUpTime

`func (o *BillingMetricesFields) GetTotalUpTime() float32`

GetTotalUpTime returns the TotalUpTime field if non-nil, zero value otherwise.

### GetTotalUpTimeOk

`func (o *BillingMetricesFields) GetTotalUpTimeOk() (*float32, bool)`

GetTotalUpTimeOk returns a tuple with the TotalUpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUpTime

`func (o *BillingMetricesFields) SetTotalUpTime(v float32)`

SetTotalUpTime sets TotalUpTime field to given value.

### HasTotalUpTime

`func (o *BillingMetricesFields) HasTotalUpTime() bool`

HasTotalUpTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


