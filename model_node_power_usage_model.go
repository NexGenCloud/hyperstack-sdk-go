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

// checks if the NodePowerUsageModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodePowerUsageModel{}

// NodePowerUsageModel struct for NodePowerUsageModel
type NodePowerUsageModel struct {
	Flavors       []string            `json:"flavors,omitempty"`
	NexgenName    *string             `json:"nexgen_name,omitempty"`
	OpenstackId   string              `json:"openstack_id"`
	OpenstackName *string             `json:"openstack_name,omitempty"`
	Organizations []int32             `json:"organizations,omitempty"`
	PowerUsages   *PowerUsageModel    `json:"power_usages,omitempty"`
	Projects      []string            `json:"projects,omitempty"`
	ProvisionDate *CustomTime         `json:"provision_date,omitempty"`
	Status        *string             `json:"status,omitempty"`
	Stocks        []NodeStocksPayload `json:"stocks,omitempty"`
	SunsetDate    *CustomTime         `json:"sunset_date,omitempty"`
}

type _NodePowerUsageModel NodePowerUsageModel

// NewNodePowerUsageModel instantiates a new NodePowerUsageModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodePowerUsageModel(openstackId string) *NodePowerUsageModel {
	this := NodePowerUsageModel{}
	this.OpenstackId = openstackId
	return &this
}

// NewNodePowerUsageModelWithDefaults instantiates a new NodePowerUsageModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodePowerUsageModelWithDefaults() *NodePowerUsageModel {
	this := NodePowerUsageModel{}
	return &this
}

// GetFlavors returns the Flavors field value if set, zero value otherwise.
func (o *NodePowerUsageModel) GetFlavors() []string {
	if o == nil || IsNil(o.Flavors) {
		var ret []string
		return ret
	}
	return o.Flavors
}

// GetFlavorsOk returns a tuple with the Flavors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePowerUsageModel) GetFlavorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flavors) {
		return nil, false
	}
	return o.Flavors, true
}

// HasFlavors returns a boolean if a field has been set.
func (o *NodePowerUsageModel) HasFlavors() bool {
	if o != nil && !IsNil(o.Flavors) {
		return true
	}

	return false
}

// SetFlavors gets a reference to the given []string and assigns it to the Flavors field.
func (o *NodePowerUsageModel) SetFlavors(v []string) {
	o.Flavors = v
}

// GetNexgenName returns the NexgenName field value if set, zero value otherwise.
func (o *NodePowerUsageModel) GetNexgenName() string {
	if o == nil || IsNil(o.NexgenName) {
		var ret string
		return ret
	}
	return *o.NexgenName
}

// GetNexgenNameOk returns a tuple with the NexgenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePowerUsageModel) GetNexgenNameOk() (*string, bool) {
	if o == nil || IsNil(o.NexgenName) {
		return nil, false
	}
	return o.NexgenName, true
}

// HasNexgenName returns a boolean if a field has been set.
func (o *NodePowerUsageModel) HasNexgenName() bool {
	if o != nil && !IsNil(o.NexgenName) {
		return true
	}

	return false
}

// SetNexgenName gets a reference to the given string and assigns it to the NexgenName field.
func (o *NodePowerUsageModel) SetNexgenName(v string) {
	o.NexgenName = &v
}

// GetOpenstackId returns the OpenstackId field value
func (o *NodePowerUsageModel) GetOpenstackId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OpenstackId
}

// GetOpenstackIdOk returns a tuple with the OpenstackId field value
// and a boolean to check if the value has been set.
func (o *NodePowerUsageModel) GetOpenstackIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpenstackId, true
}

// SetOpenstackId sets field value
func (o *NodePowerUsageModel) SetOpenstackId(v string) {
	o.OpenstackId = v
}

// GetOpenstackName returns the OpenstackName field value if set, zero value otherwise.
func (o *NodePowerUsageModel) GetOpenstackName() string {
	if o == nil || IsNil(o.OpenstackName) {
		var ret string
		return ret
	}
	return *o.OpenstackName
}

// GetOpenstackNameOk returns a tuple with the OpenstackName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePowerUsageModel) GetOpenstackNameOk() (*string, bool) {
	if o == nil || IsNil(o.OpenstackName) {
		return nil, false
	}
	return o.OpenstackName, true
}

// HasOpenstackName returns a boolean if a field has been set.
func (o *NodePowerUsageModel) HasOpenstackName() bool {
	if o != nil && !IsNil(o.OpenstackName) {
		return true
	}

	return false
}

// SetOpenstackName gets a reference to the given string and assigns it to the OpenstackName field.
func (o *NodePowerUsageModel) SetOpenstackName(v string) {
	o.OpenstackName = &v
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *NodePowerUsageModel) GetOrganizations() []int32 {
	if o == nil || IsNil(o.Organizations) {
		var ret []int32
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePowerUsageModel) GetOrganizationsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Organizations) {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *NodePowerUsageModel) HasOrganizations() bool {
	if o != nil && !IsNil(o.Organizations) {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []int32 and assigns it to the Organizations field.
func (o *NodePowerUsageModel) SetOrganizations(v []int32) {
	o.Organizations = v
}

// GetPowerUsages returns the PowerUsages field value if set, zero value otherwise.
func (o *NodePowerUsageModel) GetPowerUsages() PowerUsageModel {
	if o == nil || IsNil(o.PowerUsages) {
		var ret PowerUsageModel
		return ret
	}
	return *o.PowerUsages
}

// GetPowerUsagesOk returns a tuple with the PowerUsages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePowerUsageModel) GetPowerUsagesOk() (*PowerUsageModel, bool) {
	if o == nil || IsNil(o.PowerUsages) {
		return nil, false
	}
	return o.PowerUsages, true
}

// HasPowerUsages returns a boolean if a field has been set.
func (o *NodePowerUsageModel) HasPowerUsages() bool {
	if o != nil && !IsNil(o.PowerUsages) {
		return true
	}

	return false
}

// SetPowerUsages gets a reference to the given PowerUsageModel and assigns it to the PowerUsages field.
func (o *NodePowerUsageModel) SetPowerUsages(v PowerUsageModel) {
	o.PowerUsages = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *NodePowerUsageModel) GetProjects() []string {
	if o == nil || IsNil(o.Projects) {
		var ret []string
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePowerUsageModel) GetProjectsOk() ([]string, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *NodePowerUsageModel) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []string and assigns it to the Projects field.
func (o *NodePowerUsageModel) SetProjects(v []string) {
	o.Projects = v
}

// GetProvisionDate returns the ProvisionDate field value if set, zero value otherwise.
func (o *NodePowerUsageModel) GetProvisionDate() CustomTime {
	if o == nil || IsNil(o.ProvisionDate) {
		var ret CustomTime
		return ret
	}
	return *o.ProvisionDate
}

// GetProvisionDateOk returns a tuple with the ProvisionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePowerUsageModel) GetProvisionDateOk() (*CustomTime, bool) {
	if o == nil || IsNil(o.ProvisionDate) {
		return nil, false
	}
	return o.ProvisionDate, true
}

// HasProvisionDate returns a boolean if a field has been set.
func (o *NodePowerUsageModel) HasProvisionDate() bool {
	if o != nil && !IsNil(o.ProvisionDate) {
		return true
	}

	return false
}

// SetProvisionDate gets a reference to the given CustomTime and assigns it to the ProvisionDate field.
func (o *NodePowerUsageModel) SetProvisionDate(v CustomTime) {
	o.ProvisionDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NodePowerUsageModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePowerUsageModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NodePowerUsageModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NodePowerUsageModel) SetStatus(v string) {
	o.Status = &v
}

// GetStocks returns the Stocks field value if set, zero value otherwise.
func (o *NodePowerUsageModel) GetStocks() []NodeStocksPayload {
	if o == nil || IsNil(o.Stocks) {
		var ret []NodeStocksPayload
		return ret
	}
	return o.Stocks
}

// GetStocksOk returns a tuple with the Stocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePowerUsageModel) GetStocksOk() ([]NodeStocksPayload, bool) {
	if o == nil || IsNil(o.Stocks) {
		return nil, false
	}
	return o.Stocks, true
}

// HasStocks returns a boolean if a field has been set.
func (o *NodePowerUsageModel) HasStocks() bool {
	if o != nil && !IsNil(o.Stocks) {
		return true
	}

	return false
}

// SetStocks gets a reference to the given []NodeStocksPayload and assigns it to the Stocks field.
func (o *NodePowerUsageModel) SetStocks(v []NodeStocksPayload) {
	o.Stocks = v
}

// GetSunsetDate returns the SunsetDate field value if set, zero value otherwise.
func (o *NodePowerUsageModel) GetSunsetDate() CustomTime {
	if o == nil || IsNil(o.SunsetDate) {
		var ret CustomTime
		return ret
	}
	return *o.SunsetDate
}

// GetSunsetDateOk returns a tuple with the SunsetDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePowerUsageModel) GetSunsetDateOk() (*CustomTime, bool) {
	if o == nil || IsNil(o.SunsetDate) {
		return nil, false
	}
	return o.SunsetDate, true
}

// HasSunsetDate returns a boolean if a field has been set.
func (o *NodePowerUsageModel) HasSunsetDate() bool {
	if o != nil && !IsNil(o.SunsetDate) {
		return true
	}

	return false
}

// SetSunsetDate gets a reference to the given CustomTime and assigns it to the SunsetDate field.
func (o *NodePowerUsageModel) SetSunsetDate(v CustomTime) {
	o.SunsetDate = &v
}

func (o NodePowerUsageModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodePowerUsageModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Flavors) {
		toSerialize["flavors"] = o.Flavors
	}
	if !IsNil(o.NexgenName) {
		toSerialize["nexgen_name"] = o.NexgenName
	}
	toSerialize["openstack_id"] = o.OpenstackId
	if !IsNil(o.OpenstackName) {
		toSerialize["openstack_name"] = o.OpenstackName
	}
	if !IsNil(o.Organizations) {
		toSerialize["organizations"] = o.Organizations
	}
	if !IsNil(o.PowerUsages) {
		toSerialize["power_usages"] = o.PowerUsages
	}
	if !IsNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	if !IsNil(o.ProvisionDate) {
		toSerialize["provision_date"] = o.ProvisionDate
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Stocks) {
		toSerialize["stocks"] = o.Stocks
	}
	if !IsNil(o.SunsetDate) {
		toSerialize["sunset_date"] = o.SunsetDate
	}
	return toSerialize, nil
}

func (o *NodePowerUsageModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"openstack_id",
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

	varNodePowerUsageModel := _NodePowerUsageModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNodePowerUsageModel)

	if err != nil {
		return err
	}

	*o = NodePowerUsageModel(varNodePowerUsageModel)

	return err
}

type NullableNodePowerUsageModel struct {
	value *NodePowerUsageModel
	isSet bool
}

func (v NullableNodePowerUsageModel) Get() *NodePowerUsageModel {
	return v.value
}

func (v *NullableNodePowerUsageModel) Set(val *NodePowerUsageModel) {
	v.value = val
	v.isSet = true
}

func (v NullableNodePowerUsageModel) IsSet() bool {
	return v.isSet
}

func (v *NullableNodePowerUsageModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodePowerUsageModel(val *NodePowerUsageModel) *NullableNodePowerUsageModel {
	return &NullableNodePowerUsageModel{value: val, isSet: true}
}

func (v NullableNodePowerUsageModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodePowerUsageModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
