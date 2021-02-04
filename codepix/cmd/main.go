package main

import (
	"os"

	"github.com/erickrodrigs/codepix/codepix-go/application/grpc"
	"github.com/erickrodrigs/codepix/codepix-go/infrastructure/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
