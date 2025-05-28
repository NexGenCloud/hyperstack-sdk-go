# UsersInfoListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**UsersInfo** | Pointer to [**UsersInfoFields**](UsersInfoFields.md) |  | [optional] 

## Methods

### NewUsersInfoListResponse

`func NewUsersInfoListResponse() *UsersInfoListResponse`

NewUsersInfoListResponse instantiates a new UsersInfoListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersInfoListResponseWithDefaults

`func NewUsersInfoListResponseWithDefaults() *UsersInfoListResponse`

NewUsersInfoListResponseWithDefaults instantiates a new UsersInfoListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UsersInfoListResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UsersInfoListResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UsersInfoListResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UsersInfoListResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *UsersInfoListResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UsersInfoListResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UsersInfoListResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UsersInfoListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsersInfo

`func (o *UsersInfoListResponse) GetUsersInfo() UsersInfoFields`

GetUsersInfo returns the UsersInfo field if non-nil, zero value otherwise.

### GetUsersInfoOk

`func (o *UsersInfoListResponse) GetUsersInfoOk() (*UsersInfoFields, bool)`

GetUsersInfoOk returns a tuple with the UsersInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersInfo

`func (o *UsersInfoListResponse) SetUsersInfo(v UsersInfoFields)`

SetUsersInfo sets UsersInfo field to given value.

### HasUsersInfo

`func (o *UsersInfoListResponse) HasUsersInfo() bool`

HasUsersInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


