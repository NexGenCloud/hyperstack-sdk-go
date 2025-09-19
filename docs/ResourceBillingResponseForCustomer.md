# ResourceBillingResponseForCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalculatedResourceBills** | Pointer to [**ResourceObjectResponseForCustomer**](ResourceObjectResponseForCustomer.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewResourceBillingResponseForCustomer

`func NewResourceBillingResponseForCustomer() *ResourceBillingResponseForCustomer`

NewResourceBillingResponseForCustomer instantiates a new ResourceBillingResponseForCustomer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceBillingResponseForCustomerWithDefaults

`func NewResourceBillingResponseForCustomerWithDefaults() *ResourceBillingResponseForCustomer`

NewResourceBillingResponseForCustomerWithDefaults instantiates a new ResourceBillingResponseForCustomer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalculatedResourceBills

`func (o *ResourceBillingResponseForCustomer) GetCalculatedResourceBills() ResourceObjectResponseForCustomer`

GetCalculatedResourceBills returns the CalculatedResourceBills field if non-nil, zero value otherwise.

### GetCalculatedResourceBillsOk

`func (o *ResourceBillingResponseForCustomer) GetCalculatedResourceBillsOk() (*ResourceObjectResponseForCustomer, bool)`

GetCalculatedResourceBillsOk returns a tuple with the CalculatedResourceBills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedResourceBills

`func (o *ResourceBillingResponseForCustomer) SetCalculatedResourceBills(v ResourceObjectResponseForCustomer)`

SetCalculatedResourceBills sets CalculatedResourceBills field to given value.

### HasCalculatedResourceBills

`func (o *ResourceBillingResponseForCustomer) HasCalculatedResourceBills() bool`

HasCalculatedResourceBills returns a boolean if a field has been set.

### GetMessage

`func (o *ResourceBillingResponseForCustomer) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResourceBillingResponseForCustomer) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResourceBillingResponseForCustomer) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResourceBillingResponseForCustomer) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ResourceBillingResponseForCustomer) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceBillingResponseForCustomer) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceBillingResponseForCustomer) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResourceBillingResponseForCustomer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


