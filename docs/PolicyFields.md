# PolicyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**[]PolicyPermissionFields**](PolicyPermissionFields.md) |  | [optional] 

## Methods

### NewPolicyFields

`func NewPolicyFields() *PolicyFields`

NewPolicyFields instantiates a new PolicyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyFieldsWithDefaults

`func NewPolicyFieldsWithDefaults() *PolicyFields`

NewPolicyFieldsWithDefaults instantiates a new PolicyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PolicyFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PolicyFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PolicyFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PolicyFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyFields) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyFields) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *PolicyFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PolicyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *PolicyFields) GetPermissions() []PolicyPermissionFields`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PolicyFields) GetPermissionsOk() (*[]PolicyPermissionFields, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PolicyFields) SetPermissions(v []PolicyPermissionFields)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PolicyFields) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


