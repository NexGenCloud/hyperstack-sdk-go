# PermissionFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Permission** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 

## Methods

### NewPermissionFields

`func NewPermissionFields() *PermissionFields`

NewPermissionFields instantiates a new PermissionFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionFieldsWithDefaults

`func NewPermissionFieldsWithDefaults() *PermissionFields`

NewPermissionFieldsWithDefaults instantiates a new PermissionFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PermissionFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PermissionFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PermissionFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PermissionFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEndpoint

`func (o *PermissionFields) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *PermissionFields) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *PermissionFields) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *PermissionFields) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetId

`func (o *PermissionFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PermissionFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PermissionFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PermissionFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMethod

`func (o *PermissionFields) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PermissionFields) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PermissionFields) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PermissionFields) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPermission

`func (o *PermissionFields) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *PermissionFields) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *PermissionFields) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *PermissionFields) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetResource

`func (o *PermissionFields) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *PermissionFields) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *PermissionFields) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *PermissionFields) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


