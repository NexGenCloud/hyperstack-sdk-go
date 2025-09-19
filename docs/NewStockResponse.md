# NewStockResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Models** | Pointer to [**[]NewModelResponse**](NewModelResponse.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**StockType** | Pointer to **string** |  | [optional] 

## Methods

### NewNewStockResponse

`func NewNewStockResponse() *NewStockResponse`

NewNewStockResponse instantiates a new NewStockResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewStockResponseWithDefaults

`func NewNewStockResponseWithDefaults() *NewStockResponse`

NewNewStockResponseWithDefaults instantiates a new NewStockResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModels

`func (o *NewStockResponse) GetModels() []NewModelResponse`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *NewStockResponse) GetModelsOk() (*[]NewModelResponse, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *NewStockResponse) SetModels(v []NewModelResponse)`

SetModels sets Models field to given value.

### HasModels

`func (o *NewStockResponse) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetRegion

`func (o *NewStockResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NewStockResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NewStockResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NewStockResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStockType

`func (o *NewStockResponse) GetStockType() string`

GetStockType returns the StockType field if non-nil, zero value otherwise.

### GetStockTypeOk

`func (o *NewStockResponse) GetStockTypeOk() (*string, bool)`

GetStockTypeOk returns a tuple with the StockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockType

`func (o *NewStockResponse) SetStockType(v string)`

SetStockType sets StockType field to given value.

### HasStockType

`func (o *NewStockResponse) HasStockType() bool`

HasStockType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


