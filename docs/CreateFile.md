# CreateFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | Pointer to [**EntityBase**](EntityBase.md) |  | [optional] 
**File** | Pointer to [***os.File**](*os.File.md) |  | [optional] 

## Methods

### NewCreateFile

`func NewCreateFile() *CreateFile`

NewCreateFile instantiates a new CreateFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFileWithDefaults

`func NewCreateFileWithDefaults() *CreateFile`

NewCreateFileWithDefaults instantiates a new CreateFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *CreateFile) GetRequest() EntityBase`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *CreateFile) GetRequestOk() (*EntityBase, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *CreateFile) SetRequest(v EntityBase)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *CreateFile) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetFile

`func (o *CreateFile) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *CreateFile) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *CreateFile) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *CreateFile) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


