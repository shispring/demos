/*
 * Strimzi Kafka Bridge API Reference
 *
 * The Strimzi Kafka Bridge provides a REST API for integrating HTTP based client applications with a Kafka cluster. You can use the API to create and manage consumers and send and receive records over HTTP rather than the native Kafka protocol. 
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Consumer struct {
	// The unique name for the consumer instance. The name is unique within the scope of the consumer group. The name is used in URLs. If a name is not specified, a randomly generated name is assigned.
	Name string `json:"name,omitempty"`
	// The allowable message format for the consumer, which can be `binary` (default) or `json`. The messages are converted into a JSON format. 
	Format string `json:"format,omitempty"`
	// Resets the offset position for the consumer. If set to `latest` (default), messages are read from the latest offset. If set to `earliest`, messages are read from the first offset.
	AutoOffsetReset string `json:"auto.offset.reset,omitempty"`
	// Sets the minimum amount of data, in bytes, for the consumer to receive. The broker waits until the data to send exceeds this amount. Default is `1` byte.
	FetchMinBytes int32 `json:"fetch.min.bytes,omitempty"`
	// Sets the maximum amount of time, in milliseconds, for the consumer to wait for messages for a request. If the timeout period is reached without a response, an error is returned. Default is `30000` (30 seconds).
	ConsumerRequestTimeoutMs int32 `json:"consumer.request.timeout.ms,omitempty"`
	// If set to `true` (default), message offsets are committed automatically for the consumer. If set to `false`, message offsets must be committed manually.
	EnableAutoCommit bool `json:"enable.auto.commit,omitempty"`
	// If set to `read_uncommitted` (default), all transaction records are retrieved, indpendent of any transaction outcome. If set to `read_committed`, the records from committed transactions are retrieved.
	IsolationLevel string `json:"isolation.level,omitempty"`
}