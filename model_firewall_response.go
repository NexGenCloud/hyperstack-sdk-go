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

// checks if the FirewallResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirewallResponse{}

// FirewallResponse struct for FirewallResponse
type FirewallResponse struct {
	Firewall *FirewallFields `json:"firewall,omitempty"`
	Message  *string         `json:"message,omitempty"`
	Status   *bool           `json:"status,omitempty"`
}

// NewFirewallResponse instantiates a new FirewallResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallResponse() *FirewallResponse {
	this := FirewallResponse{}
	return &this
}

// NewFirewallResponseWithDefaults instantiates a new FirewallResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallResponseWithDefaults() *FirewallResponse {
	this := FirewallResponse{}
	return &this
}

// GetFirewall returns the Firewall field value if set, zero value otherwise.
func (o *FirewallResponse) GetFirewall() FirewallFields {
	if o == nil || IsNil(o.Firewall) {
		var ret FirewallFields
		return ret
	}
	return *o.Firewall
}

// GetFirewallOk returns a tuple with the Firewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallResponse) GetFirewallOk() (*FirewallFields, bool) {
	if o == nil || IsNil(o.Firewall) {
		return nil, false
	}
	return o.Firewall, true
}

// HasFirewall returns a boolean if a field has been set.
func (o *FirewallResponse) HasFirewall() bool {
	if o != nil && !IsNil(o.Firewall) {
		return true
	}

	return false
}

// SetFirewall gets a reference to the given FirewallFields and assigns it to the Firewall field.
func (o *FirewallResponse) SetFirewall(v FirewallFields) {
	o.Firewall = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *FirewallResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *FirewallResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *FirewallResponse) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FirewallResponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallResponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FirewallResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *FirewallResponse) SetStatus(v bool) {
	o.Status = &v
}

func (o FirewallResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirewallResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Firewall) {
		toSerialize["firewall"] = o.Firewall
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableFirewallResponse struct {
	value *FirewallResponse
	isSet bool
}

func (v NullableFirewallResponse) Get() *FirewallResponse {
	return v.value
}

func (v *NullableFirewallResponse) Set(val *FirewallResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallResponse(val *FirewallResponse) *NullableFirewallResponse {
	return &NullableFirewallResponse{value: val, isSet: true}
}

func (v NullableFirewallResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
