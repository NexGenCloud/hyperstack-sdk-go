# GetAllContractsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contracts** | Pointer to [**[]GetAllContractFields**](GetAllContractFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetAllContractsResponseModel

`func NewGetAllContractsResponseModel() *GetAllContractsResponseModel`

NewGetAllContractsResponseModel instantiates a new GetAllContractsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllContractsResponseModelWithDefaults

`func NewGetAllContractsResponseModelWithDefaults() *GetAllContractsResponseModel`

NewGetAllContractsResponseModelWithDefaults instantiates a new GetAllContractsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContracts

`func (o *GetAllContractsResponseModel) GetContracts() []GetAllContractFields`

GetContracts returns the Contracts field if non-nil, zero value otherwise.

### GetContractsOk

`func (o *GetAllContractsResponseModel) GetContractsOk() (*[]GetAllContractFields, bool)`

GetContractsOk returns a tuple with the Contracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContracts

`func (o *GetAllContractsResponseModel) SetContracts(v []GetAllContractFields)`

SetContracts sets Contracts field to given value.

### HasContracts

`func (o *GetAllContractsResponseModel) HasContracts() bool`

HasContracts returns a boolean if a field has been set.

### GetMessage

`func (o *GetAllContractsResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetAllContractsResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetAllContractsResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetAllContractsResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *GetAllContractsResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAllContractsResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAllContractsResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetAllContractsResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


