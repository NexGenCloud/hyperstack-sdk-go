# VerifyApiKeyResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to [**ApiKeyVerifyFields**](ApiKeyVerifyFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewVerifyApiKeyResponseModel

`func NewVerifyApiKeyResponseModel() *VerifyApiKeyResponseModel`

NewVerifyApiKeyResponseModel instantiates a new VerifyApiKeyResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyApiKeyResponseModelWithDefaults

`func NewVerifyApiKeyResponseModelWithDefaults() *VerifyApiKeyResponseModel`

NewVerifyApiKeyResponseModelWithDefaults instantiates a new VerifyApiKeyResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *VerifyApiKeyResponseModel) GetApiKey() ApiKeyVerifyFields`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *VerifyApiKeyResponseModel) GetApiKeyOk() (*ApiKeyVerifyFields, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *VerifyApiKeyResponseModel) SetApiKey(v ApiKeyVerifyFields)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *VerifyApiKeyResponseModel) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetMessage

`func (o *VerifyApiKeyResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VerifyApiKeyResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VerifyApiKeyResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *VerifyApiKeyResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *VerifyApiKeyResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VerifyApiKeyResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VerifyApiKeyResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VerifyApiKeyResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


