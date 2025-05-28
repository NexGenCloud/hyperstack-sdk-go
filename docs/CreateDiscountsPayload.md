# CreateDiscountsPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customers** | [**[]CustomerPayload**](CustomerPayload.md) |  | 
**DiscountResources** | [**[]DiscountResourcePayload**](DiscountResourcePayload.md) |  | 
**DiscountStatus** | **string** |  | 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreateDiscountsPayload

`func NewCreateDiscountsPayload(customers []CustomerPayload, discountResources []DiscountResourcePayload, discountStatus string, ) *CreateDiscountsPayload`

NewCreateDiscountsPayload instantiates a new CreateDiscountsPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDiscountsPayloadWithDefaults

`func NewCreateDiscountsPayloadWithDefaults() *CreateDiscountsPayload`

NewCreateDiscountsPayloadWithDefaults instantiates a new CreateDiscountsPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomers

`func (o *CreateDiscountsPayload) GetCustomers() []CustomerPayload`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *CreateDiscountsPayload) GetCustomersOk() (*[]CustomerPayload, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *CreateDiscountsPayload) SetCustomers(v []CustomerPayload)`

SetCustomers sets Customers field to given value.


### GetDiscountResources

`func (o *CreateDiscountsPayload) GetDiscountResources() []DiscountResourcePayload`

GetDiscountResources returns the DiscountResources field if non-nil, zero value otherwise.

### GetDiscountResourcesOk

`func (o *CreateDiscountsPayload) GetDiscountResourcesOk() (*[]DiscountResourcePayload, bool)`

GetDiscountResourcesOk returns a tuple with the DiscountResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountResources

`func (o *CreateDiscountsPayload) SetDiscountResources(v []DiscountResourcePayload)`

SetDiscountResources sets DiscountResources field to given value.


### GetDiscountStatus

`func (o *CreateDiscountsPayload) GetDiscountStatus() string`

GetDiscountStatus returns the DiscountStatus field if non-nil, zero value otherwise.

### GetDiscountStatusOk

`func (o *CreateDiscountsPayload) GetDiscountStatusOk() (*string, bool)`

GetDiscountStatusOk returns a tuple with the DiscountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountStatus

`func (o *CreateDiscountsPayload) SetDiscountStatus(v string)`

SetDiscountStatus sets DiscountStatus field to given value.


### GetEndDate

`func (o *CreateDiscountsPayload) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CreateDiscountsPayload) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CreateDiscountsPayload) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CreateDiscountsPayload) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDate

`func (o *CreateDiscountsPayload) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateDiscountsPayload) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateDiscountsPayload) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreateDiscountsPayload) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


