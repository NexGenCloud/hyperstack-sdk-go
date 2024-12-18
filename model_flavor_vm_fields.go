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

// checks if the FlavorVMFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlavorVMFields{}

// FlavorVMFields struct for FlavorVMFields
type FlavorVMFields struct {
	CreatedAt   *CustomTime `json:"created_at,omitempty"`
	Host        *string     `json:"host,omitempty"`
	Id          *int32      `json:"id,omitempty"`
	Name        *string     `json:"name,omitempty"`
	OpenstackId *string     `json:"openstack_id,omitempty"`
	OrgId       *int32      `json:"org_id,omitempty"`
	Status      *string     `json:"status,omitempty"`
}

// NewFlavorVMFields instantiates a new FlavorVMFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlavorVMFields() *FlavorVMFields {
	this := FlavorVMFields{}
	return &this
}

// NewFlavorVMFieldsWithDefaults instantiates a new FlavorVMFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlavorVMFieldsWithDefaults() *FlavorVMFields {
	this := FlavorVMFields{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FlavorVMFields) GetCreatedAt() CustomTime {
	if o == nil || IsNil(o.CreatedAt) {
		var ret CustomTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorVMFields) GetCreatedAtOk() (*CustomTime, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FlavorVMFields) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given CustomTime and assigns it to the CreatedAt field.
func (o *FlavorVMFields) SetCreatedAt(v CustomTime) {
	o.CreatedAt = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *FlavorVMFields) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorVMFields) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *FlavorVMFields) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *FlavorVMFields) SetHost(v string) {
	o.Host = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FlavorVMFields) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorVMFields) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FlavorVMFields) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FlavorVMFields) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FlavorVMFields) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorVMFields) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FlavorVMFields) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FlavorVMFields) SetName(v string) {
	o.Name = &v
}

// GetOpenstackId returns the OpenstackId field value if set, zero value otherwise.
func (o *FlavorVMFields) GetOpenstackId() string {
	if o == nil || IsNil(o.OpenstackId) {
		var ret string
		return ret
	}
	return *o.OpenstackId
}

// GetOpenstackIdOk returns a tuple with the OpenstackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorVMFields) GetOpenstackIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenstackId) {
		return nil, false
	}
	return o.OpenstackId, true
}

// HasOpenstackId returns a boolean if a field has been set.
func (o *FlavorVMFields) HasOpenstackId() bool {
	if o != nil && !IsNil(o.OpenstackId) {
		return true
	}

	return false
}

// SetOpenstackId gets a reference to the given string and assigns it to the OpenstackId field.
func (o *FlavorVMFields) SetOpenstackId(v string) {
	o.OpenstackId = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *FlavorVMFields) GetOrgId() int32 {
	if o == nil || IsNil(o.OrgId) {
		var ret int32
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorVMFields) GetOrgIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *FlavorVMFields) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given int32 and assigns it to the OrgId field.
func (o *FlavorVMFields) SetOrgId(v int32) {
	o.OrgId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FlavorVMFields) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorVMFields) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FlavorVMFields) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FlavorVMFields) SetStatus(v string) {
	o.Status = &v
}

func (o FlavorVMFields) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlavorVMFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OpenstackId) {
		toSerialize["openstack_id"] = o.OpenstackId
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableFlavorVMFields struct {
	value *FlavorVMFields
	isSet bool
}

func (v NullableFlavorVMFields) Get() *FlavorVMFields {
	return v.value
}

func (v *NullableFlavorVMFields) Set(val *FlavorVMFields) {
	v.value = val
	v.isSet = true
}

func (v NullableFlavorVMFields) IsSet() bool {
	return v.isSet
}

func (v *NullableFlavorVMFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlavorVMFields(val *FlavorVMFields) *NullableFlavorVMFields {
	return &NullableFlavorVMFields{value: val, isSet: true}
}

func (v NullableFlavorVMFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlavorVMFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
