# FutureNodeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectedProvisionDate** | **time.Time** | Date and time in the format YYYY-MM-DD HH:mm:ss | 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**NexgenName** | Pointer to **string** |  | [optional] 
**OpenstackName** | Pointer to **string** |  | [optional] 
**Stocks** | Pointer to [**[]FutureNodeStockModel**](FutureNodeStockModel.md) |  | [optional] 

## Methods

### NewFutureNodeModel

`func NewFutureNodeModel(expectedProvisionDate time.Time, ) *FutureNodeModel`

NewFutureNodeModel instantiates a new FutureNodeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFutureNodeModelWithDefaults

`func NewFutureNodeModelWithDefaults() *FutureNodeModel`

NewFutureNodeModelWithDefaults instantiates a new FutureNodeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpectedProvisionDate

`func (o *FutureNodeModel) GetExpectedProvisionDate() time.Time`

GetExpectedProvisionDate returns the ExpectedProvisionDate field if non-nil, zero value otherwise.

### GetExpectedProvisionDateOk

`func (o *FutureNodeModel) GetExpectedProvisionDateOk() (*time.Time, bool)`

GetExpectedProvisionDateOk returns a tuple with the ExpectedProvisionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedProvisionDate

`func (o *FutureNodeModel) SetExpectedProvisionDate(v time.Time)`

SetExpectedProvisionDate sets ExpectedProvisionDate field to given value.


### GetId

`func (o *FutureNodeModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FutureNodeModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FutureNodeModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FutureNodeModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNexgenName

`func (o *FutureNodeModel) GetNexgenName() string`

GetNexgenName returns the NexgenName field if non-nil, zero value otherwise.

### GetNexgenNameOk

`func (o *FutureNodeModel) GetNexgenNameOk() (*string, bool)`

GetNexgenNameOk returns a tuple with the NexgenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexgenName

`func (o *FutureNodeModel) SetNexgenName(v string)`

SetNexgenName sets NexgenName field to given value.

### HasNexgenName

`func (o *FutureNodeModel) HasNexgenName() bool`

HasNexgenName returns a boolean if a field has been set.

### GetOpenstackName

`func (o *FutureNodeModel) GetOpenstackName() string`

GetOpenstackName returns the OpenstackName field if non-nil, zero value otherwise.

### GetOpenstackNameOk

`func (o *FutureNodeModel) GetOpenstackNameOk() (*string, bool)`

GetOpenstackNameOk returns a tuple with the OpenstackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackName

`func (o *FutureNodeModel) SetOpenstackName(v string)`

SetOpenstackName sets OpenstackName field to given value.

### HasOpenstackName

`func (o *FutureNodeModel) HasOpenstackName() bool`

HasOpenstackName returns a boolean if a field has been set.

### GetStocks

`func (o *FutureNodeModel) GetStocks() []FutureNodeStockModel`

GetStocks returns the Stocks field if non-nil, zero value otherwise.

### GetStocksOk

`func (o *FutureNodeModel) GetStocksOk() (*[]FutureNodeStockModel, bool)`

GetStocksOk returns a tuple with the Stocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStocks

`func (o *FutureNodeModel) SetStocks(v []FutureNodeStockModel)`

SetStocks sets Stocks field to given value.

### HasStocks

`func (o *FutureNodeModel) HasStocks() bool`

HasStocks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


