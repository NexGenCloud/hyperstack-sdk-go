# AddUserInfoSuccessResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**UsersInfoFields**](UsersInfoFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddUserInfoSuccessResponseModel

`func NewAddUserInfoSuccessResponseModel() *AddUserInfoSuccessResponseModel`

NewAddUserInfoSuccessResponseModel instantiates a new AddUserInfoSuccessResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUserInfoSuccessResponseModelWithDefaults

`func NewAddUserInfoSuccessResponseModelWithDefaults() *AddUserInfoSuccessResponseModel`

NewAddUserInfoSuccessResponseModelWithDefaults instantiates a new AddUserInfoSuccessResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AddUserInfoSuccessResponseModel) GetData() UsersInfoFields`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AddUserInfoSuccessResponseModel) GetDataOk() (*UsersInfoFields, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AddUserInfoSuccessResponseModel) SetData(v UsersInfoFields)`

SetData sets Data field to given value.

### HasData

`func (o *AddUserInfoSuccessResponseModel) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *AddUserInfoSuccessResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AddUserInfoSuccessResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AddUserInfoSuccessResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AddUserInfoSuccessResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *AddUserInfoSuccessResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddUserInfoSuccessResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddUserInfoSuccessResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddUserInfoSuccessResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


