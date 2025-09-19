# GetUserPermissionsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**[]UserPermissionFields**](UserPermissionFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetUserPermissionsResponseModel

`func NewGetUserPermissionsResponseModel() *GetUserPermissionsResponseModel`

NewGetUserPermissionsResponseModel instantiates a new GetUserPermissionsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserPermissionsResponseModelWithDefaults

`func NewGetUserPermissionsResponseModelWithDefaults() *GetUserPermissionsResponseModel`

NewGetUserPermissionsResponseModelWithDefaults instantiates a new GetUserPermissionsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetUserPermissionsResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetUserPermissionsResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetUserPermissionsResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetUserPermissionsResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPermissions

`func (o *GetUserPermissionsResponseModel) GetPermissions() []UserPermissionFields`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GetUserPermissionsResponseModel) GetPermissionsOk() (*[]UserPermissionFields, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GetUserPermissionsResponseModel) SetPermissions(v []UserPermissionFields)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GetUserPermissionsResponseModel) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetStatus

`func (o *GetUserPermissionsResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetUserPermissionsResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetUserPermissionsResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetUserPermissionsResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


