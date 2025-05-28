# AuthRequestLoginResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**AuthRequestLoginFields**](AuthRequestLoginFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewAuthRequestLoginResponseModel

`func NewAuthRequestLoginResponseModel() *AuthRequestLoginResponseModel`

NewAuthRequestLoginResponseModel instantiates a new AuthRequestLoginResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthRequestLoginResponseModelWithDefaults

`func NewAuthRequestLoginResponseModelWithDefaults() *AuthRequestLoginResponseModel`

NewAuthRequestLoginResponseModelWithDefaults instantiates a new AuthRequestLoginResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AuthRequestLoginResponseModel) GetData() AuthRequestLoginFields`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuthRequestLoginResponseModel) GetDataOk() (*AuthRequestLoginFields, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuthRequestLoginResponseModel) SetData(v AuthRequestLoginFields)`

SetData sets Data field to given value.

### HasData

`func (o *AuthRequestLoginResponseModel) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *AuthRequestLoginResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuthRequestLoginResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuthRequestLoginResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AuthRequestLoginResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *AuthRequestLoginResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthRequestLoginResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthRequestLoginResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthRequestLoginResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


