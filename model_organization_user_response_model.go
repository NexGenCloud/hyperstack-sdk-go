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

// checks if the OrganizationUserResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationUserResponseModel{}

// OrganizationUserResponseModel struct for OrganizationUserResponseModel
type OrganizationUserResponseModel struct {
	Email     *string         `json:"email,omitempty"`
	Id        *int32          `json:"id,omitempty"`
	JoinedAt  *CustomTime     `json:"joined_at,omitempty"`
	LastLogin *CustomTime     `json:"last_login,omitempty"`
	Name      *string         `json:"name,omitempty"`
	RbacRoles []RbacRoleField `json:"rbac_roles,omitempty"`
	Role      *string         `json:"role,omitempty"`
	Sub       *string         `json:"sub,omitempty"`
	Username  *string         `json:"username,omitempty"`
}

// NewOrganizationUserResponseModel instantiates a new OrganizationUserResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationUserResponseModel() *OrganizationUserResponseModel {
	this := OrganizationUserResponseModel{}
	return &this
}

// NewOrganizationUserResponseModelWithDefaults instantiates a new OrganizationUserResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationUserResponseModelWithDefaults() *OrganizationUserResponseModel {
	this := OrganizationUserResponseModel{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OrganizationUserResponseModel) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUserResponseModel) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OrganizationUserResponseModel) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *OrganizationUserResponseModel) SetEmail(v string) {
	o.Email = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationUserResponseModel) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUserResponseModel) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationUserResponseModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OrganizationUserResponseModel) SetId(v int32) {
	o.Id = &v
}

// GetJoinedAt returns the JoinedAt field value if set, zero value otherwise.
func (o *OrganizationUserResponseModel) GetJoinedAt() CustomTime {
	if o == nil || IsNil(o.JoinedAt) {
		var ret CustomTime
		return ret
	}
	return *o.JoinedAt
}

// GetJoinedAtOk returns a tuple with the JoinedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUserResponseModel) GetJoinedAtOk() (*CustomTime, bool) {
	if o == nil || IsNil(o.JoinedAt) {
		return nil, false
	}
	return o.JoinedAt, true
}

// HasJoinedAt returns a boolean if a field has been set.
func (o *OrganizationUserResponseModel) HasJoinedAt() bool {
	if o != nil && !IsNil(o.JoinedAt) {
		return true
	}

	return false
}

// SetJoinedAt gets a reference to the given CustomTime and assigns it to the JoinedAt field.
func (o *OrganizationUserResponseModel) SetJoinedAt(v CustomTime) {
	o.JoinedAt = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise.
func (o *OrganizationUserResponseModel) GetLastLogin() CustomTime {
	if o == nil || IsNil(o.LastLogin) {
		var ret CustomTime
		return ret
	}
	return *o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUserResponseModel) GetLastLoginOk() (*CustomTime, bool) {
	if o == nil || IsNil(o.LastLogin) {
		return nil, false
	}
	return o.LastLogin, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *OrganizationUserResponseModel) HasLastLogin() bool {
	if o != nil && !IsNil(o.LastLogin) {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given CustomTime and assigns it to the LastLogin field.
func (o *OrganizationUserResponseModel) SetLastLogin(v CustomTime) {
	o.LastLogin = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationUserResponseModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUserResponseModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationUserResponseModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationUserResponseModel) SetName(v string) {
	o.Name = &v
}

// GetRbacRoles returns the RbacRoles field value if set, zero value otherwise.
func (o *OrganizationUserResponseModel) GetRbacRoles() []RbacRoleField {
	if o == nil || IsNil(o.RbacRoles) {
		var ret []RbacRoleField
		return ret
	}
	return o.RbacRoles
}

// GetRbacRolesOk returns a tuple with the RbacRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUserResponseModel) GetRbacRolesOk() ([]RbacRoleField, bool) {
	if o == nil || IsNil(o.RbacRoles) {
		return nil, false
	}
	return o.RbacRoles, true
}

// HasRbacRoles returns a boolean if a field has been set.
func (o *OrganizationUserResponseModel) HasRbacRoles() bool {
	if o != nil && !IsNil(o.RbacRoles) {
		return true
	}

	return false
}

// SetRbacRoles gets a reference to the given []RbacRoleField and assigns it to the RbacRoles field.
func (o *OrganizationUserResponseModel) SetRbacRoles(v []RbacRoleField) {
	o.RbacRoles = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *OrganizationUserResponseModel) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUserResponseModel) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *OrganizationUserResponseModel) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *OrganizationUserResponseModel) SetRole(v string) {
	o.Role = &v
}

// GetSub returns the Sub field value if set, zero value otherwise.
func (o *OrganizationUserResponseModel) GetSub() string {
	if o == nil || IsNil(o.Sub) {
		var ret string
		return ret
	}
	return *o.Sub
}

// GetSubOk returns a tuple with the Sub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUserResponseModel) GetSubOk() (*string, bool) {
	if o == nil || IsNil(o.Sub) {
		return nil, false
	}
	return o.Sub, true
}

// HasSub returns a boolean if a field has been set.
func (o *OrganizationUserResponseModel) HasSub() bool {
	if o != nil && !IsNil(o.Sub) {
		return true
	}

	return false
}

// SetSub gets a reference to the given string and assigns it to the Sub field.
func (o *OrganizationUserResponseModel) SetSub(v string) {
	o.Sub = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *OrganizationUserResponseModel) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUserResponseModel) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *OrganizationUserResponseModel) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *OrganizationUserResponseModel) SetUsername(v string) {
	o.Username = &v
}

func (o OrganizationUserResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationUserResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.JoinedAt) {
		toSerialize["joined_at"] = o.JoinedAt
	}
	if !IsNil(o.LastLogin) {
		toSerialize["last_login"] = o.LastLogin
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RbacRoles) {
		toSerialize["rbac_roles"] = o.RbacRoles
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Sub) {
		toSerialize["sub"] = o.Sub
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableOrganizationUserResponseModel struct {
	value *OrganizationUserResponseModel
	isSet bool
}

func (v NullableOrganizationUserResponseModel) Get() *OrganizationUserResponseModel {
	return v.value
}

func (v *NullableOrganizationUserResponseModel) Set(val *OrganizationUserResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationUserResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationUserResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationUserResponseModel(val *OrganizationUserResponseModel) *NullableOrganizationUserResponseModel {
	return &NullableOrganizationUserResponseModel{value: val, isSet: true}
}

func (v NullableOrganizationUserResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationUserResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
