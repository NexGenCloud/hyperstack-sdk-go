# BillingHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**Attributes**](Attributes.md) |  | [optional] 
**Metrics** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBillingHistory

`func NewBillingHistory() *BillingHistory`

NewBillingHistory instantiates a new BillingHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingHistoryWithDefaults

`func NewBillingHistoryWithDefaults() *BillingHistory`

NewBillingHistoryWithDefaults instantiates a new BillingHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *BillingHistory) GetAttributes() Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BillingHistory) GetAttributesOk() (*Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BillingHistory) SetAttributes(v Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *BillingHistory) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetMetrics

`func (o *BillingHistory) GetMetrics() map[string]interface{}`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *BillingHistory) GetMetricsOk() (*map[string]interface{}, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *BillingHistory) SetMetrics(v map[string]interface{})`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *BillingHistory) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


