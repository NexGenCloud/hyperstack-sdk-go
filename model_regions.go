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

// checks if the Regions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Regions{}

// Regions struct for Regions
type Regions struct {
	Message *string        `json:"message,omitempty"`
	Regions []RegionFields `json:"regions,omitempty"`
	Status  *bool          `json:"status,omitempty"`
}

// NewRegions instantiates a new Regions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegions() *Regions {
	this := Regions{}
	return &this
}

// NewRegionsWithDefaults instantiates a new Regions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionsWithDefaults() *Regions {
	this := Regions{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Regions) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Regions) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Regions) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Regions) SetMessage(v string) {
	o.Message = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *Regions) GetRegions() []RegionFields {
	if o == nil || IsNil(o.Regions) {
		var ret []RegionFields
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Regions) GetRegionsOk() ([]RegionFields, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *Regions) HasRegions() bool {
	if o != nil && !IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []RegionFields and assigns it to the Regions field.
func (o *Regions) SetRegions(v []RegionFields) {
	o.Regions = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Regions) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Regions) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Regions) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *Regions) SetStatus(v bool) {
	o.Status = &v
}

func (o Regions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Regions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableRegions struct {
	value *Regions
	isSet bool
}

func (v NullableRegions) Get() *Regions {
	return v.value
}

func (v *NullableRegions) Set(val *Regions) {
	v.value = val
	v.isSet = true
}

func (v NullableRegions) IsSet() bool {
	return v.isSet
}

func (v *NullableRegions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegions(val *Regions) *NullableRegions {
	return &NullableRegions{value: val, isSet: true}
}

func (v NullableRegions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}