# AddUpdateFlavorOrganizationPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | **int32** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Disk** | **int32** |  | 
**Ephemeral** | Pointer to **int32** |  | [optional] 
**GpuCount** | **int32** |  | 
**GpuId** | **int32** |  | 
**IsPublic** | **bool** |  | 
**Labels** | Pointer to **[]string** |  | [optional] 
**Name** | **string** |  | 
**OpenstackId** | **string** |  | 
**Organizations** | **[]int32** |  | 
**Ram** | **float32** |  | 
**RegionId** | **int32** |  | 

## Methods

### NewAddUpdateFlavorOrganizationPayload

`func NewAddUpdateFlavorOrganizationPayload(cpu int32, disk int32, gpuCount int32, gpuId int32, isPublic bool, name string, openstackId string, organizations []int32, ram float32, regionId int32, ) *AddUpdateFlavorOrganizationPayload`

NewAddUpdateFlavorOrganizationPayload instantiates a new AddUpdateFlavorOrganizationPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUpdateFlavorOrganizationPayloadWithDefaults

`func NewAddUpdateFlavorOrganizationPayloadWithDefaults() *AddUpdateFlavorOrganizationPayload`

NewAddUpdateFlavorOrganizationPayloadWithDefaults instantiates a new AddUpdateFlavorOrganizationPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *AddUpdateFlavorOrganizationPayload) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *AddUpdateFlavorOrganizationPayload) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *AddUpdateFlavorOrganizationPayload) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetDescription

`func (o *AddUpdateFlavorOrganizationPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddUpdateFlavorOrganizationPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddUpdateFlavorOrganizationPayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddUpdateFlavorOrganizationPayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisk

`func (o *AddUpdateFlavorOrganizationPayload) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *AddUpdateFlavorOrganizationPayload) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *AddUpdateFlavorOrganizationPayload) SetDisk(v int32)`

SetDisk sets Disk field to given value.


### GetEphemeral

`func (o *AddUpdateFlavorOrganizationPayload) GetEphemeral() int32`

GetEphemeral returns the Ephemeral field if non-nil, zero value otherwise.

### GetEphemeralOk

`func (o *AddUpdateFlavorOrganizationPayload) GetEphemeralOk() (*int32, bool)`

GetEphemeralOk returns a tuple with the Ephemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeral

`func (o *AddUpdateFlavorOrganizationPayload) SetEphemeral(v int32)`

SetEphemeral sets Ephemeral field to given value.

### HasEphemeral

`func (o *AddUpdateFlavorOrganizationPayload) HasEphemeral() bool`

HasEphemeral returns a boolean if a field has been set.

### GetGpuCount

`func (o *AddUpdateFlavorOrganizationPayload) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *AddUpdateFlavorOrganizationPayload) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *AddUpdateFlavorOrganizationPayload) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.


### GetGpuId

`func (o *AddUpdateFlavorOrganizationPayload) GetGpuId() int32`

GetGpuId returns the GpuId field if non-nil, zero value otherwise.

### GetGpuIdOk

`func (o *AddUpdateFlavorOrganizationPayload) GetGpuIdOk() (*int32, bool)`

GetGpuIdOk returns a tuple with the GpuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuId

`func (o *AddUpdateFlavorOrganizationPayload) SetGpuId(v int32)`

SetGpuId sets GpuId field to given value.


### GetIsPublic

`func (o *AddUpdateFlavorOrganizationPayload) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *AddUpdateFlavorOrganizationPayload) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *AddUpdateFlavorOrganizationPayload) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetLabels

`func (o *AddUpdateFlavorOrganizationPayload) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddUpdateFlavorOrganizationPayload) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddUpdateFlavorOrganizationPayload) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddUpdateFlavorOrganizationPayload) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *AddUpdateFlavorOrganizationPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddUpdateFlavorOrganizationPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddUpdateFlavorOrganizationPayload) SetName(v string)`

SetName sets Name field to given value.


### GetOpenstackId

`func (o *AddUpdateFlavorOrganizationPayload) GetOpenstackId() string`

GetOpenstackId returns the OpenstackId field if non-nil, zero value otherwise.

### GetOpenstackIdOk

`func (o *AddUpdateFlavorOrganizationPayload) GetOpenstackIdOk() (*string, bool)`

GetOpenstackIdOk returns a tuple with the OpenstackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackId

`func (o *AddUpdateFlavorOrganizationPayload) SetOpenstackId(v string)`

SetOpenstackId sets OpenstackId field to given value.


### GetOrganizations

`func (o *AddUpdateFlavorOrganizationPayload) GetOrganizations() []int32`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *AddUpdateFlavorOrganizationPayload) GetOrganizationsOk() (*[]int32, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *AddUpdateFlavorOrganizationPayload) SetOrganizations(v []int32)`

SetOrganizations sets Organizations field to given value.


### GetRam

`func (o *AddUpdateFlavorOrganizationPayload) GetRam() float32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *AddUpdateFlavorOrganizationPayload) GetRamOk() (*float32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *AddUpdateFlavorOrganizationPayload) SetRam(v float32)`

SetRam sets Ram field to given value.


### GetRegionId

`func (o *AddUpdateFlavorOrganizationPayload) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *AddUpdateFlavorOrganizationPayload) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *AddUpdateFlavorOrganizationPayload) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


