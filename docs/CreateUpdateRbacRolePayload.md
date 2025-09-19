# CreateUpdateRbacRolePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the role. | 
**Name** | **string** | Name of the RBAC role. | 
**Permissions** | Pointer to **[]int32** |  | [optional] 
**Policies** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewCreateUpdateRbacRolePayload

`func NewCreateUpdateRbacRolePayload(description string, name string, ) *CreateUpdateRbacRolePayload`

NewCreateUpdateRbacRolePayload instantiates a new CreateUpdateRbacRolePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdateRbacRolePayloadWithDefaults

`func NewCreateUpdateRbacRolePayloadWithDefaults() *CreateUpdateRbacRolePayload`

NewCreateUpdateRbacRolePayloadWithDefaults instantiates a new CreateUpdateRbacRolePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateUpdateRbacRolePayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateUpdateRbacRolePayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateUpdateRbacRolePayload) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *CreateUpdateRbacRolePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUpdateRbacRolePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUpdateRbacRolePayload) SetName(v string)`

SetName sets Name field to given value.


### GetPermissions

`func (o *CreateUpdateRbacRolePayload) GetPermissions() []int32`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateUpdateRbacRolePayload) GetPermissionsOk() (*[]int32, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateUpdateRbacRolePayload) SetPermissions(v []int32)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateUpdateRbacRolePayload) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPolicies

`func (o *CreateUpdateRbacRolePayload) GetPolicies() []int32`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *CreateUpdateRbacRolePayload) GetPoliciesOk() (*[]int32, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *CreateUpdateRbacRolePayload) SetPolicies(v []int32)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *CreateUpdateRbacRolePayload) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


