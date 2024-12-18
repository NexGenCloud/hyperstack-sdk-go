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

// checks if the GenerateUpdateApiKeyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateUpdateApiKeyResponseModel{}

// GenerateUpdateApiKeyResponseModel struct for GenerateUpdateApiKeyResponseModel
type GenerateUpdateApiKeyResponseModel struct {
	ApiKey  *ApiKeyFields `json:"api_key,omitempty"`
	Message *string       `json:"message,omitempty"`
	Status  *bool         `json:"status,omitempty"`
}

// NewGenerateUpdateApiKeyResponseModel instantiates a new GenerateUpdateApiKeyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateUpdateApiKeyResponseModel() *GenerateUpdateApiKeyResponseModel {
	this := GenerateUpdateApiKeyResponseModel{}
	return &this
}

// NewGenerateUpdateApiKeyResponseModelWithDefaults instantiates a new GenerateUpdateApiKeyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateUpdateApiKeyResponseModelWithDefaults() *GenerateUpdateApiKeyResponseModel {
	this := GenerateUpdateApiKeyResponseModel{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *GenerateUpdateApiKeyResponseModel) GetApiKey() ApiKeyFields {
	if o == nil || IsNil(o.ApiKey) {
		var ret ApiKeyFields
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateUpdateApiKeyResponseModel) GetApiKeyOk() (*ApiKeyFields, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *GenerateUpdateApiKeyResponseModel) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given ApiKeyFields and assigns it to the ApiKey field.
func (o *GenerateUpdateApiKeyResponseModel) SetApiKey(v ApiKeyFields) {
	o.ApiKey = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GenerateUpdateApiKeyResponseModel) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateUpdateApiKeyResponseModel) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GenerateUpdateApiKeyResponseModel) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GenerateUpdateApiKeyResponseModel) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GenerateUpdateApiKeyResponseModel) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateUpdateApiKeyResponseModel) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GenerateUpdateApiKeyResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *GenerateUpdateApiKeyResponseModel) SetStatus(v bool) {
	o.Status = &v
}

func (o GenerateUpdateApiKeyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateUpdateApiKeyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiKey) {
		toSerialize["api_key"] = o.ApiKey
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableGenerateUpdateApiKeyResponseModel struct {
	value *GenerateUpdateApiKeyResponseModel
	isSet bool
}

func (v NullableGenerateUpdateApiKeyResponseModel) Get() *GenerateUpdateApiKeyResponseModel {
	return v.value
}

func (v *NullableGenerateUpdateApiKeyResponseModel) Set(val *GenerateUpdateApiKeyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateUpdateApiKeyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateUpdateApiKeyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateUpdateApiKeyResponseModel(val *GenerateUpdateApiKeyResponseModel) *NullableGenerateUpdateApiKeyResponseModel {
	return &NullableGenerateUpdateApiKeyResponseModel{value: val, isSet: true}
}

func (v NullableGenerateUpdateApiKeyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateUpdateApiKeyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
