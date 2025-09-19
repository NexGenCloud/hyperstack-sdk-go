# RbacRoleFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**[]RolePermissionFields**](RolePermissionFields.md) |  | [optional] 
**Policies** | Pointer to [**[]RolePolicyFields**](RolePolicyFields.md) |  | [optional] 

## Methods

### NewRbacRoleFields

`func NewRbacRoleFields() *RbacRoleFields`

NewRbacRoleFields instantiates a new RbacRoleFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRbacRoleFieldsWithDefaults

`func NewRbacRoleFieldsWithDefaults() *RbacRoleFields`

NewRbacRoleFieldsWithDefaults instantiates a new RbacRoleFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *RbacRoleFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RbacRoleFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RbacRoleFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RbacRoleFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *RbacRoleFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RbacRoleFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RbacRoleFields) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RbacRoleFields) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *RbacRoleFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RbacRoleFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RbacRoleFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RbacRoleFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RbacRoleFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RbacRoleFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RbacRoleFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RbacRoleFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *RbacRoleFields) GetPermissions() []RolePermissionFields`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RbacRoleFields) GetPermissionsOk() (*[]RolePermissionFields, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RbacRoleFields) SetPermissions(v []RolePermissionFields)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *RbacRoleFields) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPolicies

`func (o *RbacRoleFields) GetPolicies() []RolePolicyFields`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *RbacRoleFields) GetPoliciesOk() (*[]RolePolicyFields, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *RbacRoleFields) SetPolicies(v []RolePolicyFields)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *RbacRoleFields) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


