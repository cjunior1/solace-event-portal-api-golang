# {{classname}}

All URIs are relative to *https://api.solace.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessagingService**](MessagingServicesApi.md#CreateMessagingService) | **Post** /api/v2/architecture/messagingServices | (Beta) Creates a messaging service object
[**DeleteMessagingService**](MessagingServicesApi.md#DeleteMessagingService) | **Delete** /api/v2/architecture/messagingServices/{id} | (Beta) Deletes a messaging service object
[**GetMessagingService**](MessagingServicesApi.md#GetMessagingService) | **Get** /api/v2/architecture/messagingServices/{id} | (Beta) Retrieves a messaging service object
[**GetMessagingServices**](MessagingServicesApi.md#GetMessagingServices) | **Get** /api/v2/architecture/messagingServices | (Beta) Retrieves a list of messaging services
[**RemoveAssociationMessagingService**](MessagingServicesApi.md#RemoveAssociationMessagingService) | **Put** /api/v2/architecture/messagingServices/{id}/removeAssociation | (Beta) Removes a messaging service&#x27;s association to the requested entity.
[**ScanStartMessagingService**](MessagingServicesApi.md#ScanStartMessagingService) | **Put** /api/v2/architecture/messagingServices/{messagingServiceId}/scanStart | (Beta) Requests a scan to run against a messaging service
[**UpdateMessagingService**](MessagingServicesApi.md#UpdateMessagingService) | **Patch** /api/v2/architecture/messagingServices/{id} | (Beta) Updates a messaging service object

# **CreateMessagingService**
> MessagingServiceResponse CreateMessagingService(ctx, body)
(Beta) Creates a messaging service object

Use this API to create a messaging service object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MessagingService**](MessagingService.md)| The messaging service object. | 

### Return type

[**MessagingServiceResponse**](MessagingServiceResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMessagingService**
> DeleteMessagingService(ctx, id)
(Beta) Deletes a messaging service object

Use this API to delete a messaging service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the messaging service object. | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessagingService**
> MessagingServiceResponse GetMessagingService(ctx, id)
(Beta) Retrieves a messaging service object

Use this API to retrieve a single messaging service by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the messaging service object. | 

### Return type

[**MessagingServiceResponse**](MessagingServiceResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessagingServices**
> MessagingServicesResponse GetMessagingServices(ctx, optional)
(Beta) Retrieves a list of messaging services

Use this API to retrieve a list of messaging services that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagingServicesApiGetMessagingServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagingServicesApiGetMessagingServicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of messaging services to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **sort** | **optional.String**| Sort based on the provided parameters. &lt;br&gt; The value can either be a standalone field name (&#x60;?sort&#x3D;&lt;field&gt;&#x60;) or a field and direction, which must be delimited by a colon (&#x60;?sort&#x3D;&lt;field&gt;:&lt;asc|desc&gt;&#x60;). If the direction is not specified, the default is ascending. | 
 **ids** | [**optional.Interface of []string**](string.md)| The IDs of the messaging services. | 
 **messagingServiceType** | **optional.String**| Match only messaging services of the given type | 
 **runtimeAgentId** | **optional.String**| Match only messaging services in the given runtimeAgentId | 
 **eventMeshId** | **optional.String**| Match only messaging services in the given eventMeshId | 
 **eventManagementAgentId** | **optional.String**|  | 

### Return type

[**MessagingServicesResponse**](MessagingServicesResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveAssociationMessagingService**
> MessagingServiceResponse RemoveAssociationMessagingService(ctx, body, id)
(Beta) Removes a messaging service's association to the requested entity.

Use this API to remove the association between a messaging service and either of EVENT_MESH or EVENT_MANAGEMENT_AGENT.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MessagingServiceRemoveAssociation**](MessagingServiceRemoveAssociation.md)| The association object with the value matching either EVENT_MESH or EVENT_MANAGEMENT_AGENT. | 
  **id** | **string**| The ID of the messaging service object. | 

### Return type

[**MessagingServiceResponse**](MessagingServiceResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScanStartMessagingService**
> MessagingServiceOperationResponse ScanStartMessagingService(ctx, body, messagingServiceId)
(Beta) Requests a scan to run against a messaging service

Use this API to make a scan request on a messaging service object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MessagingServiceOperation**](MessagingServiceOperation.md)| The messaging service object. | 
  **messagingServiceId** | **string**| The ID of the messaging service object. | 

### Return type

[**MessagingServiceOperationResponse**](MessagingServiceOperationResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMessagingService**
> MessagingServiceResponse UpdateMessagingService(ctx, body, id)
(Beta) Updates a messaging service object

Use this API to update a messaging service. You only need to specify the fields that need to be updated. However, if you want to update anything under subObjects (i.e. anything inside messagingServiceConnections object), you need to provide the original messagingServiceConnections with the updated fields instead of just providing the changed fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MessagingService**](MessagingService.md)| The messaging service object. | 
  **id** | **string**| The ID of the messaging service object to update. | 

### Return type

[**MessagingServiceResponse**](MessagingServiceResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

