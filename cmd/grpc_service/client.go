package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/lifei6671/micro-service/kafka"
	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
	"log"
)

func main() {
	//1.初始化注册中心

	//2.创建微服务
	service := k8s.NewService(
		micro.Name("go.micro.srv.kafka"),
	)
	service.Init()

	//3.初始化客户端
	client := kafka.NewKafkaProducerService("kafka.service", service.Client())

	req := &kafka.MessageRequest{
		MessageId:   uuid.New().String(),
		Version:     1,
		ClientId:    "lifei6671",
		MessageBody: []byte("使用定义protobuf文件还可以生成微服务的框架代码。"),
	}
	// 发送单个消息
	resp, err := client.SinglePublish(context.Background(), req)
	if err != nil {
		log.Printf("发送消息失败 -> %s", err)
		return
	}
	log.Printf("[Code=%d] - [Message=%s] - [MessageId=%s]", resp.Code, resp.Message, resp.MessageId)

	server, err := client.MultiPublish(context.Background())
	if err != nil {
		log.Printf("发送消息失败 -> %s", err)
		return
	}
	defer server.Close()

	log.Println("正在发送消息")
	for i := 0; i <= 10; i++ {
		req := &kafka.MessageRequest{
			MessageId:   uuid.New().String(),
			Version:     1,
			ClientId:    "lifeilin",
			MessageBody: []byte(fmt.Sprintf("使用定义protobuf文件还可以生成微服务的框架代码。 %d", i)),
		}
		if err := server.Send(req); err != nil {
			log.Printf("发送消息失败 -> %s", err)
			return
		}

		resp, err := server.Recv()
		if err != nil {
			log.Printf("接受消息失败 -> %s", err)
			return
		}
		log.Printf("[Code=%d] - [Message=%s] - [MessageId=%s]\n", resp.Code, resp.Message, resp.MessageId)
	}

}
