# ContractBillingHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**ContractBillingHistoryResponseAttributes**](ContractBillingHistoryResponseAttributes.md) |  | [optional] 
**Metrics** | Pointer to [**ContractlBillingHistoryResponseMetrics**](ContractlBillingHistoryResponseMetrics.md) |  | [optional] 

## Methods

### NewContractBillingHistory

`func NewContractBillingHistory() *ContractBillingHistory`

NewContractBillingHistory instantiates a new ContractBillingHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractBillingHistoryWithDefaults

`func NewContractBillingHistoryWithDefaults() *ContractBillingHistory`

NewContractBillingHistoryWithDefaults instantiates a new ContractBillingHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *ContractBillingHistory) GetAttributes() ContractBillingHistoryResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ContractBillingHistory) GetAttributesOk() (*ContractBillingHistoryResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ContractBillingHistory) SetAttributes(v ContractBillingHistoryResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ContractBillingHistory) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetMetrics

`func (o *ContractBillingHistory) GetMetrics() ContractlBillingHistoryResponseMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ContractBillingHistory) GetMetricsOk() (*ContractlBillingHistoryResponseMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ContractBillingHistory) SetMetrics(v ContractlBillingHistoryResponseMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *ContractBillingHistory) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


