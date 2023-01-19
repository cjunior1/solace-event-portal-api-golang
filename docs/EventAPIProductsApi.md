# {{classname}}

All URIs are relative to *https://api.solace.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateGatewayMessagingServiceToEAPVersion**](EventAPIProductsApi.md#AssociateGatewayMessagingServiceToEAPVersion) | **Post** /api/v2/architecture/eventApiProductVersions/{eventApiProductVersionId}/memAssociations | (Beta) Associate gateway messaging service to event API product version
[**CreateEventApiProduct**](EventAPIProductsApi.md#CreateEventApiProduct) | **Post** /api/v2/architecture/eventApiProducts | Creates an event API product
[**CreateEventApiProductVersion**](EventAPIProductsApi.md#CreateEventApiProductVersion) | **Post** /api/v2/architecture/eventApiProductVersions | (Beta) Creates an event API product version
[**CreateEventApiProductVersionForEventApiProduct**](EventAPIProductsApi.md#CreateEventApiProductVersionForEventApiProduct) | **Post** /api/v2/architecture/eventApiProducts/{eventApiProductId}/versions | Creates an event API product version
[**DeleteEventApiProduct**](EventAPIProductsApi.md#DeleteEventApiProduct) | **Delete** /api/v2/architecture/eventApiProducts/{id} | Deletes an event API product
[**DeleteEventApiProductVersion**](EventAPIProductsApi.md#DeleteEventApiProductVersion) | **Delete** /api/v2/architecture/eventApiProductVersions/{versionId} | (Beta) Deletes an event API product version by ID
[**DeleteEventApiProductVersionForEventApiProduct**](EventAPIProductsApi.md#DeleteEventApiProductVersionForEventApiProduct) | **Delete** /api/v2/architecture/eventApiProducts/{eventApiProductId}/versions/{id} | Deletes an event API product version
[**DisassociateGatewayMessagingServiceFromEventApiProductVersionById**](EventAPIProductsApi.md#DisassociateGatewayMessagingServiceFromEventApiProductVersionById) | **Delete** /api/v2/architecture/eventApiProductMemAssociations/{memAssociationId} | (Beta) Disassociates a gateway messaging service from an event API product version by association ID
[**DisassociateGatewayMessagingServiceToEAPVersion**](EventAPIProductsApi.md#DisassociateGatewayMessagingServiceToEAPVersion) | **Delete** /api/v2/architecture/eventApiProductVersions/{eventApiProductVersionId}/memAssociations/{memAssociationId} | (Beta) Disassociate gateway messaging service from event API product version
[**GetEventApiProduct**](EventAPIProductsApi.md#GetEventApiProduct) | **Get** /api/v2/architecture/eventApiProducts/{id} | Retrieves an event API product
[**GetEventApiProductVersion**](EventAPIProductsApi.md#GetEventApiProductVersion) | **Get** /api/v2/architecture/eventApiProductVersions/{versionId} | (Beta) Retrieves an event API product version
[**GetEventApiProductVersionForEventApiProduct**](EventAPIProductsApi.md#GetEventApiProductVersionForEventApiProduct) | **Get** /api/v2/architecture/eventApiProducts/{eventApiProductId}/versions/{id} | Retrieves an event API product version
[**GetEventApiProductVersions**](EventAPIProductsApi.md#GetEventApiProductVersions) | **Get** /api/v2/architecture/eventApiProductVersions | (Beta) Retrieves a list of event API product versions
[**GetEventApiProductVersionsForEventApiProduct**](EventAPIProductsApi.md#GetEventApiProductVersionsForEventApiProduct) | **Get** /api/v2/architecture/eventApiProducts/{eventApiProductId}/versions | Retrieves a list of event API product versions
[**GetEventApiProducts**](EventAPIProductsApi.md#GetEventApiProducts) | **Get** /api/v2/architecture/eventApiProducts | Retrieves a list of event API products
[**UpdateEventApiProduct**](EventAPIProductsApi.md#UpdateEventApiProduct) | **Patch** /api/v2/architecture/eventApiProducts/{id} | Updates an event API product
[**UpdateEventApiProductVersion**](EventAPIProductsApi.md#UpdateEventApiProductVersion) | **Patch** /api/v2/architecture/eventApiProductVersions/{versionId} | (Beta) Updates an event API product version by version ID
[**UpdateEventApiProductVersionForEventApiProduct**](EventAPIProductsApi.md#UpdateEventApiProductVersionForEventApiProduct) | **Patch** /api/v2/architecture/eventApiProducts/{eventApiProductId}/versions/{id} | Updates an event API product version
[**UpdateEventApiProductVersionState**](EventAPIProductsApi.md#UpdateEventApiProductVersionState) | **Patch** /api/v2/architecture/eventApiProductVersions/{versionId}/state | Updates the state of an event API product version by ID
[**UpdateEventApiProductVersionStateForEventApiProduct**](EventAPIProductsApi.md#UpdateEventApiProductVersionStateForEventApiProduct) | **Patch** /api/v2/architecture/eventApiProducts/{eventApiProductId}/versions/{id}/state | Updates the state of an event API product version

# **AssociateGatewayMessagingServiceToEAPVersion**
> GatewayMessagingServiceResponse AssociateGatewayMessagingServiceToEAPVersion(ctx, body, eventApiProductVersionId)
(Beta) Associate gateway messaging service to event API product version

Use this API to associate an event API product version and gateway messaging service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GatewayMessagingService**](GatewayMessagingService.md)| Gateway messaging service Id and supported Protocols | 
  **eventApiProductVersionId** | **string**| The ID of the event API product version to associate. | 

### Return type

[**GatewayMessagingServiceResponse**](GatewayMessagingServiceResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEventApiProduct**
> EventApiProductResponse CreateEventApiProduct(ctx, body)
Creates an event API product

Use this API to create an event API product.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventApiProduct**](EventApiProduct.md)| The event API product | 

### Return type

[**EventApiProductResponse**](EventApiProductResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEventApiProductVersion**
> EventApiProductVersionResponse CreateEventApiProductVersion(ctx, body)
(Beta) Creates an event API product version

Use this API to create an event API product version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventApiProductVersion**](EventApiProductVersion.md)| Event API product version | 

### Return type

[**EventApiProductVersionResponse**](EventApiProductVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEventApiProductVersionForEventApiProduct**
> EventApiProductVersionResponse CreateEventApiProductVersionForEventApiProduct(ctx, body, eventApiProductId)
Creates an event API product version

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Event%20API%20Products/createEventApiProductVersion>another endpoint.</a><br><br>*Use this API to create an event API product version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventApiProductVersion**](EventApiProductVersion.md)| Event API product version | 
  **eventApiProductId** | **string**| The ID of the parent event API product | 

### Return type

[**EventApiProductVersionResponse**](EventApiProductVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEventApiProduct**
> DeleteEventApiProduct(ctx, id)
Deletes an event API product

Use this API to delete an event API product.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the event API product. | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEventApiProductVersion**
> DeleteEventApiProductVersion(ctx, versionId)
(Beta) Deletes an event API product version by ID

Use this API to delete an event API product version by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **versionId** | **string**| The ID of the event API product version | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEventApiProductVersionForEventApiProduct**
> DeleteEventApiProductVersionForEventApiProduct(ctx, eventApiProductId, id)
Deletes an event API product version

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Event%20API%20Products/deleteEventApiProductVersion>another endpoint.</a><br><br>*Use this API to delete an event API product version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventApiProductId** | **string**| The ID of the parent event API product | 
  **id** | **string**| The ID of the event API product version | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisassociateGatewayMessagingServiceFromEventApiProductVersionById**
> DisassociateGatewayMessagingServiceFromEventApiProductVersionById(ctx, memAssociationId)
(Beta) Disassociates a gateway messaging service from an event API product version by association ID

Use this API to disassociate an event API product version and gateway messaging service by association ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **memAssociationId** | **string**| The association ID to perform the disassociation for | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisassociateGatewayMessagingServiceToEAPVersion**
> DisassociateGatewayMessagingServiceToEAPVersion(ctx, eventApiProductVersionId, memAssociationId)
(Beta) Disassociate gateway messaging service from event API product version

Use this API to disassociate an event API product version and gateway messaging service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventApiProductVersionId** | **string**| The ID of the event API product version to disassociate. | 
  **memAssociationId** | **string**| The MEM association ID to dissociate from. | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventApiProduct**
> EventApiProductResponse GetEventApiProduct(ctx, id)
Retrieves an event API product

Use this API to retrieve a single event API product by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the event API product. | 

### Return type

[**EventApiProductResponse**](EventApiProductResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventApiProductVersion**
> EventApiProductVersionResponse GetEventApiProductVersion(ctx, versionId, include, optional)
(Beta) Retrieves an event API product version

Use this API to retrieve a single event API product version by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **versionId** | **string**| The ID of the event API product version. | 
  **include** | **string**| A list of additional entities to include in the response. | 
 **optional** | ***EventAPIProductsApiGetEventApiProductVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventAPIProductsApiGetEventApiProductVersionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientAppId** | **optional.String**| Match event API product versions with the given clientAppId. | 

### Return type

[**EventApiProductVersionResponse**](EventApiProductVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventApiProductVersionForEventApiProduct**
> EventApiProductVersionResponse GetEventApiProductVersionForEventApiProduct(ctx, eventApiProductId, id)
Retrieves an event API product version

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Event%20API%20Products/getEventApiProductVersion>another endpoint.</a><br><br>*Use this API to retrieve a single event API product version using the parent ID and the version's ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventApiProductId** | **string**| The ID of the parent event API product. | 
  **id** | **string**| The ID of the event API product version. | 

### Return type

[**EventApiProductVersionResponse**](EventApiProductVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventApiProductVersions**
> EventApiProductVersionsResponse GetEventApiProductVersions(ctx, optional)
(Beta) Retrieves a list of event API product versions

Use this API to retrieve a list of event API product versions that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventAPIProductsApiGetEventApiProductVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventAPIProductsApiGetEventApiProductVersionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of results to return in one page of results. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get results from based on the page size. | [default to 1]
 **eventApiProductIds** | [**optional.Interface of []string**](string.md)| Match only event API product versions of these event API product IDs, separated by commas. | 
 **ids** | [**optional.Interface of []string**](string.md)| Match event API product versions with the given IDs, separated by commas. | 
 **include** | **optional.String**| A list of additional entities to include in the response. | 
 **stateId** | **optional.String**| Match event API product versions with the given state ID. | 
 **messagingServiceId** | **optional.String**| Match event API product versions with the given messagingServiceId. | 
 **clientAppId** | **optional.String**| Match event API product versions with the given clientAppId. | 
 **shared** | **optional.Bool**| Match event API product versions with the parent objects shared setting. | 
 **latest** | **optional.Bool**| Only return the latest version of event API products. | 
 **customAttributes** | **optional.String**| Returns the entities that match the custom attribute filter.&lt;br&gt;To filter by custom attribute name and value, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x3D;&#x3D;&lt;custom-attribute-value&gt;&#x60;. &lt;br&gt;To filter by custom attribute name, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x60;. &lt;br&gt;The filter supports the &#x60;AND&#x60; operator for multiple custom attribute definitions (not multiple values for a given definition). Use &#x60;;&#x60; (&#x60;semicolon&#x60;) to separate multiple queries with &#x60;AND&#x60; operation. &lt;br&gt;Note: the filter only supports custom attribute values containing characters in &#x60;[a-zA-Z0-9_\\-\\. ]&#x60;. | 

### Return type

[**EventApiProductVersionsResponse**](EventApiProductVersionsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventApiProductVersionsForEventApiProduct**
> EventApiProductVersionsResponse GetEventApiProductVersionsForEventApiProduct(ctx, eventApiProductId, optional)
Retrieves a list of event API product versions

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Event%20API%20Products/getEventApiProductVersions>another endpoint.</a><br><br>*Use this API to retrieve a list of event API product versions under a particular event API product matching the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventApiProductId** | **string**| The ID of the parent event API product. | 
 **optional** | ***EventAPIProductsApiGetEventApiProductVersionsForEventApiProductOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventAPIProductsApiGetEventApiProductVersionsForEventApiProductOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **displayName** | **optional.String**| Match event API product versions with the given display name. | 
 **id** | **optional.String**| The ID of the event API product version. | 
 **ids** | [**optional.Interface of []string**](string.md)| Match event API product versions with the given IDs separated by commas. | 
 **version** | **optional.String**| Match event API product versions with the given version. | 
 **stateId** | **optional.String**| Match event API product versions with the given state ID. | 
 **customAttributes** | **optional.String**| Returns the entities that match the custom attribute filter.&lt;br&gt;To filter by custom attribute name and value, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x3D;&#x3D;&lt;custom-attribute-value&gt;&#x60;. &lt;br&gt;To filter by custom attribute name, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x60;. &lt;br&gt;The filter supports the &#x60;AND&#x60; operator for multiple custom attribute definitions (not multiple values for a given definition). Use &#x60;;&#x60; (&#x60;semicolon&#x60;) to separate multiple queries with &#x60;AND&#x60; operation. &lt;br&gt;Note: the filter only supports custom attribute values containing characters in &#x60;[a-zA-Z0-9_\\-\\. ]&#x60;. | 

### Return type

[**EventApiProductVersionsResponse**](EventApiProductVersionsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventApiProducts**
> EventApiProductsResponse GetEventApiProducts(ctx, optional)
Retrieves a list of event API products

Use this API to retrieve a list of event API products that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventAPIProductsApiGetEventApiProductsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventAPIProductsApiGetEventApiProductsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of event API products to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **brokerType** | **optional.String**| Match only event API products with the given broken type. | 
 **name** | **optional.String**| Name of the event API product to match on. | 
 **ids** | [**optional.Interface of []string**](string.md)| Match only event API products with the given IDs separated by commas. | 
 **applicationDomainId** | **optional.String**| Match only event API products in the given application domain. | 
 **applicationDomainIds** | [**optional.Interface of []string**](string.md)| Match only event API products in the given application domains. | 
 **shared** | **optional.Bool**| Match only with shared or unshared event API products. | 
 **sort** | **optional.String**| Sort based on the provided parameters. &lt;br&gt; The value can either be a standalone field name (&#x60;?sort&#x3D;&lt;field&gt;&#x60;) or a field and direction, which must be delimited by a colon (&#x60;?sort&#x3D;&lt;field&gt;:&lt;asc|desc&gt;&#x60;). If the direction is not specified, the default is ascending. | 
 **customAttributes** | **optional.String**| Returns the entities that match the custom attribute filter.&lt;br&gt;To filter by custom attribute name and value, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x3D;&#x3D;&lt;custom-attribute-value&gt;&#x60;. &lt;br&gt;To filter by custom attribute name, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x60;. &lt;br&gt;The filter supports the &#x60;AND&#x60; operator for multiple custom attribute definitions (not multiple values for a given definition). Use &#x60;;&#x60; (&#x60;semicolon&#x60;) to separate multiple queries with &#x60;AND&#x60; operation. &lt;br&gt;Note: the filter only supports custom attribute values containing characters in &#x60;[a-zA-Z0-9_\\-\\. ]&#x60;. | 

### Return type

[**EventApiProductsResponse**](EventApiProductsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventApiProduct**
> EventApiProductResponse UpdateEventApiProduct(ctx, body, id)
Updates an event API product

Use this API to update an event API product. You only need to specify the fields that need to be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventApiProduct**](EventApiProduct.md)| The event API product | 
  **id** | **string**| The ID of the event API product to update. | 

### Return type

[**EventApiProductResponse**](EventApiProductResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventApiProductVersion**
> EventApiProductVersionResponse UpdateEventApiProductVersion(ctx, body, versionId)
(Beta) Updates an event API product version by version ID

Use this API to update an event API product version. You only need to specify the fields that need to be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventApiProductVersion**](EventApiProductVersion.md)| The event API product version. | 
  **versionId** | **string**| The ID of the event API product version. | 

### Return type

[**EventApiProductVersionResponse**](EventApiProductVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventApiProductVersionForEventApiProduct**
> EventApiProductVersionResponse UpdateEventApiProductVersionForEventApiProduct(ctx, body, eventApiProductId, id)
Updates an event API product version

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Event%20API%20Products/updateEventApiProductVersion>another endpoint.</a><br><br>*Use this API to update an event API product version. You only need to specify the fields that need to be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventApiProductVersion**](EventApiProductVersion.md)| The event API product version. | 
  **eventApiProductId** | **string**| The ID of the parent event API product. | 
  **id** | **string**| The ID of the event API product version to update. | 

### Return type

[**EventApiProductVersionResponse**](EventApiProductVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventApiProductVersionState**
> StateChangeRequestResponse UpdateEventApiProductVersionState(ctx, body, versionId)
Updates the state of an event API product version by ID

Use this API to update the state of an event API product version. You only need to specify the state ID field with the desired state ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventApiProductVersion**](EventApiProductVersion.md)| The event API product version. | 
  **versionId** | **string**| The ID of the event API product version. | 

### Return type

[**StateChangeRequestResponse**](StateChangeRequestResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventApiProductVersionStateForEventApiProduct**
> VersionedObjectStateChangeRequest UpdateEventApiProductVersionStateForEventApiProduct(ctx, body, eventApiProductId, id)
Updates the state of an event API product version

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Event%20API%20Products/updateEventApiProductVersionState>another endpoint.</a><br><br>*Use this API to update the state of an event API product version. You only need to specify the state ID field with the desired state ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventApiProductVersion**](EventApiProductVersion.md)| The event API product version. | 
  **eventApiProductId** | **string**| The ID of the parent event API product. | 
  **id** | **string**| The ID of the event API product version to update. | 

### Return type

[**VersionedObjectStateChangeRequest**](VersionedObjectStateChangeRequest.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

