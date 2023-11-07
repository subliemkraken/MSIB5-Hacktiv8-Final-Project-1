package main

import "MSIB5-Hacktiv8-FinalProject1/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
