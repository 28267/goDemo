package main

import router "github.com/28267/goDemo/gin-demo01/router"

func main() {

	r := router.Router()
	r.Run(":8888")

}
