/*
 * Event Portal Open API
 *
 * REST API Concepts  Solace PubSub+ Event Portal provides REST APIs that you can use to manage your data in PubSub+ Cloud.  The REST APIs allow you to model your event-driven architectures from your own client applications.  The following document describes the public REST APIs available for use in Event Portal 2.0. Objects created in this  version will not be available in Event Portal 1.0. APIs that display (Beta) in their summary are provided as-is and  are subject to change or removal. They may not be of the quality expected for generally available APIs and have no  guaranteed forward compatibility with the generally available version of the API.  
 *
 * API version: 2.0.11
 * Generated by: Swagger Codegen (https://github.com/eventportal-api/eventportal-codegen.git)
 */
package eventportal

type MessagingServiceScan struct {
	CreatedTime string `json:"createdTime,omitempty"`
	UpdatedTime string `json:"updatedTime,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	ChangedBy string `json:"changedBy,omitempty"`
	// Primary key set by the server.
	Id string `json:"id,omitempty"`
	// The status of the messaging service scan.
	Status string `json:"status,omitempty"`
	// The description of the messaging service scan status.
	StatusDescription string `json:"statusDescription,omitempty"`
	// The messagingServiceId of the scan.
	MessagingServiceId string `json:"messagingServiceId,omitempty"`
	// The messagingServiceName of the scan.
	MessagingServiceName string `json:"messagingServiceName,omitempty"`
	// The scanTypes that were requested for the scan.
	ScanTypes string `json:"scanTypes,omitempty"`
	// The destinations which EMA will send the scan results.
	Destinations string `json:"destinations,omitempty"`
	Type_ string `json:"type,omitempty"`
}
