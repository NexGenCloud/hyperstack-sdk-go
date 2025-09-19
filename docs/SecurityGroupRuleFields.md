# SecurityGroupRuleFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Ethertype** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**PortRangeMax** | Pointer to **int32** |  | [optional] 
**PortRangeMin** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**RemoteIpPrefix** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewSecurityGroupRuleFields

`func NewSecurityGroupRuleFields() *SecurityGroupRuleFields`

NewSecurityGroupRuleFields instantiates a new SecurityGroupRuleFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupRuleFieldsWithDefaults

`func NewSecurityGroupRuleFieldsWithDefaults() *SecurityGroupRuleFields`

NewSecurityGroupRuleFieldsWithDefaults instantiates a new SecurityGroupRuleFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SecurityGroupRuleFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecurityGroupRuleFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecurityGroupRuleFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SecurityGroupRuleFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDirection

`func (o *SecurityGroupRuleFields) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SecurityGroupRuleFields) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SecurityGroupRuleFields) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *SecurityGroupRuleFields) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetEthertype

`func (o *SecurityGroupRuleFields) GetEthertype() string`

GetEthertype returns the Ethertype field if non-nil, zero value otherwise.

### GetEthertypeOk

`func (o *SecurityGroupRuleFields) GetEthertypeOk() (*string, bool)`

GetEthertypeOk returns a tuple with the Ethertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthertype

`func (o *SecurityGroupRuleFields) SetEthertype(v string)`

SetEthertype sets Ethertype field to given value.

### HasEthertype

`func (o *SecurityGroupRuleFields) HasEthertype() bool`

HasEthertype returns a boolean if a field has been set.

### GetId

`func (o *SecurityGroupRuleFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityGroupRuleFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityGroupRuleFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityGroupRuleFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPortRangeMax

`func (o *SecurityGroupRuleFields) GetPortRangeMax() int32`

GetPortRangeMax returns the PortRangeMax field if non-nil, zero value otherwise.

### GetPortRangeMaxOk

`func (o *SecurityGroupRuleFields) GetPortRangeMaxOk() (*int32, bool)`

GetPortRangeMaxOk returns a tuple with the PortRangeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMax

`func (o *SecurityGroupRuleFields) SetPortRangeMax(v int32)`

SetPortRangeMax sets PortRangeMax field to given value.

### HasPortRangeMax

`func (o *SecurityGroupRuleFields) HasPortRangeMax() bool`

HasPortRangeMax returns a boolean if a field has been set.

### GetPortRangeMin

`func (o *SecurityGroupRuleFields) GetPortRangeMin() int32`

GetPortRangeMin returns the PortRangeMin field if non-nil, zero value otherwise.

### GetPortRangeMinOk

`func (o *SecurityGroupRuleFields) GetPortRangeMinOk() (*int32, bool)`

GetPortRangeMinOk returns a tuple with the PortRangeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMin

`func (o *SecurityGroupRuleFields) SetPortRangeMin(v int32)`

SetPortRangeMin sets PortRangeMin field to given value.

### HasPortRangeMin

`func (o *SecurityGroupRuleFields) HasPortRangeMin() bool`

HasPortRangeMin returns a boolean if a field has been set.

### GetProtocol

`func (o *SecurityGroupRuleFields) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SecurityGroupRuleFields) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SecurityGroupRuleFields) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SecurityGroupRuleFields) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRemoteIpPrefix

`func (o *SecurityGroupRuleFields) GetRemoteIpPrefix() string`

GetRemoteIpPrefix returns the RemoteIpPrefix field if non-nil, zero value otherwise.

### GetRemoteIpPrefixOk

`func (o *SecurityGroupRuleFields) GetRemoteIpPrefixOk() (*string, bool)`

GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpPrefix

`func (o *SecurityGroupRuleFields) SetRemoteIpPrefix(v string)`

SetRemoteIpPrefix sets RemoteIpPrefix field to given value.

### HasRemoteIpPrefix

`func (o *SecurityGroupRuleFields) HasRemoteIpPrefix() bool`

HasRemoteIpPrefix returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityGroupRuleFields) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityGroupRuleFields) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityGroupRuleFields) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecurityGroupRuleFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


