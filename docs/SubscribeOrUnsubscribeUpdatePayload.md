# SubscribeOrUnsubscribeUpdatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscribe** | **bool** | &#x60;false&#x60; indicates that the user will no longer receive notifications for this specific threshold, whereas &#x60;true&#x60; signifies that the user will receive notification emails. | 

## Methods

### NewSubscribeOrUnsubscribeUpdatePayload

`func NewSubscribeOrUnsubscribeUpdatePayload(subscribe bool, ) *SubscribeOrUnsubscribeUpdatePayload`

NewSubscribeOrUnsubscribeUpdatePayload instantiates a new SubscribeOrUnsubscribeUpdatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscribeOrUnsubscribeUpdatePayloadWithDefaults

`func NewSubscribeOrUnsubscribeUpdatePayloadWithDefaults() *SubscribeOrUnsubscribeUpdatePayload`

NewSubscribeOrUnsubscribeUpdatePayloadWithDefaults instantiates a new SubscribeOrUnsubscribeUpdatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscribe

`func (o *SubscribeOrUnsubscribeUpdatePayload) GetSubscribe() bool`

GetSubscribe returns the Subscribe field if non-nil, zero value otherwise.

### GetSubscribeOk

`func (o *SubscribeOrUnsubscribeUpdatePayload) GetSubscribeOk() (*bool, bool)`

GetSubscribeOk returns a tuple with the Subscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribe

`func (o *SubscribeOrUnsubscribeUpdatePayload) SetSubscribe(v bool)`

SetSubscribe sets Subscribe field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


