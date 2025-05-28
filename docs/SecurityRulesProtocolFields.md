# SecurityRulesProtocolFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Protocols** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewSecurityRulesProtocolFields

`func NewSecurityRulesProtocolFields() *SecurityRulesProtocolFields`

NewSecurityRulesProtocolFields instantiates a new SecurityRulesProtocolFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityRulesProtocolFieldsWithDefaults

`func NewSecurityRulesProtocolFieldsWithDefaults() *SecurityRulesProtocolFields`

NewSecurityRulesProtocolFieldsWithDefaults instantiates a new SecurityRulesProtocolFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SecurityRulesProtocolFields) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SecurityRulesProtocolFields) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SecurityRulesProtocolFields) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SecurityRulesProtocolFields) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetProtocols

`func (o *SecurityRulesProtocolFields) GetProtocols() []string`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *SecurityRulesProtocolFields) GetProtocolsOk() (*[]string, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *SecurityRulesProtocolFields) SetProtocols(v []string)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *SecurityRulesProtocolFields) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityRulesProtocolFields) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityRulesProtocolFields) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityRulesProtocolFields) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecurityRulesProtocolFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


