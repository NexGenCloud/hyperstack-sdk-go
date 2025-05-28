# CreateDiscountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscountPlan** | Pointer to [**InsertDiscountPlanFields**](InsertDiscountPlanFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateDiscountResponse

`func NewCreateDiscountResponse() *CreateDiscountResponse`

NewCreateDiscountResponse instantiates a new CreateDiscountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDiscountResponseWithDefaults

`func NewCreateDiscountResponseWithDefaults() *CreateDiscountResponse`

NewCreateDiscountResponseWithDefaults instantiates a new CreateDiscountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscountPlan

`func (o *CreateDiscountResponse) GetDiscountPlan() InsertDiscountPlanFields`

GetDiscountPlan returns the DiscountPlan field if non-nil, zero value otherwise.

### GetDiscountPlanOk

`func (o *CreateDiscountResponse) GetDiscountPlanOk() (*InsertDiscountPlanFields, bool)`

GetDiscountPlanOk returns a tuple with the DiscountPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPlan

`func (o *CreateDiscountResponse) SetDiscountPlan(v InsertDiscountPlanFields)`

SetDiscountPlan sets DiscountPlan field to given value.

### HasDiscountPlan

`func (o *CreateDiscountResponse) HasDiscountPlan() bool`

HasDiscountPlan returns a boolean if a field has been set.

### GetMessage

`func (o *CreateDiscountResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateDiscountResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateDiscountResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateDiscountResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *CreateDiscountResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateDiscountResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateDiscountResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateDiscountResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


