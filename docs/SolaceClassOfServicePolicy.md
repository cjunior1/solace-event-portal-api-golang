# SolaceClassOfServicePolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID value of the object | [optional] [default to null]
**Type_** | **string** | The type of payload | [optional] [default to solaceClassOfServicePolicy]
**MessageDeliveryMode** | **string** | The mode that will be used for message delivery (ex: &#x60;guaranteed&#x60; uses a queue) | [optional] [default to MESSAGE_DELIVERY_MODE.DIRECT]
**AccessType** | **string** |  | [optional] [default to null]
**MaximumTimeToLive** | **int32** | Duration in seconds of how long a message can live in a queue | [optional] [default to null]
**QueueType** | **string** | The arrangement of queues on a broker used for message delivery (ex: &#x60;single&#x60; uses one queue per event API version in this event API product) | [optional] [default to QUEUE_TYPE.COMBINED]
**MaxMsgSpoolUsage** | **int32** | Total number of MBs available for the queue to use | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

