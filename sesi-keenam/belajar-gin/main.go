package main

import (
	"belajar-gin/routers"
)


func main() {
	routers.StartServer().Run(":8080")
}