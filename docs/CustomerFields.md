# CustomerFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**PlanType** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomerFields

`func NewCustomerFields() *CustomerFields`

NewCustomerFields instantiates a new CustomerFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerFieldsWithDefaults

`func NewCustomerFieldsWithDefaults() *CustomerFields`

NewCustomerFieldsWithDefaults instantiates a new CustomerFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlanType

`func (o *CustomerFields) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *CustomerFields) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *CustomerFields) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.

### HasPlanType

`func (o *CustomerFields) HasPlanType() bool`

HasPlanType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


