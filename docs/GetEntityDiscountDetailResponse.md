# GetEntityDiscountDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to [**[]DiscountPlanFields**](DiscountPlanFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**VirtualMachine** | Pointer to [**[]DiscountPlanFields**](DiscountPlanFields.md) |  | [optional] 

## Methods

### NewGetEntityDiscountDetailResponse

`func NewGetEntityDiscountDetailResponse() *GetEntityDiscountDetailResponse`

NewGetEntityDiscountDetailResponse instantiates a new GetEntityDiscountDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntityDiscountDetailResponseWithDefaults

`func NewGetEntityDiscountDetailResponseWithDefaults() *GetEntityDiscountDetailResponse`

NewGetEntityDiscountDetailResponseWithDefaults instantiates a new GetEntityDiscountDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetEntityDiscountDetailResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetEntityDiscountDetailResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetEntityDiscountDetailResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetEntityDiscountDetailResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOrganization

`func (o *GetEntityDiscountDetailResponse) GetOrganization() []DiscountPlanFields`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *GetEntityDiscountDetailResponse) GetOrganizationOk() (*[]DiscountPlanFields, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *GetEntityDiscountDetailResponse) SetOrganization(v []DiscountPlanFields)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *GetEntityDiscountDetailResponse) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetStatus

`func (o *GetEntityDiscountDetailResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEntityDiscountDetailResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEntityDiscountDetailResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetEntityDiscountDetailResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVirtualMachine

`func (o *GetEntityDiscountDetailResponse) GetVirtualMachine() []DiscountPlanFields`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *GetEntityDiscountDetailResponse) GetVirtualMachineOk() (*[]DiscountPlanFields, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *GetEntityDiscountDetailResponse) SetVirtualMachine(v []DiscountPlanFields)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *GetEntityDiscountDetailResponse) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


