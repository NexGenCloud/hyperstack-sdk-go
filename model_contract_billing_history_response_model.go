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

// checks if the ContractBillingHistoryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractBillingHistoryResponseModel{}

// ContractBillingHistoryResponseModel struct for ContractBillingHistoryResponseModel
type ContractBillingHistoryResponseModel struct {
	BillingHistoryContract *ContractBillingHistory `json:"billing_history_contract,omitempty"`
	Message                *string                 `json:"message,omitempty"`
	Status                 *bool                   `json:"status,omitempty"`
}

// NewContractBillingHistoryResponseModel instantiates a new ContractBillingHistoryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractBillingHistoryResponseModel() *ContractBillingHistoryResponseModel {
	this := ContractBillingHistoryResponseModel{}
	return &this
}

// NewContractBillingHistoryResponseModelWithDefaults instantiates a new ContractBillingHistoryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractBillingHistoryResponseModelWithDefaults() *ContractBillingHistoryResponseModel {
	this := ContractBillingHistoryResponseModel{}
	return &this
}

// GetBillingHistoryContract returns the BillingHistoryContract field value if set, zero value otherwise.
func (o *ContractBillingHistoryResponseModel) GetBillingHistoryContract() ContractBillingHistory {
	if o == nil || IsNil(o.BillingHistoryContract) {
		var ret ContractBillingHistory
		return ret
	}
	return *o.BillingHistoryContract
}

// GetBillingHistoryContractOk returns a tuple with the BillingHistoryContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractBillingHistoryResponseModel) GetBillingHistoryContractOk() (*ContractBillingHistory, bool) {
	if o == nil || IsNil(o.BillingHistoryContract) {
		return nil, false
	}
	return o.BillingHistoryContract, true
}

// HasBillingHistoryContract returns a boolean if a field has been set.
func (o *ContractBillingHistoryResponseModel) HasBillingHistoryContract() bool {
	if o != nil && !IsNil(o.BillingHistoryContract) {
		return true
	}

	return false
}

// SetBillingHistoryContract gets a reference to the given ContractBillingHistory and assigns it to the BillingHistoryContract field.
func (o *ContractBillingHistoryResponseModel) SetBillingHistoryContract(v ContractBillingHistory) {
	o.BillingHistoryContract = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ContractBillingHistoryResponseModel) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractBillingHistoryResponseModel) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ContractBillingHistoryResponseModel) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ContractBillingHistoryResponseModel) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ContractBillingHistoryResponseModel) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractBillingHistoryResponseModel) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ContractBillingHistoryResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *ContractBillingHistoryResponseModel) SetStatus(v bool) {
	o.Status = &v
}

func (o ContractBillingHistoryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractBillingHistoryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingHistoryContract) {
		toSerialize["billing_history_contract"] = o.BillingHistoryContract
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableContractBillingHistoryResponseModel struct {
	value *ContractBillingHistoryResponseModel
	isSet bool
}

func (v NullableContractBillingHistoryResponseModel) Get() *ContractBillingHistoryResponseModel {
	return v.value
}

func (v *NullableContractBillingHistoryResponseModel) Set(val *ContractBillingHistoryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableContractBillingHistoryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableContractBillingHistoryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractBillingHistoryResponseModel(val *ContractBillingHistoryResponseModel) *NullableContractBillingHistoryResponseModel {
	return &NullableContractBillingHistoryResponseModel{value: val, isSet: true}
}

func (v NullableContractBillingHistoryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractBillingHistoryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}