package config

var (
	// GetRabbitMQConnectionString returns connection string from rabbitMQ section in toml file
	GetRabbitMQConnectionString = getRabbitMQConnectionString
)

func getRabbitMQConnectionString() string {
	return getConfigString("rabbitmq.connection_string")
}
