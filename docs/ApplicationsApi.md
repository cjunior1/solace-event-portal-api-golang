# {{classname}}

All URIs are relative to *https://api.solace.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplication**](ApplicationsApi.md#CreateApplication) | **Post** /api/v2/architecture/applications | Creates an application object
[**CreateApplicationVersion**](ApplicationsApi.md#CreateApplicationVersion) | **Post** /api/v2/architecture/applicationVersions | Creates an application version object
[**CreateApplicationVersionForApplication**](ApplicationsApi.md#CreateApplicationVersionForApplication) | **Post** /api/v2/architecture/applications/{applicationId}/versions | Creates an application version object
[**DeleteApplication**](ApplicationsApi.md#DeleteApplication) | **Delete** /api/v2/architecture/applications/{id} | Deletes an application object
[**DeleteApplicationVersion**](ApplicationsApi.md#DeleteApplicationVersion) | **Delete** /api/v2/architecture/applicationVersions/{versionId} | Deletes an application version object
[**DeleteApplicationVersionForApplication**](ApplicationsApi.md#DeleteApplicationVersionForApplication) | **Delete** /api/v2/architecture/applications/{applicationId}/versions/{id} | Deletes an application version object
[**GetApplication**](ApplicationsApi.md#GetApplication) | **Get** /api/v2/architecture/applications/{id} | Retrieves an application object
[**GetApplicationVersion**](ApplicationsApi.md#GetApplicationVersion) | **Get** /api/v2/architecture/applicationVersions/{versionId} | Retrieves an application version object
[**GetApplicationVersionAsyncApiForApplication**](ApplicationsApi.md#GetApplicationVersionAsyncApiForApplication) | **Get** /api/v2/architecture/applications/{applicationId}/versions/{id}/asyncApi | Retrieves the AsyncAPI specification for an application version
[**GetApplicationVersionForApplication**](ApplicationsApi.md#GetApplicationVersionForApplication) | **Get** /api/v2/architecture/applications/{applicationId}/versions/{id} | Retrieves an application version object
[**GetApplicationVersions**](ApplicationsApi.md#GetApplicationVersions) | **Get** /api/v2/architecture/applicationVersions | Gets the application version objects
[**GetApplicationVersionsForApplication**](ApplicationsApi.md#GetApplicationVersionsForApplication) | **Get** /api/v2/architecture/applications/{applicationId}/versions | Gets application version objects
[**GetApplications**](ApplicationsApi.md#GetApplications) | **Get** /api/v2/architecture/applications | Gets the application objects
[**GetAsyncApiForApplicationVersion**](ApplicationsApi.md#GetAsyncApiForApplicationVersion) | **Get** /api/v2/architecture/applicationVersions/{applicationVersionId}/asyncApi | Retrieves the AsyncAPI specification for an application version
[**UpdateApplication**](ApplicationsApi.md#UpdateApplication) | **Patch** /api/v2/architecture/applications/{id} | Updates an application object
[**UpdateApplicationVersion**](ApplicationsApi.md#UpdateApplicationVersion) | **Patch** /api/v2/architecture/applicationVersions/{versionId} | Updates an application version object
[**UpdateApplicationVersionForApplication**](ApplicationsApi.md#UpdateApplicationVersionForApplication) | **Patch** /api/v2/architecture/applications/{applicationId}/versions/{id} | Updates an application version object
[**UpdateApplicationVersionState**](ApplicationsApi.md#UpdateApplicationVersionState) | **Patch** /api/v2/architecture/applicationVersions/{versionId}/state | Updates the state of an application version object
[**UpdateApplicationVersionStateForApplication**](ApplicationsApi.md#UpdateApplicationVersionStateForApplication) | **Patch** /api/v2/architecture/applications/{applicationId}/versions/{id}/state | Updates the state of an application version object
[**UpdateMsgSvcAssociationForAppVersion**](ApplicationsApi.md#UpdateMsgSvcAssociationForAppVersion) | **Put** /api/v2/architecture/applicationVersions/{versionId}/messagingServices | Update messaging service association for an application version object

# **CreateApplication**
> ApplicationResponse CreateApplication(ctx, body)
Creates an application object

To model your event-driven architecture, applications are a fundamental building block for modelling the producers and consumers of events. Use this API to create applications and model the events they produce and consume.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Application**](Application.md)| Applications have a name and live within an application domain. Events can be added to the application as produced or consumed. | 

### Return type

[**ApplicationResponse**](ApplicationResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateApplicationVersion**
> ApplicationVersionResponse CreateApplicationVersion(ctx, body)
Creates an application version object

Creates an application version object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplicationVersion**](ApplicationVersion.md)| App version request body description | 

### Return type

[**ApplicationVersionResponse**](ApplicationVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateApplicationVersionForApplication**
> ApplicationVersionResponse CreateApplicationVersionForApplication(ctx, body, applicationId)
Creates an application version object

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Applications/createApplicationVersion>another endpoint.</a><br><br>*Creates an application version object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplicationVersion**](ApplicationVersion.md)| App version request body description | 
  **applicationId** | **string**| The ID of the parent application | 

### Return type

[**ApplicationVersionResponse**](ApplicationVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApplication**
> DeleteApplication(ctx, id)
Deletes an application object

Use this API to delete an application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the application | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApplicationVersion**
> DeleteApplicationVersion(ctx, versionId)
Deletes an application version object

Use this API to delete an application version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **versionId** | **string**| The ID of the application version | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApplicationVersionForApplication**
> DeleteApplicationVersionForApplication(ctx, applicationId, id)
Deletes an application version object

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Applications/deleteApplicationVersion>another endpoint.</a><br><br>*Use this API to delete an application version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**| The ID of the parent application | 
  **id** | **string**| The ID of the application version | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplication**
> ApplicationResponse GetApplication(ctx, id)
Retrieves an application object

Use this API to retrieve a single application by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the application object. | 

### Return type

[**ApplicationResponse**](ApplicationResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationVersion**
> ApplicationVersionResponse GetApplicationVersion(ctx, versionId)
Retrieves an application version object

Use this API to retrieve a single application version by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **versionId** | **string**| The ID of the application version object. | 

### Return type

[**ApplicationVersionResponse**](ApplicationVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationVersionAsyncApiForApplication**
> string GetApplicationVersionAsyncApiForApplication(ctx, applicationId, id, optional)
Retrieves the AsyncAPI specification for an application version

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Applications/getAsyncApiForApplicationVersion>another endpoint.</a><br><br>*Use this API to retrieve the AsyncAPI specification for an application version using the parent ID and the version's ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**| The ID of the parent application. | 
  **id** | **string**| The ID of the application version. | 
 **optional** | ***ApplicationsApiGetApplicationVersionAsyncApiForApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiGetApplicationVersionAsyncApiForApplicationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **showVersioning** | **optional.Bool**| Include versions in each AsyncAPI object&#x27;s name when only one version is present | [default to false]
 **includedExtensions** | **optional.String**| The event portal database keys to include for each AsyncAPI object. | [default to all]
 **asyncApiVersion** | **optional.String**| The version of AsyncAPI to use | [default to 2.5.0]
 **format** | **optional.String**| The format in which to retrieve the AsyncAPI specification. Possible values are yaml and json. | [default to json]
 **messagingServiceId** | **optional.String**| Applies bindings from consumed events that are published in this messaging service&#x27;s modeled event mesh. | 

### Return type

**string**

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationVersionForApplication**
> ApplicationVersionResponse GetApplicationVersionForApplication(ctx, applicationId, id)
Retrieves an application version object

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Applications/getApplicationVersion>another endpoint.</a><br><br>*Use this API to retrieve a single application version by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**| The ID of the parent application. | 
  **id** | **string**| The ID of the application version. | 

### Return type

[**ApplicationVersionResponse**](ApplicationVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationVersions**
> ApplicationVersionsResponse GetApplicationVersions(ctx, optional)
Gets the application version objects

Use this API to retrieve a list of application versions that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationsApiGetApplicationVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiGetApplicationVersionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of application versions to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **applicationIds** | [**optional.Interface of []string**](string.md)| Match only application versions of these application IDs, separated by commas. | 
 **ids** | [**optional.Interface of []string**](string.md)| Match only application versions with the given IDs, separated by commas. | 
 **messagingServiceIds** | [**optional.Interface of []string**](string.md)| Match only application versions with the given messaging service IDs, separated by commas. | 
 **customAttributes** | **optional.String**| Returns the entities that match the custom attribute filter.&lt;br&gt;To filter by custom attribute name and value, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x3D;&#x3D;&lt;custom-attribute-value&gt;&#x60;. &lt;br&gt;To filter by custom attribute name, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x60;. &lt;br&gt;The filter supports the &#x60;AND&#x60; operator for multiple custom attribute definitions (not multiple values for a given definition). Use &#x60;;&#x60; (&#x60;semicolon&#x60;) to separate multiple queries with &#x60;AND&#x60; operation. &lt;br&gt;Note: the filter only supports custom attribute values containing characters in &#x60;[a-zA-Z0-9_\\-\\. ]&#x60;. | 

### Return type

[**ApplicationVersionsResponse**](ApplicationVersionsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationVersionsForApplication**
> ApplicationVersionsResponse GetApplicationVersionsForApplication(ctx, applicationId, optional)
Gets application version objects

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Applications/getApplicationVersions>another endpoint.</a><br><br>*Use this API to retrieve a list of application versions that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**| The ID of the parent application. | 
 **optional** | ***ApplicationsApiGetApplicationVersionsForApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiGetApplicationVersionsForApplicationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **displayName** | **optional.String**| Match application versions with the given display name. | 
 **ids** | [**optional.Interface of []string**](string.md)| Match application versions with the given IDs separated by commas. | 
 **version** | **optional.String**| Match application version objects with the given version. | 
 **customAttributes** | **optional.String**| Returns the entities that match the custom attribute filter.&lt;br&gt;To filter by custom attribute name and value, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x3D;&#x3D;&lt;custom-attribute-value&gt;&#x60;. &lt;br&gt;To filter by custom attribute name, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x60;. &lt;br&gt;The filter supports the &#x60;AND&#x60; operator for multiple custom attribute definitions (not multiple values for a given definition). Use &#x60;;&#x60; (&#x60;semicolon&#x60;) to separate multiple queries with &#x60;AND&#x60; operation. &lt;br&gt;Note: the filter only supports custom attribute values containing characters in &#x60;[a-zA-Z0-9_\\-\\. ]&#x60;. | 

### Return type

[**ApplicationVersionsResponse**](ApplicationVersionsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplications**
> ApplicationsResponse GetApplications(ctx, optional)
Gets the application objects

Use this API to retrieve a list of applications that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationsApiGetApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiGetApplicationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of applications to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **name** | **optional.String**| Name of the application to match on. | 
 **applicationDomainId** | **optional.String**| Match only applications in the given application domain. | 
 **ids** | [**optional.Interface of []string**](string.md)| Match only applications with the given IDs separated by commas. | 
 **sort** | **optional.String**| Sort based on the provided parameters. &lt;br&gt; The value can either be a standalone field name (&#x60;?sort&#x3D;&lt;field&gt;&#x60;) or a field and direction, which must be delimited by a colon (&#x60;?sort&#x3D;&lt;field&gt;:&lt;asc|desc&gt;&#x60;). If the direction is not specified, the default is ascending. | 
 **customAttributes** | **optional.String**| Returns the entities that match the custom attribute filter.&lt;br&gt;To filter by custom attribute name and value, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x3D;&#x3D;&lt;custom-attribute-value&gt;&#x60;. &lt;br&gt;To filter by custom attribute name, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x60;. &lt;br&gt;The filter supports the &#x60;AND&#x60; operator for multiple custom attribute definitions (not multiple values for a given definition). Use &#x60;;&#x60; (&#x60;semicolon&#x60;) to separate multiple queries with &#x60;AND&#x60; operation. &lt;br&gt;Note: the filter only supports custom attribute values containing characters in &#x60;[a-zA-Z0-9_\\-\\. ]&#x60;. | 
 **applicationType** | **optional.String**| Match only applications with the given applicationType. | 

### Return type

[**ApplicationsResponse**](ApplicationsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAsyncApiForApplicationVersion**
> string GetAsyncApiForApplicationVersion(ctx, applicationVersionId, optional)
Retrieves the AsyncAPI specification for an application version

Use this API to retrieve the AsyncAPI specification for an application version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationVersionId** | **string**| The ID of the application version. | 
 **optional** | ***ApplicationsApiGetAsyncApiForApplicationVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiGetAsyncApiForApplicationVersionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**| The format in which to retrieve the AsyncAPI specification. Possible values are yaml and json. | [default to json]
 **showVersioning** | **optional.Bool**| Include versions in each AsyncAPI object&#x27;s name when only one version is present | [default to false]
 **includedExtensions** | **optional.String**| The event portal database keys to include for each AsyncAPI object. | [default to all]
 **asyncApiVersion** | **optional.String**| The version of AsyncAPI to use. | [default to 2.5.0]
 **messagingServiceId** | **optional.String**| Applies bindings from consumed events that are published in this messaging service&#x27;s modeled event mesh. | 

### Return type

**string**

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApplication**
> ApplicationResponse UpdateApplication(ctx, body, id)
Updates an application object

Use this API to update an application. You only need to specify the fields that need to be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Application**](Application.md)| The application object. | 
  **id** | **string**| The ID of the application object to update. | 

### Return type

[**ApplicationResponse**](ApplicationResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApplicationVersion**
> ApplicationVersionResponse UpdateApplicationVersion(ctx, body, versionId, optional)
Updates an application version object

Use this API to update an application version. You only need to specify the fields that need to be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplicationVersion**](ApplicationVersion.md)| The application version object. | 
  **versionId** | **string**| The ID of the application version object to update. | 
 **optional** | ***ApplicationsApiUpdateApplicationVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiUpdateApplicationVersionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **include** | [**optional.Interface of []string**](string.md)|  | 
 **relationsBrokerType** | **optional.**|  | 

### Return type

[**ApplicationVersionResponse**](ApplicationVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApplicationVersionForApplication**
> ApplicationVersionResponse UpdateApplicationVersionForApplication(ctx, body, applicationId, id)
Updates an application version object

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Applications/updateApplicationVersion>another endpoint.</a><br><br>*Use this API to update an application version. You only need to specify the fields that need to be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplicationVersion**](ApplicationVersion.md)| The application version object. | 
  **applicationId** | **string**| The ID of the parent application object. | 
  **id** | **string**| The ID of the application version object to update. | 

### Return type

[**ApplicationVersionResponse**](ApplicationVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApplicationVersionState**
> StateChangeRequestResponse UpdateApplicationVersionState(ctx, body, versionId)
Updates the state of an application version object

Use this API to update the state of an application version. You only need to specify the target stateId field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VersionedObjectStateChangeRequest**](VersionedObjectStateChangeRequest.md)| The state change object. | 
  **versionId** | **string**| The ID of the application version object to update. | 

### Return type

[**StateChangeRequestResponse**](StateChangeRequestResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApplicationVersionStateForApplication**
> VersionedObjectStateChangeRequest UpdateApplicationVersionStateForApplication(ctx, body, applicationId, id)
Updates the state of an application version object

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Applications/updateApplicationVersionState>another endpoint.</a><br><br>*Use this API to update the state of an application version. You only need to specify the target stateId field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplicationVersion**](ApplicationVersion.md)| The application version object. | 
  **applicationId** | **string**| The ID of the parent application object. | 
  **id** | **string**| The ID of the application version object to update. | 

### Return type

[**VersionedObjectStateChangeRequest**](VersionedObjectStateChangeRequest.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgSvcAssociationForAppVersion**
> MessagingServiceAssociationResponse UpdateMsgSvcAssociationForAppVersion(ctx, body, versionId)
Update messaging service association for an application version object

Use this API to update the messaging service association for an application version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MessagingServiceAssociationDto**](MessagingServiceAssociationDto.md)| The messaging service association object | 
  **versionId** | **string**| The ID of the application version | 

### Return type

[**MessagingServiceAssociationResponse**](MessagingServiceAssociationResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

