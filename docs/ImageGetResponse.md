# ImageGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | Pointer to [**[]ImageFields**](ImageFields.md) |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**RegionName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewImageGetResponse

`func NewImageGetResponse() *ImageGetResponse`

NewImageGetResponse instantiates a new ImageGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageGetResponseWithDefaults

`func NewImageGetResponseWithDefaults() *ImageGetResponse`

NewImageGetResponseWithDefaults instantiates a new ImageGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *ImageGetResponse) GetImages() []ImageFields`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ImageGetResponse) GetImagesOk() (*[]ImageFields, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ImageGetResponse) SetImages(v []ImageFields)`

SetImages sets Images field to given value.

### HasImages

`func (o *ImageGetResponse) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetLogo

`func (o *ImageGetResponse) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ImageGetResponse) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ImageGetResponse) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ImageGetResponse) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetRegionName

`func (o *ImageGetResponse) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ImageGetResponse) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ImageGetResponse) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ImageGetResponse) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetType

`func (o *ImageGetResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImageGetResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImageGetResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ImageGetResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


