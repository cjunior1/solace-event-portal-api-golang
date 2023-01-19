# {{classname}}

All URIs are relative to *https://api.solace.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSchema**](SchemasApi.md#CreateSchema) | **Post** /api/v2/architecture/schemas | Creates a schema object
[**CreateSchemaVersion**](SchemasApi.md#CreateSchemaVersion) | **Post** /api/v2/architecture/schemaVersions | Creates a schema version object
[**CreateSchemaVersionForSchema**](SchemasApi.md#CreateSchemaVersionForSchema) | **Post** /api/v2/architecture/schemas/{schemaId}/versions | Creates a schema version object
[**DeleteSchema**](SchemasApi.md#DeleteSchema) | **Delete** /api/v2/architecture/schemas/{id} | Deletes a schema object
[**DeleteSchemaVersion**](SchemasApi.md#DeleteSchemaVersion) | **Delete** /api/v2/architecture/schemaVersions/{id} | Deletes a schema version object
[**DeleteSchemaVersionForSchema**](SchemasApi.md#DeleteSchemaVersionForSchema) | **Delete** /api/v2/architecture/schemas/{schemaId}/versions/{id} | Deletes a schema version object
[**GetSchema**](SchemasApi.md#GetSchema) | **Get** /api/v2/architecture/schemas/{id} | Retrieves a schema object
[**GetSchemaVersion**](SchemasApi.md#GetSchemaVersion) | **Get** /api/v2/architecture/schemaVersions/{versionId} | Retrieves a schema version object
[**GetSchemaVersionForSchema**](SchemasApi.md#GetSchemaVersionForSchema) | **Get** /api/v2/architecture/schemas/{schemaId}/versions/{id} | Retrieves a schema version object
[**GetSchemaVersions**](SchemasApi.md#GetSchemaVersions) | **Get** /api/v2/architecture/schemaVersions | Retrieves a list of schema version objects
[**GetSchemaVersionsForSchema**](SchemasApi.md#GetSchemaVersionsForSchema) | **Get** /api/v2/architecture/schemas/{schemaId}/versions | Gets the schema version objects for a schema
[**GetSchemas**](SchemasApi.md#GetSchemas) | **Get** /api/v2/architecture/schemas | Gets the schema objects
[**UpdateSchema**](SchemasApi.md#UpdateSchema) | **Patch** /api/v2/architecture/schemas/{id} | Updates a schema object
[**UpdateSchemaVersion**](SchemasApi.md#UpdateSchemaVersion) | **Patch** /api/v2/architecture/schemaVersions/{id} | Updates a schema version object
[**UpdateSchemaVersionForSchema**](SchemasApi.md#UpdateSchemaVersionForSchema) | **Patch** /api/v2/architecture/schemas/{schemaId}/versions/{id} | Updates a schema version object
[**UpdateSchemaVersionState**](SchemasApi.md#UpdateSchemaVersionState) | **Patch** /api/v2/architecture/schemaVersions/{id}/state | Updates the state of a schema version object
[**UpdateSchemaVersionStateForSchema**](SchemasApi.md#UpdateSchemaVersionStateForSchema) | **Patch** /api/v2/architecture/schemas/{schemaId}/versions/{id}/state | Updates the state of a schema version object

# **CreateSchema**
> SchemaResponse CreateSchema(ctx, body)
Creates a schema object

To model your event-driven architecture, schemas are a fundamental building block for modelling the payloads of the events flowing through your system. Use this API to create schemas that can later be referenced by events.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SchemaObject**](SchemaObject.md)| The schema requires a name, an application domain, a schema type and a content type. | 

### Return type

[**SchemaResponse**](SchemaResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSchemaVersion**
> SchemaVersionResponse CreateSchemaVersion(ctx, body)
Creates a schema version object

Creates a schema version object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SchemaVersion**](SchemaVersion.md)| schema version details | 

### Return type

[**SchemaVersionResponse**](SchemaVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSchemaVersionForSchema**
> SchemaVersionResponse CreateSchemaVersionForSchema(ctx, body, schemaId)
Creates a schema version object

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Schemas/createSchemaVersion>another endpoint.</a><br><br>*Creates a schema version object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SchemaVersion**](SchemaVersion.md)| schema version details | 
  **schemaId** | **string**| The ID of the parent schema | 

### Return type

[**SchemaVersionResponse**](SchemaVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSchema**
> DeleteSchema(ctx, id)
Deletes a schema object

Use this API to delete a schema. The schema must not be in use by any events else it cannot be deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the schema object. | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSchemaVersion**
> DeleteSchemaVersion(ctx, id)
Deletes a schema version object

Use this API to delete a schema version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the schema version. | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSchemaVersionForSchema**
> DeleteSchemaVersionForSchema(ctx, schemaId, id)
Deletes a schema version object

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Schemas/deleteSchemaVersion>another endpoint.</a><br><br>*Use this API to delete a schema version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **schemaId** | **string**| The ID of the schema object. | 
  **id** | **string**| The ID of the schema version. | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSchema**
> SchemaResponse GetSchema(ctx, id)
Retrieves a schema object

Use this API to retrieve a single schema by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the schema object. | 

### Return type

[**SchemaResponse**](SchemaResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSchemaVersion**
> SchemaVersionResponse GetSchemaVersion(ctx, versionId)
Retrieves a schema version object

Use this API to retrieve a single schema version by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **versionId** | **string**| The ID of the schema version object. | 

### Return type

[**SchemaVersionResponse**](SchemaVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSchemaVersionForSchema**
> SchemaVersionResponse GetSchemaVersionForSchema(ctx, schemaId, id)
Retrieves a schema version object

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Schemas/getSchemaVersion>another endpoint.</a><br><br>*Use this API to retrieve a single schema by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **schemaId** | **string**| The ID of the schema object. | 
  **id** | **string**| The ID of the schema version. | 

### Return type

[**SchemaVersionResponse**](SchemaVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSchemaVersions**
> SchemaVersionsResponse GetSchemaVersions(ctx, optional)
Retrieves a list of schema version objects

Use this API to retrieve a list of schema versions that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SchemasApiGetSchemaVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SchemasApiGetSchemaVersionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of schemas to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **schemaIds** | [**optional.Interface of []string**](string.md)| Match only schema versions of these schema IDs, separated by commas. | 
 **ids** | [**optional.Interface of []string**](string.md)| Match only schema versions with the given IDs, separated by commas. | 
 **customAttributes** | **optional.String**| Returns the entities that match the custom attribute filter.&lt;br&gt;To filter by custom attribute name and value, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x3D;&#x3D;&lt;custom-attribute-value&gt;&#x60;. &lt;br&gt;To filter by custom attribute name, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x60;. &lt;br&gt;The filter supports the &#x60;AND&#x60; operator for multiple custom attribute definitions (not multiple values for a given definition). Use &#x60;;&#x60; (&#x60;semicolon&#x60;) to separate multiple queries with &#x60;AND&#x60; operation. &lt;br&gt;Note: the filter only supports custom attribute values containing characters in &#x60;[a-zA-Z0-9_\\-\\. ]&#x60;. | 
 **include** | **optional.String**|  | 

### Return type

[**SchemaVersionsResponse**](SchemaVersionsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSchemaVersionsForSchema**
> SchemaVersionsResponse GetSchemaVersionsForSchema(ctx, schemaId, optional)
Gets the schema version objects for a schema

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Schemas/getSchemaVersions>another endpoint.</a><br><br>*Use this API to retrieve a list of schema versions that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **schemaId** | **string**| The ID of the schema object. | 
 **optional** | ***SchemasApiGetSchemaVersionsForSchemaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SchemasApiGetSchemaVersionsForSchemaOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **pageSize** | **optional.Int32**| The number of schemas to get per page. | 
 **versions** | [**optional.Interface of []string**](string.md)| Match only with schema versions. | 
 **displayName** | **optional.String**| Match only schema versions with the given display name. | 
 **ids** | [**optional.Interface of []string**](string.md)| Match only schema versions with the given IDs separated by commas. | 
 **customAttributes** | **optional.String**| Returns the entities that match the custom attribute filter.&lt;br&gt;To filter by custom attribute name and value, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x3D;&#x3D;&lt;custom-attribute-value&gt;&#x60;. &lt;br&gt;To filter by custom attribute name, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x60;. &lt;br&gt;The filter supports the &#x60;AND&#x60; operator for multiple custom attribute definitions (not multiple values for a given definition). Use &#x60;;&#x60; (&#x60;semicolon&#x60;) to separate multiple queries with &#x60;AND&#x60; operation. &lt;br&gt;Note: the filter only supports custom attribute values containing characters in &#x60;[a-zA-Z0-9_\\-\\. ]&#x60;. | 

### Return type

[**SchemaVersionsResponse**](SchemaVersionsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSchemas**
> SchemasResponse GetSchemas(ctx, optional)
Gets the schema objects

Use this API to retrieve a list of schemas that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SchemasApiGetSchemasOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SchemasApiGetSchemasOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of schemas to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **name** | **optional.String**| Name of the schema to match on. | 
 **shared** | **optional.Bool**| Match only with shared or unshared schemas. | 
 **applicationDomainId** | **optional.String**| Match only schemas in the given application domain. | 
 **applicationDomainIds** | [**optional.Interface of []string**](string.md)| Match only schemas in the given application domain ids. | 
 **ids** | [**optional.Interface of []string**](string.md)| Match only schemas with the given IDs separated by commas. | 
 **sort** | **optional.String**| Sort based on the provided parameters. &lt;br&gt; The value can either be a standalone field name (&#x60;?sort&#x3D;&lt;field&gt;&#x60;) or a field and direction, which must be delimited by a colon (&#x60;?sort&#x3D;&lt;field&gt;:&lt;asc|desc&gt;&#x60;). If the direction is not specified, the default is ascending. | 
 **customAttributes** | **optional.String**| Returns the entities that match the custom attribute filter.&lt;br&gt;To filter by custom attribute name and value, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x3D;&#x3D;&lt;custom-attribute-value&gt;&#x60;. &lt;br&gt;To filter by custom attribute name, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x60;. &lt;br&gt;The filter supports the &#x60;AND&#x60; operator for multiple custom attribute definitions (not multiple values for a given definition). Use &#x60;;&#x60; (&#x60;semicolon&#x60;) to separate multiple queries with &#x60;AND&#x60; operation. &lt;br&gt;Note: the filter only supports custom attribute values containing characters in &#x60;[a-zA-Z0-9_\\-\\. ]&#x60;. | 

### Return type

[**SchemasResponse**](SchemasResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSchema**
> SchemaResponse UpdateSchema(ctx, body, id)
Updates a schema object

Update a schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SchemaObject**](SchemaObject.md)| The schema requires a name, an application domain, a schema type and a content type. | 
  **id** | **string**| The ID of the schema object. | 

### Return type

[**SchemaResponse**](SchemaResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSchemaVersion**
> SchemaVersionResponse UpdateSchemaVersion(ctx, body, id)
Updates a schema version object

Use this API to update a schema version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SchemaVersion**](SchemaVersion.md)|  | 
  **id** | **string**| The ID of the schema version. | 

### Return type

[**SchemaVersionResponse**](SchemaVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSchemaVersionForSchema**
> SchemaVersionResponse UpdateSchemaVersionForSchema(ctx, body, schemaId, id)
Updates a schema version object

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Schemas/updateSchemaVersion>another endpoint.</a><br><br>*Use this API to update a schema version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SchemaVersion**](SchemaVersion.md)|  | 
  **schemaId** | **string**| The ID of the schema object. | 
  **id** | **string**| The ID of the schema version. | 

### Return type

[**SchemaVersionResponse**](SchemaVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSchemaVersionState**
> StateChangeRequestResponse UpdateSchemaVersionState(ctx, body, id)
Updates the state of a schema version object

Use this API to update the state of a schema version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VersionedObjectStateChangeRequest**](VersionedObjectStateChangeRequest.md)| The state change object. | 
  **id** | **string**| The ID of the schema version. | 

### Return type

[**StateChangeRequestResponse**](StateChangeRequestResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSchemaVersionStateForSchema**
> VersionedObjectStateChangeRequest UpdateSchemaVersionStateForSchema(ctx, body, schemaId, id)
Updates the state of a schema version object

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Schemas/updateSchemaVersionState>another endpoint.</a><br><br>*Use this API to update the state of a schema version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VersionedObjectStateChangeRequest**](VersionedObjectStateChangeRequest.md)| The state change object. | 
  **schemaId** | **string**| The ID of the schema object. | 
  **id** | **string**| The ID of the schema version. | 

### Return type

[**VersionedObjectStateChangeRequest**](VersionedObjectStateChangeRequest.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

