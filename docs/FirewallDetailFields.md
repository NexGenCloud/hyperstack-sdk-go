# FirewallDetailFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to [**[]FirewallAttachmentModel**](FirewallAttachmentModel.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to [**FirewallEnvironmentFields**](FirewallEnvironmentFields.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]SecurityGroupRuleFields**](SecurityGroupRuleFields.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewFirewallDetailFields

`func NewFirewallDetailFields() *FirewallDetailFields`

NewFirewallDetailFields instantiates a new FirewallDetailFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallDetailFieldsWithDefaults

`func NewFirewallDetailFieldsWithDefaults() *FirewallDetailFields`

NewFirewallDetailFieldsWithDefaults instantiates a new FirewallDetailFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *FirewallDetailFields) GetAttachments() []FirewallAttachmentModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *FirewallDetailFields) GetAttachmentsOk() (*[]FirewallAttachmentModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *FirewallDetailFields) SetAttachments(v []FirewallAttachmentModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *FirewallDetailFields) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FirewallDetailFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FirewallDetailFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FirewallDetailFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FirewallDetailFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *FirewallDetailFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirewallDetailFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirewallDetailFields) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FirewallDetailFields) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *FirewallDetailFields) GetEnvironment() FirewallEnvironmentFields`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FirewallDetailFields) GetEnvironmentOk() (*FirewallEnvironmentFields, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FirewallDetailFields) SetEnvironment(v FirewallEnvironmentFields)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FirewallDetailFields) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *FirewallDetailFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallDetailFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallDetailFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FirewallDetailFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FirewallDetailFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallDetailFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallDetailFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirewallDetailFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRules

`func (o *FirewallDetailFields) GetRules() []SecurityGroupRuleFields`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FirewallDetailFields) GetRulesOk() (*[]SecurityGroupRuleFields, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FirewallDetailFields) SetRules(v []SecurityGroupRuleFields)`

SetRules sets Rules field to given value.

### HasRules

`func (o *FirewallDetailFields) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetStatus

`func (o *FirewallDetailFields) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirewallDetailFields) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirewallDetailFields) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FirewallDetailFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


