# Deployments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployments** | Pointer to [**[]DeploymentFields**](DeploymentFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewDeployments

`func NewDeployments() *Deployments`

NewDeployments instantiates a new Deployments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentsWithDefaults

`func NewDeploymentsWithDefaults() *Deployments`

NewDeploymentsWithDefaults instantiates a new Deployments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployments

`func (o *Deployments) GetDeployments() []DeploymentFields`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *Deployments) GetDeploymentsOk() (*[]DeploymentFields, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *Deployments) SetDeployments(v []DeploymentFields)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *Deployments) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.

### GetMessage

`func (o *Deployments) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Deployments) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Deployments) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Deployments) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *Deployments) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Deployments) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Deployments) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Deployments) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


