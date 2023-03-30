package main

import (
	"tugas8/database"
	"tugas8/routers"
)

func main(){
	database.StartDB()
	var PORT = ":8080"

	routers.StartServer().Run(PORT)

}