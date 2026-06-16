# AutoTopupStatusSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Auto top-up status: active, disabled, pending_setup, cancelled, or null | [optional] 
**ThresholdAmount** | Pointer to **float32** | Balance threshold that triggers auto top-up | [optional] 
**TopupAmount** | Pointer to **float32** | Amount to top up when threshold is reached | [optional] 

## Methods

### NewAutoTopupStatusSchema

`func NewAutoTopupStatusSchema() *AutoTopupStatusSchema`

NewAutoTopupStatusSchema instantiates a new AutoTopupStatusSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTopupStatusSchemaWithDefaults

`func NewAutoTopupStatusSchemaWithDefaults() *AutoTopupStatusSchema`

NewAutoTopupStatusSchemaWithDefaults instantiates a new AutoTopupStatusSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AutoTopupStatusSchema) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutoTopupStatusSchema) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutoTopupStatusSchema) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AutoTopupStatusSchema) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThresholdAmount

`func (o *AutoTopupStatusSchema) GetThresholdAmount() float32`

GetThresholdAmount returns the ThresholdAmount field if non-nil, zero value otherwise.

### GetThresholdAmountOk

`func (o *AutoTopupStatusSchema) GetThresholdAmountOk() (*float32, bool)`

GetThresholdAmountOk returns a tuple with the ThresholdAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdAmount

`func (o *AutoTopupStatusSchema) SetThresholdAmount(v float32)`

SetThresholdAmount sets ThresholdAmount field to given value.

### HasThresholdAmount

`func (o *AutoTopupStatusSchema) HasThresholdAmount() bool`

HasThresholdAmount returns a boolean if a field has been set.

### GetTopupAmount

`func (o *AutoTopupStatusSchema) GetTopupAmount() float32`

GetTopupAmount returns the TopupAmount field if non-nil, zero value otherwise.

### GetTopupAmountOk

`func (o *AutoTopupStatusSchema) GetTopupAmountOk() (*float32, bool)`

GetTopupAmountOk returns a tuple with the TopupAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopupAmount

`func (o *AutoTopupStatusSchema) SetTopupAmount(v float32)`

SetTopupAmount sets TopupAmount field to given value.

### HasTopupAmount

`func (o *AutoTopupStatusSchema) HasTopupAmount() bool`

HasTopupAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


