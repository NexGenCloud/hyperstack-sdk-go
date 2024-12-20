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

// checks if the Templates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Templates{}

// Templates struct for Templates
type Templates struct {
	Message   *string          `json:"message,omitempty"`
	Status    *bool            `json:"status,omitempty"`
	Templates []TemplateFields `json:"templates,omitempty"`
}

// NewTemplates instantiates a new Templates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplates() *Templates {
	this := Templates{}
	return &this
}

// NewTemplatesWithDefaults instantiates a new Templates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplatesWithDefaults() *Templates {
	this := Templates{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Templates) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Templates) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Templates) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Templates) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Templates) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Templates) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Templates) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *Templates) SetStatus(v bool) {
	o.Status = &v
}

// GetTemplates returns the Templates field value if set, zero value otherwise.
func (o *Templates) GetTemplates() []TemplateFields {
	if o == nil || IsNil(o.Templates) {
		var ret []TemplateFields
		return ret
	}
	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Templates) GetTemplatesOk() ([]TemplateFields, bool) {
	if o == nil || IsNil(o.Templates) {
		return nil, false
	}
	return o.Templates, true
}

// HasTemplates returns a boolean if a field has been set.
func (o *Templates) HasTemplates() bool {
	if o != nil && !IsNil(o.Templates) {
		return true
	}

	return false
}

// SetTemplates gets a reference to the given []TemplateFields and assigns it to the Templates field.
func (o *Templates) SetTemplates(v []TemplateFields) {
	o.Templates = v
}

func (o Templates) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Templates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Templates) {
		toSerialize["templates"] = o.Templates
	}
	return toSerialize, nil
}

type NullableTemplates struct {
	value *Templates
	isSet bool
}

func (v NullableTemplates) Get() *Templates {
	return v.value
}

func (v *NullableTemplates) Set(val *Templates) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplates) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplates(val *Templates) *NullableTemplates {
	return &NullableTemplates{value: val, isSet: true}
}

func (v NullableTemplates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
