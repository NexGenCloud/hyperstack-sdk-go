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

// checks if the GetApiKeysResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetApiKeysResponseModel{}

// GetApiKeysResponseModel struct for GetApiKeysResponseModel
type GetApiKeysResponseModel struct {
	ApiKeys []ApiKeyFields `json:"api_keys,omitempty"`
	Message *string        `json:"message,omitempty"`
	Status  *bool          `json:"status,omitempty"`
}

// NewGetApiKeysResponseModel instantiates a new GetApiKeysResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetApiKeysResponseModel() *GetApiKeysResponseModel {
	this := GetApiKeysResponseModel{}
	return &this
}

// NewGetApiKeysResponseModelWithDefaults instantiates a new GetApiKeysResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetApiKeysResponseModelWithDefaults() *GetApiKeysResponseModel {
	this := GetApiKeysResponseModel{}
	return &this
}

// GetApiKeys returns the ApiKeys field value if set, zero value otherwise.
func (o *GetApiKeysResponseModel) GetApiKeys() []ApiKeyFields {
	if o == nil || IsNil(o.ApiKeys) {
		var ret []ApiKeyFields
		return ret
	}
	return o.ApiKeys
}

// GetApiKeysOk returns a tuple with the ApiKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeysResponseModel) GetApiKeysOk() ([]ApiKeyFields, bool) {
	if o == nil || IsNil(o.ApiKeys) {
		return nil, false
	}
	return o.ApiKeys, true
}

// HasApiKeys returns a boolean if a field has been set.
func (o *GetApiKeysResponseModel) HasApiKeys() bool {
	if o != nil && !IsNil(o.ApiKeys) {
		return true
	}

	return false
}

// SetApiKeys gets a reference to the given []ApiKeyFields and assigns it to the ApiKeys field.
func (o *GetApiKeysResponseModel) SetApiKeys(v []ApiKeyFields) {
	o.ApiKeys = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetApiKeysResponseModel) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeysResponseModel) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetApiKeysResponseModel) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetApiKeysResponseModel) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetApiKeysResponseModel) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeysResponseModel) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetApiKeysResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *GetApiKeysResponseModel) SetStatus(v bool) {
	o.Status = &v
}

func (o GetApiKeysResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetApiKeysResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiKeys) {
		toSerialize["api_keys"] = o.ApiKeys
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableGetApiKeysResponseModel struct {
	value *GetApiKeysResponseModel
	isSet bool
}

func (v NullableGetApiKeysResponseModel) Get() *GetApiKeysResponseModel {
	return v.value
}

func (v *NullableGetApiKeysResponseModel) Set(val *GetApiKeysResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGetApiKeysResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGetApiKeysResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetApiKeysResponseModel(val *GetApiKeysResponseModel) *NullableGetApiKeysResponseModel {
	return &NullableGetApiKeysResponseModel{value: val, isSet: true}
}

func (v NullableGetApiKeysResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetApiKeysResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
