# FirewallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirewallRule** | Pointer to [**SecurityGroupRuleFields**](SecurityGroupRuleFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewFirewallRule

`func NewFirewallRule() *FirewallRule`

NewFirewallRule instantiates a new FirewallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleWithDefaults

`func NewFirewallRuleWithDefaults() *FirewallRule`

NewFirewallRuleWithDefaults instantiates a new FirewallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirewallRule

`func (o *FirewallRule) GetFirewallRule() SecurityGroupRuleFields`

GetFirewallRule returns the FirewallRule field if non-nil, zero value otherwise.

### GetFirewallRuleOk

`func (o *FirewallRule) GetFirewallRuleOk() (*SecurityGroupRuleFields, bool)`

GetFirewallRuleOk returns a tuple with the FirewallRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallRule

`func (o *FirewallRule) SetFirewallRule(v SecurityGroupRuleFields)`

SetFirewallRule sets FirewallRule field to given value.

### HasFirewallRule

`func (o *FirewallRule) HasFirewallRule() bool`

HasFirewallRule returns a boolean if a field has been set.

### GetMessage

`func (o *FirewallRule) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FirewallRule) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FirewallRule) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FirewallRule) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *FirewallRule) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirewallRule) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirewallRule) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FirewallRule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


