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

// checks if the InstanceEventsFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceEventsFields{}

// InstanceEventsFields struct for InstanceEventsFields
type InstanceEventsFields struct {
	InstanceId *int32      `json:"instance_id,omitempty"`
	Message    *string     `json:"message,omitempty"`
	Object     *string     `json:"object,omitempty"`
	OrgId      *int32      `json:"org_id,omitempty"`
	Reason     *string     `json:"reason,omitempty"`
	Time       *CustomTime `json:"time,omitempty"`
	Type       *string     `json:"type,omitempty"`
	UserId     *int32      `json:"user_id,omitempty"`
}

// NewInstanceEventsFields instantiates a new InstanceEventsFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceEventsFields() *InstanceEventsFields {
	this := InstanceEventsFields{}
	return &this
}

// NewInstanceEventsFieldsWithDefaults instantiates a new InstanceEventsFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceEventsFieldsWithDefaults() *InstanceEventsFields {
	this := InstanceEventsFields{}
	return &this
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *InstanceEventsFields) GetInstanceId() int32 {
	if o == nil || IsNil(o.InstanceId) {
		var ret int32
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceEventsFields) GetInstanceIdOk() (*int32, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *InstanceEventsFields) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given int32 and assigns it to the InstanceId field.
func (o *InstanceEventsFields) SetInstanceId(v int32) {
	o.InstanceId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InstanceEventsFields) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceEventsFields) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InstanceEventsFields) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InstanceEventsFields) SetMessage(v string) {
	o.Message = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *InstanceEventsFields) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceEventsFields) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *InstanceEventsFields) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *InstanceEventsFields) SetObject(v string) {
	o.Object = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *InstanceEventsFields) GetOrgId() int32 {
	if o == nil || IsNil(o.OrgId) {
		var ret int32
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceEventsFields) GetOrgIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *InstanceEventsFields) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given int32 and assigns it to the OrgId field.
func (o *InstanceEventsFields) SetOrgId(v int32) {
	o.OrgId = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *InstanceEventsFields) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceEventsFields) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *InstanceEventsFields) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *InstanceEventsFields) SetReason(v string) {
	o.Reason = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *InstanceEventsFields) GetTime() CustomTime {
	if o == nil || IsNil(o.Time) {
		var ret CustomTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceEventsFields) GetTimeOk() (*CustomTime, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *InstanceEventsFields) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given CustomTime and assigns it to the Time field.
func (o *InstanceEventsFields) SetTime(v CustomTime) {
	o.Time = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InstanceEventsFields) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceEventsFields) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InstanceEventsFields) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InstanceEventsFields) SetType(v string) {
	o.Type = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *InstanceEventsFields) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceEventsFields) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *InstanceEventsFields) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *InstanceEventsFields) SetUserId(v int32) {
	o.UserId = &v
}

func (o InstanceEventsFields) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceEventsFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstanceId) {
		toSerialize["instance_id"] = o.InstanceId
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableInstanceEventsFields struct {
	value *InstanceEventsFields
	isSet bool
}

func (v NullableInstanceEventsFields) Get() *InstanceEventsFields {
	return v.value
}

func (v *NullableInstanceEventsFields) Set(val *InstanceEventsFields) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceEventsFields) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceEventsFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceEventsFields(val *InstanceEventsFields) *NullableInstanceEventsFields {
	return &NullableInstanceEventsFields{value: val, isSet: true}
}

func (v NullableInstanceEventsFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceEventsFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}