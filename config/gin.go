package config

var (
	GetGinMode = getGinMode
	GetGinUrl  = getGinUrl
)

func getGinMode() string {
	return getConfigString("order.gin.mode")

}
func getGinUrl() string {
	return getConfigString("order.gin.url")

}
