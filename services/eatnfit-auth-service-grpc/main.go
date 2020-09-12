package main

import "github.com/vivaldy22/eatnfit-auth-service-grpc/config"

func main() {
	db, _ := config.InitDB()
	config.RunServer(db)
}
