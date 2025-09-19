# CreateVolumePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUrl** | Pointer to **string** | A URL that can be attached to the volume you are creating. This &#x60;callback_url&#x60; will post any action events that occur to your volume to the provided URL. | [optional] 
**Description** | Pointer to **string** | A brief description or comment about the volume. | [optional] 
**EnvironmentName** | **string** | The name of the [environment](https://docs.hyperstack.cloud/docs/api-reference/core-resources/environments/) within which the volume is being created. | 
**ImageId** | Pointer to **int32** | The ID of the operating system image that will be associated with the volume. By providing an &#x60;image_id&#x60; in the create volume request, you will create a bootable volume. | [optional] 
**Name** | **string** | The name of the volume being created. | 
**Size** | **int32** | The size of the volume in GB. 1048576GB storage capacity per volume. | 
**VolumeType** | **string** | Specifies the type of volume being created, which determines the storage technology it will use. Call the [List volume types](https://infrahub-api-doc.nexgencloud.com/#get-/core/volumes) endpoint to retrieve a list of available volume model types. | 

## Methods

### NewCreateVolumePayload

`func NewCreateVolumePayload(environmentName string, name string, size int32, volumeType string, ) *CreateVolumePayload`

NewCreateVolumePayload instantiates a new CreateVolumePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVolumePayloadWithDefaults

`func NewCreateVolumePayloadWithDefaults() *CreateVolumePayload`

NewCreateVolumePayloadWithDefaults instantiates a new CreateVolumePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUrl

`func (o *CreateVolumePayload) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *CreateVolumePayload) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *CreateVolumePayload) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *CreateVolumePayload) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetDescription

`func (o *CreateVolumePayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateVolumePayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateVolumePayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateVolumePayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentName

`func (o *CreateVolumePayload) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *CreateVolumePayload) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *CreateVolumePayload) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.


### GetImageId

`func (o *CreateVolumePayload) GetImageId() int32`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CreateVolumePayload) GetImageIdOk() (*int32, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CreateVolumePayload) SetImageId(v int32)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *CreateVolumePayload) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetName

`func (o *CreateVolumePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVolumePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVolumePayload) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *CreateVolumePayload) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateVolumePayload) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateVolumePayload) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVolumeType

`func (o *CreateVolumePayload) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *CreateVolumePayload) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *CreateVolumePayload) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


