# AuthGetTokenResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to [**AccessTokenField**](AccessTokenField.md) |  | [optional] 

## Methods

### NewAuthGetTokenResponseModel

`func NewAuthGetTokenResponseModel() *AuthGetTokenResponseModel`

NewAuthGetTokenResponseModel instantiates a new AuthGetTokenResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthGetTokenResponseModelWithDefaults

`func NewAuthGetTokenResponseModelWithDefaults() *AuthGetTokenResponseModel`

NewAuthGetTokenResponseModelWithDefaults instantiates a new AuthGetTokenResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AuthGetTokenResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuthGetTokenResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuthGetTokenResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AuthGetTokenResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *AuthGetTokenResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthGetTokenResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthGetTokenResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthGetTokenResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *AuthGetTokenResponseModel) GetToken() AccessTokenField`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthGetTokenResponseModel) GetTokenOk() (*AccessTokenField, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthGetTokenResponseModel) SetToken(v AccessTokenField)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthGetTokenResponseModel) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


