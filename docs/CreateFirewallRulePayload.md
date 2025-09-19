# CreateFirewallRulePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | **string** | The direction of traffic that the firewall rule applies to. | 
**Ethertype** | **string** | The Ethernet type associated with the rule. | 
**PortRangeMax** | Pointer to **int32** | The maximum port number in the range of ports to be allowed by the firewall rule. | [optional] 
**PortRangeMin** | Pointer to **int32** | The minimum port number in the range of ports to be allowed by the firewall rule. | [optional] 
**Protocol** | **string** | The network protocol associated with the rule. Call the [&#x60;GET /core/sg-rules-protocols&#x60;](https://infrahub-api-doc.nexgencloud.com/#get-/core/sg-rules-protocols) endpoint to retrieve a list of permitted network protocols. | 
**RemoteIpPrefix** | **string** | The IP address range that is allowed to access the specified port. Use \&quot;0.0.0.0/0\&quot; to allow any IP address. | 

## Methods

### NewCreateFirewallRulePayload

`func NewCreateFirewallRulePayload(direction string, ethertype string, protocol string, remoteIpPrefix string, ) *CreateFirewallRulePayload`

NewCreateFirewallRulePayload instantiates a new CreateFirewallRulePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirewallRulePayloadWithDefaults

`func NewCreateFirewallRulePayloadWithDefaults() *CreateFirewallRulePayload`

NewCreateFirewallRulePayloadWithDefaults instantiates a new CreateFirewallRulePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *CreateFirewallRulePayload) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CreateFirewallRulePayload) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CreateFirewallRulePayload) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetEthertype

`func (o *CreateFirewallRulePayload) GetEthertype() string`

GetEthertype returns the Ethertype field if non-nil, zero value otherwise.

### GetEthertypeOk

`func (o *CreateFirewallRulePayload) GetEthertypeOk() (*string, bool)`

GetEthertypeOk returns a tuple with the Ethertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthertype

`func (o *CreateFirewallRulePayload) SetEthertype(v string)`

SetEthertype sets Ethertype field to given value.


### GetPortRangeMax

`func (o *CreateFirewallRulePayload) GetPortRangeMax() int32`

GetPortRangeMax returns the PortRangeMax field if non-nil, zero value otherwise.

### GetPortRangeMaxOk

`func (o *CreateFirewallRulePayload) GetPortRangeMaxOk() (*int32, bool)`

GetPortRangeMaxOk returns a tuple with the PortRangeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMax

`func (o *CreateFirewallRulePayload) SetPortRangeMax(v int32)`

SetPortRangeMax sets PortRangeMax field to given value.

### HasPortRangeMax

`func (o *CreateFirewallRulePayload) HasPortRangeMax() bool`

HasPortRangeMax returns a boolean if a field has been set.

### GetPortRangeMin

`func (o *CreateFirewallRulePayload) GetPortRangeMin() int32`

GetPortRangeMin returns the PortRangeMin field if non-nil, zero value otherwise.

### GetPortRangeMinOk

`func (o *CreateFirewallRulePayload) GetPortRangeMinOk() (*int32, bool)`

GetPortRangeMinOk returns a tuple with the PortRangeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMin

`func (o *CreateFirewallRulePayload) SetPortRangeMin(v int32)`

SetPortRangeMin sets PortRangeMin field to given value.

### HasPortRangeMin

`func (o *CreateFirewallRulePayload) HasPortRangeMin() bool`

HasPortRangeMin returns a boolean if a field has been set.

### GetProtocol

`func (o *CreateFirewallRulePayload) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreateFirewallRulePayload) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreateFirewallRulePayload) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetRemoteIpPrefix

`func (o *CreateFirewallRulePayload) GetRemoteIpPrefix() string`

GetRemoteIpPrefix returns the RemoteIpPrefix field if non-nil, zero value otherwise.

### GetRemoteIpPrefixOk

`func (o *CreateFirewallRulePayload) GetRemoteIpPrefixOk() (*string, bool)`

GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpPrefix

`func (o *CreateFirewallRulePayload) SetRemoteIpPrefix(v string)`

SetRemoteIpPrefix sets RemoteIpPrefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


