# PaymentDetailsFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**GatewayResponse** | Pointer to **string** |  | [optional] 
**Invoice** | Pointer to **string** |  | [optional] 
**PaidFrom** | Pointer to **string** |  | [optional] 
**PaymentId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TaxAmount** | Pointer to **float32** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPaymentDetailsFields

`func NewPaymentDetailsFields() *PaymentDetailsFields`

NewPaymentDetailsFields instantiates a new PaymentDetailsFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentDetailsFieldsWithDefaults

`func NewPaymentDetailsFieldsWithDefaults() *PaymentDetailsFields`

NewPaymentDetailsFieldsWithDefaults instantiates a new PaymentDetailsFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PaymentDetailsFields) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentDetailsFields) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentDetailsFields) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentDetailsFields) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaymentDetailsFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentDetailsFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentDetailsFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaymentDetailsFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentDetailsFields) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentDetailsFields) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentDetailsFields) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentDetailsFields) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *PaymentDetailsFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentDetailsFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentDetailsFields) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentDetailsFields) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGatewayResponse

`func (o *PaymentDetailsFields) GetGatewayResponse() string`

GetGatewayResponse returns the GatewayResponse field if non-nil, zero value otherwise.

### GetGatewayResponseOk

`func (o *PaymentDetailsFields) GetGatewayResponseOk() (*string, bool)`

GetGatewayResponseOk returns a tuple with the GatewayResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayResponse

`func (o *PaymentDetailsFields) SetGatewayResponse(v string)`

SetGatewayResponse sets GatewayResponse field to given value.

### HasGatewayResponse

`func (o *PaymentDetailsFields) HasGatewayResponse() bool`

HasGatewayResponse returns a boolean if a field has been set.

### GetInvoice

`func (o *PaymentDetailsFields) GetInvoice() string`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *PaymentDetailsFields) GetInvoiceOk() (*string, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *PaymentDetailsFields) SetInvoice(v string)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *PaymentDetailsFields) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetPaidFrom

`func (o *PaymentDetailsFields) GetPaidFrom() string`

GetPaidFrom returns the PaidFrom field if non-nil, zero value otherwise.

### GetPaidFromOk

`func (o *PaymentDetailsFields) GetPaidFromOk() (*string, bool)`

GetPaidFromOk returns a tuple with the PaidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidFrom

`func (o *PaymentDetailsFields) SetPaidFrom(v string)`

SetPaidFrom sets PaidFrom field to given value.

### HasPaidFrom

`func (o *PaymentDetailsFields) HasPaidFrom() bool`

HasPaidFrom returns a boolean if a field has been set.

### GetPaymentId

`func (o *PaymentDetailsFields) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentDetailsFields) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentDetailsFields) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *PaymentDetailsFields) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentDetailsFields) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentDetailsFields) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentDetailsFields) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentDetailsFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaxAmount

`func (o *PaymentDetailsFields) GetTaxAmount() float32`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *PaymentDetailsFields) GetTaxAmountOk() (*float32, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *PaymentDetailsFields) SetTaxAmount(v float32)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *PaymentDetailsFields) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTransactionId

`func (o *PaymentDetailsFields) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *PaymentDetailsFields) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *PaymentDetailsFields) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *PaymentDetailsFields) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PaymentDetailsFields) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaymentDetailsFields) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaymentDetailsFields) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PaymentDetailsFields) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


