# CustomerContractDetailResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contracts** | Pointer to [**CustomerContractFields**](CustomerContractFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustomerContractDetailResponseModel

`func NewCustomerContractDetailResponseModel() *CustomerContractDetailResponseModel`

NewCustomerContractDetailResponseModel instantiates a new CustomerContractDetailResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerContractDetailResponseModelWithDefaults

`func NewCustomerContractDetailResponseModelWithDefaults() *CustomerContractDetailResponseModel`

NewCustomerContractDetailResponseModelWithDefaults instantiates a new CustomerContractDetailResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContracts

`func (o *CustomerContractDetailResponseModel) GetContracts() CustomerContractFields`

GetContracts returns the Contracts field if non-nil, zero value otherwise.

### GetContractsOk

`func (o *CustomerContractDetailResponseModel) GetContractsOk() (*CustomerContractFields, bool)`

GetContractsOk returns a tuple with the Contracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContracts

`func (o *CustomerContractDetailResponseModel) SetContracts(v CustomerContractFields)`

SetContracts sets Contracts field to given value.

### HasContracts

`func (o *CustomerContractDetailResponseModel) HasContracts() bool`

HasContracts returns a boolean if a field has been set.

### GetMessage

`func (o *CustomerContractDetailResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CustomerContractDetailResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CustomerContractDetailResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CustomerContractDetailResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *CustomerContractDetailResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerContractDetailResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerContractDetailResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomerContractDetailResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


