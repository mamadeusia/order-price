package config

var (
	GetRedisUrl = getRedisUrl

	GetRedisPass = getRedisPass
)

func getRedisUrl() string {
	return getConfigString("order.redis.url")
}
func getRedisPass() string {
	return getConfigString("order.redis.pass")
}
