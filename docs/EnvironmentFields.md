# EnvironmentFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Features** | Pointer to [**EnvironmentFeatures**](EnvironmentFeatures.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 

## Methods

### NewEnvironmentFields

`func NewEnvironmentFields() *EnvironmentFields`

NewEnvironmentFields instantiates a new EnvironmentFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentFieldsWithDefaults

`func NewEnvironmentFieldsWithDefaults() *EnvironmentFields`

NewEnvironmentFieldsWithDefaults instantiates a new EnvironmentFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EnvironmentFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EnvironmentFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFeatures

`func (o *EnvironmentFields) GetFeatures() EnvironmentFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *EnvironmentFields) GetFeaturesOk() (*EnvironmentFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *EnvironmentFields) SetFeatures(v EnvironmentFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *EnvironmentFields) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetId

`func (o *EnvironmentFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EnvironmentFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *EnvironmentFields) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EnvironmentFields) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EnvironmentFields) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *EnvironmentFields) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


