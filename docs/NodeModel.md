# NodeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flavors** | Pointer to **[]string** |  | [optional] 
**NexgenName** | Pointer to **string** |  | [optional] 
**OpenstackId** | **string** |  | 
**OpenstackName** | Pointer to **string** |  | [optional] 
**Organizations** | Pointer to **[]int32** |  | [optional] 
**Projects** | Pointer to **[]string** |  | [optional] 
**ProvisionDate** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Stocks** | Pointer to [**[]NodeStocksPayload**](NodeStocksPayload.md) |  | [optional] 
**SunsetDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewNodeModel

`func NewNodeModel(openstackId string, ) *NodeModel`

NewNodeModel instantiates a new NodeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeModelWithDefaults

`func NewNodeModelWithDefaults() *NodeModel`

NewNodeModelWithDefaults instantiates a new NodeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlavors

`func (o *NodeModel) GetFlavors() []string`

GetFlavors returns the Flavors field if non-nil, zero value otherwise.

### GetFlavorsOk

`func (o *NodeModel) GetFlavorsOk() (*[]string, bool)`

GetFlavorsOk returns a tuple with the Flavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavors

`func (o *NodeModel) SetFlavors(v []string)`

SetFlavors sets Flavors field to given value.

### HasFlavors

`func (o *NodeModel) HasFlavors() bool`

HasFlavors returns a boolean if a field has been set.

### GetNexgenName

`func (o *NodeModel) GetNexgenName() string`

GetNexgenName returns the NexgenName field if non-nil, zero value otherwise.

### GetNexgenNameOk

`func (o *NodeModel) GetNexgenNameOk() (*string, bool)`

GetNexgenNameOk returns a tuple with the NexgenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexgenName

`func (o *NodeModel) SetNexgenName(v string)`

SetNexgenName sets NexgenName field to given value.

### HasNexgenName

`func (o *NodeModel) HasNexgenName() bool`

HasNexgenName returns a boolean if a field has been set.

### GetOpenstackId

`func (o *NodeModel) GetOpenstackId() string`

GetOpenstackId returns the OpenstackId field if non-nil, zero value otherwise.

### GetOpenstackIdOk

`func (o *NodeModel) GetOpenstackIdOk() (*string, bool)`

GetOpenstackIdOk returns a tuple with the OpenstackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackId

`func (o *NodeModel) SetOpenstackId(v string)`

SetOpenstackId sets OpenstackId field to given value.


### GetOpenstackName

`func (o *NodeModel) GetOpenstackName() string`

GetOpenstackName returns the OpenstackName field if non-nil, zero value otherwise.

### GetOpenstackNameOk

`func (o *NodeModel) GetOpenstackNameOk() (*string, bool)`

GetOpenstackNameOk returns a tuple with the OpenstackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackName

`func (o *NodeModel) SetOpenstackName(v string)`

SetOpenstackName sets OpenstackName field to given value.

### HasOpenstackName

`func (o *NodeModel) HasOpenstackName() bool`

HasOpenstackName returns a boolean if a field has been set.

### GetOrganizations

`func (o *NodeModel) GetOrganizations() []int32`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *NodeModel) GetOrganizationsOk() (*[]int32, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *NodeModel) SetOrganizations(v []int32)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *NodeModel) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetProjects

`func (o *NodeModel) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *NodeModel) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *NodeModel) SetProjects(v []string)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *NodeModel) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetProvisionDate

`func (o *NodeModel) GetProvisionDate() time.Time`

GetProvisionDate returns the ProvisionDate field if non-nil, zero value otherwise.

### GetProvisionDateOk

`func (o *NodeModel) GetProvisionDateOk() (*time.Time, bool)`

GetProvisionDateOk returns a tuple with the ProvisionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionDate

`func (o *NodeModel) SetProvisionDate(v time.Time)`

SetProvisionDate sets ProvisionDate field to given value.

### HasProvisionDate

`func (o *NodeModel) HasProvisionDate() bool`

HasProvisionDate returns a boolean if a field has been set.

### GetStatus

`func (o *NodeModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NodeModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStocks

`func (o *NodeModel) GetStocks() []NodeStocksPayload`

GetStocks returns the Stocks field if non-nil, zero value otherwise.

### GetStocksOk

`func (o *NodeModel) GetStocksOk() (*[]NodeStocksPayload, bool)`

GetStocksOk returns a tuple with the Stocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStocks

`func (o *NodeModel) SetStocks(v []NodeStocksPayload)`

SetStocks sets Stocks field to given value.

### HasStocks

`func (o *NodeModel) HasStocks() bool`

HasStocks returns a boolean if a field has been set.

### GetSunsetDate

`func (o *NodeModel) GetSunsetDate() time.Time`

GetSunsetDate returns the SunsetDate field if non-nil, zero value otherwise.

### GetSunsetDateOk

`func (o *NodeModel) GetSunsetDateOk() (*time.Time, bool)`

GetSunsetDateOk returns a tuple with the SunsetDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunsetDate

`func (o *NodeModel) SetSunsetDate(v time.Time)`

SetSunsetDate sets SunsetDate field to given value.

### HasSunsetDate

`func (o *NodeModel) HasSunsetDate() bool`

HasSunsetDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


