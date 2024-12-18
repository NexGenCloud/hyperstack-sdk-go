# CreateClusterPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentName** | **string** |  | 
**ImageName** | **string** |  | 
**KeypairName** | **string** |  | 
**KubernetesVersion** | **string** |  | 
**MasterFlavorName** | **string** |  | 
**Name** | **string** |  | 
**NodeCount** | **int32** |  | 
**NodeFlavorName** | **string** |  | 

## Methods

### NewCreateClusterPayload

`func NewCreateClusterPayload(environmentName string, imageName string, keypairName string, kubernetesVersion string, masterFlavorName string, name string, nodeCount int32, nodeFlavorName string, ) *CreateClusterPayload`

NewCreateClusterPayload instantiates a new CreateClusterPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterPayloadWithDefaults

`func NewCreateClusterPayloadWithDefaults() *CreateClusterPayload`

NewCreateClusterPayloadWithDefaults instantiates a new CreateClusterPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentName

`func (o *CreateClusterPayload) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *CreateClusterPayload) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *CreateClusterPayload) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.


### GetImageName

`func (o *CreateClusterPayload) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *CreateClusterPayload) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *CreateClusterPayload) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetKeypairName

`func (o *CreateClusterPayload) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *CreateClusterPayload) GetKeypairNameOk() (*string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairName

`func (o *CreateClusterPayload) SetKeypairName(v string)`

SetKeypairName sets KeypairName field to given value.


### GetKubernetesVersion

`func (o *CreateClusterPayload) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *CreateClusterPayload) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *CreateClusterPayload) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.


### GetMasterFlavorName

`func (o *CreateClusterPayload) GetMasterFlavorName() string`

GetMasterFlavorName returns the MasterFlavorName field if non-nil, zero value otherwise.

### GetMasterFlavorNameOk

`func (o *CreateClusterPayload) GetMasterFlavorNameOk() (*string, bool)`

GetMasterFlavorNameOk returns a tuple with the MasterFlavorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterFlavorName

`func (o *CreateClusterPayload) SetMasterFlavorName(v string)`

SetMasterFlavorName sets MasterFlavorName field to given value.


### GetName

`func (o *CreateClusterPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateClusterPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateClusterPayload) SetName(v string)`

SetName sets Name field to given value.


### GetNodeCount

`func (o *CreateClusterPayload) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *CreateClusterPayload) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *CreateClusterPayload) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.


### GetNodeFlavorName

`func (o *CreateClusterPayload) GetNodeFlavorName() string`

GetNodeFlavorName returns the NodeFlavorName field if non-nil, zero value otherwise.

### GetNodeFlavorNameOk

`func (o *CreateClusterPayload) GetNodeFlavorNameOk() (*string, bool)`

GetNodeFlavorNameOk returns a tuple with the NodeFlavorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeFlavorName

`func (o *CreateClusterPayload) SetNodeFlavorName(v string)`

SetNodeFlavorName sets NodeFlavorName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


