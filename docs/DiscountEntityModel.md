# DiscountEntityModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]DiscountPlanFields**](DiscountPlanFields.md) |  | [optional] 
**Entity** | Pointer to **string** |  | [optional] 

## Methods

### NewDiscountEntityModel

`func NewDiscountEntityModel() *DiscountEntityModel`

NewDiscountEntityModel instantiates a new DiscountEntityModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountEntityModelWithDefaults

`func NewDiscountEntityModelWithDefaults() *DiscountEntityModel`

NewDiscountEntityModelWithDefaults instantiates a new DiscountEntityModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DiscountEntityModel) GetData() []DiscountPlanFields`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DiscountEntityModel) GetDataOk() (*[]DiscountPlanFields, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DiscountEntityModel) SetData(v []DiscountPlanFields)`

SetData sets Data field to given value.

### HasData

`func (o *DiscountEntityModel) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEntity

`func (o *DiscountEntityModel) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *DiscountEntityModel) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *DiscountEntityModel) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *DiscountEntityModel) HasEntity() bool`

HasEntity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


