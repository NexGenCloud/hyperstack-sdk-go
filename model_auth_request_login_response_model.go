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

// checks if the AuthRequestLoginResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthRequestLoginResponseModel{}

// AuthRequestLoginResponseModel struct for AuthRequestLoginResponseModel
type AuthRequestLoginResponseModel struct {
	Data    *AuthRequestLoginFields `json:"data,omitempty"`
	Message *string                 `json:"message,omitempty"`
	Status  *bool                   `json:"status,omitempty"`
}

// NewAuthRequestLoginResponseModel instantiates a new AuthRequestLoginResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthRequestLoginResponseModel() *AuthRequestLoginResponseModel {
	this := AuthRequestLoginResponseModel{}
	return &this
}

// NewAuthRequestLoginResponseModelWithDefaults instantiates a new AuthRequestLoginResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthRequestLoginResponseModelWithDefaults() *AuthRequestLoginResponseModel {
	this := AuthRequestLoginResponseModel{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AuthRequestLoginResponseModel) GetData() AuthRequestLoginFields {
	if o == nil || IsNil(o.Data) {
		var ret AuthRequestLoginFields
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestLoginResponseModel) GetDataOk() (*AuthRequestLoginFields, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AuthRequestLoginResponseModel) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AuthRequestLoginFields and assigns it to the Data field.
func (o *AuthRequestLoginResponseModel) SetData(v AuthRequestLoginFields) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AuthRequestLoginResponseModel) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestLoginResponseModel) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AuthRequestLoginResponseModel) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AuthRequestLoginResponseModel) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AuthRequestLoginResponseModel) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestLoginResponseModel) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AuthRequestLoginResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *AuthRequestLoginResponseModel) SetStatus(v bool) {
	o.Status = &v
}

func (o AuthRequestLoginResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthRequestLoginResponseModel) ToMap() (map[string]interface{}, error) {
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

type NullableAuthRequestLoginResponseModel struct {
	value *AuthRequestLoginResponseModel
	isSet bool
}

func (v NullableAuthRequestLoginResponseModel) Get() *AuthRequestLoginResponseModel {
	return v.value
}

func (v *NullableAuthRequestLoginResponseModel) Set(val *AuthRequestLoginResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthRequestLoginResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthRequestLoginResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthRequestLoginResponseModel(val *AuthRequestLoginResponseModel) *NullableAuthRequestLoginResponseModel {
	return &NullableAuthRequestLoginResponseModel{value: val, isSet: true}
}

func (v NullableAuthRequestLoginResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthRequestLoginResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
