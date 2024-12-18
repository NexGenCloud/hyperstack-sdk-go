# SubResourcesGraphResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**SubResourcesCosts** | Pointer to [**SubResourcesCostsResponseModel**](SubResourcesCostsResponseModel.md) |  | [optional] 

## Methods

### NewSubResourcesGraphResponseModel

`func NewSubResourcesGraphResponseModel() *SubResourcesGraphResponseModel`

NewSubResourcesGraphResponseModel instantiates a new SubResourcesGraphResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubResourcesGraphResponseModelWithDefaults

`func NewSubResourcesGraphResponseModelWithDefaults() *SubResourcesGraphResponseModel`

NewSubResourcesGraphResponseModelWithDefaults instantiates a new SubResourcesGraphResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SubResourcesGraphResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SubResourcesGraphResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SubResourcesGraphResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SubResourcesGraphResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *SubResourcesGraphResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubResourcesGraphResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubResourcesGraphResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubResourcesGraphResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubResourcesCosts

`func (o *SubResourcesGraphResponseModel) GetSubResourcesCosts() SubResourcesCostsResponseModel`

GetSubResourcesCosts returns the SubResourcesCosts field if non-nil, zero value otherwise.

### GetSubResourcesCostsOk

`func (o *SubResourcesGraphResponseModel) GetSubResourcesCostsOk() (*SubResourcesCostsResponseModel, bool)`

GetSubResourcesCostsOk returns a tuple with the SubResourcesCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubResourcesCosts

`func (o *SubResourcesGraphResponseModel) SetSubResourcesCosts(v SubResourcesCostsResponseModel)`

SetSubResourcesCosts sets SubResourcesCosts field to given value.

### HasSubResourcesCosts

`func (o *SubResourcesGraphResponseModel) HasSubResourcesCosts() bool`

HasSubResourcesCosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


