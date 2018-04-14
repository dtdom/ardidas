package main

import "ardidas/config"

func main() {
	r := Runner{}
	r.Initialize()
	r.Run(":" + config.MainConfig.Server.Port)
}
