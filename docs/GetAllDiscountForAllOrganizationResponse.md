# GetAllDiscountForAllOrganizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscountPlans** | Pointer to [**[]GetAllDiscountsFields**](GetAllDiscountsFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetAllDiscountForAllOrganizationResponse

`func NewGetAllDiscountForAllOrganizationResponse() *GetAllDiscountForAllOrganizationResponse`

NewGetAllDiscountForAllOrganizationResponse instantiates a new GetAllDiscountForAllOrganizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllDiscountForAllOrganizationResponseWithDefaults

`func NewGetAllDiscountForAllOrganizationResponseWithDefaults() *GetAllDiscountForAllOrganizationResponse`

NewGetAllDiscountForAllOrganizationResponseWithDefaults instantiates a new GetAllDiscountForAllOrganizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscountPlans

`func (o *GetAllDiscountForAllOrganizationResponse) GetDiscountPlans() []GetAllDiscountsFields`

GetDiscountPlans returns the DiscountPlans field if non-nil, zero value otherwise.

### GetDiscountPlansOk

`func (o *GetAllDiscountForAllOrganizationResponse) GetDiscountPlansOk() (*[]GetAllDiscountsFields, bool)`

GetDiscountPlansOk returns a tuple with the DiscountPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPlans

`func (o *GetAllDiscountForAllOrganizationResponse) SetDiscountPlans(v []GetAllDiscountsFields)`

SetDiscountPlans sets DiscountPlans field to given value.

### HasDiscountPlans

`func (o *GetAllDiscountForAllOrganizationResponse) HasDiscountPlans() bool`

HasDiscountPlans returns a boolean if a field has been set.

### GetMessage

`func (o *GetAllDiscountForAllOrganizationResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetAllDiscountForAllOrganizationResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetAllDiscountForAllOrganizationResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetAllDiscountForAllOrganizationResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *GetAllDiscountForAllOrganizationResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAllDiscountForAllOrganizationResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAllDiscountForAllOrganizationResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetAllDiscountForAllOrganizationResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


