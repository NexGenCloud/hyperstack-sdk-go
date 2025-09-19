# GetCustomerContractsListResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contracts** | Pointer to [**[]CustomerContractFields**](CustomerContractFields.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetCustomerContractsListResponseModel

`func NewGetCustomerContractsListResponseModel() *GetCustomerContractsListResponseModel`

NewGetCustomerContractsListResponseModel instantiates a new GetCustomerContractsListResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomerContractsListResponseModelWithDefaults

`func NewGetCustomerContractsListResponseModelWithDefaults() *GetCustomerContractsListResponseModel`

NewGetCustomerContractsListResponseModelWithDefaults instantiates a new GetCustomerContractsListResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContracts

`func (o *GetCustomerContractsListResponseModel) GetContracts() []CustomerContractFields`

GetContracts returns the Contracts field if non-nil, zero value otherwise.

### GetContractsOk

`func (o *GetCustomerContractsListResponseModel) GetContractsOk() (*[]CustomerContractFields, bool)`

GetContractsOk returns a tuple with the Contracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContracts

`func (o *GetCustomerContractsListResponseModel) SetContracts(v []CustomerContractFields)`

SetContracts sets Contracts field to given value.

### HasContracts

`func (o *GetCustomerContractsListResponseModel) HasContracts() bool`

HasContracts returns a boolean if a field has been set.

### GetCount

`func (o *GetCustomerContractsListResponseModel) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetCustomerContractsListResponseModel) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetCustomerContractsListResponseModel) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetCustomerContractsListResponseModel) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetMessage

`func (o *GetCustomerContractsListResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetCustomerContractsListResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetCustomerContractsListResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetCustomerContractsListResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *GetCustomerContractsListResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCustomerContractsListResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCustomerContractsListResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetCustomerContractsListResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


