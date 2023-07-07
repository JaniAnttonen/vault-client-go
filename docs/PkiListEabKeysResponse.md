# PkiListEabKeysResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyInfo** | Pointer to **map[string]interface{}** | EAB details keyed by the eab key id | [optional] 
**Keys** | Pointer to **[]string** | A list of unused eab keys | [optional] 



## Methods


### NewPkiListEabKeysResponse

`func NewPkiListEabKeysResponse() *PkiListEabKeysResponse`

NewPkiListEabKeysResponse instantiates a new PkiListEabKeysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiListEabKeysResponseWithDefaults

`func NewPkiListEabKeysResponseWithDefaults() *PkiListEabKeysResponse`

NewPkiListEabKeysResponseWithDefaults instantiates a new PkiListEabKeysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetKeyInfo

`func (o *PkiListEabKeysResponse) GetKeyInfo() map[string]interface{}`

GetKeyInfo returns the KeyInfo field if non-nil, zero value otherwise.

### GetKeyInfoOk

`func (o *PkiListEabKeysResponse) GetKeyInfoOk() (*map[string]interface{}, bool)`

GetKeyInfoOk returns a tuple with the KeyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyInfo

`func (o *PkiListEabKeysResponse) SetKeyInfo(v map[string]interface{})`

SetKeyInfo sets KeyInfo field to given value.


### HasKeyInfo

`func (o *PkiListEabKeysResponse) HasKeyInfo() bool`

HasKeyInfo returns a boolean if a field has been set.




### GetKeys

`func (o *PkiListEabKeysResponse) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *PkiListEabKeysResponse) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *PkiListEabKeysResponse) SetKeys(v []string)`

SetKeys sets Keys field to given value.


### HasKeys

`func (o *PkiListEabKeysResponse) HasKeys() bool`

HasKeys returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

