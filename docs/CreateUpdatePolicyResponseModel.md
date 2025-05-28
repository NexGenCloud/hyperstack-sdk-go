# CreateUpdatePolicyResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Policy** | Pointer to [**PolicyFields**](PolicyFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateUpdatePolicyResponseModel

`func NewCreateUpdatePolicyResponseModel() *CreateUpdatePolicyResponseModel`

NewCreateUpdatePolicyResponseModel instantiates a new CreateUpdatePolicyResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdatePolicyResponseModelWithDefaults

`func NewCreateUpdatePolicyResponseModelWithDefaults() *CreateUpdatePolicyResponseModel`

NewCreateUpdatePolicyResponseModelWithDefaults instantiates a new CreateUpdatePolicyResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CreateUpdatePolicyResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateUpdatePolicyResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateUpdatePolicyResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateUpdatePolicyResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPolicy

`func (o *CreateUpdatePolicyResponseModel) GetPolicy() PolicyFields`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *CreateUpdatePolicyResponseModel) GetPolicyOk() (*PolicyFields, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *CreateUpdatePolicyResponseModel) SetPolicy(v PolicyFields)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *CreateUpdatePolicyResponseModel) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetStatus

`func (o *CreateUpdatePolicyResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateUpdatePolicyResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateUpdatePolicyResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateUpdatePolicyResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


