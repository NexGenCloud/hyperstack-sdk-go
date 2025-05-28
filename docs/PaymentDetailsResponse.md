# PaymentDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**PaymentDetailsFields**](PaymentDetailsFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewPaymentDetailsResponse

`func NewPaymentDetailsResponse() *PaymentDetailsResponse`

NewPaymentDetailsResponse instantiates a new PaymentDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentDetailsResponseWithDefaults

`func NewPaymentDetailsResponseWithDefaults() *PaymentDetailsResponse`

NewPaymentDetailsResponseWithDefaults instantiates a new PaymentDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PaymentDetailsResponse) GetData() PaymentDetailsFields`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentDetailsResponse) GetDataOk() (*PaymentDetailsFields, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentDetailsResponse) SetData(v PaymentDetailsFields)`

SetData sets Data field to given value.

### HasData

`func (o *PaymentDetailsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *PaymentDetailsResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PaymentDetailsResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PaymentDetailsResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PaymentDetailsResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentDetailsResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentDetailsResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentDetailsResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentDetailsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


