package main

import (
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQl()
	utils.InitRedis()
	r := router.Router()
	r.Run(":8083") //8080
}
