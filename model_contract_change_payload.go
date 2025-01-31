/*
Infrahub-API

Leverage the Infrahub API and Hyperstack platform to easily create, manage, and scale powerful GPU virtual machines and their associated resources.   Access this SDK to automate the deployment of your workloads and streamline your infrastructure management.  To contribute, please raise an issue with a bug report, feature request, feedback, or general inquiry.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hyperstack

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ContractChangePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractChangePayload{}

// ContractChangePayload struct for ContractChangePayload
type ContractChangePayload struct {
	// List of field changes for 'updated' type
	Changes []FieldChange `json:"changes,omitempty"`
	// The ID of the contract
	Id int32 `json:"id"`
	// The ORG ID of the contract
	OrgId int32 `json:"org_id"`
	// Purpose of the change: created, deleted, expired, or updated
	Type string `json:"type"`
}

type _ContractChangePayload ContractChangePayload

// NewContractChangePayload instantiates a new ContractChangePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractChangePayload(id int32, orgId int32, type_ string) *ContractChangePayload {
	this := ContractChangePayload{}
	this.Id = id
	this.OrgId = orgId
	this.Type = type_
	return &this
}

// NewContractChangePayloadWithDefaults instantiates a new ContractChangePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractChangePayloadWithDefaults() *ContractChangePayload {
	this := ContractChangePayload{}
	return &this
}

// GetChanges returns the Changes field value if set, zero value otherwise.
func (o *ContractChangePayload) GetChanges() []FieldChange {
	if o == nil || IsNil(o.Changes) {
		var ret []FieldChange
		return ret
	}
	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractChangePayload) GetChangesOk() ([]FieldChange, bool) {
	if o == nil || IsNil(o.Changes) {
		return nil, false
	}
	return o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *ContractChangePayload) HasChanges() bool {
	if o != nil && !IsNil(o.Changes) {
		return true
	}

	return false
}

// SetChanges gets a reference to the given []FieldChange and assigns it to the Changes field.
func (o *ContractChangePayload) SetChanges(v []FieldChange) {
	o.Changes = v
}

// GetId returns the Id field value
func (o *ContractChangePayload) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ContractChangePayload) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ContractChangePayload) SetId(v int32) {
	o.Id = v
}

// GetOrgId returns the OrgId field value
func (o *ContractChangePayload) GetOrgId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *ContractChangePayload) GetOrgIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *ContractChangePayload) SetOrgId(v int32) {
	o.OrgId = v
}

// GetType returns the Type field value
func (o *ContractChangePayload) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContractChangePayload) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContractChangePayload) SetType(v string) {
	o.Type = v
}

func (o ContractChangePayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractChangePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Changes) {
		toSerialize["changes"] = o.Changes
	}
	toSerialize["id"] = o.Id
	toSerialize["org_id"] = o.OrgId
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ContractChangePayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"org_id",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varContractChangePayload := _ContractChangePayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContractChangePayload)

	if err != nil {
		return err
	}

	*o = ContractChangePayload(varContractChangePayload)

	return err
}

type NullableContractChangePayload struct {
	value *ContractChangePayload
	isSet bool
}

func (v NullableContractChangePayload) Get() *ContractChangePayload {
	return v.value
}

func (v *NullableContractChangePayload) Set(val *ContractChangePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableContractChangePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableContractChangePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractChangePayload(val *ContractChangePayload) *NullableContractChangePayload {
	return &NullableContractChangePayload{value: val, isSet: true}
}

func (v NullableContractChangePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractChangePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
