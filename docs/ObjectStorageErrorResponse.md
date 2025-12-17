# ObjectStorageErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorReason** | **string** |  | 
**Message** | **string** |  | 
**Status** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewObjectStorageErrorResponse

`func NewObjectStorageErrorResponse(errorReason string, message string, ) *ObjectStorageErrorResponse`

NewObjectStorageErrorResponse instantiates a new ObjectStorageErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageErrorResponseWithDefaults

`func NewObjectStorageErrorResponseWithDefaults() *ObjectStorageErrorResponse`

NewObjectStorageErrorResponseWithDefaults instantiates a new ObjectStorageErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorReason

`func (o *ObjectStorageErrorResponse) GetErrorReason() string`

GetErrorReason returns the ErrorReason field if non-nil, zero value otherwise.

### GetErrorReasonOk

`func (o *ObjectStorageErrorResponse) GetErrorReasonOk() (*string, bool)`

GetErrorReasonOk returns a tuple with the ErrorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorReason

`func (o *ObjectStorageErrorResponse) SetErrorReason(v string)`

SetErrorReason sets ErrorReason field to given value.


### GetMessage

`func (o *ObjectStorageErrorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ObjectStorageErrorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ObjectStorageErrorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStatus

`func (o *ObjectStorageErrorResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ObjectStorageErrorResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ObjectStorageErrorResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ObjectStorageErrorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


