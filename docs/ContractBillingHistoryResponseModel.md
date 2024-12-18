# ContractBillingHistoryResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingHistoryContract** | Pointer to [**ContractBillingHistory**](ContractBillingHistory.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewContractBillingHistoryResponseModel

`func NewContractBillingHistoryResponseModel() *ContractBillingHistoryResponseModel`

NewContractBillingHistoryResponseModel instantiates a new ContractBillingHistoryResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractBillingHistoryResponseModelWithDefaults

`func NewContractBillingHistoryResponseModelWithDefaults() *ContractBillingHistoryResponseModel`

NewContractBillingHistoryResponseModelWithDefaults instantiates a new ContractBillingHistoryResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingHistoryContract

`func (o *ContractBillingHistoryResponseModel) GetBillingHistoryContract() ContractBillingHistory`

GetBillingHistoryContract returns the BillingHistoryContract field if non-nil, zero value otherwise.

### GetBillingHistoryContractOk

`func (o *ContractBillingHistoryResponseModel) GetBillingHistoryContractOk() (*ContractBillingHistory, bool)`

GetBillingHistoryContractOk returns a tuple with the BillingHistoryContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingHistoryContract

`func (o *ContractBillingHistoryResponseModel) SetBillingHistoryContract(v ContractBillingHistory)`

SetBillingHistoryContract sets BillingHistoryContract field to given value.

### HasBillingHistoryContract

`func (o *ContractBillingHistoryResponseModel) HasBillingHistoryContract() bool`

HasBillingHistoryContract returns a boolean if a field has been set.

### GetMessage

`func (o *ContractBillingHistoryResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ContractBillingHistoryResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ContractBillingHistoryResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ContractBillingHistoryResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ContractBillingHistoryResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContractBillingHistoryResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContractBillingHistoryResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ContractBillingHistoryResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


