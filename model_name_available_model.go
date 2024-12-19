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

// checks if the NameAvailableModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NameAvailableModel{}

// NameAvailableModel struct for NameAvailableModel
type NameAvailableModel struct {
	Available *bool   `json:"available,omitempty"`
	Message   *string `json:"message,omitempty"`
	Name      *string `json:"name,omitempty"`
}

// NewNameAvailableModel instantiates a new NameAvailableModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNameAvailableModel() *NameAvailableModel {
	this := NameAvailableModel{}
	return &this
}

// NewNameAvailableModelWithDefaults instantiates a new NameAvailableModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNameAvailableModelWithDefaults() *NameAvailableModel {
	this := NameAvailableModel{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *NameAvailableModel) GetAvailable() bool {
	if o == nil || IsNil(o.Available) {
		var ret bool
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameAvailableModel) GetAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *NameAvailableModel) HasAvailable() bool {
	if o != nil && !IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given bool and assigns it to the Available field.
func (o *NameAvailableModel) SetAvailable(v bool) {
	o.Available = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *NameAvailableModel) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameAvailableModel) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *NameAvailableModel) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *NameAvailableModel) SetMessage(v string) {
	o.Message = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NameAvailableModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameAvailableModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NameAvailableModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NameAvailableModel) SetName(v string) {
	o.Name = &v
}

func (o NameAvailableModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NameAvailableModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableNameAvailableModel struct {
	value *NameAvailableModel
	isSet bool
}

func (v NullableNameAvailableModel) Get() *NameAvailableModel {
	return v.value
}

func (v *NullableNameAvailableModel) Set(val *NameAvailableModel) {
	v.value = val
	v.isSet = true
}

func (v NullableNameAvailableModel) IsSet() bool {
	return v.isSet
}

func (v *NullableNameAvailableModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNameAvailableModel(val *NameAvailableModel) *NullableNameAvailableModel {
	return &NullableNameAvailableModel{value: val, isSet: true}
}

func (v NullableNameAvailableModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNameAvailableModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}