package config

var (
	GetGrpcConnectionString = getGrpcConnectionString
	// GetUserGRPCConnectionString returns user grpc connection string from grpc section in toml file
	GetUserGRPCConnectionString = getUserGRPCConnectionString
)

func getGrpcConnectionString() string {
	return getConfigString("grpc.connection_string")
}

func getUserGRPCConnectionString() string {
	return getConfigString("grpc.user_grpc_connection_string")
}
