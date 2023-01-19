# {{classname}}

All URIs are relative to *https://api.solace.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEventMesh**](EventMeshesApi.md#CreateEventMesh) | **Post** /api/v2/architecture/eventMeshes | Creates an event mesh object
[**DeleteEventMesh**](EventMeshesApi.md#DeleteEventMesh) | **Delete** /api/v2/architecture/eventMeshes/{id} | Deletes an event mesh object
[**GetEventMesh**](EventMeshesApi.md#GetEventMesh) | **Get** /api/v2/architecture/eventMeshes/{id} | Retrieves an event mesh object
[**GetEventMeshes**](EventMeshesApi.md#GetEventMeshes) | **Get** /api/v2/architecture/eventMeshes | Gets the event mesh objects
[**UpdateEventMesh**](EventMeshesApi.md#UpdateEventMesh) | **Patch** /api/v2/architecture/eventMeshes/{id} | Updates an event mesh object

# **CreateEventMesh**
> EventMeshResponse CreateEventMesh(ctx, body)
Creates an event mesh object

Creates an event mesh object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventMesh**](EventMesh.md)| Event mesh. | 

### Return type

[**EventMeshResponse**](EventMeshResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEventMesh**
> DeleteEventMesh(ctx, id)
Deletes an event mesh object

Use this API to delete an event mesh.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the event mesh object. | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventMesh**
> EventMeshResponse GetEventMesh(ctx, id)
Retrieves an event mesh object

Retrieves a single event mesh by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the event mesh object. | 

### Return type

[**EventMeshResponse**](EventMeshResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventMeshes**
> EventMeshesResponse GetEventMeshes(ctx, optional)
Gets the event mesh objects

Use this API to retrieve a list of event meshes that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventMeshesApiGetEventMeshesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventMeshesApiGetEventMeshesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of event meshes to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **name** | **optional.String**| Name of the event mesh to match on. | 
 **environmentId** | **optional.String**| Match only event meshes in the given environment | 

### Return type

[**EventMeshesResponse**](EventMeshesResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventMesh**
> EventMeshResponse UpdateEventMesh(ctx, body, id)
Updates an event mesh object

Use this API to update an event mesh. You only need to specify the fields that need to be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventMesh**](EventMesh.md)| The event mesh object. | 
  **id** | **string**| The ID of the event mesh object to update. | 

### Return type

[**EventMeshResponse**](EventMeshResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

