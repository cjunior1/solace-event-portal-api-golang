# {{classname}}

All URIs are relative to *https://api.solace.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomAttributeDefinition**](CustomAttributeDefinitionsApi.md#CreateCustomAttributeDefinition) | **Post** /api/v2/architecture/customAttributeDefinitions | Creates a custom attribute definition object
[**CreateCustomAttributeDefinitionByApplicationDomain**](CustomAttributeDefinitionsApi.md#CreateCustomAttributeDefinitionByApplicationDomain) | **Post** /api/v2/architecture/applicationDomains/{applicationDomainId}/customAttributeDefinitions | Creates a custom attribute definition object for provided application domain
[**DeleteCustomAttributeDefinition**](CustomAttributeDefinitionsApi.md#DeleteCustomAttributeDefinition) | **Delete** /api/v2/architecture/customAttributeDefinitions/{id} | Deletes a custom attribute definition object
[**DeleteCustomAttributeDefinitionByApplicationDomain**](CustomAttributeDefinitionsApi.md#DeleteCustomAttributeDefinitionByApplicationDomain) | **Delete** /api/v2/architecture/applicationDomains/{applicationDomainId}/customAttributeDefinitions | Deletes a custom attribute definition object by application domain
[**DeleteCustomAttributeDefinitionOfApplicationDomain**](CustomAttributeDefinitionsApi.md#DeleteCustomAttributeDefinitionOfApplicationDomain) | **Delete** /api/v2/architecture/applicationDomains/{applicationDomainId}/customAttributeDefinitions/{customAttributeId} | Deletes a custom attribute definition object of application domain
[**GetCustomAttributeDefinition**](CustomAttributeDefinitionsApi.md#GetCustomAttributeDefinition) | **Get** /api/v2/architecture/customAttributeDefinitions/{id} | Retrieves a custom attribute definition object
[**GetCustomAttributeDefinitions**](CustomAttributeDefinitionsApi.md#GetCustomAttributeDefinitions) | **Get** /api/v2/architecture/customAttributeDefinitions | Gets the custom attribute definition objects
[**GetCustomAttributeDefinitionsByApplicationDomain**](CustomAttributeDefinitionsApi.md#GetCustomAttributeDefinitionsByApplicationDomain) | **Get** /api/v2/architecture/applicationDomains/{applicationDomainId}/customAttributeDefinitions | Gets the custom attribute definition objects by Application domain
[**UpdateCustomAttributeDefinition**](CustomAttributeDefinitionsApi.md#UpdateCustomAttributeDefinition) | **Patch** /api/v2/architecture/customAttributeDefinitions/{id} | Updates a custom attribute definition object
[**UpdateCustomAttributeDefinitionByApplicationDomain**](CustomAttributeDefinitionsApi.md#UpdateCustomAttributeDefinitionByApplicationDomain) | **Patch** /api/v2/architecture/applicationDomains/{applicationDomainId}/customAttributeDefinitions/{customAttributeId} | Updates a custom attribute definition object for provided application domain

# **CreateCustomAttributeDefinition**
> CustomAttributeDefinitionResponse CreateCustomAttributeDefinition(ctx, body)
Creates a custom attribute definition object

Use this API to create a custom attribute definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CustomAttributeDefinition**](CustomAttributeDefinition.md)| The custom attribute object. | 

### Return type

[**CustomAttributeDefinitionResponse**](CustomAttributeDefinitionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCustomAttributeDefinitionByApplicationDomain**
> CustomAttributeDefinitionResponse CreateCustomAttributeDefinitionByApplicationDomain(ctx, body, applicationDomainId)
Creates a custom attribute definition object for provided application domain

Use this API to create a custom attribute definition for provided application domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CustomAttributeDefinition**](CustomAttributeDefinition.md)| The custom attribute object. | 
  **applicationDomainId** | **string**| The ID of the application domain | 

### Return type

[**CustomAttributeDefinitionResponse**](CustomAttributeDefinitionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCustomAttributeDefinition**
> DeleteCustomAttributeDefinition(ctx, id)
Deletes a custom attribute definition object

Use this API to delete a custom attribute definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the custom attribute definition | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCustomAttributeDefinitionByApplicationDomain**
> DeleteCustomAttributeDefinitionByApplicationDomain(ctx, applicationDomainId)
Deletes a custom attribute definition object by application domain

Use this API to delete a custom attribute definition by given application domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationDomainId** | **string**| The ID of the application domain | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCustomAttributeDefinitionOfApplicationDomain**
> DeleteCustomAttributeDefinitionOfApplicationDomain(ctx, applicationDomainId, customAttributeId)
Deletes a custom attribute definition object of application domain

Use this API to delete a custom attribute definition of given application domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationDomainId** | **string**| The ID of the application domain | 
  **customAttributeId** | **string**| The ID of the custom attribute | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomAttributeDefinition**
> CustomAttributeDefinitionResponse GetCustomAttributeDefinition(ctx, id)
Retrieves a custom attribute definition object

Use this API to retrieve a single custom attribute by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the custom attribute object. | 

### Return type

[**CustomAttributeDefinitionResponse**](CustomAttributeDefinitionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomAttributeDefinitions**
> CustomAttributeDefinitionsResponse GetCustomAttributeDefinitions(ctx, optional)
Gets the custom attribute definition objects

Use this API to retrieve a list of custom attributes that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomAttributeDefinitionsApiGetCustomAttributeDefinitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomAttributeDefinitionsApiGetCustomAttributeDefinitionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of custom attribute definitions to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **associatedEntityTypes** | [**optional.Interface of []string**](string.md)| Match only custom attribute definitions with the given associated entity type names separated by commas. | 

### Return type

[**CustomAttributeDefinitionsResponse**](CustomAttributeDefinitionsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomAttributeDefinitionsByApplicationDomain**
> CustomAttributeDefinitionsResponse GetCustomAttributeDefinitionsByApplicationDomain(ctx, applicationDomainId, optional)
Gets the custom attribute definition objects by Application domain

Use this API to retrieve a list of custom attributes that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationDomainId** | **string**| Match only custom attribute definitions with the given application domain Id  | 
 **optional** | ***CustomAttributeDefinitionsApiGetCustomAttributeDefinitionsByApplicationDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomAttributeDefinitionsApiGetCustomAttributeDefinitionsByApplicationDomainOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| The number of custom attribute definitions to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]

### Return type

[**CustomAttributeDefinitionsResponse**](CustomAttributeDefinitionsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomAttributeDefinition**
> CustomAttributeDefinitionResponse UpdateCustomAttributeDefinition(ctx, id, optional)
Updates a custom attribute definition object

Use this API to update a custom attribute definition. You can only update the associated entity types.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the custom attribute object to update. | 
 **optional** | ***CustomAttributeDefinitionsApiUpdateCustomAttributeDefinitionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomAttributeDefinitionsApiUpdateCustomAttributeDefinitionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CustomAttributeDefinition**](CustomAttributeDefinition.md)| The custom attribute definition object. | 

### Return type

[**CustomAttributeDefinitionResponse**](CustomAttributeDefinitionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomAttributeDefinitionByApplicationDomain**
> CustomAttributeDefinitionResponse UpdateCustomAttributeDefinitionByApplicationDomain(ctx, body, applicationDomainId, customAttributeId)
Updates a custom attribute definition object for provided application domain

Use this API to update a custom attribute definition for provided application domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CustomAttributeDefinition**](CustomAttributeDefinition.md)| The custom attribute object. | 
  **applicationDomainId** | **string**| The ID of the application domain | 
  **customAttributeId** | **string**| The ID of the custom attribute | 

### Return type

[**CustomAttributeDefinitionResponse**](CustomAttributeDefinitionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

