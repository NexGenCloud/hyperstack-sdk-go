# GetRbacRolesResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to [**[]RbacRoleFields**](RbacRoleFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetRbacRolesResponseModel

`func NewGetRbacRolesResponseModel() *GetRbacRolesResponseModel`

NewGetRbacRolesResponseModel instantiates a new GetRbacRolesResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRbacRolesResponseModelWithDefaults

`func NewGetRbacRolesResponseModelWithDefaults() *GetRbacRolesResponseModel`

NewGetRbacRolesResponseModelWithDefaults instantiates a new GetRbacRolesResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetRbacRolesResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetRbacRolesResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetRbacRolesResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetRbacRolesResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRoles

`func (o *GetRbacRolesResponseModel) GetRoles() []RbacRoleFields`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GetRbacRolesResponseModel) GetRolesOk() (*[]RbacRoleFields, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GetRbacRolesResponseModel) SetRoles(v []RbacRoleFields)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *GetRbacRolesResponseModel) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetStatus

`func (o *GetRbacRolesResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRbacRolesResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRbacRolesResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetRbacRolesResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


