# FutureNodeUpdateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectedProvisionDate** | **time.Time** |  | 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**NexgenName** | **string** |  | 
**OpenstackName** | **string** |  | 
**RegionId** | Pointer to **int32** |  | [optional] 

## Methods

### NewFutureNodeUpdateModel

`func NewFutureNodeUpdateModel(expectedProvisionDate time.Time, nexgenName string, openstackName string, ) *FutureNodeUpdateModel`

NewFutureNodeUpdateModel instantiates a new FutureNodeUpdateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFutureNodeUpdateModelWithDefaults

`func NewFutureNodeUpdateModelWithDefaults() *FutureNodeUpdateModel`

NewFutureNodeUpdateModelWithDefaults instantiates a new FutureNodeUpdateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpectedProvisionDate

`func (o *FutureNodeUpdateModel) GetExpectedProvisionDate() time.Time`

GetExpectedProvisionDate returns the ExpectedProvisionDate field if non-nil, zero value otherwise.

### GetExpectedProvisionDateOk

`func (o *FutureNodeUpdateModel) GetExpectedProvisionDateOk() (*time.Time, bool)`

GetExpectedProvisionDateOk returns a tuple with the ExpectedProvisionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedProvisionDate

`func (o *FutureNodeUpdateModel) SetExpectedProvisionDate(v time.Time)`

SetExpectedProvisionDate sets ExpectedProvisionDate field to given value.


### GetId

`func (o *FutureNodeUpdateModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FutureNodeUpdateModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FutureNodeUpdateModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FutureNodeUpdateModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNexgenName

`func (o *FutureNodeUpdateModel) GetNexgenName() string`

GetNexgenName returns the NexgenName field if non-nil, zero value otherwise.

### GetNexgenNameOk

`func (o *FutureNodeUpdateModel) GetNexgenNameOk() (*string, bool)`

GetNexgenNameOk returns a tuple with the NexgenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexgenName

`func (o *FutureNodeUpdateModel) SetNexgenName(v string)`

SetNexgenName sets NexgenName field to given value.


### GetOpenstackName

`func (o *FutureNodeUpdateModel) GetOpenstackName() string`

GetOpenstackName returns the OpenstackName field if non-nil, zero value otherwise.

### GetOpenstackNameOk

`func (o *FutureNodeUpdateModel) GetOpenstackNameOk() (*string, bool)`

GetOpenstackNameOk returns a tuple with the OpenstackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackName

`func (o *FutureNodeUpdateModel) SetOpenstackName(v string)`

SetOpenstackName sets OpenstackName field to given value.


### GetRegionId

`func (o *FutureNodeUpdateModel) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *FutureNodeUpdateModel) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *FutureNodeUpdateModel) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *FutureNodeUpdateModel) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


