# Billingmetricesfields

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

### NewBillingmetricesfields

`func NewBillingmetricesfields() *Billingmetricesfields`

NewBillingmetricesfields instantiates a new Billingmetricesfields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingmetricesfieldsWithDefaults

`func NewBillingmetricesfieldsWithDefaults() *Billingmetricesfields`

NewBillingmetricesfieldsWithDefaults instantiates a new Billingmetricesfields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *Billingmetricesfields) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Billingmetricesfields) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Billingmetricesfields) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Billingmetricesfields) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBillPerMinute

`func (o *Billingmetricesfields) GetBillPerMinute() float32`

GetBillPerMinute returns the BillPerMinute field if non-nil, zero value otherwise.

### GetBillPerMinuteOk

`func (o *Billingmetricesfields) GetBillPerMinuteOk() (*float32, bool)`

GetBillPerMinuteOk returns a tuple with the BillPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillPerMinute

`func (o *Billingmetricesfields) SetBillPerMinute(v float32)`

SetBillPerMinute sets BillPerMinute field to given value.

### HasBillPerMinute

`func (o *Billingmetricesfields) HasBillPerMinute() bool`

HasBillPerMinute returns a boolean if a field has been set.

### GetCreateTime

`func (o *Billingmetricesfields) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Billingmetricesfields) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Billingmetricesfields) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *Billingmetricesfields) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetExcludeBilling

`func (o *Billingmetricesfields) GetExcludeBilling() bool`

GetExcludeBilling returns the ExcludeBilling field if non-nil, zero value otherwise.

### GetExcludeBillingOk

`func (o *Billingmetricesfields) GetExcludeBillingOk() (*bool, bool)`

GetExcludeBillingOk returns a tuple with the ExcludeBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeBilling

`func (o *Billingmetricesfields) SetExcludeBilling(v bool)`

SetExcludeBilling sets ExcludeBilling field to given value.

### HasExcludeBilling

`func (o *Billingmetricesfields) HasExcludeBilling() bool`

HasExcludeBilling returns a boolean if a field has been set.

### GetName

`func (o *Billingmetricesfields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Billingmetricesfields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Billingmetricesfields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Billingmetricesfields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Billingmetricesfields) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Billingmetricesfields) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Billingmetricesfields) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Billingmetricesfields) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetResourceId

`func (o *Billingmetricesfields) GetResourceId() int32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Billingmetricesfields) GetResourceIdOk() (*int32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Billingmetricesfields) SetResourceId(v int32)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *Billingmetricesfields) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *Billingmetricesfields) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *Billingmetricesfields) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *Billingmetricesfields) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *Billingmetricesfields) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetTerminateTime

`func (o *Billingmetricesfields) GetTerminateTime() time.Time`

GetTerminateTime returns the TerminateTime field if non-nil, zero value otherwise.

### GetTerminateTimeOk

`func (o *Billingmetricesfields) GetTerminateTimeOk() (*time.Time, bool)`

GetTerminateTimeOk returns a tuple with the TerminateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminateTime

`func (o *Billingmetricesfields) SetTerminateTime(v time.Time)`

SetTerminateTime sets TerminateTime field to given value.

### HasTerminateTime

`func (o *Billingmetricesfields) HasTerminateTime() bool`

HasTerminateTime returns a boolean if a field has been set.

### GetTotalBill

`func (o *Billingmetricesfields) GetTotalBill() float32`

GetTotalBill returns the TotalBill field if non-nil, zero value otherwise.

### GetTotalBillOk

`func (o *Billingmetricesfields) GetTotalBillOk() (*float32, bool)`

GetTotalBillOk returns a tuple with the TotalBill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBill

`func (o *Billingmetricesfields) SetTotalBill(v float32)`

SetTotalBill sets TotalBill field to given value.

### HasTotalBill

`func (o *Billingmetricesfields) HasTotalBill() bool`

HasTotalBill returns a boolean if a field has been set.

### GetTotalUpTime

`func (o *Billingmetricesfields) GetTotalUpTime() float32`

GetTotalUpTime returns the TotalUpTime field if non-nil, zero value otherwise.

### GetTotalUpTimeOk

`func (o *Billingmetricesfields) GetTotalUpTimeOk() (*float32, bool)`

GetTotalUpTimeOk returns a tuple with the TotalUpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUpTime

`func (o *Billingmetricesfields) SetTotalUpTime(v float32)`

SetTotalUpTime sets TotalUpTime field to given value.

### HasTotalUpTime

`func (o *Billingmetricesfields) HasTotalUpTime() bool`

HasTotalUpTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


