# ProfileListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Profiles** | Pointer to [**[]ProfileFields**](ProfileFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewProfileListResponse

`func NewProfileListResponse() *ProfileListResponse`

NewProfileListResponse instantiates a new ProfileListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileListResponseWithDefaults

`func NewProfileListResponseWithDefaults() *ProfileListResponse`

NewProfileListResponseWithDefaults instantiates a new ProfileListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ProfileListResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProfileListResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProfileListResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProfileListResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetProfiles

`func (o *ProfileListResponse) GetProfiles() []ProfileFields`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *ProfileListResponse) GetProfilesOk() (*[]ProfileFields, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *ProfileListResponse) SetProfiles(v []ProfileFields)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *ProfileListResponse) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetStatus

`func (o *ProfileListResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProfileListResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProfileListResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProfileListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


