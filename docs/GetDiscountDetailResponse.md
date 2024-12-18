# GetDiscountDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscountsEntity** | Pointer to [**DiscountEntityModel**](DiscountEntityModel.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetDiscountDetailResponse

`func NewGetDiscountDetailResponse() *GetDiscountDetailResponse`

NewGetDiscountDetailResponse instantiates a new GetDiscountDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDiscountDetailResponseWithDefaults

`func NewGetDiscountDetailResponseWithDefaults() *GetDiscountDetailResponse`

NewGetDiscountDetailResponseWithDefaults instantiates a new GetDiscountDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscountsEntity

`func (o *GetDiscountDetailResponse) GetDiscountsEntity() DiscountEntityModel`

GetDiscountsEntity returns the DiscountsEntity field if non-nil, zero value otherwise.

### GetDiscountsEntityOk

`func (o *GetDiscountDetailResponse) GetDiscountsEntityOk() (*DiscountEntityModel, bool)`

GetDiscountsEntityOk returns a tuple with the DiscountsEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountsEntity

`func (o *GetDiscountDetailResponse) SetDiscountsEntity(v DiscountEntityModel)`

SetDiscountsEntity sets DiscountsEntity field to given value.

### HasDiscountsEntity

`func (o *GetDiscountDetailResponse) HasDiscountsEntity() bool`

HasDiscountsEntity returns a boolean if a field has been set.

### GetMessage

`func (o *GetDiscountDetailResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetDiscountDetailResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetDiscountDetailResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetDiscountDetailResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *GetDiscountDetailResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDiscountDetailResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDiscountDetailResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetDiscountDetailResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


