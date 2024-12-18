# CustomerPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**PlanType** | **string** |  | 

## Methods

### NewCustomerPayload

`func NewCustomerPayload(id int32, planType string, ) *CustomerPayload`

NewCustomerPayload instantiates a new CustomerPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPayloadWithDefaults

`func NewCustomerPayloadWithDefaults() *CustomerPayload`

NewCustomerPayloadWithDefaults instantiates a new CustomerPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerPayload) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerPayload) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerPayload) SetId(v int32)`

SetId sets Id field to given value.


### GetPlanType

`func (o *CustomerPayload) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *CustomerPayload) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *CustomerPayload) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


