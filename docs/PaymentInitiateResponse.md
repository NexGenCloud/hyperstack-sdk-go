# PaymentInitiateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**PaymentInitiateFields**](PaymentInitiateFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewPaymentInitiateResponse

`func NewPaymentInitiateResponse() *PaymentInitiateResponse`

NewPaymentInitiateResponse instantiates a new PaymentInitiateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiateResponseWithDefaults

`func NewPaymentInitiateResponseWithDefaults() *PaymentInitiateResponse`

NewPaymentInitiateResponseWithDefaults instantiates a new PaymentInitiateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PaymentInitiateResponse) GetData() PaymentInitiateFields`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentInitiateResponse) GetDataOk() (*PaymentInitiateFields, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentInitiateResponse) SetData(v PaymentInitiateFields)`

SetData sets Data field to given value.

### HasData

`func (o *PaymentInitiateResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *PaymentInitiateResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PaymentInitiateResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PaymentInitiateResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PaymentInitiateResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentInitiateResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentInitiateResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentInitiateResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentInitiateResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


