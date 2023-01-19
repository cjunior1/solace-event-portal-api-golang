# {{classname}}

All URIs are relative to *https://api.solace.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationDomain**](ApplicationDomainsApi.md#CreateApplicationDomain) | **Post** /api/v2/architecture/applicationDomains | Creates an application domain object
[**DeleteApplicationDomain**](ApplicationDomainsApi.md#DeleteApplicationDomain) | **Delete** /api/v2/architecture/applicationDomains/{id} | Deletes an application domain object
[**GetApplicationDomain**](ApplicationDomainsApi.md#GetApplicationDomain) | **Get** /api/v2/architecture/applicationDomains/{id} | Retrieves an application domain object
[**GetApplicationDomains**](ApplicationDomainsApi.md#GetApplicationDomains) | **Get** /api/v2/architecture/applicationDomains | Gets the application domain objects
[**ImportApplicationDomains**](ApplicationDomainsApi.md#ImportApplicationDomains) | **Post** /api/v2/architecture/applicationDomains/import | (Beta) Import application domains and their entities
[**UpdateApplicationDomain**](ApplicationDomainsApi.md#UpdateApplicationDomain) | **Patch** /api/v2/architecture/applicationDomains/{id} | Updates an application domain object

# **CreateApplicationDomain**
> ApplicationDomainResponse CreateApplicationDomain(ctx, body)
Creates an application domain object

To help keep your event-driven architecture organized, use application domains to create namespaces for your applications, events and schemas.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplicationDomain**](ApplicationDomain.md)| Application domains have a name and topic domain. | 

### Return type

[**ApplicationDomainResponse**](ApplicationDomainResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApplicationDomain**
> DeleteApplicationDomain(ctx, id)
Deletes an application domain object

Use this API to delete an application domain. This action also deletes all applications, events, and schemas in the application domain. You cannot undo this operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the application domain object. | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationDomain**
> ApplicationDomainResponse GetApplicationDomain(ctx, id, optional)
Retrieves an application domain object

Use this API to retrieve a single application domain by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the application domain object. | 
 **optional** | ***ApplicationDomainsApiGetApplicationDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationDomainsApiGetApplicationDomainOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | [**optional.Interface of []string**](string.md)| Specify extra data to be included, options are: stats | 

### Return type

[**ApplicationDomainResponse**](ApplicationDomainResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationDomains**
> ApplicationDomainsResponse GetApplicationDomains(ctx, optional)
Gets the application domain objects

Use this API to retrieve a list of application domains that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationDomainsApiGetApplicationDomainsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationDomainsApiGetApplicationDomainsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of application domains to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **name** | **optional.String**| Name to be used to match the application domain. | 
 **ids** | [**optional.Interface of []string**](string.md)| Match only application domains with the given IDs separated by commas. | 
 **include** | [**optional.Interface of []string**](string.md)| Specify extra data to be included, options are: stats | 

### Return type

[**ApplicationDomainsResponse**](ApplicationDomainsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportApplicationDomains**
> ImportApplicationDomains(ctx, body)
(Beta) Import application domains and their entities

Use this API to import application domains and their nested entities. Please note that this endpoint is in beta and could be subject to change in the future

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplicationDomainImportDto**](ApplicationDomainImportDto.md)| Application domain import file | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApplicationDomain**
> ApplicationDomainResponse UpdateApplicationDomain(ctx, body, id)
Updates an application domain object

Use this API to update an application domain. You only need to specify the fields that need to be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplicationDomain**](ApplicationDomain.md)| The application domain object. | 
  **id** | **string**| The ID of the application domain object. | 

### Return type

[**ApplicationDomainResponse**](ApplicationDomainResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

