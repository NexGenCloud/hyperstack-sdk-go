# UserDefaultChoicesForUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**UserDefaultChoices** | Pointer to [**[]UserDefaultChoiceForUserFields**](UserDefaultChoiceForUserFields.md) |  | [optional] 

## Methods

### NewUserDefaultChoicesForUserResponse

`func NewUserDefaultChoicesForUserResponse() *UserDefaultChoicesForUserResponse`

NewUserDefaultChoicesForUserResponse instantiates a new UserDefaultChoicesForUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDefaultChoicesForUserResponseWithDefaults

`func NewUserDefaultChoicesForUserResponseWithDefaults() *UserDefaultChoicesForUserResponse`

NewUserDefaultChoicesForUserResponseWithDefaults instantiates a new UserDefaultChoicesForUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UserDefaultChoicesForUserResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UserDefaultChoicesForUserResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UserDefaultChoicesForUserResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UserDefaultChoicesForUserResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *UserDefaultChoicesForUserResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserDefaultChoicesForUserResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserDefaultChoicesForUserResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserDefaultChoicesForUserResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserDefaultChoices

`func (o *UserDefaultChoicesForUserResponse) GetUserDefaultChoices() []UserDefaultChoiceForUserFields`

GetUserDefaultChoices returns the UserDefaultChoices field if non-nil, zero value otherwise.

### GetUserDefaultChoicesOk

`func (o *UserDefaultChoicesForUserResponse) GetUserDefaultChoicesOk() (*[]UserDefaultChoiceForUserFields, bool)`

GetUserDefaultChoicesOk returns a tuple with the UserDefaultChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefaultChoices

`func (o *UserDefaultChoicesForUserResponse) SetUserDefaultChoices(v []UserDefaultChoiceForUserFields)`

SetUserDefaultChoices sets UserDefaultChoices field to given value.

### HasUserDefaultChoices

`func (o *UserDefaultChoicesForUserResponse) HasUserDefaultChoices() bool`

HasUserDefaultChoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


