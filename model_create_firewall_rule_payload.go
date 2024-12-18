/*
Infrahub-API

Leverage the Infrahub API and Hyperstack platform to easily create, manage, and scale powerful GPU virtual machines and their associated resources.   Access this SDK to automate the deployment of your workloads and streamline your infrastructure management.  To contribute, please raise an issue with a bug report, feature request, feedback, or general inquiry.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hyperstack

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CreateFirewallRulePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFirewallRulePayload{}

// CreateFirewallRulePayload struct for CreateFirewallRulePayload
type CreateFirewallRulePayload struct {
	// The direction of traffic that the firewall rule applies to.
	Direction string `json:"direction"`
	// The Ethernet type associated with the rule.
	Ethertype string `json:"ethertype"`
	// The maximum port number in the range of ports to be allowed by the firewall rule.
	PortRangeMax *int32 `json:"port_range_max,omitempty"`
	// The minimum port number in the range of ports to be allowed by the firewall rule.
	PortRangeMin *int32 `json:"port_range_min,omitempty"`
	// The network protocol associated with the rule. Call the [`GET /core/sg-rules-protocols`](https://infrahub-api-doc.nexgencloud.com/#get-/core/sg-rules-protocols) endpoint to retrieve a list of permitted network protocols.
	Protocol string `json:"protocol"`
	// The IP address range that is allowed to access the specified port. Use \"0.0.0.0/0\" to allow any IP address.
	RemoteIpPrefix string `json:"remote_ip_prefix"`
}

type _CreateFirewallRulePayload CreateFirewallRulePayload

// NewCreateFirewallRulePayload instantiates a new CreateFirewallRulePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFirewallRulePayload(direction string, ethertype string, protocol string, remoteIpPrefix string) *CreateFirewallRulePayload {
	this := CreateFirewallRulePayload{}
	this.Direction = direction
	this.Ethertype = ethertype
	this.Protocol = protocol
	this.RemoteIpPrefix = remoteIpPrefix
	return &this
}

// NewCreateFirewallRulePayloadWithDefaults instantiates a new CreateFirewallRulePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFirewallRulePayloadWithDefaults() *CreateFirewallRulePayload {
	this := CreateFirewallRulePayload{}
	return &this
}

// GetDirection returns the Direction field value
func (o *CreateFirewallRulePayload) GetDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *CreateFirewallRulePayload) GetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *CreateFirewallRulePayload) SetDirection(v string) {
	o.Direction = v
}

// GetEthertype returns the Ethertype field value
func (o *CreateFirewallRulePayload) GetEthertype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ethertype
}

// GetEthertypeOk returns a tuple with the Ethertype field value
// and a boolean to check if the value has been set.
func (o *CreateFirewallRulePayload) GetEthertypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ethertype, true
}

// SetEthertype sets field value
func (o *CreateFirewallRulePayload) SetEthertype(v string) {
	o.Ethertype = v
}

// GetPortRangeMax returns the PortRangeMax field value if set, zero value otherwise.
func (o *CreateFirewallRulePayload) GetPortRangeMax() int32 {
	if o == nil || IsNil(o.PortRangeMax) {
		var ret int32
		return ret
	}
	return *o.PortRangeMax
}

// GetPortRangeMaxOk returns a tuple with the PortRangeMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirewallRulePayload) GetPortRangeMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.PortRangeMax) {
		return nil, false
	}
	return o.PortRangeMax, true
}

// HasPortRangeMax returns a boolean if a field has been set.
func (o *CreateFirewallRulePayload) HasPortRangeMax() bool {
	if o != nil && !IsNil(o.PortRangeMax) {
		return true
	}

	return false
}

// SetPortRangeMax gets a reference to the given int32 and assigns it to the PortRangeMax field.
func (o *CreateFirewallRulePayload) SetPortRangeMax(v int32) {
	o.PortRangeMax = &v
}

// GetPortRangeMin returns the PortRangeMin field value if set, zero value otherwise.
func (o *CreateFirewallRulePayload) GetPortRangeMin() int32 {
	if o == nil || IsNil(o.PortRangeMin) {
		var ret int32
		return ret
	}
	return *o.PortRangeMin
}

// GetPortRangeMinOk returns a tuple with the PortRangeMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirewallRulePayload) GetPortRangeMinOk() (*int32, bool) {
	if o == nil || IsNil(o.PortRangeMin) {
		return nil, false
	}
	return o.PortRangeMin, true
}

// HasPortRangeMin returns a boolean if a field has been set.
func (o *CreateFirewallRulePayload) HasPortRangeMin() bool {
	if o != nil && !IsNil(o.PortRangeMin) {
		return true
	}

	return false
}

// SetPortRangeMin gets a reference to the given int32 and assigns it to the PortRangeMin field.
func (o *CreateFirewallRulePayload) SetPortRangeMin(v int32) {
	o.PortRangeMin = &v
}

// GetProtocol returns the Protocol field value
func (o *CreateFirewallRulePayload) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *CreateFirewallRulePayload) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *CreateFirewallRulePayload) SetProtocol(v string) {
	o.Protocol = v
}

// GetRemoteIpPrefix returns the RemoteIpPrefix field value
func (o *CreateFirewallRulePayload) GetRemoteIpPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteIpPrefix
}

// GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field value
// and a boolean to check if the value has been set.
func (o *CreateFirewallRulePayload) GetRemoteIpPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteIpPrefix, true
}

// SetRemoteIpPrefix sets field value
func (o *CreateFirewallRulePayload) SetRemoteIpPrefix(v string) {
	o.RemoteIpPrefix = v
}

func (o CreateFirewallRulePayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFirewallRulePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["direction"] = o.Direction
	toSerialize["ethertype"] = o.Ethertype
	if !IsNil(o.PortRangeMax) {
		toSerialize["port_range_max"] = o.PortRangeMax
	}
	if !IsNil(o.PortRangeMin) {
		toSerialize["port_range_min"] = o.PortRangeMin
	}
	toSerialize["protocol"] = o.Protocol
	toSerialize["remote_ip_prefix"] = o.RemoteIpPrefix
	return toSerialize, nil
}

func (o *CreateFirewallRulePayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"direction",
		"ethertype",
		"protocol",
		"remote_ip_prefix",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateFirewallRulePayload := _CreateFirewallRulePayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateFirewallRulePayload)

	if err != nil {
		return err
	}

	*o = CreateFirewallRulePayload(varCreateFirewallRulePayload)

	return err
}

type NullableCreateFirewallRulePayload struct {
	value *CreateFirewallRulePayload
	isSet bool
}

func (v NullableCreateFirewallRulePayload) Get() *CreateFirewallRulePayload {
	return v.value
}

func (v *NullableCreateFirewallRulePayload) Set(val *CreateFirewallRulePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFirewallRulePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFirewallRulePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFirewallRulePayload(val *CreateFirewallRulePayload) *NullableCreateFirewallRulePayload {
	return &NullableCreateFirewallRulePayload{value: val, isSet: true}
}

func (v NullableCreateFirewallRulePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFirewallRulePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
