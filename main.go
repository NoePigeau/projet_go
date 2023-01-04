package main

import (
	"log"
	"net"
	"os"
	"project-go/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	pb "project-go/utils/grpc"
)

func gcrpServerRun(db *gorm.DB) {
	grpcListener, _ := net.Listen("tcp", ":5000")

	s := grpc.NewServer()
	grpcServer := pb.GetGrpcServer(db)

	pb.RegisterGRPCServer(s, grpcServer)

	s.Serve(grpcListener)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "user:password@tcp(127.0.0.1:3306)/project-go?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	r := gin.Default()
	routes.InitRoutes(r, db)

	go gcrpServerRun(db)
	r.Run(":3000")

}
