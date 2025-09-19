# UserPermissionFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Permission** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 

## Methods

### NewUserPermissionFields

`func NewUserPermissionFields() *UserPermissionFields`

NewUserPermissionFields instantiates a new UserPermissionFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPermissionFieldsWithDefaults

`func NewUserPermissionFieldsWithDefaults() *UserPermissionFields`

NewUserPermissionFieldsWithDefaults instantiates a new UserPermissionFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserPermissionFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserPermissionFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserPermissionFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserPermissionFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPermission

`func (o *UserPermissionFields) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *UserPermissionFields) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *UserPermissionFields) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *UserPermissionFields) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetResource

`func (o *UserPermissionFields) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *UserPermissionFields) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *UserPermissionFields) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *UserPermissionFields) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


