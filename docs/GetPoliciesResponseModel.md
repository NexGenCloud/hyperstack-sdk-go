# GetPoliciesResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Policies** | Pointer to [**[]PolicyFields**](PolicyFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetPoliciesResponseModel

`func NewGetPoliciesResponseModel() *GetPoliciesResponseModel`

NewGetPoliciesResponseModel instantiates a new GetPoliciesResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPoliciesResponseModelWithDefaults

`func NewGetPoliciesResponseModelWithDefaults() *GetPoliciesResponseModel`

NewGetPoliciesResponseModelWithDefaults instantiates a new GetPoliciesResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetPoliciesResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetPoliciesResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetPoliciesResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetPoliciesResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPolicies

`func (o *GetPoliciesResponseModel) GetPolicies() []PolicyFields`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *GetPoliciesResponseModel) GetPoliciesOk() (*[]PolicyFields, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *GetPoliciesResponseModel) SetPolicies(v []PolicyFields)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *GetPoliciesResponseModel) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetStatus

`func (o *GetPoliciesResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetPoliciesResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetPoliciesResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetPoliciesResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


