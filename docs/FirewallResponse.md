# FirewallResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Firewall** | Pointer to [**FirewallFields**](FirewallFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewFirewallResponse

`func NewFirewallResponse() *FirewallResponse`

NewFirewallResponse instantiates a new FirewallResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallResponseWithDefaults

`func NewFirewallResponseWithDefaults() *FirewallResponse`

NewFirewallResponseWithDefaults instantiates a new FirewallResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirewall

`func (o *FirewallResponse) GetFirewall() FirewallFields`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *FirewallResponse) GetFirewallOk() (*FirewallFields, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *FirewallResponse) SetFirewall(v FirewallFields)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *FirewallResponse) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.

### GetMessage

`func (o *FirewallResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FirewallResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FirewallResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FirewallResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *FirewallResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirewallResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirewallResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FirewallResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


