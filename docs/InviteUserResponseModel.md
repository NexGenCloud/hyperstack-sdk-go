# InviteUserResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invite** | Pointer to [**InviteFields**](InviteFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewInviteUserResponseModel

`func NewInviteUserResponseModel() *InviteUserResponseModel`

NewInviteUserResponseModel instantiates a new InviteUserResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteUserResponseModelWithDefaults

`func NewInviteUserResponseModelWithDefaults() *InviteUserResponseModel`

NewInviteUserResponseModelWithDefaults instantiates a new InviteUserResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvite

`func (o *InviteUserResponseModel) GetInvite() InviteFields`

GetInvite returns the Invite field if non-nil, zero value otherwise.

### GetInviteOk

`func (o *InviteUserResponseModel) GetInviteOk() (*InviteFields, bool)`

GetInviteOk returns a tuple with the Invite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvite

`func (o *InviteUserResponseModel) SetInvite(v InviteFields)`

SetInvite sets Invite field to given value.

### HasInvite

`func (o *InviteUserResponseModel) HasInvite() bool`

HasInvite returns a boolean if a field has been set.

### GetMessage

`func (o *InviteUserResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InviteUserResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InviteUserResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InviteUserResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *InviteUserResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InviteUserResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InviteUserResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InviteUserResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


