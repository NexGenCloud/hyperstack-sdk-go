# AttachFirewallsToVMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Firewalls** | **[]int32** | Ids of the firewalls to be attached with a virtual machine. | 

## Methods

### NewAttachFirewallsToVMPayload

`func NewAttachFirewallsToVMPayload(firewalls []int32, ) *AttachFirewallsToVMPayload`

NewAttachFirewallsToVMPayload instantiates a new AttachFirewallsToVMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachFirewallsToVMPayloadWithDefaults

`func NewAttachFirewallsToVMPayloadWithDefaults() *AttachFirewallsToVMPayload`

NewAttachFirewallsToVMPayloadWithDefaults instantiates a new AttachFirewallsToVMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirewalls

`func (o *AttachFirewallsToVMPayload) GetFirewalls() []int32`

GetFirewalls returns the Firewalls field if non-nil, zero value otherwise.

### GetFirewallsOk

`func (o *AttachFirewallsToVMPayload) GetFirewallsOk() (*[]int32, bool)`

GetFirewallsOk returns a tuple with the Firewalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewalls

`func (o *AttachFirewallsToVMPayload) SetFirewalls(v []int32)`

SetFirewalls sets Firewalls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


