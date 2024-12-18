# UpdateDiscountsPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscountResources** | [**[]DiscountResourcePayload**](DiscountResourcePayload.md) |  | 
**DiscountStatus** | **string** |  | 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateDiscountsPayload

`func NewUpdateDiscountsPayload(discountResources []DiscountResourcePayload, discountStatus string, ) *UpdateDiscountsPayload`

NewUpdateDiscountsPayload instantiates a new UpdateDiscountsPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDiscountsPayloadWithDefaults

`func NewUpdateDiscountsPayloadWithDefaults() *UpdateDiscountsPayload`

NewUpdateDiscountsPayloadWithDefaults instantiates a new UpdateDiscountsPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscountResources

`func (o *UpdateDiscountsPayload) GetDiscountResources() []DiscountResourcePayload`

GetDiscountResources returns the DiscountResources field if non-nil, zero value otherwise.

### GetDiscountResourcesOk

`func (o *UpdateDiscountsPayload) GetDiscountResourcesOk() (*[]DiscountResourcePayload, bool)`

GetDiscountResourcesOk returns a tuple with the DiscountResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountResources

`func (o *UpdateDiscountsPayload) SetDiscountResources(v []DiscountResourcePayload)`

SetDiscountResources sets DiscountResources field to given value.


### GetDiscountStatus

`func (o *UpdateDiscountsPayload) GetDiscountStatus() string`

GetDiscountStatus returns the DiscountStatus field if non-nil, zero value otherwise.

### GetDiscountStatusOk

`func (o *UpdateDiscountsPayload) GetDiscountStatusOk() (*string, bool)`

GetDiscountStatusOk returns a tuple with the DiscountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountStatus

`func (o *UpdateDiscountsPayload) SetDiscountStatus(v string)`

SetDiscountStatus sets DiscountStatus field to given value.


### GetEndDate

`func (o *UpdateDiscountsPayload) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UpdateDiscountsPayload) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UpdateDiscountsPayload) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UpdateDiscountsPayload) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDate

`func (o *UpdateDiscountsPayload) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdateDiscountsPayload) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UpdateDiscountsPayload) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UpdateDiscountsPayload) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


