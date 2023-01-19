# {{classname}}

All URIs are relative to *https://api.solace.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEventApi**](EventAPIsApi.md#CreateEventApi) | **Post** /api/v2/architecture/eventApis | Creates an event API
[**CreateEventApiVersion**](EventAPIsApi.md#CreateEventApiVersion) | **Post** /api/v2/architecture/eventApiVersions | Creates an event API version
[**CreateEventApiVersionForEventApi**](EventAPIsApi.md#CreateEventApiVersionForEventApi) | **Post** /api/v2/architecture/eventApis/{eventApiId}/versions | Creates an event API version
[**DeleteEventApi**](EventAPIsApi.md#DeleteEventApi) | **Delete** /api/v2/architecture/eventApis/{id} | Deletes an event API
[**DeleteEventApiVersion**](EventAPIsApi.md#DeleteEventApiVersion) | **Delete** /api/v2/architecture/eventApiVersions/{versionId} | Deletes an event API version
[**DeleteEventApiVersionForEventApi**](EventAPIsApi.md#DeleteEventApiVersionForEventApi) | **Delete** /api/v2/architecture/eventApis/{eventApiId}/versions/{id} | Deletes an event API version
[**GetAsyncApiForEventApiVersion**](EventAPIsApi.md#GetAsyncApiForEventApiVersion) | **Get** /api/v2/architecture/eventApiVersions/{eventApiVersionId}/asyncApi | Retrieves the AsyncAPI specification for an event API version
[**GetEventApi**](EventAPIsApi.md#GetEventApi) | **Get** /api/v2/architecture/eventApis/{id} | Retrieves an event API
[**GetEventApiVersion**](EventAPIsApi.md#GetEventApiVersion) | **Get** /api/v2/architecture/eventApiVersions/{versionId} | Retrieves an event API version
[**GetEventApiVersionAsyncApiForEventApi**](EventAPIsApi.md#GetEventApiVersionAsyncApiForEventApi) | **Get** /api/v2/architecture/eventApis/{eventApiId}/versions/{id}/asyncApi | Retrieves the AsyncAPI specification for an event API version
[**GetEventApiVersionForEventApi**](EventAPIsApi.md#GetEventApiVersionForEventApi) | **Get** /api/v2/architecture/eventApis/{eventApiId}/versions/{id} | Retrieves an event API version
[**GetEventApiVersions**](EventAPIsApi.md#GetEventApiVersions) | **Get** /api/v2/architecture/eventApiVersions | Retrieves a list of event API versions
[**GetEventApiVersionsForEventApi**](EventAPIsApi.md#GetEventApiVersionsForEventApi) | **Get** /api/v2/architecture/eventApis/{eventApiId}/versions | Retrieves a list of event API versions
[**GetEventApis**](EventAPIsApi.md#GetEventApis) | **Get** /api/v2/architecture/eventApis | Retrieves a list of event APIs
[**UpdateEventApi**](EventAPIsApi.md#UpdateEventApi) | **Patch** /api/v2/architecture/eventApis/{id} | Updates an event API
[**UpdateEventApiVersion**](EventAPIsApi.md#UpdateEventApiVersion) | **Patch** /api/v2/architecture/eventApiVersions/{versionId} | Updates an event API by event API version ID
[**UpdateEventApiVersionForEventApi**](EventAPIsApi.md#UpdateEventApiVersionForEventApi) | **Patch** /api/v2/architecture/eventApis/{eventApiId}/versions/{id} | Updates an event API
[**UpdateEventApiVersionState**](EventAPIsApi.md#UpdateEventApiVersionState) | **Patch** /api/v2/architecture/eventApiVersions/{versionId}/state | Updates the state of an event API version by event API version ID
[**UpdateEventApiVersionStateForEventApi**](EventAPIsApi.md#UpdateEventApiVersionStateForEventApi) | **Patch** /api/v2/architecture/eventApis/{eventApiId}/versions/{id}/state | Updates the state of an event API version

# **CreateEventApi**
> EventApiResponse CreateEventApi(ctx, body)
Creates an event API

Use this API to create an event API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventApi**](EventApi.md)| The event API | 

### Return type

[**EventApiResponse**](EventApiResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEventApiVersion**
> EventApiVersionResponse CreateEventApiVersion(ctx, body)
Creates an event API version

Use this API to create an event API version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventApiVersion**](EventApiVersion.md)| Event API version | 

### Return type

[**EventApiVersionResponse**](EventApiVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEventApiVersionForEventApi**
> EventApiVersionResponse CreateEventApiVersionForEventApi(ctx, body, eventApiId)
Creates an event API version

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Event%20APIs/createEventApiVersion>another endpoint.</a><br><br>*Use this API to create an event API version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventApiVersion**](EventApiVersion.md)| Event API version | 
  **eventApiId** | **string**| The ID of the parent event API | 

### Return type

[**EventApiVersionResponse**](EventApiVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEventApi**
> DeleteEventApi(ctx, id)
Deletes an event API

Use this API to delete an event API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the event API. | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEventApiVersion**
> DeleteEventApiVersion(ctx, versionId)
Deletes an event API version

Use this API to delete an event API version by event API version ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **versionId** | **string**| The ID of the event API version | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEventApiVersionForEventApi**
> DeleteEventApiVersionForEventApi(ctx, eventApiId, id)
Deletes an event API version

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Event%20APIs/deleteEventApiVersion>another endpoint.</a><br><br>*Use this API to delete an event API version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventApiId** | **string**| The ID of the parent event API | 
  **id** | **string**| The ID of the event API version | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAsyncApiForEventApiVersion**
> string GetAsyncApiForEventApiVersion(ctx, eventApiVersionId, optional)
Retrieves the AsyncAPI specification for an event API version

Use this API to retrieve the AsyncAPI specification for an event API version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventApiVersionId** | **string**| The ID of the event API version. | 
 **optional** | ***EventAPIsApiGetAsyncApiForEventApiVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventAPIsApiGetAsyncApiForEventApiVersionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showVersioning** | **optional.Bool**| Include versions in each AsyncAPI object&#x27;s name when only one version is present | [default to false]
 **format** | **optional.String**| The format in which to retrieve the AsyncAPI specification. Possible values are yaml and json. | [default to json]
 **includedExtensions** | **optional.String**| The event portal database keys to include for each AsyncAPI object. | [default to all]
 **version** | **optional.String**| The version of AsyncAPI to use. | [default to 2.5.0]
 **asyncApiVersion** | **optional.String**| The version of AsyncAPI to use. | 
 **eventApiProductVersionId** | **optional.String**| The ID of the event API Product Version to use for generating bindings. | 
 **planId** | **optional.String**| The ID of the plan to use for generating bindings. | 
 **gatewayMessagingServiceIds** | [**optional.Interface of []string**](string.md)| The list IDs of gateway messaging services for generating bindings. | 

### Return type

**string**

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventApi**
> EventApiResponse GetEventApi(ctx, id)
Retrieves an event API

Use this API to retrieve a single event API by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the event API. | 

### Return type

[**EventApiResponse**](EventApiResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventApiVersion**
> EventApiVersionResponse GetEventApiVersion(ctx, versionId, include)
Retrieves an event API version

Use this API to retrieve a single event API version by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **versionId** | **string**| The ID of the event API version. | 
  **include** | **string**| A list of additional entities to include in the response. | 

### Return type

[**EventApiVersionResponse**](EventApiVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventApiVersionAsyncApiForEventApi**
> string GetEventApiVersionAsyncApiForEventApi(ctx, eventApiId, id, optional)
Retrieves the AsyncAPI specification for an event API version

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Event%20APIs/getAsyncApiForEventApiVersion>another endpoint.</a><br><br>*Use this API to retrieve the AsyncAPI specification for an event API version using the parent ID and the version's ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventApiId** | **string**| The ID of the parent event API. | 
  **id** | **string**| The ID of the event API version. | 
 **optional** | ***EventAPIsApiGetEventApiVersionAsyncApiForEventApiOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventAPIsApiGetEventApiVersionAsyncApiForEventApiOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **showVersioning** | **optional.Bool**| Include versions in each AsyncAPI object&#x27;s name when only one version is present | [default to false]
 **includedExtensions** | **optional.String**| The event portal database keys to include for each AsyncAPI object. | [default to all]
 **format** | **optional.String**| The format in which to retrieve the AsyncAPI specification. Possible values are yaml and json. | [default to json]
 **version** | **optional.String**| The version of AsyncAPI to use | [default to 2.5.0]
 **eventApiProductVersionId** | **optional.String**| The ID of the event API Product Version to use for generating bindings. | 
 **planId** | **optional.String**| The ID of the plan to use for generating bindings. | 
 **gatewayMessagingServiceIds** | [**optional.Interface of []string**](string.md)| The list IDs of gateway messaging services for generating bindings. | 

### Return type

**string**

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventApiVersionForEventApi**
> EventApiVersionResponse GetEventApiVersionForEventApi(ctx, eventApiId, id)
Retrieves an event API version

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Event%20APIs/getEventApiVersion>another endpoint.</a><br><br>*Use this API to retrieve a single event API version using the parent ID and the version's ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventApiId** | **string**| The ID of the parent event API. | 
  **id** | **string**| The ID of the event API version. | 

### Return type

[**EventApiVersionResponse**](EventApiVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventApiVersions**
> EventApiVersionsResponse GetEventApiVersions(ctx, optional)
Retrieves a list of event API versions

Use this API to retrieve a list of event API versions that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventAPIsApiGetEventApiVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventAPIsApiGetEventApiVersionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of results to return in one page of results. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get results from based on the page size. | [default to 1]
 **eventApiIds** | [**optional.Interface of []string**](string.md)| Match only event API versions of these event API IDs, separated by commas. | 
 **ids** | [**optional.Interface of []string**](string.md)| Match event API versions with the given IDs, separated by commas. | 
 **include** | **optional.String**| A list of additional entities to include in the response. | 
 **stateId** | **optional.String**| Match event API versions with the given state ID. | 
 **customAttributes** | **optional.String**| Returns the entities that match the custom attribute filter.&lt;br&gt;To filter by custom attribute name and value, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x3D;&#x3D;&lt;custom-attribute-value&gt;&#x60;. &lt;br&gt;To filter by custom attribute name, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x60;. &lt;br&gt;The filter supports the &#x60;AND&#x60; operator for multiple custom attribute definitions (not multiple values for a given definition). Use &#x60;;&#x60; (&#x60;semicolon&#x60;) to separate multiple queries with &#x60;AND&#x60; operation. &lt;br&gt;Note: the filter only supports custom attribute values containing characters in &#x60;[a-zA-Z0-9_\\-\\. ]&#x60;. | 

### Return type

[**EventApiVersionsResponse**](EventApiVersionsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventApiVersionsForEventApi**
> EventApiVersionsResponse GetEventApiVersionsForEventApi(ctx, eventApiId, id, optional)
Retrieves a list of event API versions

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Event%20APIs/getEventApiVersions>another endpoint.</a><br><br>*Use this API to retrieve a list of event API versions under a particular event API matching the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventApiId** | **string**| The ID of the parent event API. | 
  **id** | **string**| The ID of the event API version. | 
 **optional** | ***EventAPIsApiGetEventApiVersionsForEventApiOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventAPIsApiGetEventApiVersionsForEventApiOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **displayName** | **optional.String**| Match event API versions with the given display name. | 
 **ids** | [**optional.Interface of []string**](string.md)| Match event API versions with the given IDs separated by commas. | 
 **version** | **optional.String**| Match event API versions with the given version. | 
 **stateId** | **optional.String**| Match event API versions with the given state ID. | 
 **customAttributes** | **optional.String**| Returns the entities that match the custom attribute filter.&lt;br&gt;To filter by custom attribute name and value, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x3D;&#x3D;&lt;custom-attribute-value&gt;&#x60;. &lt;br&gt;To filter by custom attribute name, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x60;. &lt;br&gt;The filter supports the &#x60;AND&#x60; operator for multiple custom attribute definitions (not multiple values for a given definition). Use &#x60;;&#x60; (&#x60;semicolon&#x60;) to separate multiple queries with &#x60;AND&#x60; operation. &lt;br&gt;Note: the filter only supports custom attribute values containing characters in &#x60;[a-zA-Z0-9_\\-\\. ]&#x60;. | 

### Return type

[**EventApiVersionsResponse**](EventApiVersionsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventApis**
> EventApisResponse GetEventApis(ctx, optional)
Retrieves a list of event APIs

Use this API to retrieve a list of event APIs that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventAPIsApiGetEventApisOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventAPIsApiGetEventApisOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of event APIs to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **name** | **optional.String**| Name of the event API to match on. | 
 **ids** | [**optional.Interface of []string**](string.md)| Match only event APIs with the given IDs separated by commas. | 
 **applicationDomainId** | **optional.String**| Match only event APIs in the given application domain. | 
 **applicationDomainIds** | [**optional.Interface of []string**](string.md)| Match only event APIs in the given application domains. | 
 **eventApiVersionIds** | [**optional.Interface of []string**](string.md)| Match only event APIs in the given event API version ids. | 
 **availableWithinApplicationDomainIds** | **optional.Bool**| Additionally match any shared event APIs in any application domain. | 
 **shared** | **optional.Bool**| Match only with shared or unshared event APIs. | 
 **brokerType** | **optional.String**| Match only event APIs with the given broker type. | 
 **sort** | **optional.String**| Sort based on the provided parameters. &lt;br&gt; The value can either be a standalone field name (&#x60;?sort&#x3D;&lt;field&gt;&#x60;) or a field and direction, which must be delimited by a colon (&#x60;?sort&#x3D;&lt;field&gt;:&lt;asc|desc&gt;&#x60;). If the direction is not specified, the default is ascending. | 
 **customAttributes** | **optional.String**| Returns the entities that match the custom attribute filter.&lt;br&gt;To filter by custom attribute name and value, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x3D;&#x3D;&lt;custom-attribute-value&gt;&#x60;. &lt;br&gt;To filter by custom attribute name, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x60;. &lt;br&gt;The filter supports the &#x60;AND&#x60; operator for multiple custom attribute definitions (not multiple values for a given definition). Use &#x60;;&#x60; (&#x60;semicolon&#x60;) to separate multiple queries with &#x60;AND&#x60; operation. &lt;br&gt;Note: the filter only supports custom attribute values containing characters in &#x60;[a-zA-Z0-9_\\-\\. ]&#x60;. | 

### Return type

[**EventApisResponse**](EventApisResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventApi**
> EventApiResponse UpdateEventApi(ctx, body, id)
Updates an event API

Use this API to update an event API. You only need to specify the fields that need to be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventApi**](EventApi.md)| The event API | 
  **id** | **string**| The ID of the event API to update. | 

### Return type

[**EventApiResponse**](EventApiResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventApiVersion**
> EventApiVersionResponse UpdateEventApiVersion(ctx, body, versionId)
Updates an event API by event API version ID

Use this API to update an event API version by event API version ID.You only need to specify the fields that need to be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventApiVersion**](EventApiVersion.md)| The event API version. | 
  **versionId** | **string**| The ID of the event API version. | 

### Return type

[**EventApiVersionResponse**](EventApiVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventApiVersionForEventApi**
> EventApiVersionResponse UpdateEventApiVersionForEventApi(ctx, body, eventApiId, id)
Updates an event API

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Event%20APIs/updateEventApiVersion>another endpoint.</a><br><br>*Use this API to update an event API version. You only need to specify the fields that need to be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventApiVersion**](EventApiVersion.md)| The event API version. | 
  **eventApiId** | **string**| The ID of the parent event API. | 
  **id** | **string**| The ID of the event API version to update. | 

### Return type

[**EventApiVersionResponse**](EventApiVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventApiVersionState**
> StateChangeRequestResponse UpdateEventApiVersionState(ctx, body, versionId)
Updates the state of an event API version by event API version ID

Use this API to update the state of an event API version. You only need to specify the state ID field with the desired state ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventApiVersion**](EventApiVersion.md)| The Event API version. | 
  **versionId** | **string**| The ID of the event API version. | 

### Return type

[**StateChangeRequestResponse**](StateChangeRequestResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventApiVersionStateForEventApi**
> VersionedObjectStateChangeRequest UpdateEventApiVersionStateForEventApi(ctx, body, eventApiId, id)
Updates the state of an event API version

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Event%20APIs/updateEventApiVersionState>another endpoint.</a><br><br>*Use this API to update the state of an event API version. You only need to specify the state ID field with the desired state ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventApiVersion**](EventApiVersion.md)| The Event API version. | 
  **eventApiId** | **string**| The ID of the parent event API. | 
  **id** | **string**| The ID of the event API version to update. | 

### Return type

[**VersionedObjectStateChangeRequest**](VersionedObjectStateChangeRequest.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

