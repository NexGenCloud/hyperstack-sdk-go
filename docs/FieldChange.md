# FieldChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | Pointer to **string** | The name of the field that was changed | [optional] 
**NewValue** | Pointer to **string** | The new value of the field | [optional] 
**OldValue** | Pointer to **string** | The old value of the field | [optional] 

## Methods

### NewFieldChange

`func NewFieldChange() *FieldChange`

NewFieldChange instantiates a new FieldChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldChangeWithDefaults

`func NewFieldChangeWithDefaults() *FieldChange`

NewFieldChangeWithDefaults instantiates a new FieldChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *FieldChange) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *FieldChange) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *FieldChange) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *FieldChange) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetNewValue

`func (o *FieldChange) GetNewValue() string`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *FieldChange) GetNewValueOk() (*string, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *FieldChange) SetNewValue(v string)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *FieldChange) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### GetOldValue

`func (o *FieldChange) GetOldValue() string`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *FieldChange) GetOldValueOk() (*string, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *FieldChange) SetOldValue(v string)`

SetOldValue sets OldValue field to given value.

### HasOldValue

`func (o *FieldChange) HasOldValue() bool`

HasOldValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


