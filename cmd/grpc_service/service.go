package main

import (
	"github.com/lifei6671/micro-service/kafka"
	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
	"log"
)

func main() {
	service := k8s.NewService(
		micro.Name("go.micro.srv.kafka"),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	if err := kafka.RegisterKafkaProducerHandler(service.Server(), new(kafka.ProducerService)); err != nil {
		log.Fatalf("注册服务失败 ->%s", err)
	}

	// Run server
	if err := service.Run(); err != nil {
		log.Fatalf("user service error: %v\n", err)
	}
}
