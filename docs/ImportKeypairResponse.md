# ImportKeypairResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keypair** | Pointer to [**KeypairFields**](KeypairFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewImportKeypairResponse

`func NewImportKeypairResponse() *ImportKeypairResponse`

NewImportKeypairResponse instantiates a new ImportKeypairResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportKeypairResponseWithDefaults

`func NewImportKeypairResponseWithDefaults() *ImportKeypairResponse`

NewImportKeypairResponseWithDefaults instantiates a new ImportKeypairResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeypair

`func (o *ImportKeypairResponse) GetKeypair() KeypairFields`

GetKeypair returns the Keypair field if non-nil, zero value otherwise.

### GetKeypairOk

`func (o *ImportKeypairResponse) GetKeypairOk() (*KeypairFields, bool)`

GetKeypairOk returns a tuple with the Keypair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypair

`func (o *ImportKeypairResponse) SetKeypair(v KeypairFields)`

SetKeypair sets Keypair field to given value.

### HasKeypair

`func (o *ImportKeypairResponse) HasKeypair() bool`

HasKeypair returns a boolean if a field has been set.

### GetMessage

`func (o *ImportKeypairResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ImportKeypairResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ImportKeypairResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ImportKeypairResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ImportKeypairResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImportKeypairResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImportKeypairResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ImportKeypairResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


