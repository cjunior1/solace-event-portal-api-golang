# MessagingServiceAuthentication

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | **string** |  | [optional] [default to null]
**UpdatedTime** | **string** |  | [optional] [default to null]
**CreatedBy** | **string** |  | [optional] [default to null]
**ChangedBy** | **string** |  | [optional] [default to null]
**Id** | **string** | Primary key set by the server. | [optional] [default to null]
**MessagingServiceConnectionId** | **string** | The ID of the connection object associated to the authentication object. | [optional] [default to null]
**Name** | **string** | The name of the authentication object. | [default to null]
**AuthenticationType** | **string** | The type of the authentication object. | [default to null]
**AuthenticationDetails** | [**map[string]interface{}**](interface{}.md) | A JSON map containing a map of extra details for the authentication. | [optional] [default to null]
**MessagingServiceCredentials** | [**[]MessagingServiceCredentials**](MessagingServiceCredentials.md) |  | [optional] [default to null]
**Type_** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

