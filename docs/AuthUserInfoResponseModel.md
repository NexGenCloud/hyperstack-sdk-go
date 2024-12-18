# AuthUserInfoResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**User** | Pointer to [**AuthUserFields**](AuthUserFields.md) |  | [optional] 

## Methods

### NewAuthUserInfoResponseModel

`func NewAuthUserInfoResponseModel() *AuthUserInfoResponseModel`

NewAuthUserInfoResponseModel instantiates a new AuthUserInfoResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthUserInfoResponseModelWithDefaults

`func NewAuthUserInfoResponseModelWithDefaults() *AuthUserInfoResponseModel`

NewAuthUserInfoResponseModelWithDefaults instantiates a new AuthUserInfoResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AuthUserInfoResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuthUserInfoResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuthUserInfoResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AuthUserInfoResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *AuthUserInfoResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthUserInfoResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthUserInfoResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthUserInfoResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUser

`func (o *AuthUserInfoResponseModel) GetUser() AuthUserFields`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuthUserInfoResponseModel) GetUserOk() (*AuthUserFields, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuthUserInfoResponseModel) SetUser(v AuthUserFields)`

SetUser sets User field to given value.

### HasUser

`func (o *AuthUserInfoResponseModel) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


