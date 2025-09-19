# NodePowerUsageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flavors** | Pointer to **[]string** |  | [optional] 
**NexgenName** | Pointer to **string** |  | [optional] 
**OpenstackId** | **string** |  | 
**OpenstackName** | Pointer to **string** |  | [optional] 
**Organizations** | Pointer to **[]int32** |  | [optional] 
**PowerUsages** | Pointer to [**PowerUsageModel**](PowerUsageModel.md) |  | [optional] 
**Projects** | Pointer to **[]string** |  | [optional] 
**ProvisionDate** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Stocks** | Pointer to [**[]NodeStocksPayload**](NodeStocksPayload.md) |  | [optional] 
**SunsetDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewNodePowerUsageModel

`func NewNodePowerUsageModel(openstackId string, ) *NodePowerUsageModel`

NewNodePowerUsageModel instantiates a new NodePowerUsageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePowerUsageModelWithDefaults

`func NewNodePowerUsageModelWithDefaults() *NodePowerUsageModel`

NewNodePowerUsageModelWithDefaults instantiates a new NodePowerUsageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlavors

`func (o *NodePowerUsageModel) GetFlavors() []string`

GetFlavors returns the Flavors field if non-nil, zero value otherwise.

### GetFlavorsOk

`func (o *NodePowerUsageModel) GetFlavorsOk() (*[]string, bool)`

GetFlavorsOk returns a tuple with the Flavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavors

`func (o *NodePowerUsageModel) SetFlavors(v []string)`

SetFlavors sets Flavors field to given value.

### HasFlavors

`func (o *NodePowerUsageModel) HasFlavors() bool`

HasFlavors returns a boolean if a field has been set.

### GetNexgenName

`func (o *NodePowerUsageModel) GetNexgenName() string`

GetNexgenName returns the NexgenName field if non-nil, zero value otherwise.

### GetNexgenNameOk

`func (o *NodePowerUsageModel) GetNexgenNameOk() (*string, bool)`

GetNexgenNameOk returns a tuple with the NexgenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexgenName

`func (o *NodePowerUsageModel) SetNexgenName(v string)`

SetNexgenName sets NexgenName field to given value.

### HasNexgenName

`func (o *NodePowerUsageModel) HasNexgenName() bool`

HasNexgenName returns a boolean if a field has been set.

### GetOpenstackId

`func (o *NodePowerUsageModel) GetOpenstackId() string`

GetOpenstackId returns the OpenstackId field if non-nil, zero value otherwise.

### GetOpenstackIdOk

`func (o *NodePowerUsageModel) GetOpenstackIdOk() (*string, bool)`

GetOpenstackIdOk returns a tuple with the OpenstackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackId

`func (o *NodePowerUsageModel) SetOpenstackId(v string)`

SetOpenstackId sets OpenstackId field to given value.


### GetOpenstackName

`func (o *NodePowerUsageModel) GetOpenstackName() string`

GetOpenstackName returns the OpenstackName field if non-nil, zero value otherwise.

### GetOpenstackNameOk

`func (o *NodePowerUsageModel) GetOpenstackNameOk() (*string, bool)`

GetOpenstackNameOk returns a tuple with the OpenstackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackName

`func (o *NodePowerUsageModel) SetOpenstackName(v string)`

SetOpenstackName sets OpenstackName field to given value.

### HasOpenstackName

`func (o *NodePowerUsageModel) HasOpenstackName() bool`

HasOpenstackName returns a boolean if a field has been set.

### GetOrganizations

`func (o *NodePowerUsageModel) GetOrganizations() []int32`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *NodePowerUsageModel) GetOrganizationsOk() (*[]int32, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *NodePowerUsageModel) SetOrganizations(v []int32)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *NodePowerUsageModel) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetPowerUsages

`func (o *NodePowerUsageModel) GetPowerUsages() PowerUsageModel`

GetPowerUsages returns the PowerUsages field if non-nil, zero value otherwise.

### GetPowerUsagesOk

`func (o *NodePowerUsageModel) GetPowerUsagesOk() (*PowerUsageModel, bool)`

GetPowerUsagesOk returns a tuple with the PowerUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerUsages

`func (o *NodePowerUsageModel) SetPowerUsages(v PowerUsageModel)`

SetPowerUsages sets PowerUsages field to given value.

### HasPowerUsages

`func (o *NodePowerUsageModel) HasPowerUsages() bool`

HasPowerUsages returns a boolean if a field has been set.

### GetProjects

`func (o *NodePowerUsageModel) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *NodePowerUsageModel) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *NodePowerUsageModel) SetProjects(v []string)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *NodePowerUsageModel) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetProvisionDate

`func (o *NodePowerUsageModel) GetProvisionDate() time.Time`

GetProvisionDate returns the ProvisionDate field if non-nil, zero value otherwise.

### GetProvisionDateOk

`func (o *NodePowerUsageModel) GetProvisionDateOk() (*time.Time, bool)`

GetProvisionDateOk returns a tuple with the ProvisionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionDate

`func (o *NodePowerUsageModel) SetProvisionDate(v time.Time)`

SetProvisionDate sets ProvisionDate field to given value.

### HasProvisionDate

`func (o *NodePowerUsageModel) HasProvisionDate() bool`

HasProvisionDate returns a boolean if a field has been set.

### GetStatus

`func (o *NodePowerUsageModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodePowerUsageModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodePowerUsageModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NodePowerUsageModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStocks

`func (o *NodePowerUsageModel) GetStocks() []NodeStocksPayload`

GetStocks returns the Stocks field if non-nil, zero value otherwise.

### GetStocksOk

`func (o *NodePowerUsageModel) GetStocksOk() (*[]NodeStocksPayload, bool)`

GetStocksOk returns a tuple with the Stocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStocks

`func (o *NodePowerUsageModel) SetStocks(v []NodeStocksPayload)`

SetStocks sets Stocks field to given value.

### HasStocks

`func (o *NodePowerUsageModel) HasStocks() bool`

HasStocks returns a boolean if a field has been set.

### GetSunsetDate

`func (o *NodePowerUsageModel) GetSunsetDate() time.Time`

GetSunsetDate returns the SunsetDate field if non-nil, zero value otherwise.

### GetSunsetDateOk

`func (o *NodePowerUsageModel) GetSunsetDateOk() (*time.Time, bool)`

GetSunsetDateOk returns a tuple with the SunsetDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunsetDate

`func (o *NodePowerUsageModel) SetSunsetDate(v time.Time)`

SetSunsetDate sets SunsetDate field to given value.

### HasSunsetDate

`func (o *NodePowerUsageModel) HasSunsetDate() bool`

HasSunsetDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


