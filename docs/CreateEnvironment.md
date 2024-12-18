# CreateEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the environment being created. | 
**Region** | **string** | The geographic location of the data center where the environment is being created. To learn more about regions, [**click here**](https://infrahub-doc.nexgencloud.com/docs/features/regions). | 

## Methods

### NewCreateEnvironment

`func NewCreateEnvironment(name string, region string, ) *CreateEnvironment`

NewCreateEnvironment instantiates a new CreateEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEnvironmentWithDefaults

`func NewCreateEnvironmentWithDefaults() *CreateEnvironment`

NewCreateEnvironmentWithDefaults instantiates a new CreateEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEnvironment) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *CreateEnvironment) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateEnvironment) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateEnvironment) SetRegion(v string)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


