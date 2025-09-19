# GetDiscountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscountEntites** | Pointer to [**[]DiscountFields**](DiscountFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetDiscountResponse

`func NewGetDiscountResponse() *GetDiscountResponse`

NewGetDiscountResponse instantiates a new GetDiscountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDiscountResponseWithDefaults

`func NewGetDiscountResponseWithDefaults() *GetDiscountResponse`

NewGetDiscountResponseWithDefaults instantiates a new GetDiscountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscountEntites

`func (o *GetDiscountResponse) GetDiscountEntites() []DiscountFields`

GetDiscountEntites returns the DiscountEntites field if non-nil, zero value otherwise.

### GetDiscountEntitesOk

`func (o *GetDiscountResponse) GetDiscountEntitesOk() (*[]DiscountFields, bool)`

GetDiscountEntitesOk returns a tuple with the DiscountEntites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountEntites

`func (o *GetDiscountResponse) SetDiscountEntites(v []DiscountFields)`

SetDiscountEntites sets DiscountEntites field to given value.

### HasDiscountEntites

`func (o *GetDiscountResponse) HasDiscountEntites() bool`

HasDiscountEntites returns a boolean if a field has been set.

### GetMessage

`func (o *GetDiscountResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetDiscountResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetDiscountResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetDiscountResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *GetDiscountResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDiscountResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDiscountResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetDiscountResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


