# SecurityGroupRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**SecurityRule** | Pointer to [**SecurityGroupRuleFields**](SecurityGroupRuleFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewSecurityGroupRule

`func NewSecurityGroupRule() *SecurityGroupRule`

NewSecurityGroupRule instantiates a new SecurityGroupRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupRuleWithDefaults

`func NewSecurityGroupRuleWithDefaults() *SecurityGroupRule`

NewSecurityGroupRuleWithDefaults instantiates a new SecurityGroupRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SecurityGroupRule) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SecurityGroupRule) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SecurityGroupRule) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SecurityGroupRule) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSecurityRule

`func (o *SecurityGroupRule) GetSecurityRule() SecurityGroupRuleFields`

GetSecurityRule returns the SecurityRule field if non-nil, zero value otherwise.

### GetSecurityRuleOk

`func (o *SecurityGroupRule) GetSecurityRuleOk() (*SecurityGroupRuleFields, bool)`

GetSecurityRuleOk returns a tuple with the SecurityRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityRule

`func (o *SecurityGroupRule) SetSecurityRule(v SecurityGroupRuleFields)`

SetSecurityRule sets SecurityRule field to given value.

### HasSecurityRule

`func (o *SecurityGroupRule) HasSecurityRule() bool`

HasSecurityRule returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityGroupRule) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityGroupRule) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityGroupRule) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecurityGroupRule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


