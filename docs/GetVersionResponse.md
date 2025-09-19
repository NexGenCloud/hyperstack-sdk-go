# GetVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewGetVersionResponse

`func NewGetVersionResponse() *GetVersionResponse`

NewGetVersionResponse instantiates a new GetVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVersionResponseWithDefaults

`func NewGetVersionResponseWithDefaults() *GetVersionResponse`

NewGetVersionResponseWithDefaults instantiates a new GetVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetVersionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetVersionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetVersionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetVersionResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *GetVersionResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetVersionResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetVersionResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetVersionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *GetVersionResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetVersionResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetVersionResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetVersionResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


