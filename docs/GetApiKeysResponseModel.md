# GetApiKeysResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeys** | Pointer to [**[]ApiKeyFields**](ApiKeyFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetApiKeysResponseModel

`func NewGetApiKeysResponseModel() *GetApiKeysResponseModel`

NewGetApiKeysResponseModel instantiates a new GetApiKeysResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiKeysResponseModelWithDefaults

`func NewGetApiKeysResponseModelWithDefaults() *GetApiKeysResponseModel`

NewGetApiKeysResponseModelWithDefaults instantiates a new GetApiKeysResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeys

`func (o *GetApiKeysResponseModel) GetApiKeys() []ApiKeyFields`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *GetApiKeysResponseModel) GetApiKeysOk() (*[]ApiKeyFields, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *GetApiKeysResponseModel) SetApiKeys(v []ApiKeyFields)`

SetApiKeys sets ApiKeys field to given value.

### HasApiKeys

`func (o *GetApiKeysResponseModel) HasApiKeys() bool`

HasApiKeys returns a boolean if a field has been set.

### GetMessage

`func (o *GetApiKeysResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetApiKeysResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetApiKeysResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetApiKeysResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *GetApiKeysResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetApiKeysResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetApiKeysResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetApiKeysResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


