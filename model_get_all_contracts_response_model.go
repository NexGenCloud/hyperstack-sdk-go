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

// checks if the GetAllContractsResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllContractsResponseModel{}

// GetAllContractsResponseModel struct for GetAllContractsResponseModel
type GetAllContractsResponseModel struct {
	Contracts []GetAllContractFields `json:"contracts,omitempty"`
	Message   *string                `json:"message,omitempty"`
	Status    *bool                  `json:"status,omitempty"`
}

// NewGetAllContractsResponseModel instantiates a new GetAllContractsResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllContractsResponseModel() *GetAllContractsResponseModel {
	this := GetAllContractsResponseModel{}
	return &this
}

// NewGetAllContractsResponseModelWithDefaults instantiates a new GetAllContractsResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllContractsResponseModelWithDefaults() *GetAllContractsResponseModel {
	this := GetAllContractsResponseModel{}
	return &this
}

// GetContracts returns the Contracts field value if set, zero value otherwise.
func (o *GetAllContractsResponseModel) GetContracts() []GetAllContractFields {
	if o == nil || IsNil(o.Contracts) {
		var ret []GetAllContractFields
		return ret
	}
	return o.Contracts
}

// GetContractsOk returns a tuple with the Contracts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllContractsResponseModel) GetContractsOk() ([]GetAllContractFields, bool) {
	if o == nil || IsNil(o.Contracts) {
		return nil, false
	}
	return o.Contracts, true
}

// HasContracts returns a boolean if a field has been set.
func (o *GetAllContractsResponseModel) HasContracts() bool {
	if o != nil && !IsNil(o.Contracts) {
		return true
	}

	return false
}

// SetContracts gets a reference to the given []GetAllContractFields and assigns it to the Contracts field.
func (o *GetAllContractsResponseModel) SetContracts(v []GetAllContractFields) {
	o.Contracts = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetAllContractsResponseModel) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllContractsResponseModel) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetAllContractsResponseModel) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetAllContractsResponseModel) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetAllContractsResponseModel) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllContractsResponseModel) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetAllContractsResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *GetAllContractsResponseModel) SetStatus(v bool) {
	o.Status = &v
}

func (o GetAllContractsResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllContractsResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Contracts) {
		toSerialize["contracts"] = o.Contracts
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableGetAllContractsResponseModel struct {
	value *GetAllContractsResponseModel
	isSet bool
}

func (v NullableGetAllContractsResponseModel) Get() *GetAllContractsResponseModel {
	return v.value
}

func (v *NullableGetAllContractsResponseModel) Set(val *GetAllContractsResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllContractsResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllContractsResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllContractsResponseModel(val *GetAllContractsResponseModel) *NullableGetAllContractsResponseModel {
	return &NullableGetAllContractsResponseModel{value: val, isSet: true}
}

func (v NullableGetAllContractsResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllContractsResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
