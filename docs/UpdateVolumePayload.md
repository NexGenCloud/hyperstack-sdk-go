# UpdateVolumePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentName** | **string** | The name of the target environment to move the volume to. The target environment must be in the same region as the current environment. | 

## Methods

### NewUpdateVolumePayload

`func NewUpdateVolumePayload(environmentName string, ) *UpdateVolumePayload`

NewUpdateVolumePayload instantiates a new UpdateVolumePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVolumePayloadWithDefaults

`func NewUpdateVolumePayloadWithDefaults() *UpdateVolumePayload`

NewUpdateVolumePayloadWithDefaults instantiates a new UpdateVolumePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentName

`func (o *UpdateVolumePayload) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *UpdateVolumePayload) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *UpdateVolumePayload) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


