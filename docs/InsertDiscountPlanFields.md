# InsertDiscountPlanFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customers** | Pointer to [**[]CustomerFields**](CustomerFields.md) |  | [optional] 
**DiscountResources** | Pointer to [**[]DiscountResourceFields**](DiscountResourceFields.md) |  | [optional] 
**DiscountStatus** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInsertDiscountPlanFields

`func NewInsertDiscountPlanFields() *InsertDiscountPlanFields`

NewInsertDiscountPlanFields instantiates a new InsertDiscountPlanFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsertDiscountPlanFieldsWithDefaults

`func NewInsertDiscountPlanFieldsWithDefaults() *InsertDiscountPlanFields`

NewInsertDiscountPlanFieldsWithDefaults instantiates a new InsertDiscountPlanFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomers

`func (o *InsertDiscountPlanFields) GetCustomers() []CustomerFields`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *InsertDiscountPlanFields) GetCustomersOk() (*[]CustomerFields, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *InsertDiscountPlanFields) SetCustomers(v []CustomerFields)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *InsertDiscountPlanFields) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.

### GetDiscountResources

`func (o *InsertDiscountPlanFields) GetDiscountResources() []DiscountResourceFields`

GetDiscountResources returns the DiscountResources field if non-nil, zero value otherwise.

### GetDiscountResourcesOk

`func (o *InsertDiscountPlanFields) GetDiscountResourcesOk() (*[]DiscountResourceFields, bool)`

GetDiscountResourcesOk returns a tuple with the DiscountResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountResources

`func (o *InsertDiscountPlanFields) SetDiscountResources(v []DiscountResourceFields)`

SetDiscountResources sets DiscountResources field to given value.

### HasDiscountResources

`func (o *InsertDiscountPlanFields) HasDiscountResources() bool`

HasDiscountResources returns a boolean if a field has been set.

### GetDiscountStatus

`func (o *InsertDiscountPlanFields) GetDiscountStatus() string`

GetDiscountStatus returns the DiscountStatus field if non-nil, zero value otherwise.

### GetDiscountStatusOk

`func (o *InsertDiscountPlanFields) GetDiscountStatusOk() (*string, bool)`

GetDiscountStatusOk returns a tuple with the DiscountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountStatus

`func (o *InsertDiscountPlanFields) SetDiscountStatus(v string)`

SetDiscountStatus sets DiscountStatus field to given value.

### HasDiscountStatus

`func (o *InsertDiscountPlanFields) HasDiscountStatus() bool`

HasDiscountStatus returns a boolean if a field has been set.

### GetEndDate

`func (o *InsertDiscountPlanFields) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InsertDiscountPlanFields) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InsertDiscountPlanFields) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InsertDiscountPlanFields) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDate

`func (o *InsertDiscountPlanFields) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InsertDiscountPlanFields) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InsertDiscountPlanFields) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InsertDiscountPlanFields) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


