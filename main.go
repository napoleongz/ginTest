package main

import (
	"ginTest/config"
	"ginTest/cmd"
)

func main() {
	config.InitContentConfig()


	cmd.API()

}
