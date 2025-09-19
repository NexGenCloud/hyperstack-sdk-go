# InstanceResizePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flavor** | Pointer to [**FlavorObjectFields**](FlavorObjectFields.md) |  | [optional] 
**FlavorName** | Pointer to **string** |  | [optional] 

## Methods

### NewInstanceResizePayload

`func NewInstanceResizePayload() *InstanceResizePayload`

NewInstanceResizePayload instantiates a new InstanceResizePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceResizePayloadWithDefaults

`func NewInstanceResizePayloadWithDefaults() *InstanceResizePayload`

NewInstanceResizePayloadWithDefaults instantiates a new InstanceResizePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlavor

`func (o *InstanceResizePayload) GetFlavor() FlavorObjectFields`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *InstanceResizePayload) GetFlavorOk() (*FlavorObjectFields, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *InstanceResizePayload) SetFlavor(v FlavorObjectFields)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *InstanceResizePayload) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### GetFlavorName

`func (o *InstanceResizePayload) GetFlavorName() string`

GetFlavorName returns the FlavorName field if non-nil, zero value otherwise.

### GetFlavorNameOk

`func (o *InstanceResizePayload) GetFlavorNameOk() (*string, bool)`

GetFlavorNameOk returns a tuple with the FlavorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorName

`func (o *InstanceResizePayload) SetFlavorName(v string)`

SetFlavorName sets FlavorName field to given value.

### HasFlavorName

`func (o *InstanceResizePayload) HasFlavorName() bool`

HasFlavorName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


