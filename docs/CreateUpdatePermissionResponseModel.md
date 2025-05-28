# CreateUpdatePermissionResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Permission** | Pointer to [**PermissionFields**](PermissionFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateUpdatePermissionResponseModel

`func NewCreateUpdatePermissionResponseModel() *CreateUpdatePermissionResponseModel`

NewCreateUpdatePermissionResponseModel instantiates a new CreateUpdatePermissionResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdatePermissionResponseModelWithDefaults

`func NewCreateUpdatePermissionResponseModelWithDefaults() *CreateUpdatePermissionResponseModel`

NewCreateUpdatePermissionResponseModelWithDefaults instantiates a new CreateUpdatePermissionResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CreateUpdatePermissionResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateUpdatePermissionResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateUpdatePermissionResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateUpdatePermissionResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPermission

`func (o *CreateUpdatePermissionResponseModel) GetPermission() PermissionFields`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *CreateUpdatePermissionResponseModel) GetPermissionOk() (*PermissionFields, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *CreateUpdatePermissionResponseModel) SetPermission(v PermissionFields)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *CreateUpdatePermissionResponseModel) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetStatus

`func (o *CreateUpdatePermissionResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateUpdatePermissionResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateUpdatePermissionResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateUpdatePermissionResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


