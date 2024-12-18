# OverviewInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to [**ContainerOverviewFields**](ContainerOverviewFields.md) |  | [optional] 
**Instance** | Pointer to [**InstanceOverviewFields**](InstanceOverviewFields.md) |  | [optional] 
**Volume** | Pointer to [**VolumeOverviewFields**](VolumeOverviewFields.md) |  | [optional] 

## Methods

### NewOverviewInfo

`func NewOverviewInfo() *OverviewInfo`

NewOverviewInfo instantiates a new OverviewInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewInfoWithDefaults

`func NewOverviewInfoWithDefaults() *OverviewInfo`

NewOverviewInfoWithDefaults instantiates a new OverviewInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *OverviewInfo) GetContainer() ContainerOverviewFields`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *OverviewInfo) GetContainerOk() (*ContainerOverviewFields, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *OverviewInfo) SetContainer(v ContainerOverviewFields)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *OverviewInfo) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetInstance

`func (o *OverviewInfo) GetInstance() InstanceOverviewFields`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *OverviewInfo) GetInstanceOk() (*InstanceOverviewFields, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *OverviewInfo) SetInstance(v InstanceOverviewFields)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *OverviewInfo) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetVolume

`func (o *OverviewInfo) GetVolume() VolumeOverviewFields`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *OverviewInfo) GetVolumeOk() (*VolumeOverviewFields, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *OverviewInfo) SetVolume(v VolumeOverviewFields)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *OverviewInfo) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


