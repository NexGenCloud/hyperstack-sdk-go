# OrganizationFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Credit** | Pointer to **int32** |  | [optional] 
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Threshold** | Pointer to **int32** |  | [optional] 
**TotalClusters** | Pointer to **int32** |  | [optional] 
**TotalContainers** | Pointer to **int32** |  | [optional] 
**TotalInstances** | Pointer to **int32** |  | [optional] 
**TotalVolumes** | Pointer to **int32** |  | [optional] 
**Users** | Pointer to [**[]OrganizationUserResponseModel**](OrganizationUserResponseModel.md) |  | [optional] 

## Methods

### NewOrganizationFields

`func NewOrganizationFields(id int32, name string, ) *OrganizationFields`

NewOrganizationFields instantiates a new OrganizationFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationFieldsWithDefaults

`func NewOrganizationFieldsWithDefaults() *OrganizationFields`

NewOrganizationFieldsWithDefaults instantiates a new OrganizationFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *OrganizationFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrganizationFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredit

`func (o *OrganizationFields) GetCredit() int32`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *OrganizationFields) GetCreditOk() (*int32, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *OrganizationFields) SetCredit(v int32)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *OrganizationFields) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetId

`func (o *OrganizationFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationFields) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *OrganizationFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationFields) SetName(v string)`

SetName sets Name field to given value.


### GetThreshold

`func (o *OrganizationFields) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *OrganizationFields) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *OrganizationFields) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *OrganizationFields) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetTotalClusters

`func (o *OrganizationFields) GetTotalClusters() int32`

GetTotalClusters returns the TotalClusters field if non-nil, zero value otherwise.

### GetTotalClustersOk

`func (o *OrganizationFields) GetTotalClustersOk() (*int32, bool)`

GetTotalClustersOk returns a tuple with the TotalClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalClusters

`func (o *OrganizationFields) SetTotalClusters(v int32)`

SetTotalClusters sets TotalClusters field to given value.

### HasTotalClusters

`func (o *OrganizationFields) HasTotalClusters() bool`

HasTotalClusters returns a boolean if a field has been set.

### GetTotalContainers

`func (o *OrganizationFields) GetTotalContainers() int32`

GetTotalContainers returns the TotalContainers field if non-nil, zero value otherwise.

### GetTotalContainersOk

`func (o *OrganizationFields) GetTotalContainersOk() (*int32, bool)`

GetTotalContainersOk returns a tuple with the TotalContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalContainers

`func (o *OrganizationFields) SetTotalContainers(v int32)`

SetTotalContainers sets TotalContainers field to given value.

### HasTotalContainers

`func (o *OrganizationFields) HasTotalContainers() bool`

HasTotalContainers returns a boolean if a field has been set.

### GetTotalInstances

`func (o *OrganizationFields) GetTotalInstances() int32`

GetTotalInstances returns the TotalInstances field if non-nil, zero value otherwise.

### GetTotalInstancesOk

`func (o *OrganizationFields) GetTotalInstancesOk() (*int32, bool)`

GetTotalInstancesOk returns a tuple with the TotalInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInstances

`func (o *OrganizationFields) SetTotalInstances(v int32)`

SetTotalInstances sets TotalInstances field to given value.

### HasTotalInstances

`func (o *OrganizationFields) HasTotalInstances() bool`

HasTotalInstances returns a boolean if a field has been set.

### GetTotalVolumes

`func (o *OrganizationFields) GetTotalVolumes() int32`

GetTotalVolumes returns the TotalVolumes field if non-nil, zero value otherwise.

### GetTotalVolumesOk

`func (o *OrganizationFields) GetTotalVolumesOk() (*int32, bool)`

GetTotalVolumesOk returns a tuple with the TotalVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolumes

`func (o *OrganizationFields) SetTotalVolumes(v int32)`

SetTotalVolumes sets TotalVolumes field to given value.

### HasTotalVolumes

`func (o *OrganizationFields) HasTotalVolumes() bool`

HasTotalVolumes returns a boolean if a field has been set.

### GetUsers

`func (o *OrganizationFields) GetUsers() []OrganizationUserResponseModel`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *OrganizationFields) GetUsersOk() (*[]OrganizationUserResponseModel, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *OrganizationFields) SetUsers(v []OrganizationUserResponseModel)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *OrganizationFields) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


