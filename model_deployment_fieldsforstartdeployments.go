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

// checks if the DeploymentFieldsforstartdeployments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentFieldsforstartdeployments{}

// DeploymentFieldsforstartdeployments struct for DeploymentFieldsforstartdeployments
type DeploymentFieldsforstartdeployments struct {
	CreatedAt   *CustomTime `json:"created_at,omitempty"`
	Description *string     `json:"description,omitempty"`
	Id          *int32      `json:"id,omitempty"`
	Name        *string     `json:"name,omitempty"`
	Status      *string     `json:"status,omitempty"`
	Template    *string     `json:"template,omitempty"`
	Variables   *string     `json:"variables,omitempty"`
}

// NewDeploymentFieldsforstartdeployments instantiates a new DeploymentFieldsforstartdeployments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentFieldsforstartdeployments() *DeploymentFieldsforstartdeployments {
	this := DeploymentFieldsforstartdeployments{}
	return &this
}

// NewDeploymentFieldsforstartdeploymentsWithDefaults instantiates a new DeploymentFieldsforstartdeployments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentFieldsforstartdeploymentsWithDefaults() *DeploymentFieldsforstartdeployments {
	this := DeploymentFieldsforstartdeployments{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DeploymentFieldsforstartdeployments) GetCreatedAt() CustomTime {
	if o == nil || IsNil(o.CreatedAt) {
		var ret CustomTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentFieldsforstartdeployments) GetCreatedAtOk() (*CustomTime, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DeploymentFieldsforstartdeployments) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given CustomTime and assigns it to the CreatedAt field.
func (o *DeploymentFieldsforstartdeployments) SetCreatedAt(v CustomTime) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeploymentFieldsforstartdeployments) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentFieldsforstartdeployments) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeploymentFieldsforstartdeployments) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeploymentFieldsforstartdeployments) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeploymentFieldsforstartdeployments) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentFieldsforstartdeployments) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeploymentFieldsforstartdeployments) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DeploymentFieldsforstartdeployments) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentFieldsforstartdeployments) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentFieldsforstartdeployments) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentFieldsforstartdeployments) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentFieldsforstartdeployments) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeploymentFieldsforstartdeployments) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentFieldsforstartdeployments) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploymentFieldsforstartdeployments) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DeploymentFieldsforstartdeployments) SetStatus(v string) {
	o.Status = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *DeploymentFieldsforstartdeployments) GetTemplate() string {
	if o == nil || IsNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentFieldsforstartdeployments) GetTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *DeploymentFieldsforstartdeployments) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *DeploymentFieldsforstartdeployments) SetTemplate(v string) {
	o.Template = &v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *DeploymentFieldsforstartdeployments) GetVariables() string {
	if o == nil || IsNil(o.Variables) {
		var ret string
		return ret
	}
	return *o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentFieldsforstartdeployments) GetVariablesOk() (*string, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *DeploymentFieldsforstartdeployments) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given string and assigns it to the Variables field.
func (o *DeploymentFieldsforstartdeployments) SetVariables(v string) {
	o.Variables = &v
}

func (o DeploymentFieldsforstartdeployments) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentFieldsforstartdeployments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	return toSerialize, nil
}

type NullableDeploymentFieldsforstartdeployments struct {
	value *DeploymentFieldsforstartdeployments
	isSet bool
}

func (v NullableDeploymentFieldsforstartdeployments) Get() *DeploymentFieldsforstartdeployments {
	return v.value
}

func (v *NullableDeploymentFieldsforstartdeployments) Set(val *DeploymentFieldsforstartdeployments) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentFieldsforstartdeployments) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentFieldsforstartdeployments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentFieldsforstartdeployments(val *DeploymentFieldsforstartdeployments) *NullableDeploymentFieldsforstartdeployments {
	return &NullableDeploymentFieldsforstartdeployments{value: val, isSet: true}
}

func (v NullableDeploymentFieldsforstartdeployments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentFieldsforstartdeployments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
