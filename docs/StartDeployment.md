# StartDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployment** | Pointer to [**DeploymentFieldsForStartDeployments**](DeploymentFieldsForStartDeployments.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewStartDeployment

`func NewStartDeployment() *StartDeployment`

NewStartDeployment instantiates a new StartDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartDeploymentWithDefaults

`func NewStartDeploymentWithDefaults() *StartDeployment`

NewStartDeploymentWithDefaults instantiates a new StartDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployment

`func (o *StartDeployment) GetDeployment() DeploymentFieldsForStartDeployments`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *StartDeployment) GetDeploymentOk() (*DeploymentFieldsForStartDeployments, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *StartDeployment) SetDeployment(v DeploymentFieldsForStartDeployments)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *StartDeployment) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### GetMessage

`func (o *StartDeployment) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StartDeployment) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StartDeployment) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *StartDeployment) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *StartDeployment) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StartDeployment) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StartDeployment) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StartDeployment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


