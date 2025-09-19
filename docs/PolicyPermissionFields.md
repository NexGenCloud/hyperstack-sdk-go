# PolicyPermissionFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Permission** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 

## Methods

### NewPolicyPermissionFields

`func NewPolicyPermissionFields() *PolicyPermissionFields`

NewPolicyPermissionFields instantiates a new PolicyPermissionFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyPermissionFieldsWithDefaults

`func NewPolicyPermissionFieldsWithDefaults() *PolicyPermissionFields`

NewPolicyPermissionFieldsWithDefaults instantiates a new PolicyPermissionFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PolicyPermissionFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyPermissionFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyPermissionFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyPermissionFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPermission

`func (o *PolicyPermissionFields) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *PolicyPermissionFields) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *PolicyPermissionFields) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *PolicyPermissionFields) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetResource

`func (o *PolicyPermissionFields) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *PolicyPermissionFields) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *PolicyPermissionFields) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *PolicyPermissionFields) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


