package kafka

import (
	"context"
	"fmt"
	"io"
	"log"
)

type ProducerService struct {
}

func (s *ProducerService) SinglePublish(ctx context.Context, in *MessageRequest, out *MessageResponse) error {
	fmt.Printf("[MessageId=%s] - [Version=%d] - [ClientId=%s] - [Key=%s] - [MessageBody=%s]\n", in.MessageId, in.Version, in.ClientId, in.Key, string(in.MessageBody))

	out.Code = 0
	out.MessageId = in.MessageId
	out.Message = "OK"
	return nil
}

func (s *ProducerService) MultiPublish(ctx context.Context, stream KafkaProducer_MultiPublishStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("出现未处理异常 ->%s\n", err)
			return err
		}
		fmt.Printf("[MessageId=%s] - [Version=%d] - [ClientId=%s] - [Key=%s] - [MessageBody=%s]\n", req.MessageId, req.Version, req.ClientId, req.Key, string(req.MessageBody))

		err = stream.Send(&MessageResponse{Code: 0, MessageId: req.MessageId, Message: "OK"})
		if err != nil {
			log.Printf("响应 gRPC 客户端失败 ->%s", err)
		}
	}
}
