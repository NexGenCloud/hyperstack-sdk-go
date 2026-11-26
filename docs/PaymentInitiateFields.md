# PaymentInitiateFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoTopupLinked** | Pointer to **bool** | Whether the purchase is linked to a pending auto top up. If true, the checkout saves the card for off-session use (setup_future_usage&#x3D;off_session). | [optional] 
**CheckoutUrl** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**PaymentId** | Pointer to **string** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentInitiateFields

`func NewPaymentInitiateFields() *PaymentInitiateFields`

NewPaymentInitiateFields instantiates a new PaymentInitiateFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiateFieldsWithDefaults

`func NewPaymentInitiateFieldsWithDefaults() *PaymentInitiateFields`

NewPaymentInitiateFieldsWithDefaults instantiates a new PaymentInitiateFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoTopupLinked

`func (o *PaymentInitiateFields) GetAutoTopupLinked() bool`

GetAutoTopupLinked returns the AutoTopupLinked field if non-nil, zero value otherwise.

### GetAutoTopupLinkedOk

`func (o *PaymentInitiateFields) GetAutoTopupLinkedOk() (*bool, bool)`

GetAutoTopupLinkedOk returns a tuple with the AutoTopupLinked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopupLinked

`func (o *PaymentInitiateFields) SetAutoTopupLinked(v bool)`

SetAutoTopupLinked sets AutoTopupLinked field to given value.

### HasAutoTopupLinked

`func (o *PaymentInitiateFields) HasAutoTopupLinked() bool`

HasAutoTopupLinked returns a boolean if a field has been set.

### GetCheckoutUrl

`func (o *PaymentInitiateFields) GetCheckoutUrl() string`

GetCheckoutUrl returns the CheckoutUrl field if non-nil, zero value otherwise.

### GetCheckoutUrlOk

`func (o *PaymentInitiateFields) GetCheckoutUrlOk() (*string, bool)`

GetCheckoutUrlOk returns a tuple with the CheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUrl

`func (o *PaymentInitiateFields) SetCheckoutUrl(v string)`

SetCheckoutUrl sets CheckoutUrl field to given value.

### HasCheckoutUrl

`func (o *PaymentInitiateFields) HasCheckoutUrl() bool`

HasCheckoutUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PaymentInitiateFields) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PaymentInitiateFields) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PaymentInitiateFields) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PaymentInitiateFields) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetPaymentId

`func (o *PaymentInitiateFields) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentInitiateFields) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentInitiateFields) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *PaymentInitiateFields) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetSessionId

`func (o *PaymentInitiateFields) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *PaymentInitiateFields) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *PaymentInitiateFields) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *PaymentInitiateFields) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


