/*
Infrahub-API

Leverage the Infrahub API and Hyperstack platform to easily create, manage, and scale powerful GPU virtual machines and their associated resources.   Access this SDK to automate the deployment of your workloads and streamline your infrastructure management.  To contribute, please raise an issue with a bug report, feature request, feedback, or general inquiry.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hyperstack

import (
	"encoding/json"
)

// checks if the FirewallRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirewallRule{}

// FirewallRule struct for FirewallRule
type FirewallRule struct {
	FirewallRule *SecurityGroupRuleFields `json:"firewall_rule,omitempty"`
	Message      *string                  `json:"message,omitempty"`
	Status       *bool                    `json:"status,omitempty"`
}

// NewFirewallRule instantiates a new FirewallRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallRule() *FirewallRule {
	this := FirewallRule{}
	return &this
}

// NewFirewallRuleWithDefaults instantiates a new FirewallRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallRuleWithDefaults() *FirewallRule {
	this := FirewallRule{}
	return &this
}

// GetFirewallRule returns the FirewallRule field value if set, zero value otherwise.
func (o *FirewallRule) GetFirewallRule() SecurityGroupRuleFields {
	if o == nil || IsNil(o.FirewallRule) {
		var ret SecurityGroupRuleFields
		return ret
	}
	return *o.FirewallRule
}

// GetFirewallRuleOk returns a tuple with the FirewallRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetFirewallRuleOk() (*SecurityGroupRuleFields, bool) {
	if o == nil || IsNil(o.FirewallRule) {
		return nil, false
	}
	return o.FirewallRule, true
}

// HasFirewallRule returns a boolean if a field has been set.
func (o *FirewallRule) HasFirewallRule() bool {
	if o != nil && !IsNil(o.FirewallRule) {
		return true
	}

	return false
}

// SetFirewallRule gets a reference to the given SecurityGroupRuleFields and assigns it to the FirewallRule field.
func (o *FirewallRule) SetFirewallRule(v SecurityGroupRuleFields) {
	o.FirewallRule = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *FirewallRule) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *FirewallRule) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *FirewallRule) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FirewallRule) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FirewallRule) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *FirewallRule) SetStatus(v bool) {
	o.Status = &v
}

func (o FirewallRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirewallRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirewallRule) {
		toSerialize["firewall_rule"] = o.FirewallRule
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableFirewallRule struct {
	value *FirewallRule
	isSet bool
}

func (v NullableFirewallRule) Get() *FirewallRule {
	return v.value
}

func (v *NullableFirewallRule) Set(val *FirewallRule) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallRule) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallRule(val *FirewallRule) *NullableFirewallRule {
	return &NullableFirewallRule{value: val, isSet: true}
}

func (v NullableFirewallRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
