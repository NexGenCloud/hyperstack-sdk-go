# CreateProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Profile** | Pointer to [**ProfileFields**](ProfileFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateProfileResponse

`func NewCreateProfileResponse() *CreateProfileResponse`

NewCreateProfileResponse instantiates a new CreateProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProfileResponseWithDefaults

`func NewCreateProfileResponseWithDefaults() *CreateProfileResponse`

NewCreateProfileResponseWithDefaults instantiates a new CreateProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CreateProfileResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateProfileResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateProfileResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateProfileResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetProfile

`func (o *CreateProfileResponse) GetProfile() ProfileFields`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *CreateProfileResponse) GetProfileOk() (*ProfileFields, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *CreateProfileResponse) SetProfile(v ProfileFields)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *CreateProfileResponse) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetStatus

`func (o *CreateProfileResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateProfileResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateProfileResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateProfileResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


