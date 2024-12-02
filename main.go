package main

import (
	"ginTest/cmd"
	"ginTest/config"
)

func main1() {
	config.InitContentConfig()

	cmd.API()

}
