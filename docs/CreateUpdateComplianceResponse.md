# CreateUpdateComplianceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compliance** | Pointer to [**ComplianceModelFields**](ComplianceModelFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateUpdateComplianceResponse

`func NewCreateUpdateComplianceResponse() *CreateUpdateComplianceResponse`

NewCreateUpdateComplianceResponse instantiates a new CreateUpdateComplianceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdateComplianceResponseWithDefaults

`func NewCreateUpdateComplianceResponseWithDefaults() *CreateUpdateComplianceResponse`

NewCreateUpdateComplianceResponseWithDefaults instantiates a new CreateUpdateComplianceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompliance

`func (o *CreateUpdateComplianceResponse) GetCompliance() ComplianceModelFields`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *CreateUpdateComplianceResponse) GetComplianceOk() (*ComplianceModelFields, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *CreateUpdateComplianceResponse) SetCompliance(v ComplianceModelFields)`

SetCompliance sets Compliance field to given value.

### HasCompliance

`func (o *CreateUpdateComplianceResponse) HasCompliance() bool`

HasCompliance returns a boolean if a field has been set.

### GetMessage

`func (o *CreateUpdateComplianceResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateUpdateComplianceResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateUpdateComplianceResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateUpdateComplianceResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *CreateUpdateComplianceResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateUpdateComplianceResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateUpdateComplianceResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateUpdateComplianceResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


