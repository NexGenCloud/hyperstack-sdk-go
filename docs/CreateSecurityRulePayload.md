# CreateSecurityRulePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | **string** | The direction of traffic that the firewall rule applies to. | 
**Ethertype** | **string** | The Ethernet type associated with the rule. | 
**Protocol** | **string** | The network protocol associated with the rule. Call the [&#x60;GET /core/sg-rules-protocols&#x60;](https://infrahub-api-doc.nexgencloud.com/#get-/core/sg-rules-protocols) endpoint to retrieve a list of permitted network protocols. | 
**RemoteIpPrefix** | **string** | The IP address range that is allowed to access the specified port. Use \&quot;0.0.0.0/0\&quot; to allow any IP address. | 
**PortRangeMin** | Pointer to **int32** |  | [optional] 
**PortRangeMax** | Pointer to **int32** |  | [optional] 

## Methods

### NewCreateSecurityRulePayload

`func NewCreateSecurityRulePayload(direction string, ethertype string, protocol string, remoteIpPrefix string, ) *CreateSecurityRulePayload`

NewCreateSecurityRulePayload instantiates a new CreateSecurityRulePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecurityRulePayloadWithDefaults

`func NewCreateSecurityRulePayloadWithDefaults() *CreateSecurityRulePayload`

NewCreateSecurityRulePayloadWithDefaults instantiates a new CreateSecurityRulePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *CreateSecurityRulePayload) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CreateSecurityRulePayload) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CreateSecurityRulePayload) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetEthertype

`func (o *CreateSecurityRulePayload) GetEthertype() string`

GetEthertype returns the Ethertype field if non-nil, zero value otherwise.

### GetEthertypeOk

`func (o *CreateSecurityRulePayload) GetEthertypeOk() (*string, bool)`

GetEthertypeOk returns a tuple with the Ethertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthertype

`func (o *CreateSecurityRulePayload) SetEthertype(v string)`

SetEthertype sets Ethertype field to given value.


### GetProtocol

`func (o *CreateSecurityRulePayload) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreateSecurityRulePayload) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreateSecurityRulePayload) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetRemoteIpPrefix

`func (o *CreateSecurityRulePayload) GetRemoteIpPrefix() string`

GetRemoteIpPrefix returns the RemoteIpPrefix field if non-nil, zero value otherwise.

### GetRemoteIpPrefixOk

`func (o *CreateSecurityRulePayload) GetRemoteIpPrefixOk() (*string, bool)`

GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpPrefix

`func (o *CreateSecurityRulePayload) SetRemoteIpPrefix(v string)`

SetRemoteIpPrefix sets RemoteIpPrefix field to given value.


### GetPortRangeMin

`func (o *CreateSecurityRulePayload) GetPortRangeMin() int32`

GetPortRangeMin returns the PortRangeMin field if non-nil, zero value otherwise.

### GetPortRangeMinOk

`func (o *CreateSecurityRulePayload) GetPortRangeMinOk() (*int32, bool)`

GetPortRangeMinOk returns a tuple with the PortRangeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMin

`func (o *CreateSecurityRulePayload) SetPortRangeMin(v int32)`

SetPortRangeMin sets PortRangeMin field to given value.

### HasPortRangeMin

`func (o *CreateSecurityRulePayload) HasPortRangeMin() bool`

HasPortRangeMin returns a boolean if a field has been set.

### GetPortRangeMax

`func (o *CreateSecurityRulePayload) GetPortRangeMax() int32`

GetPortRangeMax returns the PortRangeMax field if non-nil, zero value otherwise.

### GetPortRangeMaxOk

`func (o *CreateSecurityRulePayload) GetPortRangeMaxOk() (*int32, bool)`

GetPortRangeMaxOk returns a tuple with the PortRangeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMax

`func (o *CreateSecurityRulePayload) SetPortRangeMax(v int32)`

SetPortRangeMax sets PortRangeMax field to given value.

### HasPortRangeMax

`func (o *CreateSecurityRulePayload) HasPortRangeMax() bool`

HasPortRangeMax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


