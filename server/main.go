package main

import "pass-crypto/initialize"

func main() {
	router := initialize.InitRouter()

	router.Run("0.0.0.0:8888")
}
