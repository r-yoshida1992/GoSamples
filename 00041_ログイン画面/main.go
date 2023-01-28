package main

import (
	"00041_login/app/router"
)

func main() {
	r := router.GetRouter()
	r.Run(":8080")
}
