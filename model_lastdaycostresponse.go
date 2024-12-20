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

// checks if the Lastdaycostresponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Lastdaycostresponse{}

// Lastdaycostresponse struct for Lastdaycostresponse
type Lastdaycostresponse struct {
	Data    *Lastdaycostfields `json:"data,omitempty"`
	Message *string            `json:"message,omitempty"`
	Status  *bool              `json:"status,omitempty"`
}

// NewLastdaycostresponse instantiates a new Lastdaycostresponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLastdaycostresponse() *Lastdaycostresponse {
	this := Lastdaycostresponse{}
	return &this
}

// NewLastdaycostresponseWithDefaults instantiates a new Lastdaycostresponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLastdaycostresponseWithDefaults() *Lastdaycostresponse {
	this := Lastdaycostresponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Lastdaycostresponse) GetData() Lastdaycostfields {
	if o == nil || IsNil(o.Data) {
		var ret Lastdaycostfields
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lastdaycostresponse) GetDataOk() (*Lastdaycostfields, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Lastdaycostresponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Lastdaycostfields and assigns it to the Data field.
func (o *Lastdaycostresponse) SetData(v Lastdaycostfields) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Lastdaycostresponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lastdaycostresponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Lastdaycostresponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Lastdaycostresponse) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Lastdaycostresponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lastdaycostresponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Lastdaycostresponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *Lastdaycostresponse) SetStatus(v bool) {
	o.Status = &v
}

func (o Lastdaycostresponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Lastdaycostresponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableLastdaycostresponse struct {
	value *Lastdaycostresponse
	isSet bool
}

func (v NullableLastdaycostresponse) Get() *Lastdaycostresponse {
	return v.value
}

func (v *NullableLastdaycostresponse) Set(val *Lastdaycostresponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLastdaycostresponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLastdaycostresponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLastdaycostresponse(val *Lastdaycostresponse) *NullableLastdaycostresponse {
	return &NullableLastdaycostresponse{value: val, isSet: true}
}

func (v NullableLastdaycostresponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLastdaycostresponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
