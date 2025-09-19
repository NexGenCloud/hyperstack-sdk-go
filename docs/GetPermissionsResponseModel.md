# GetPermissionsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**[]PermissionFields**](PermissionFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetPermissionsResponseModel

`func NewGetPermissionsResponseModel() *GetPermissionsResponseModel`

NewGetPermissionsResponseModel instantiates a new GetPermissionsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPermissionsResponseModelWithDefaults

`func NewGetPermissionsResponseModelWithDefaults() *GetPermissionsResponseModel`

NewGetPermissionsResponseModelWithDefaults instantiates a new GetPermissionsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetPermissionsResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetPermissionsResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetPermissionsResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetPermissionsResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPermissions

`func (o *GetPermissionsResponseModel) GetPermissions() []PermissionFields`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GetPermissionsResponseModel) GetPermissionsOk() (*[]PermissionFields, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GetPermissionsResponseModel) SetPermissions(v []PermissionFields)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GetPermissionsResponseModel) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetStatus

`func (o *GetPermissionsResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetPermissionsResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetPermissionsResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetPermissionsResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


